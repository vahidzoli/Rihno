package routes

import (
	"Rihno/controllers"
	authController "Rihno/controllers/auth"
	"Rihno/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Errors())
	grp1 := r.Group("/api")
	{
		grp1.POST("register", authController.Register)
		grp1.POST("login", authController.Login)

		grp2 := grp1.Group("/users")
		grp2.Use(middleware.CheckLogin())
		{
			grp2.GET("/", controllers.GetUsers)
			grp2.GET("user/:id", controllers.GetUserByID)
			grp2.PATCH("user/:id", controllers.UpdateUser)
			grp2.DELETE("user/:id", controllers.DeleteUser)
		}

		grp3 := grp1.Group("/projects")
		grp3.Use(middleware.CheckLogin())
		{
			grp3.GET("/", controllers.GetProjects)
			grp3.POST("/project", controllers.CreateProject)
			grp3.GET("project/:id", controllers.GetProjectByID)
			grp3.PATCH("project/:id", controllers.UpdateProject)
			grp3.DELETE("project/:id", controllers.DeleteProject)

			grp3.GET("project/:id/plans", controllers.GetPlans)
			grp3.POST("project/:id/plan", controllers.CreatePlan)
			grp3.GET("project/:id/plan/:pid", controllers.GetPlanByID)
			grp3.PATCH("project/:id/plan/:pid", controllers.UpdatePlan)
			grp3.DELETE("project/:id/plan/:pid", controllers.DeletePlan)

			grp3.GET("project/:id/resolutions", controllers.GetResolutions)
		}
	}
	return r
}
