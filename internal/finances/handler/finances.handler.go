package finances

import (
	"net/http"
	finances "peru/internal/finances/service"
	"peru/internal/interfaces"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create New Finances
// @Description Create New Finances
// @Tags Finances
// @Accept json
// @Produce json
// @Param request body interfaces.InputFinances true "Data for create new Finances"
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.Finances "List of All Assets"
// @Router /api/finances [post]
func CreateFinances(c *gin.Context) {
	var financesInput interfaces.InputFinances

	if err := c.ShouldBindJSON(&financesInput); err != nil {
		c.Set("Error", "Parameters are invalid, need a JSON")
		c.Status(http.StatusInternalServerError)
		return
	}
	if err := c.ShouldBindJSON(&financesInput); err != nil {

		return
	}

	if err := finances.CreateFinancesServices(c, financesInput); err != nil {

		return
	}
	if err := finances.CreateFinancesServices(c, financesInput); err != nil {
		c.Set("Error", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", "Finances created successfully")
	c.Status(http.StatusOK)

}

// @Summary Retrieve All Finances
// @Description Retrieve all Finances
// @Tags Finances
// @Accept json
// @Produce json
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.Finances "List of All Finances"
// @Router /api/all-finances [get]
func PullAllFinances(c *gin.Context) {
	finances.PullAllFinancesServices(c)
}

// @Summary Retrieve Finances by ID
// @Description Retrieve an Finances by its ID
// @Tags Finances
// @Accept json
// @Produce json
// @Param id path int true "Finances ID"
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.Finances "Finances Details"
// @Router /api/finances/{id} [get]
func PullFinancesId(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Set("Error", "Invalid asset ID")
		c.Status(http.StatusBadRequest)
		return
	}
	finances.PullFinancesServicesId(c, id)
}

// @Summary Retrieve Finances by ID
// @Description Retrieve an Finances by its ID
// @Tags Finances
// @Accept json
// @Produce json
// @Param request body interfaces.InputFinancesEdit true "Edit Finances"
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.Finances "Finances Details"
// @Router /api/finances [put]
func EditFinances(c *gin.Context) {
	var financesInput interfaces.InputFinancesEdit
	if err := c.ShouldBindJSON(&financesInput); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Parameters are invalid, need a JSON"})
		return
	}

	if err := finances.EditFinancesService(c, financesInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Set("Response", "Event created successfully")
	c.Status(http.StatusOK)
}
