package health

import (
	"net/http"
	health "peru/internal/health/service"
	"peru/internal/interfaces"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create New Health
// @Description Create New Health
// @Tags Health
// @Accept json
// @Produce json
// @Param request body interfaces.InputHealth true "Data for create new Health"
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.Health "List of All Health"
// @Router /api/health [post]
func CreateHealth(c *gin.Context) {
	var healthInput interfaces.InputHealth

	if err := c.ShouldBindJSON(&healthInput); err != nil {
		c.Set("Error", "Parameters are invalid, need a JSON")
		c.Status(http.StatusInternalServerError)
		return
	}
	if err := c.ShouldBindJSON(&healthInput); err != nil {

		return
	}

	if err := health.CreateHealthServices(c, healthInput); err != nil {

		return
	}
	if err := health.CreateHealthServices(c, healthInput); err != nil {
		c.Set("Error", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", "health created successfully")
	c.Status(http.StatusOK)

}

// @Summary Retrieve All health
// @Description Retrieve all health
// @Tags health
// @Accept json
// @Produce json
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.Health "List of All Health"
// @Router /api/all-health [get]
func PullAllHealth(c *gin.Context) {
	health.PullAllHealthServices(c)
}

// @Summary Retrieve health by ID
// @Description Retrieve an health by its ID
// @Tags health
// @Accept json
// @Produce json
// @Param id path int true "health ID"
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.Health "Health Details"
// @Router /api/health/{id} [get]
func PullHealthId(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Set("Error", "Invalid asset ID")
		c.Status(http.StatusBadRequest)
		return
	}
	health.PullHealthServicesId(c, id)
}

// @Summary Retrieve Health by ID
// @Description Retrieve an Health by its ID
// @Tags Health
// @Accept json
// @Produce json
// @Param request body interfaces.InputHealthEdit true "Edit Health"
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.Health "Health Details"
// @Router /api/Health [put]
func EditHealth(c *gin.Context) {
	var HealthInput interfaces.InputHealthEdit
	if err := c.ShouldBindJSON(&HealthInput); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Parameters are invalid, need a JSON"})
		return
	}

	if err := health.EditHealthService(c, HealthInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Set("Response", "Event created successfully")
	c.Status(http.StatusOK)
}
