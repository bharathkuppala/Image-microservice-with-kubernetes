package utility

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	model "github/image-service/microservice/models"
)

// Response ...
type Response struct {
	Ok   bool
	Data interface{}
}

// SendResponse ...
func SendResponse(w http.ResponseWriter, ok bool, data interface{}) {
	response := Response{
		Ok:   ok,
		Data: data,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
	}
}

// ToMap ...
func ToMap(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	var resp map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		log.Println("error in decoding request body", err)
		return nil, err
	}
	return resp, nil
}

// CheckKeyExist ...
func CheckKeyExist(m1, m2 map[string][]model.Image, k1, k2 string) (
	v1, v2 []model.Image, ok1, ok2 bool) {

	v1, ok1 = m1[k1]
	v2, ok2 = m2[k2]
	return
}

// ToUnmarshall ...
func ToUnmarshall(fileName string) (map[string][]model.Image, error) {
	file, err := ioutil.ReadFile("image-album-db.json")
	if err != nil {
		fmt.Println("error in reading json file", err)
		return nil, err
	}

	data := make(map[string][]model.Image, 0)
	if err = json.Unmarshal(file, &data); err != nil {
		log.Println("error with unmarshalling", err)
		return nil, err
	}
	return data, nil
}

// ToMarshalIndent ...
func ToMarshalIndent(data map[string][]model.Image) ([]byte, error) {
	fileData, err := json.MarshalIndent(data, " ", "\t")
	if err != nil {
		log.Println("marshal error", err)
		return nil, err
	}
	return fileData, nil
}
