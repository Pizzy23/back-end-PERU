package welfare

import (
	"errors"
	"net/http"
	"peru/db"
	"peru/internal/interfaces"
	score "peru/internal/score/service"
	"time"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func CreateWelfareServices(c *gin.Context, data interfaces.InputWelfare) error {
	engine, exists := c.Get("db")
	if !exists {
		return errors.New("database connection not found")
	}

	welfareInput := db.Welfare{
		UserID:              data.UserID,
		Steps:               data.Steps,
		GoalSteps:           data.GoalSteps,
		HoursSlept:          data.HoursSlept,
		GoalHoursSlept:      data.GoalHoursSlept,
		MaxSmokingFrequency: data.MaxSmokingFrequency,
		SmokingFrequency:    data.SmokingFrequency,
		Score:               data.Score,
		Best:                data.Best,
		Date:                time.Now(),
	}
	newScore := score.CalculateWelfareScore(welfareInput)
	welfareInput.Score = newScore

	errChan := make(chan error, 1)

	go func() {
		err := db.Create(engine.(*xorm.Engine), &welfareInput)
		errChan <- err
	}()

	if err := <-errChan; err != nil {
		return err
	}

	return nil
}

func PullAllWelfareServices(c *gin.Context) {
	var welfare []db.Welfare
	engine, exists := c.Get("db")
	if !exists {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	if err := db.GetAll(engine.(*xorm.Engine), &welfare); err != nil {
		c.Set("Error", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", welfare)
	c.Status(http.StatusOK)
}

func PullWelfareServicesId(c *gin.Context, id int) {
	var welfare db.Welfare
	engine, exists := c.Get("db")
	if !exists {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	found, err := db.GetByID(engine.(*xorm.Engine), &welfare, int64(id))
	if err != nil {
		c.Set("Error", "Error retrieving Welfare")
		c.Status(http.StatusInternalServerError)
		return
	}
	if !found {
		c.Set("Error", "Welfare not found")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", welfare)
	c.Status(http.StatusOK)
}

func EditWelfareService(c *gin.Context, data interfaces.InputWelfareEdit) error {
	welfareInput := db.Welfare{
		UserID:              data.UserID,
		Steps:               data.Steps,
		GoalSteps:           data.GoalSteps,
		HoursSlept:          data.HoursSlept,
		GoalHoursSlept:      data.GoalHoursSlept,
		MaxSmokingFrequency: data.MaxSmokingFrequency,
		SmokingFrequency:    data.SmokingFrequency,
		Score:               data.Score,
		Best:                data.Best,
		Date:                time.Now(),
	}

	engine, exists := c.Get("db")
	if !exists {
		return errors.New("database connection not found")
	}

	condition := &db.Welfare{ID: data.ID}
	if err := db.Update(engine.(*xorm.Engine), &welfareInput, condition); err != nil {
		return err
	}
	return nil
}
