package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// Middleware ..
func Middleware(w http.ResponseWriter, r *http.Request) {
	AUTHTOKEN := r.Header.Get("Authorization")
	fmt.Println("TOKEN IS", AUTHTOKEN)
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(AUTHTOKEN, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("wqdsfsdchbaskfkjsdbvkjsdbdaskhbvsdhk"), nil
	})
	// ... error handling
	if err != nil {
		fmt.Println("error occured in token decode", err)
	}
	fmt.Println("toke decoded", token)
	// do something with decoded claims
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}

}
