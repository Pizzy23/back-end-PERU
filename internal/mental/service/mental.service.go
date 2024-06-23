package mental

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

func CreateMentalServices(c *gin.Context, data interfaces.InputMental) error {
	engine, exists := c.Get("db")
	if !exists {
		return errors.New("database connection not found")
	}

	MentalInput := db.Mental{
		UserID:           data.UserID,
		QualityOfSleep:   data.QualityOfSleep,
		StressScore:      data.SmokingFrequency,
		DepressionScore:  data.DepressionScore,
		MindfulnessScore: data.MindfulnessScore,
		SmokingFrequency: data.SmokingFrequency,
		Score:            data.Score,
		Best:             data.Best,
		Date:             time.Now(),
	}
	newScore := score.CalculateMentalScore(MentalInput)
	MentalInput.Score = newScore

	errChan := make(chan error, 1)

	go func() {
		err := db.Create(engine.(*xorm.Engine), &MentalInput)
		errChan <- err
	}()

	if err := <-errChan; err != nil {
		return err
	}

	return nil
}

func PullAllMentalServices(c *gin.Context) {
	var Mental []db.Mental
	engine, exists := c.Get("db")
	if !exists {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	if err := db.GetAll(engine.(*xorm.Engine), &Mental); err != nil {
		c.Set("Error", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", Mental)
	c.Status(http.StatusOK)
}

func PullMentalServicesId(c *gin.Context, id int) {
	var Mental db.Mental
	engine, exists := c.Get("db")
	if !exists {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	found, err := db.GetByID(engine.(*xorm.Engine), &Mental, int64(id))
	if err != nil {
		c.Set("Error", "Error retrieving Mental")
		c.Status(http.StatusInternalServerError)
		return
	}
	if !found {
		c.Set("Error", "Mental not found")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", Mental)
	c.Status(http.StatusOK)
}

func EditMentalService(c *gin.Context, data interfaces.InputMentalEdit) error {
	MentalInput := db.Mental{
		UserID:           data.UserID,
		QualityOfSleep:   data.QualityOfSleep,
		StressScore:      data.SmokingFrequency,
		DepressionScore:  data.DepressionScore,
		MindfulnessScore: data.MindfulnessScore,
		SmokingFrequency: data.SmokingFrequency,
		Score:            data.Score,
		Best:             data.Best,
		Date:             time.Now(),
	}
	newScore := score.CalculateMentalScore(MentalInput)
	MentalInput.Score = newScore

	engine, exists := c.Get("db")
	if !exists {
		return errors.New("database connection not found")
	}

	condition := &db.Mental{ID: data.ID}
	if err := db.Update(engine.(*xorm.Engine), &MentalInput, condition); err != nil {
		return err
	}
	return nil
}
