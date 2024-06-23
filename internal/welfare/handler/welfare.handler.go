package welfare

import (
	"net/http"
	"peru/internal/interfaces"
	welfare "peru/internal/welfare/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create New welfare
// @Description Create New welfare
// @Tags welfare
// @Accept json
// @Produce json
// @Param request body interfaces.InputWelfare true "Data for create new welfare"
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.Welfare "List of All welfare"
// @Router /api/welfare [post]
func CreateWelfare(c *gin.Context) {
	var welfareInput interfaces.InputWelfare

	if err := c.ShouldBindJSON(&welfareInput); err != nil {
		c.Set("Error", "Parameters are invalid, need a JSON")
		c.Status(http.StatusInternalServerError)
		return
	}
	if err := c.ShouldBindJSON(&welfareInput); err != nil {

		return
	}

	if err := welfare.CreateWelfareServices(c, welfareInput); err != nil {

		return
	}
	if err := welfare.CreateWelfareServices(c, welfareInput); err != nil {
		c.Set("Error", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", "welfare created successfully")
	c.Status(http.StatusOK)

}

// @Summary Retrieve All welfare
// @Description Retrieve all welfare
// @Tags welfare
// @Accept json
// @Produce json
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.Welfare "List of All welfare"
// @Router /api/all-welfare [get]
func PullAllWelfare(c *gin.Context) {
	welfare.PullAllWelfareServices(c)
}

// @Summary Retrieve welfare by ID
// @Description Retrieve an welfare by its ID
// @Tags welfare
// @Accept json
// @Produce json
// @Param id path int true "welfare ID"
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.Welfare "welfare Details"
// @Router /api/welfare/{id} [get]
func PullWelfareId(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Set("Error", "Invalid asset ID")
		c.Status(http.StatusBadRequest)
		return
	}
	welfare.PullWelfareServicesId(c, id)
}

// @Summary Retrieve welfare by ID
// @Description Retrieve an welfare by its ID
// @Tags welfare
// @Accept json
// @Produce json
// @Param request body interfaces.InputWelfareEdit true "Edit welfare"
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.Welfare "welfare Details"
// @Router /api/welfare [put]
func EditWelfare(c *gin.Context) {
	var welfareInput interfaces.InputWelfareEdit
	if err := c.ShouldBindJSON(&welfareInput); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Parameters are invalid, need a JSON"})
		return
	}

	if err := welfare.EditWelfareService(c, welfareInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Set("Response", "Event created successfully")
	c.Status(http.StatusOK)
}
