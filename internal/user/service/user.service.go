package user

import (
	"errors"
	"net/http"
	"peru/db"
	"peru/internal/interfaces"
	score "peru/internal/score/service"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

func CreateUserServices(c *gin.Context, data interfaces.InputUser) error {
	engine, exists := c.Get("db")
	if !exists {
		return errors.New("database connection not found")
	}
	newBMI := score.CalculateBMI(data.Weight, data.Height)

	UserInput := db.User{
		Name:     data.Name,
		Email:    data.Email,
		Finances: 0,
		Health:   0,
		Mental:   0,
		Welfare:  0,
		Score:    0,
		Weight:   data.Weight,
		Height:   data.Height,
		BMI:      newBMI,
	}

	errChan := make(chan error, 1)

	go func() {
		err := db.Create(engine.(*xorm.Engine), &UserInput)
		errChan <- err
	}()

	if err := <-errChan; err != nil {
		return err
	}

	return nil
}

func PullAllUserServices(c *gin.Context) {
	var User []db.User
	engine, exists := c.Get("db")
	if !exists {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	if err := db.GetAll(engine.(*xorm.Engine), &User); err != nil {
		c.Set("Error", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", User)
	c.Status(http.StatusOK)
}

func PullUserServicesId(c *gin.Context, id int) {
	var User db.User
	engine, exists := c.Get("db")
	if !exists {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	found, err := db.GetByID(engine.(*xorm.Engine), &User, int64(id))
	if err != nil {
		c.Set("Error", "Error retrieving User")
		c.Status(http.StatusInternalServerError)
		return
	}
	if !found {
		c.Set("Error", "User not found")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", User)
	c.Status(http.StatusOK)
}

func EditUserService(c *gin.Context, data interfaces.InputUserEdit) error {
	UserInput := db.User{
		Name:     data.Name,
		Email:    data.Email,
		Finances: data.Finances,
		Health:   data.Health,
		Mental:   data.Mental,
		Welfare:  data.Welfare,
		Score:    data.Score,
		Weight:   data.Weight,
		Height:   data.Height,
		BMI:      data.BMI,
	}

	newScore := score.CalculateUserScore(UserInput.Finances, UserInput.Health, UserInput.Mental, UserInput.Welfare)
	UserInput.Score = newScore

	engine, exists := c.Get("db")
	if !exists {
		return errors.New("database connection not found")
	}

	condition := &db.User{ID: data.ID}
	if err := db.Update(engine.(*xorm.Engine), &UserInput, condition); err != nil {
		return err
	}
	return nil
}

func CheckBMI(c *gin.Context, bmiInput interfaces.InputBMI) {
	check := score.BMICategory(bmiInput.BMI)
	c.Set("Response", check)
	c.Status(http.StatusOK)
}
