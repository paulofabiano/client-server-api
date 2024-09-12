package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type ServerQuotationResponse struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/quotation", nil)
	if err != nil {
		log.Fatalf("Error creating server api request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Fatalf("Timeout on server api request: %v", err)
			return
		}

		log.Fatalf("Error fetching server api: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var serverQuotationResponse ServerQuotationResponse
	err = json.Unmarshal(body, &serverQuotationResponse)
	if err != nil {
		log.Fatalf("Error unmarshalling response body: %v", err)
	}

	err = SaveQuotationInFile(serverQuotationResponse.Bid)
	if err != nil {
		log.Fatalf("Error saving quotation in file: %v", err)
	}

	fmt.Printf("USD quotation is %s and is saved successfully in quotations file\n", serverQuotationResponse.Bid)
}

func SaveQuotationInFile(quotation string) error {
	err := os.WriteFile("quotations.txt", []byte("Dólar: "+quotation), 0644)
	if err != nil {
		return err
	}

	return nil
}
