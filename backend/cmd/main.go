package main

import (
	"log"
	"net/http"

	"certificate-ledger/api/handler"
	"certificate-ledger/blockchain"
	"certificate-ledger/repository"
	"certificate-ledger/service"
	"github.com/gorilla/mux"
)

func main() {
	bc := blockchain.NewBlockchain()

	certRepo := repository.NewCertificateRepository()
	userRepo := repository.NewUserRepository()

	certService := service.NewCertificateService(certRepo, bc)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)

	certHandler := handler.NewCertificateHandler(certService)
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService)

	r := mux.NewRouter()

	r.HandleFunc("/api/certificates", certHandler.CreateCertificate).Methods("POST")
	r.HandleFunc("/api/certificates", certHandler.GetAllCertificates).Methods("GET")
	r.HandleFunc("/api/certificates/{id}", certHandler.GetCertificate).Methods("GET")
	r.HandleFunc("/api/certificates/verify/{hash}", certHandler.VerifyCertificate).Methods("GET")

	r.HandleFunc("/api/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/api/users/{id}/certificates", userHandler.GetUserCertificates).Methods("GET")

	r.HandleFunc("/api/auth/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/api/auth/register", authHandler.Register).Methods("POST")

	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	handler := corsMiddleware(r)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
