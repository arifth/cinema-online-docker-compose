package jwtToken

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/k0kubun/pp"
)

var secretKey = os.Getenv("SECRET_KEY")

// inside this package ,we have 3 funcs: GenerateToken(),
//              						 VerifyToken(),
//										 DecodeToken()

func GenerateToken(claims *jwt.MapClaims) (string, error) {

	// generate token with sha256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	pp.Println(secretKey)

	// sign token with secret value so it add up the security

	webToken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return webToken, nil

}

func VerifyToken(tokenstring string) (*jwt.Token, error) {

	// put jwt value inside param, then parse with followign method
	//  jwt.parse() take 2 values, first string that will be parsed, 2nd is callback func for handle the parse
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {

	// verify validity of token
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err

	}

	claims, isOk := token.Claims.(jwt.MapClaims)

	if isOk == token.Valid {
		return claims, nil
	}

	// if token doesn't pass validity check above return error and invalid tokens to stdout

	return nil, fmt.Errorf("invalid token")

}
