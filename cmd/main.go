package main

import (
	"go-profile-service-magangku/config"
	"go-profile-service-magangku/internal/handler"
	"go-profile-service-magangku/internal/middleware"
	"go-profile-service-magangku/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := config.ConnectDB()

	profileRepo := repository.NewProfileRepository(db)
	profileHandler := handler.NewProfileHandler(profileRepo)

	r := gin.Default()

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		userProfile := api.Group("/profile")
		userProfile.Use(middleware.UserOnly())
		{
			userProfile.GET("", profileHandler.GetMyProfile)
			userProfile.POST("", profileHandler.CreateMyProfile)
			userProfile.PUT("", profileHandler.UpdateMyProfile)
		}
	}

	r.Run(":" + config.AppConfig.AppPort)
}
