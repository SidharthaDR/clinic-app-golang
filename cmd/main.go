package main

import (
	"github.com/SidharthaDR/golang-clinic-app/internal/controllers"
	"github.com/SidharthaDR/golang-clinic-app/internal/middleware"
	"github.com/SidharthaDR/golang-clinic-app/internal/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.ConnectDatabase()

	r := gin.Default()

	//public routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	//pvt routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/profile", func(c *gin.Context) {
		userID := c.GetUint("user_id")
		role := c.GetString("role")
		c.JSON(200, gin.H{
			"user_id": userID,
			"role":    role,
			"message": "Protected route access granted ✅",
		})
	})

	// Receptionist-only routes (Create and Delete patient)
	receptionistRoutes := protected.Group("/")
	receptionistRoutes.Use(middleware.RoleAuthorization("receptionist"))
	receptionistRoutes.POST("/patients", controllers.CreatePatient)
	receptionistRoutes.DELETE("/patients/:id", controllers.DeletePatient)

	// Doctor and Receptionist routes (Read and Update patient)
	doctorReceptionistRoutes := protected.Group("/")
	doctorReceptionistRoutes.Use(middleware.RoleAuthorization("doctor", "receptionist"))
	doctorReceptionistRoutes.GET("/patients", controllers.GetPatients)
	doctorReceptionistRoutes.GET("/patients/:id", controllers.GetPatient)
	doctorReceptionistRoutes.PUT("/patients/:id", controllers.UpdatePatient)

	r.Run(":8080")
}
