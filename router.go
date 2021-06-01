package main

import (
	"fmt"
	"github.com/Bastorx/test_lbc/controllers"
	"github.com/Bastorx/test_lbc/logging"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(logging.GinLogHandler(), gin.Recovery())

	// TEMP
	gin.SetMode(gin.DebugMode)

	/* Monitoring */
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.GET("/logout", controllers.Logout)
	}

	user := r.Group("/user", isLogged)
	{
		user.GET("/:userId", controllers.GetUser)
		user.GET("/", controllers.GetAllUsers)
		user.POST("/", controllers.CreateUser)
		user.PUT("/:userId", controllers.UpdateUser)
		user.DELETE("/:userId", controllers.DeleteUser)
	}

	category := r.Group("/category", isLogged)
	{
		category.GET("/", controllers.GetAllCategories)
		category.POST("/", controllers.CreateCategory)
		category.PUT("/:categoryId", controllers.UpdateCategory)
		category.DELETE("/:categoryId", controllers.DeleteCategory)
	}

	advert := r.Group("/advert", isLogged)
	{
		advert.GET("/", controllers.GetAllAdverts)
		advert.POST("/", controllers.CreateAdvert)
		advert.PUT("/:categoryId", controllers.UpdateAdvert)
		advert.DELETE("/:categoryId", controllers.DeleteAdvert)
	}

	return r
}

func isLogged(c *gin.Context) {
	//clientIP := c.ClientIP()
	//logging.Logger.WithField("clientIP", clientIP).Info("Checking localhost by IP")
	//if clientIP != "127.0.0.1" && clientIP != "::1" {
	//	err := fmt.Errorf("%s is not authorized for this action", clientIP)
	//	abortWithError(c, http.StatusUnauthorized, err, "Unauthorized IP")
	//}
}

func abortWithError(c *gin.Context, status int, err error, errTitle string) {
	errDetail := ""
	if err != nil {
		errDetail = fmt.Sprintf("%s", err)
	}
	logging.Logger.WithField("error", err).Error(errTitle)
	c.AbortWithStatusJSON(status, gin.H{
		"error": gin.H{
			"title":  errTitle,
			"detail": errDetail,
		},
	})
}