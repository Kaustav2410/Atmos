package main

import (
	"errors"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"os"
	"github.com/Kaustav2410/backend/internal"
)

var logger *log.Logger

func main(){
	logger = log.Default()
	var mux= http.NewServeMux()
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
	routes.V1Routes(mux)
	err:= http.ListenAndServe(PORT,c.Handler(mux));
	if(errors.Is(err,http.ErrServerClosed)){
		logger.Println("Server closed");
	}else if(err!=nil){
		logger.Fatalln("Error occured while starting the server",err)
	}
}