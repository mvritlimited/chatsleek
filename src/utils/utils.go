package utils

import (
	"encoding/json"
	"fmt"
)

// MARSHAL ..
func MARSHAL(input interface{}) []byte {
	output, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Error Occured")
	}
	return output

}
