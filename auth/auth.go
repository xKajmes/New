package auth

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
)

// Client ...

// TokenDetails ...
type TokenDetails struct {
	AccessToken  string
	AccessUUID   string
	AtExpires    int64
  }
// AccessDetails ...
  type AccessDetails struct {
    AccessUUID string
    UserID   uint
}
// CreateToken ...
func CreateToken(userid uint) (*TokenDetails, error) {
	
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 150).Unix()
	td.AccessUUID = uuid.NewV4().String()
	ad := &AccessDetails{}
	ad.AccessUUID = td.AccessUUID
	ad.UserID = userid

  
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
	   return nil, err
	}
	
	
	
	return td, nil
  }
  // CreateAuth ...


// func ExtractToken(r string) string {

// 	bearToken := r
// 	//normally Authorization the_token_xxx
// 	strArr := strings.Split(bearToken, " ")
// 	if len(strArr) == 2 {
// 	   return strArr[1]
// 	}
// 	return ""
// }

// VerifyToken ...
func VerifyToken(r string) (*jwt.Token, error) {
	tokenString := r
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   //Make sure that the token method conform to "SigningMethodHMAC"
	   if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		  return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	   }
	   return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
	   return nil, err
	}
	return token, nil
}
// TokenValid ...
func TokenValid(r string) error {
	token, err := VerifyToken(r)
	if err != nil {
	   return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
	   return err
	}
	return nil
  }
  // ExtractTokenMetadata ...
  func ExtractTokenMetadata(r string) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
	   return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
	   accessuuid, ok := claims["access_uuid"].(string)
	   if !ok {
		  return nil, err
	   }
	   userid1, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 0)
	   userid := uint(userid1)
	   if err != nil {
		  return nil, err
	   }
	   return &AccessDetails{
		AccessUUID: accessuuid,
		UserID:   userid,
	   }, nil
	}
	return nil, err
  }

// DeleteAuth ...
//   func DeleteAuth(givenUuid string) (int64,error) {
// 	deleted, err := c.Cookie("token") //Client.Del(givenUuid).Result()
// 	if err != nil {
// 	   return 0, err
// 	}
// 	return deleted, nil
//   }

  // Logout ...
  func Logout(c *gin.Context) {

	// cookie, err := c.Cookie("token")
	// if err != nil {
	//  c.JSON(303, "Nie masz cookies jwt")
 	// }
	// au, err := ExtractTokenMetadata(cookie)
	// if err != nil {
	//    c.JSON(http.StatusUnauthorized, "unauthorized")
	//    return
	// }
	// deleted, delErr := DeleteAuth(au.AccessUUID)
	// if delErr != nil || deleted == 0 { //if any goes wrong
	//    c.JSON(http.StatusUnauthorized, "unauthorized")
	//    return
	// }
	c.JSON(http.StatusOK, "Successfully logged out")
  }