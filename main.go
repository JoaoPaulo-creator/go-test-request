package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ResponseBody struct {
	Cep          string `json:"cep"`
	Address_name string `json:"address_name"`
	Address_type string `json:"address_type"`
	District     string `json:"district"`
	City         string `json:"city"`
}

func main() {
	endpoint := "https://cep.awesomeapi.com.br/json/80010010"
	response, err := http.Get(endpoint)

	if err != nil {
		log.Fatal("Deu ruim na request")
	}

	defer response.Body.Close()

	fmt.Printf("Status do response %d \n", response.StatusCode)

	res := ResponseBody{}
	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		log.Fatal("Deu ruim no decode do response body")
	}

	fmt.Printf("{cep: %s, address_name: %s, address_type: %s, cep: %s, city: %s, district: %s}", res.Cep, res.Address_name, res.Address_type, res.Cep, res.City, res.District)
}
