package health

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

func CreateHealthServices(c *gin.Context, data interfaces.InputHealth) error {
	engine, exists := c.Get("db")
	if !exists {
		return errors.New("database connection not found")
	}

	healthInput := db.Health{
		UserID:         data.UserID,
		CaloriesBurned: data.CaloriesBurned,
		GoalCalories:   data.GoalCalories,
		HealthyFood:    data.HealthyFood,
		UnhealthyFood:  data.UnhealthyFood,
		Score:          data.Score,
		Best:           data.Best,
		Date:           time.Now(),
	}

	newScore := score.CalculateHealthScore(healthInput)
	healthInput.Score = newScore

	errChan := make(chan error, 1)

	go func() {
		err := db.Create(engine.(*xorm.Engine), &healthInput)
		errChan <- err
	}()

	if err := <-errChan; err != nil {
		return err
	}

	return nil
}

func PullAllHealthServices(c *gin.Context) {
	var health []db.Health
	engine, exists := c.Get("db")
	if !exists {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	if err := db.GetAll(engine.(*xorm.Engine), &health); err != nil {
		c.Set("Error", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", health)
	c.Status(http.StatusOK)
}

func PullHealthServicesId(c *gin.Context, id int) {
	var health db.Health
	engine, exists := c.Get("db")
	if !exists {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	found, err := db.GetByID(engine.(*xorm.Engine), &health, int64(id))
	if err != nil {
		c.Set("Error", "Error retrieving health")
		c.Status(http.StatusInternalServerError)
		return
	}
	if !found {
		c.Set("Error", "Health not found")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", health)
	c.Status(http.StatusOK)
}

func EditHealthService(c *gin.Context, data interfaces.InputHealthEdit) error {
	healthInput := db.Health{
		UserID:         data.UserID,
		CaloriesBurned: data.CaloriesBurned,
		GoalCalories:   data.GoalCalories,
		HealthyFood:    data.HealthyFood,
		UnhealthyFood:  data.UnhealthyFood,
		Score:          data.Score,
		Best:           data.Best,
		Date:           time.Now(),
	}

	newScore := score.CalculateHealthScore(healthInput)
	healthInput.Score = newScore

	engine, exists := c.Get("db")
	if !exists {
		return errors.New("database connection not found")
	}

	condition := &db.Health{ID: data.ID}
	if err := db.Update(engine.(*xorm.Engine), &healthInput, condition); err != nil {
		return err
	}
	return nil
}
