package authetication

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"src/m/v2/src/utils"

	"github.com/gorilla/mux"
)

// RegisterAuthRoutes ..
func RegisterAuthRoutes(router *mux.Router) {

	router.HandleFunc("/login", handlelogin).Methods("POST")
}
func handlelogin(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	type REQUEST struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var result REQUEST
	json.Unmarshal(body, &result)
	fmt.Println("Repsults in Body are", result)
	type REPONSE struct {
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
		ERROR   string      `json:"error"`
	}
	var RESP REPONSE
	RESP.Data = result.Email
	RESP.Message = "Login Successfully"
	OUTPUT := utils.MARSHAL(RESP)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(OUTPUT))
}
