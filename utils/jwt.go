package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "osahenrumwen"

func GenerateToken(email string , userId int64) (string , error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":email,
		"userId":userId,
		"exp":time.Now().Add(time.Hour ).Unix(),
	})

	return token.SignedString([]byte(SECRET_KEY))
}

func VerifyToken(token string) (int64, error){
	
parsedToken , err := jwt.Parse(token,func (token *jwt.Token)  (interface{},error){
		 _,ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok{
            // return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"]) 
		return nil, errors.New( "Unexpected signing method")
        }
		return []byte(SECRET_KEY) , nil
	})

	if err!= nil {
		return 0,errors.New("Could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("Invalid Token!")
	}

	claims , ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok ||!tokenIsValid{

        return 0, errors.New("Invalid Token Claims!")

    }
	userId := int64(claims["userId"].(float64))


return userId, nil
}