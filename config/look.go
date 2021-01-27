package config

import (
	"fmt"
	"new/auth"

	"github.com/gin-gonic/gin"
)

// LookToken ...
func LookToken(c *gin.Context) (uint) {	// Strona rejestracji
	
	cookie, _ := c.Cookie("token")
	iduser, _ := auth.ExtractTokenMetadata(cookie)
	fmt.Println(iduser.UserID)
	return iduser.UserID
}
// LookRole ...
func LookRole(c *gin.Context) (string) {
	iduser := LookToken(c)
	var role string
	DB.Raw("SELECT role FROM User WHERE id = ?", &iduser).Scan(&role) // Pobranie passwordhash z database
	//fmt.Println(role)
	return role
} 