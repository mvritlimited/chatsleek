package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../config"
	"../model"
	"github.com/dgrijalva/jwt-go"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// RegistrationHandler User singup..
func RegistrationHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user model.User
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
	// fmt.Println([]byte(`{user:"kavin"}`)
	err := json.Unmarshal(body, &user)
	var res model.ResponseResult
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	collection, err := config.GetDBCollection()

	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	var result model.User
	err = collection.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&result)
	if user.Password != "" && user.Username != "" && user.Email != "" {

		if err != nil {
			if err.Error() == "mongo: no documents in result" && user.Password != "" {

				hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

				if err != nil {
					res.Error = "Error While Hashing Password, Try Again"
					json.NewEncoder(w).Encode(res)
					return
				}
				user.Password = string(hash)

				_, err = collection.InsertOne(context.TODO(), user)
				if err != nil {
					res.Error = "Error While Creating User, Try Again"
					json.NewEncoder(w).Encode(res)
					return
				}
				res.Result = "Registration Successful"
				json.NewEncoder(w).Encode(res)
				return
			}

			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}
	} else {
		res.Result = "Please Fill All MAndatory Details"
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Result = "Username already Exists!!"
	json.NewEncoder(w).Encode(res)
	return
}

// LoginHandler ..
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var user model.User

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	collection, err := config.GetDBCollection()

	if err != nil {
		log.Fatal(err)
	}

	var result model.User
	var res model.ResponseResult

	err = collection.FindOne(context.TODO(), bson.D{{"email", user.Email}}).Decode(&result)

	if err != nil {
		res.Error = "Invalid email"
		json.NewEncoder(w).Encode(res)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))

	if err != nil {
		res.Error = "Invalid password"
		json.NewEncoder(w).Encode(res)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": result.Username,
		"email":    result.Email,
	})

	tokenString, err := token.SignedString([]byte("wqdsfsdchbaskfkjsdbvkjsdbdaskhbvsdhk"))

	if err != nil {
		res.Error = "Error while generating token,Try again"
		json.NewEncoder(w).Encode(res)
		return
	}

	result.Token = tokenString
	result.Password = ""

	json.NewEncoder(w).Encode(result)

}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
