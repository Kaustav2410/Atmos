package main

import (
	"errors"
	"log"
	"net/http"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"os"
)

var logger *log.Logger
func health(w http.ResponseWriter, r *http.Request){
     logger.Println("Health method was invoked");
	 response := map[string]string{"message":"Server is healthy"}
	 w.Header().Set("Content-Type", "application/json")
	 w.WriteHeader(http.StatusOK)
	 json.NewEncoder(w).Encode(response)
}

func main(){
	logger = log.Default()
	mux:= http.NewServeMux()
	_ = godotenv.Load()
	// if envErr != nil {
	// 	logger.Fatalln("Error loading .env file")
	// }
	PORT := os.Getenv("PORT")
	FRONTEND_URL := os.Getenv("FRONTEND_URL")
	if(PORT == "" || FRONTEND_URL == ""){
		logger.Fatalln("PORT or FRONTEND_URL environment variable is not set")
	}
	c := cors.New(cors.Options{
        AllowedOrigins:   []string{FRONTEND_URL}, // Your React URL
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    })
	mux.HandleFunc("/health",health);
	err:= http.ListenAndServe(PORT,c.Handler(mux));
	if(errors.Is(err,http.ErrServerClosed)){
		logger.Println("Server closed");
	}else if(err!=nil){
		logger.Fatalln("Error occured while starting the server",err)
	}
}