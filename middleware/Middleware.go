package middleware

// middleware.proces.authMiddleware
import (
	"fmt"
	"new/auth"

	"new/controllers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// TokenAuthMiddleware ...
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	   //cookie, err := c.Cookie("token")
	   auu, err := c.Request.Cookie("Auth")
	   if err != nil {
		//
		c.HTML(200, "access.tmpl", gin.H{
			
		})
		c.JSON(303, "Nie masz dostępu do tej strony!")
	}
		cookie := auu.Value
	   	error := auth.TokenValid(cookie)
	   	if error != nil {
		  	c.HTML(200, "access.tmpl", gin.H{
			})
			c.JSON(303, error.Error())
		  	c.Abort()
		  	return
	   	}
	   token, err := jwt.Parse(cookie, nil)
	   if token == nil {
		   return 
	   }
	   claims, _ := token.Claims.(jwt.MapClaims)
	   fmt.Printf("USERID= %s \n", claims) 
	   c.Next()
	}
  }
  

// AuthTeacher ...
func AuthTeacher() gin.HandlerFunc {
	return func(c *gin.Context) {
	   cookie, err := c.Cookie("Auth")
	   if err != nil {
		c.JSON(303, "Nie masz cookies jwt")
	}
	   error := auth.TokenValid(cookie)
	   if error != nil {
		  c.JSON(303, error.Error())
		  c.Abort()
		  return
	   }
	   role := controllers.LookRole(c)
		
	   if(role=="teacher"||role=="dziekan"){
		   c.Next()
	   }else{
		c.HTML(200, "access.tmpl", gin.H{
		})
		c.JSON(303, error.Error())
		  c.Abort()
		  return
	   }
	   
	}
  }
// AuthDziekan ...
func AuthDziekan() gin.HandlerFunc {
	return func(c *gin.Context) {
	   cookie, err := c.Cookie("Auth")
	   if err != nil {
		c.JSON(303, "Nie masz cookies jwt")
	}
	   	error := auth.TokenValid(cookie)
	   	if error != nil {
		  	c.JSON(303, error.Error())
		  	c.Abort()
		  	return
	   	}
	   	role := controllers.LookRole(c)
		
	   	if(role=="dziekan"){
		   c.Next()
	   	}else{
			c.HTML(200, "access.tmpl", gin.H{
			})
			c.JSON(303, error.Error())
		 	c.Abort()
		  	return
	   }
	   
	}
  }

// AuthMiddleware ...
// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		err := auth.TokenValid(c.Request)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, "You need to be authorized to access this route")
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}
// }
// 	"fmt"
// 	"net/http"

// 	//"github.com/dgrijalva/jwt-go"
// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"

// // AuthorizeJWT ..
// func AuthorizeJWT() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		const BEARER_SCHEMA = "Bearer"
// 		authHeader := c.GetHeader("Authorization")
// 		tokenString := authHeader[len(BEARER_SCHEMA):]
// 		token, err := service.JWTAuthService().ValidateToken(tokenString)
// 		if token.Valid {
// 			claims := token.Claims.(jwt.MapClaims)
// 			fmt.Println(claims)
// 		} else {
// 			fmt.Println(err)
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 		}

// 	}
// }