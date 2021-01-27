package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateVerify ...
func CreateVerify(userid uint) (string, error) {
	
  
	var err error
	//Creating Access Token
	os.Setenv("VERIFY_SECRET", "adsqdewdfrf") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 1440).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("VERIFY_SECRET")))
	if err != nil {
	   return "", err
	}
	
	return token, nil
	
	
  }