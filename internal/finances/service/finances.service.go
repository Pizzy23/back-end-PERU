package finances

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

func CreateFinancesServices(c *gin.Context, data interfaces.InputFinances) error {
	engine, exists := c.Get("db")
	if !exists {
		return errors.New("database connection not found")
	}

	financesInput := db.Finances{
		UserID:    data.UserID,
		Aliment:   data.Aliment,
		Life:      data.Life,
		Transport: data.Transport,
		Shopping:  data.Shopping,
		Service:   data.Service,
		Best:      data.Best,
		Date:      time.Now(),
	}

	newScore := score.CalculateFinanceScore(financesInput)
	financesInput.Score = newScore

	errChan := make(chan error, 1)

	go func() {
		err := db.Create(engine.(*xorm.Engine), &financesInput)
		errChan <- err
	}()

	if err := <-errChan; err != nil {
		return err
	}

	return nil
}

func PullAllFinancesServices(c *gin.Context) {
	var finances []db.Finances
	engine, exists := c.Get("db")
	if !exists {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	if err := db.GetAll(engine.(*xorm.Engine), &finances); err != nil {
		c.Set("Error", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", finances)
	c.Status(http.StatusOK)
}

func PullFinancesServicesId(c *gin.Context, id int) {
	var finances db.Finances
	engine, exists := c.Get("db")
	if !exists {
		c.Set("Error", "Database connection not found")
		c.Status(http.StatusInternalServerError)
		return
	}

	found, err := db.GetByID(engine.(*xorm.Engine), &finances, int64(id))
	if err != nil {
		c.Set("Error", "Error retrieving finances")
		c.Status(http.StatusInternalServerError)
		return
	}
	if !found {
		c.Set("Error", "finances not found")
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", finances)
	c.Status(http.StatusOK)
}

func EditFinancesService(c *gin.Context, data interfaces.InputFinancesEdit) error {
	financesInput := db.Finances{
		UserID:    data.UserID,
		Aliment:   data.Aliment,
		Life:      data.Life,
		Transport: data.Transport,
		Shopping:  data.Shopping,
		Service:   data.Service,
		Best:      data.Best,
		Date:      time.Now(),
	}

	newScore := score.CalculateFinanceScore(financesInput)
	financesInput.Score = newScore

	engine, exists := c.Get("db")
	if !exists {
		return errors.New("database connection not found")
	}

	condition := &db.Finances{ID: data.ID}
	if err := db.Update(engine.(*xorm.Engine), &financesInput, condition); err != nil {
		return err
	}
	return nil
}
