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
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return config.GetJWTSecret(), nil
	})

	if err != nil {
		return personModel.Person{}, errors.New("Error getting User Data")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		var c1 UserDetails
		fmt.Println("CLAIMED TOKEN", claims["data"])
		userData, err := json.Marshal(claims["data"])
		if err!=nil{
			fmt.Println(err)
			return personModel.Person{}, errors.New("Error getting User Data")
		}
		err = json.Unmarshal(userData, &c1)
		if err != nil {
  
			// if error is not nil
			// print error
			fmt.Println(err)
			return personModel.Person{}, errors.New("Error getting User Data")

		}
		fmt.Println("HELLOW LOO HERE, ", c1.Username)
		person, _ := claims["data"].(personModel.Person)
		
		return person, nil
	}
	return personModel.Person{}, errors.New("Token not valid")

}

//to return with the id. Can be merged with the above function to make it as one
func ValidateToken1(tokenString string) (UserDetails, error) {
	var c1 UserDetails
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
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
			return c1, errors.New("Error Getting the User Data")
		}
		err = json.Unmarshal(userData, &c1)
		if err != nil {
			// if error is not nil
			// print error
			fmt.Println(err)
			return c1, errors.New("Error Getting the User Data")
		}
		
		return c1, nil
	}
	return c1, errors.New("Token not valid")

}