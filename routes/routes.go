package routes

import (
	"MTBS/controllers"
	auth "MTBS/login"

	"MTBS/middleware"

	"github.com/gin-gonic/gin"
)

func SettingUpRoutes(r *gin.Engine) *gin.Engine {

	register := r.Group("/register")
	{
		register.POST("/user", controllers.RegisterCustomer)
		register.POST("/admin", controllers.RegisterAdmin)

	}
	login := r.Group("/login")
	{
		login.POST("/user", auth.LoginUser)
		login.POST("/Admin", auth.LoginAdmin)
	}

	r.POST("/admin/movies", middleware.RequirePrivilege("create_movie"), controllers.CreateMovie)
	r.POST("/admin/schedules", middleware.RequirePrivilege("add_schedule"), controllers.AddSchedule)

	r.GET("/customer/movies", middleware.RequirePrivilege("view_movies"), controllers.GetAllMovies)
	r.GET("/customer/schedules", middleware.RequirePrivilege("view_schedules"), controllers.GetSchedulesByMovie)
	r.GET("/customer/seats", middleware.RequirePrivilege("view_seatsbymovie"), controllers.ViewAvailableSeatsByMovie)
	r.POST("/customer/book", middleware.RequirePrivilege("book_ticket"), controllers.BookTicket)
	r.GET("/customer/report",middleware.RequirePrivilege("get_report"),controllers.GetReportByID)

	return r

}
