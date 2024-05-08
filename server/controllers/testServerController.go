package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Product struct {
	productName  string
	price        uint32
	rating       float32
	discount     uint16
	availability string
}

func GetProducts(company string, category string) ([]Product, error) {
	var BaseURL string = os.Getenv("BASE_URL")
	var AccessToken string = os.Getenv("ACCESS_TOKEN")

	if company != "" {
		BaseURL += "/companies/" + company
	}

	if category != "" {
		BaseURL += "/categories/" + category
	}

	BaseURL += "/products?top=10&minPrice=1&maxPrice=1000"

	log.Println("BaseURL: ", BaseURL)
	log.Println("AccessToken: ", AccessToken)

	client := &http.Client{}
	req, err := http.NewRequest("GET", BaseURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response Status: ", resp.Status)

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response Body: ", string(body))

	var products []Product
	err = json.Unmarshal(body, &products)
	if err != nil {
		log.Fatal(err)
	}

	return products, nil
}
