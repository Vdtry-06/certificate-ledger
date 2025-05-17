package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"certificate-ledger/domain"
	"certificate-ledger/repository"

	"github.com/golang-jwt/jwt/v5"
)

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value("user").(*domain.User)
		if !ok || user == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if user.Role != "admin" {
			http.Error(w, "Forbidden: Admin access required", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(userRepo *repository.UserRepository) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            authHeader := r.Header.Get("Authorization")
            if authHeader == "" {
                log.Println("Missing token")
                http.Error(w, "Missing token", http.StatusUnauthorized)
                return
            }

            tokenString := strings.TrimPrefix(authHeader, "Bearer ")
            log.Printf("Received token: %s", tokenString)
            token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                    log.Printf("Unexpected signing method: %v", token.Header["alg"])
                    return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
                }
                secret := os.Getenv("JWT_SECRET")
                log.Printf("Using secret: %s", secret)
                if secret == "" {
                    log.Println("JWT_SECRET is empty")
                }
                return []byte(secret), nil
            })
            if err != nil {
                log.Printf("Token validation error: %v", err)
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }

            if !token.Valid {
                log.Println("Token is not valid")
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }

            claims, ok := token.Claims.(jwt.MapClaims)
            if !ok {
                log.Println("Invalid token claims")
                http.Error(w, "Invalid token claims", http.StatusUnauthorized)
                return
            }

            userID, ok := claims["user_id"].(string)
            if !ok {
                log.Printf("Invalid user_id in claims: %v", claims)
                http.Error(w, "Invalid user_id in token", http.StatusUnauthorized)
                return
            }

            user, err := userRepo.FindByID(userID)
            if err != nil {
                log.Printf("User not found: %v", err)
                http.Error(w, "Invalid token", http.StatusUnauthorized)
                return
            }

            log.Printf("User found: %v", user)
            ctx := context.WithValue(r.Context(), "user", user)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}