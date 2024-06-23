package mental

import (
	"net/http"
	"peru/internal/interfaces"
	mental "peru/internal/mental/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create New Mental
// @Description Create New Mental
// @Tags Mental
// @Accept json
// @Produce json
// @Param request body interfaces.InputMental true "Data for create new Mental"
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.Mental "List of All Mental"
// @Router /api/mental [post]
func CreateMental(c *gin.Context) {
	var mentalInput interfaces.InputMental

	if err := c.ShouldBindJSON(&mentalInput); err != nil {
		c.Set("Error", "Parameters are invalid, need a JSON")
		c.Status(http.StatusInternalServerError)
		return
	}

	if err := mental.CreateMentalServices(c, mentalInput); err != nil {
		c.Set("Error", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", "mental created successfully")
	c.Status(http.StatusOK)

}

// @Summary Retrieve All Mental
// @Description Retrieve all Mental
// @Tags Mental
// @Accept json
// @Produce json
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.Mental "List of All Mental"
// @Router /api/all-mental [get]
func PullAllMental(c *gin.Context) {
	mental.PullAllMentalServices(c)
}

// @Summary Retrieve mental by ID
// @Description Retrieve an mental by its ID
// @Tags Mental
// @Accept json
// @Produce json
// @Param id path int true "mental ID"
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.Mental "mental Details"
// @Router /api/mental/{id} [get]
func PullMentalId(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Set("Error", "Invalid asset ID")
		c.Status(http.StatusBadRequest)
		return
	}
	mental.PullMentalServicesId(c, id)
}

// @Summary Retrieve mental by ID
// @Description Retrieve an mental by its ID
// @Tags Mental
// @Accept json
// @Produce json
// @Param request body interfaces.InputMentalEdit true "Edit mental"
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.Mental "mental Details"
// @Router /api/mental [put]
func EditMental(c *gin.Context) {
	var mentalInput interfaces.InputMentalEdit
	if err := c.ShouldBindJSON(&mentalInput); err != nil {
		c.Set("Response", "Parameters are invalid, need a JSON")
		c.Status(http.StatusBadRequest)
		return
	}

	if err := mental.EditMentalService(c, mentalInput); err != nil {
		c.Set("Response", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", "Event created successfully")
	c.Status(http.StatusOK)
}
