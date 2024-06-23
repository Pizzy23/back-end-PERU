package user

import (
	"net/http"
	"peru/internal/interfaces"
	user "peru/internal/user/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create New User
// @Description Create New User
// @Tags User
// @Accept json
// @Produce json
// @Param request body interfaces.InputUser true "Data for create new User"
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.User "List of All User"
// @Router /api/user [post]
func CreateUser(c *gin.Context) {
	var UserInput interfaces.InputUser

	if err := c.ShouldBindJSON(&UserInput); err != nil {
		c.Set("Error", "Parameters are invalid, need a JSON")
		c.Status(http.StatusInternalServerError)
		return
	}
	if err := c.ShouldBindJSON(&UserInput); err != nil {

		return
	}

	if err := user.CreateUserServices(c, UserInput); err != nil {

		return
	}
	if err := user.CreateUserServices(c, UserInput); err != nil {
		c.Set("Error", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Set("Response", "User created successfully")
	c.Status(http.StatusOK)

}

// @Summary Retrieve All User
// @Description Retrieve all User
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.User "List of All User"
// @Router /api/all-user [get]
func PullAllUser(c *gin.Context) {
	user.PullAllUserServices(c)
}

// @Summary Retrieve User by ID
// @Description Retrieve an User by its ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.User "User Details"
// @Router /api/user/{id} [get]
func PullUserId(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Set("Error", "Invalid asset ID")
		c.Status(http.StatusBadRequest)
		return
	}
	user.PullUserServicesId(c, id)
}

// @Summary Retrieve User by ID
// @Description Retrieve an User by its ID
// @Tags User
// @Accept json
// @Produce json
// @Param request body interfaces.InputUserEdit true "Edit User"
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.User "User Details"
// @Router /api/user [put]
func EditUser(c *gin.Context) {
	var UserInput interfaces.InputUserEdit
	if err := c.ShouldBindJSON(&UserInput); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Parameters are invalid, need a JSON"})
		return
	}

	if err := user.EditUserService(c, UserInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Set("Response", "Event created successfully")
	c.Status(http.StatusOK)
}

// @Summary BMI Checker
// @Description Check your bmi
// @Tags User
// @Accept json
// @Produce json
// @Param request body interfaces.InputBMI true "BMI :)"
// @Param Authorization header string true "Auth Token" default(Bearer <token>)
// @Success 200 {object} db.User "User Details"
// @Router /api/bmi [get]
func BMI(c *gin.Context) {
	var bmiInput interfaces.InputBMI
	if err := c.ShouldBindJSON(&bmiInput); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Parameters are invalid, need a JSON"})
		return
	}
	user.CheckBMI(c, bmiInput)
}
