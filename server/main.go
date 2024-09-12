package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/paulofabiano/client-server-api/server/api"
	"github.com/paulofabiano/client-server-api/server/database"
)

type QuotationResponse struct {
	Bid string `json:"bid"`
}

func main() {
	http.HandleFunc("/quotation", handleQuotationApi)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleQuotationApi(w http.ResponseWriter, r *http.Request) {
	db, err := database.InitDatabase()
	if err != nil {
		log.Printf("Error initializing database: %v", err)
		http.Error(w, "Error initializing database", http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Second)
	defer cancel()

	quotation, err := api.GetQuotation(ctx)
	if err != nil {
		log.Printf("Error fetching quotation: %v", err)
		http.Error(w, "Error fetching quotation", http.StatusInternalServerError)
		return
	}

	ctxDb, cancelDb := context.WithTimeout(ctx, 10*time.Second)
	defer cancelDb()

	err = database.SaveQuotation(ctxDb, db, quotation)
	if err != nil {
		log.Printf("Error saving quotation in database: %v", err)
		http.Error(w, "Error saving quotation", http.StatusInternalServerError)
		return
	}

	response := QuotationResponse{Bid: fmt.Sprintf("%v", quotation.USD.Bid)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
