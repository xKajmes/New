package controllers

import (
	"github.com/gin-gonic/gin"
	//"../Translate"
)

// HomePage ..
func HomePage(c *gin.Context) {		// Strona dla niezalagowanych
	c.HTML(200, "homepage.tmpl", gin.H{
		
	})
}
// HomeLogin ...
func HomeLogin(c *gin.Context) {	// Strona dla zalogowanych
	// token := c.Request.Header["Authorization"]
	// fmt.Println("UWAGA: /n")
	// fmt.Println(token)

	c.HTML(200, "homelogin.tmpl", gin.H{
		
	})

}
// Teacher ...
func Teacher(c *gin.Context) {	// Strona dla zalogowanych

	c.HTML(200, "teacher.tmpl", gin.H{
		
	})

}
// TeacherClass ...
func TeacherClass (c *gin.Context) {	// Strona dla zalogowanych

	c.HTML(200, "teacherclass.tmpl", gin.H{
		
	})

}
// LoginPage ..
func LoginPage(c *gin.Context) {	// Strona logowania
	// var lang string = "EN"
	// Translate.Language(lang)


	// session := sessions.Default(c)
	// session.Set("count", "ajj")
	// session.Save()



	c.HTML(200, "login.tmpl", gin.H{
		// "login": Translate.login,
		// "singup": Translate.singup,
		// "register": Translate.register,
		// "password": Translate.password,
		// "SingUpGoogle": Translate.SingUpGoogle,
	})

}
//RegisterPage ...
func RegisterPage(c *gin.Context) {	// Strona rejestracji
	c.HTML(200, "register.tmpl", gin.H{
		
	})

}
// GradesPage ...
func GradesPage(c *gin.Context) {	// Strona rejestracji
	c.HTML(200, "gradespage.tmpl", gin.H{
		
	})

}