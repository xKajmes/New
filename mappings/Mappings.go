package mappings

import (
	"new/controllers"
	"new/middleware"

	"github.com/gin-gonic/gin"
	//"github.com/gin-contrib/sessions"
	//"github.com/gin-contrib/sessions/cookie"
)

// HTML ..
func HTML() *gin.Engine {
	
	r := gin.Default()

		// Sesja

	//store := cookie.NewStore([]byte("aje"))
	//r.Use(sessions.Sessions("mysession", store))

		// Ladowanie plikow HTML 
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/html/*")
	r.StaticFile("/favicon.ico", "./favicon.ico")

		
	r.GET("/", controllers.HomePage)

	

	r.GET("/login", controllers.LoginPage)
	r.GET("/register", controllers.RegisterPage)
	r.GET("/logoutnow", controllers.LogoutNow)
	

		// Logowanie 
	auth := r.Group("/auth")
	{

		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.GET("/verify/:key", controllers.Verification)

	}

	r.GET("/home",middleware.TokenAuthMiddleware(), controllers.HomeLogin)

	

		// API

	panel := r.Group("/panel")
	// r.Use(MyMiddleware())
	{
		panel.GET("/student", controllers.ListStudent)
		//panel.GET("/logout", controllers.Logout)
		r.Use(middleware.TokenAuthMiddleware())
		panel.POST("/student", controllers.AddStudent)
		panel.GET("/student/:id", controllers.GetOne)
		panel.DELETE("/student/:id", controllers.DeleteStudent)
		panel.PUT("/student/:id", controllers.EditStudent)
		panel.POST("/completeme", controllers.CompleteMe)
		panel.GET("/lookrole", controllers.LookRolePage)
		panel.GET("/myid", controllers.MyID)
	}
	//r.GET("/teacher",middleware.AuthTeacher(), controllers.Teacher)
	teacher := r.Group("/teacher")
	r.Use(middleware.AuthTeacher())
	{
		teacher.GET("/panel", controllers.Teacher)
		teacher.GET("/class", controllers.TeacherClass)
		teacher.GET("/lookclass", controllers.LookClass) // Zobacz klase
		teacher.GET("/grades/:id", controllers.LookGrades) // Zobaczenie oceny
		teacher.GET("/grades", controllers.GradesPage) // HTML
		teacher.POST("/grades", controllers.AddGrade) // Dodawanie
		teacher.DELETE("/grades/:id", controllers.DelGrade) // Usuwanie

	}
	r.GET("/cookie", func(c *gin.Context) {
		cookie := "DZIALA"

		// set cookie outside the goroutine
		c.SetCookie("gin_cookie", cookie, 7200, "/", "", false, true)
		c.JSON(200, gin.H{"Komunikat": "aaa"})
	})


	return r
	
}