package utils

import (

	"time"

	"errors"

	"github.com/golang-jwt/jwt"
)

const secretKey = "supersecret"



func GenerateToken(email string, userid int64)(string,error) {

	// fmt.Print("jwt",userid)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":email,
		"userid":userid,
		"exp":time.Now().Add(time.Hour*2).Unix(),
	
	})
	
	return token.SignedString([]byte(secretKey))

}

func VerifyToken(token string)(int64,error){

	Parsedtoken,err := jwt.Parse(token,func(token *jwt.Token)(interface{},error){
		_,ok := token.Method.(*jwt.SigningMethodHMAC)

		if(!ok){
			return nil,errors.New("unexpected sigining Method")
		}
		return []byte(secretKey),nil
	})

	if(err != nil ){
		return 0,errors.New("could not Parse Token")
	}
	tokenIsvalid := Parsedtoken.Valid

	if(!tokenIsvalid){
		return 0,errors.New("invalid token")
	}

	claims,ok := Parsedtoken.Claims.(jwt.MapClaims)

	if(!ok){
		return 0,errors.New("invalid token")
	}
	userid := int64(claims["userid"].(float64))



	return userid,nil


}