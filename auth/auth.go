package auth

import (
	"errors"
	"fmt"
	"Golang-Graphql-withCICD-PSQL/config"
	"Golang-Graphql-withCICD-PSQL/model/Person"
	
	"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
)
type UserDetails struct {
	EmailID string
	FirstName string
	LastName string
	Password string
	Username string
}
func ValidateToken(tokenString string) (personModel.Person, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return config.GetJWTSecret(), nil
	})

	if err != nil {
		return personModel.Person{}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		var c1 UserDetails
		fmt.Println("CLAIMEDTOKEN", claims["data"])
		userData, err := json.Marshal(claims["data"])
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println("json:", string(userData), "ITS TYPE IS %T", userData)
		err = json.Unmarshal(userData, &c1)
		if err != nil {
  
			// if error is not nil
			// print error
			fmt.Println(err)
		}
		fmt.Println("HELLOW LOO HERE, ", c1.Username)
		person, _ := claims["data"].(personModel.Person)
		
		return person, nil
	}
	return personModel.Person{}, errors.New("Token not valid")

}

func ValidateToken1(tokenString string) (UserDetails, error) {
	var c1 UserDetails
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return config.GetJWTSecret(), nil
	})

	if err != nil {
		return c1, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		
		fmt.Println("CLAIMEDTOKEN", claims["data"])
		userData, err := json.Marshal(claims["data"])
		if err!=nil{
			fmt.Println(err)
		}
		// fmt.Println("json:", string(userData), "ITS TYPE IS %T", userData)
		err = json.Unmarshal(userData, &c1)
		if err != nil {
  
			// if error is not nil
			// print error
			fmt.Println(err)
		}
		fmt.Println("HELLOW LOO HERE, ", c1.Username)
		
		return c1, nil
	}
	return c1, errors.New("Token not valid")

}