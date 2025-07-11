package auth

import (
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"log/slog"
	"net/http"
	"time"
)

var jwtKey = []byte("your_secret_key") // Секретный ключ для подписи JWT

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func Auth(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Проверка логина и пароля (замените на реальную логику)
		if user.Username != "admin" || user.Password != "password" {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Создание JWT
		expirationTime := time.Now().Add(24 * time.Hour) // Токен действителен 24 часа
		claims := &Claims{
			Username: user.Username,
			Role:     "admin", // Пример роли
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			http.Error(w, "Could not generate token", http.StatusInternalServerError)
			return
		}

		// Возвращаем токен клиенту
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}

func AuthMiddleware(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				http.Error(w, "Missing token", http.StatusUnauthorized)
				return
			}

			token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(*Claims)
			if !ok || claims.ExpiresAt < time.Now().Unix() {
				http.Error(w, "Token expired", http.StatusUnauthorized)
				return
			}

			// Передаем роль в контекст
			ctx := context.WithValue(r.Context(), "role", claims.Role)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

//func loginHandler(w http.ResponseWriter, r *http.Request) {
//	var user User
//	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
//		http.Error(w, "Invalid request", http.StatusBadRequest)
//		return
//	}
//
//	// Проверка логина и пароля (замените на реальную логику)
//	if user.Username != "admin" || user.Password != "password" {
//		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
//		return
//	}
//
//	// Создание JWT
//	expirationTime := time.Now().Add(24 * time.Hour) // Токен действителен 24 часа
//	claims := &Claims{
//		Username: user.Username,
//		Role:     "admin", // Пример роли
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: expirationTime.Unix(),
//		},
//	}
//
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	tokenString, err := token.SignedString(jwtKey)
//	if err != nil {
//		http.Error(w, "Could not generate token", http.StatusInternalServerError)
//		return
//	}
//
//	// Возвращаем токен клиенту
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
//}
