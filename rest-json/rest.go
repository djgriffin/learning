package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Payload struct {
	Payload Data

}

type Data struct {
	Fruit Fruits
	Veggies Vegetables
}

type Fruits map[string]int
type Vegetables map[string]int

func serveRest(w http.ResponseWriter, r *http.Request){
	res, err := getJsonResponse()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(res))
}

func getJsonResponse() ([]byte, error){
	fruits := make(map[string]int)
	fruits["Apples"] = 12
	fruits["Oranges"] = 21

	vegetables := make(map[string]int)
	vegetables["Peppers"] = 34
	vegetables["Carrots"] = 2

	d := Data{fruits, vegetables}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}

func main() {
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("localhost:1337", nil)
}