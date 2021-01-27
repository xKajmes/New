package controllers

import (
	"fmt"
	//"net/http"
	"new/auth"
	"new/config"
	"new/models"

	"github.com/gin-gonic/gin"
	//"../Translate"
)

// Login ...
func Login(c *gin.Context) {	
	var user models.User
	c.BindJSON(&user)
	err := models.Login(&user)
	if(err == "Zalogowano"){

		ts, err := auth.CreateToken(user.ID)
		if err != nil {
			c.JSON(422, err.Error())
		  	return
	   	}
		look := models.LookVerify(user.ID)
		if look != true {
			c.JSON(303, gin.H{"error:": "Nie zwerifykiowaleś konta!"})
			fmt.Print("Dont verify")
		} 
		fmt.Print("verify")
		token := ts.AccessToken
		c.SetCookie("Auth", token, 10800, "/", "", false, true)	
		//fmt.Println(look)
		iduser, _ := auth.ExtractTokenMetadata(token)
		c.JSON(200, gin.H{
			"look": look,
			"token": token,
			"userid": iduser.UserID ,
		}) 
	}else{
		c.JSON(404, gin.H{"error": err})
	}
}
// Register ...
func Register(c *gin.Context) {	// Rejestracja
	var user models.User
	c.BindJSON(&user)
	err := models.Register(&user)
	if(err == "Zarejestrowano"){
		var student models.Student
		student.ID = user.ID
		student.Age = ""
		student.Name = ""
		student.Surname = ""
		create := models.AddStudent(&student)
		if create != nil {
			c.JSON(404, gin.H{"error:": "Nie można stowrzyć studenta"})
			//fmt.Println("Nie można stworzyć studenta")
		}else {
			verifykey , err := auth.CreateVerify(user.ID)
			if err != nil {
				c.JSON(404, gin.H{"error:": "Nie można stworzyć kodu werifykacyjnego"})
				fmt.Println("Nie można stworzyć verifykey")
			} else {
				c.JSON(200, gin.H{"Verify": verifykey }) //: ts
				var verif models.Verify
				verif.KEY = verifykey
				verif.ID = user.ID
				err := models.KeyDB(&verif)
				if err != nil {
					c.JSON(404, gin.H{"error": "Cos poszlo nie tak"})
					fmt.Println("Nie można stworzyć dodania do bazy")
				} 
				c.JSON(200, gin.H{"Komunikat": "Dodano do bazy" }) //: ts
				fmt.Println("Stworzono")
			}
		}
	}else{
		c.JSON(404, gin.H{"Komunikat": err})
		fmt.Println("Nie można zarejestrowac")
	}
}

// LookToken ...
func LookToken(c *gin.Context) (uint) {	// Strona rejestracji
	cookie, _ := c.Cookie("Auth")
	iduser, _ := auth.ExtractTokenMetadata(cookie)
	fmt.Println(iduser.UserID)
	return iduser.UserID
}

// LogoutNow ... 
func LogoutNow(c *gin.Context){

	c.SetCookie("Auth", "Logout", -1, "/", "", false, true)	
	c.JSON(200, gin.H{"Komunikat": "wylogowano" }) //: ts
	// c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	// c.Redirect(http.StatusSeeOther, "/login")
	// c.Redirect(301, "/auth/login")
	 c.Abort()
} 
// Verification ...
func Verification(c *gin.Context) {
	key := c.Params.ByName("key")
	var idtoverify int
	config.DB.Raw("SELECT ID FROM Verify WHERE key = ?", &key).Scan(&idtoverify) // Pobranie id uzytkownika do verify
	var user models.User
	var newuser models.User
	newuser.Verify = true
	err := models.EditUser(&user,&newuser,idtoverify)
	if err != nil {
		c.JSON(404, gin.H{"error": "Cos poszlo nie tak"})
	} else {
			c.JSON(200, gin.H{"Pomyslnie Zwerifykowano:": idtoverify})
		}
}