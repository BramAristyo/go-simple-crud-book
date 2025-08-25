package helper

import (
	"encoding/json"
	"net/http"
	"simple-restful-api/response"
)

func ReadRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfErr(err)
}

func WriteToResponse(w http.ResponseWriter, response response.WebResponse) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfErr(err)
}