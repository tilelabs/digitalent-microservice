package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func AddMenu(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(map[string]interface{}{
		"success": true,
	})

	if err != nil {
		fmt.Print("Failed to generate response")
		return
	}

	w.Write(response)

}
