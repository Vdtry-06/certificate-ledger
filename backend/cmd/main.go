package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"certificate-ledger/api/handler"
	"certificate-ledger/blockchain"
	"certificate-ledger/db"
	"certificate-ledger/domain"
	"certificate-ledger/repository"
	"certificate-ledger/service"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Load .env
    if err := godotenv.Load(); err != nil {
    log.Printf("Error loading .env file: %v", err)
	} else {
		log.Printf("JWT_SECRET from env: %s", os.Getenv("JWT_SECRET"))
	}

	// Khởi tạo kết nối MySQL
	db, err := db.NewDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Khởi tạo blockchain
	bc := blockchain.NewBlockchain()

	// Khởi tạo repository
	certRepo := repository.NewCertificateRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Khởi tạo service
	certService := service.NewCertificateService(certRepo, bc)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)

	// Tạo tài khoản admin nếu chưa tồn tại
	if err := createAdminUser(userRepo); err != nil {
		log.Printf("Failed to create admin user: %v", err)
	}

	// Khởi tạo handler
	certHandler := handler.NewCertificateHandler(certService)
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService)

	// Thiết lập router
	r := mux.NewRouter()

	// API công khai
	r.HandleFunc("/api/auth/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/api/auth/register", authHandler.Register).Methods("POST")

	// API yêu cầu xác thực
	protectedRouter := r.PathPrefix("/api").Subrouter()
	protectedRouter.Use(handler.AuthMiddleware(userRepo))
	protectedRouter.HandleFunc("/certificates", certHandler.CreateCertificate).Methods("POST")
	protectedRouter.HandleFunc("/certificates", certHandler.GetAllCertificates).Methods("GET")
	protectedRouter.HandleFunc("/certificates/{id}", certHandler.GetCertificate).Methods("GET")
	protectedRouter.HandleFunc("/certificates/verify/{hash}", certHandler.VerifyCertificate).Methods("GET")
	protectedRouter.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	protectedRouter.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	protectedRouter.HandleFunc("/users/{id}/certificates", userHandler.GetUserCertificates).Methods("GET")

	// API admin
	adminRouter := r.PathPrefix("/api/admin").Subrouter()
	adminRouter.Use(handler.AuthMiddleware(userRepo))
	adminRouter.Use(handler.AdminMiddleware)
	adminRouter.HandleFunc("/users", userHandler.ListUsers).Methods("GET")
	adminRouter.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	// CORS middleware
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "http://localhost:3000"
	}
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	// Tạo server với timeout
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      corsMiddleware(r),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Graceful shutdown
	go func() {
		log.Println("Server starting on port 8080...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Chờ tín hiệu dừng
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	// Tắt server an toàn
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
	log.Println("Server stopped gracefully")
}

func createAdminUser(repo *repository.UserRepository) error {
	_, err := repo.FindByEmail("tunhoipro0306@gmail.com")
	if err == nil {
		log.Println("Admin user already exists")
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash admin password: %v", err)
	}

	admin := &domain.User{
		ID:        uuid.New().String(),
		Name:      "Admin",
		Email:     "tunhoipro0306@gmail.com",
		Password:  string(hashedPassword),
		Role:      "admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := repo.Save(admin); err != nil {
		return fmt.Errorf("failed to save admin user: %v", err)
	}

	log.Println("Admin user created successfully")
	return nil
}