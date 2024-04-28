package routes

import (
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api/v1")
	{
		login := main.Group("auth")
		{
			login.POST("/login", controllers.Login)
		}
		users := main.Group("/users")
		{
			users.GET("/:id", controllers.GetUser)
			users.GET("", controllers.GetAllUsers)
			users.POST("/", controllers.CreateUser)
			users.DELETE("/:id", controllers.DeleteUser)
			users.PUT("/", controllers.UpdateUser)
		}
		thesis := main.Group("/thesis")
		{
			thesis.GET("/:id", controllers.GetThesis)
			thesis.POST("/", controllers.CreateThesis)
			thesis.GET("/own/:teacher_id", controllers.GetOwnThesis)
			thesis.DELETE("/:id", controllers.DeleteThesis)
			thesis.PUT("/", controllers.UpdateThesis)
			thesis.GET("/allrequested/:teacher_id", controllers.GetAllRequested)
		}
		thesisOffice := main.Group("/thesis-office")
		{
			thesisOffice.GET("/", controllers.GetAllthesis)
			thesisOffice.PUT("/", controllers.UpdateReqThesis)
		}
		student := main.Group("/student")
		{
			student.GET("/", controllers.GetStudentThesis)
			student.PUT("/", controllers.StudentReqThesis)
			student.GET("/:student_id", controllers.GetStudentShowReq)
		}
		process := main.Group("/process")
		{
			process.POST("/", controllers.InsertProcess)
			process.GET("/student/:student_id", controllers.GetProcessStudent)
			process.POST("/detail", controllers.InsertProcessDetail)
			process.GET("/detail/:process_id", controllers.GetProcessDetail)
			process.PUT("/student", controllers.UpdateFeedbackStudent)
		}

		processTandS := main.Group("/processTandS")
		{
			processTandS.GET("/", controllers.GetProcessAll)
			processTandS.GET("/teacher/:teacher_id", controllers.GetProcessTeacher)
			processTandS.PUT("/", controllers.UpdateProcessOne)
			processTandS.PUT("/all", controllers.UpdateProcessAll)
			processTandS.PUT("/teacher", controllers.UpdateFeedbackTeacher)
		}

		// ws := main.Group("/ws")
		// {
		// 	ws.GET("/", controllers.WShandler)
		// }
	}

	return router
}
