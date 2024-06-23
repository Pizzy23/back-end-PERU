package middleware

import (
	"peru/db"
	_ "peru/docs"
	finances "peru/internal/finances/handler"
	health "peru/internal/health/handler"
	mental "peru/internal/mental/handler"
	"peru/internal/pdf"
	user "peru/internal/user/handler"
	welfare "peru/internal/welfare/handler"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Peru Hackaton
// @version 1.0
// @description API
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /api
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Use(CORSConfig())
	r.Use(ResponseHandler())

	r.Use(func(c *gin.Context) {
		c.Set("db", db.Repo)
		c.Next()
	})

	//Use response, but not Token
	r.GET("/token", generateTokenHandler)

	auth := r.Group("/api")
	auth.Use(authMiddleware)

	//Response and token service
	auth.POST("/finances", finances.CreateFinances)
	auth.GET("/all-finances", finances.PullAllFinances)
	auth.GET("/finances/:id", finances.PullFinancesId)
	auth.PUT("/finances", finances.EditFinances)

	auth.POST("/health", health.CreateHealth)
	auth.GET("/all-health", health.PullAllHealth)
	auth.GET("/health/:id", health.PullHealthId)
	auth.PUT("/health", health.EditHealth)

	auth.POST("/mental", mental.CreateMental)
	auth.GET("/all-mental", mental.PullAllMental)
	auth.GET("/mental/:id", mental.PullMentalId)
	auth.PUT("/mental", mental.EditMental)

	auth.POST("/welfare", welfare.CreateWelfare)
	auth.GET("/all-welfare", welfare.PullAllWelfare)
	auth.GET("/welfare/:id", welfare.PullWelfareId)
	auth.PUT("/welfare", welfare.EditWelfare)

	auth.POST("/user", user.CreateUser)
	auth.GET("/all-user", user.PullAllUser)
	auth.GET("/user/:id", user.PullUserId)
	auth.GET("/bmi/:id", user.BMI)
	auth.PUT("/user", user.EditUser)

	auth.POST("/upload", pdf.UploadHandler)
	return r
}
