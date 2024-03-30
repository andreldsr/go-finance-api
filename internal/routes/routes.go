package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-finance-api/internal/controllers"
	"go-finance-api/internal/middlewares"
	"os"
	"strings"
	"time"
)

func HandleRequests() {
	r := gin.Default()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		return
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(os.Getenv("CORS_ORIGINS"), ","),
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/api/auth", controllers.Login)
	r.Use(middlewares.JWTAuthMiddleware())

	r.POST("/api/auth/register", controllers.Register)

	adminRoutes := r.Group("/api/admin")
	adminRoutes.Use(middlewares.AdminMiddleware())

	adminRoutes.GET("/user", controllers.FindUserList)
	_ = r.Run()
}
