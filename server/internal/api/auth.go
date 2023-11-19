package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ca-lee-b/golander/server/internal/repositories"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userRepository *repositories.UserRepository
}

func newAuthHandler(repo *repositories.UserRepository) *AuthHandler {
	return &AuthHandler{
		userRepository: repo,
	}
}

// JWT middleware
func (h *AuthHandler) WithToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized"))
			return
		}

		verify := verifyToken(cookie.Value)
		if !verify {
			w.WriteHeader(401)
			w.Write([]byte("Unauthorized"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Endpoints
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	type loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var data loginData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}

	if data.Email == "" {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}
	if data.Password == "" {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}

	user := h.userRepository.GetOneByEmail(data.Email)
	if user == nil {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Unauthorized"))
		return
	}

	token, err := generateToken(user.Id)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
		return
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		MaxAge:   3600, // 1 hour
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
	w.WriteHeader(200)
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	type registerData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var data registerData

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}

	if data.Email == "" {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}
	if data.Password == "" {
		w.WriteHeader(400)
		w.Write([]byte("Bad Request"))
		return
	}
	if len(data.Password) < 7 {
		w.WriteHeader(400)
		w.Write([]byte("Password must be longer than 7 characters"))
		return
	}

	user := h.userRepository.GetOneByEmail(data.Email)
	if user != nil {
		w.WriteHeader(404)
		w.Write([]byte("Email already exists"))
		return
	}

	err = h.userRepository.CreateUser(strings.ToLower(data.Email), data.Password)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Success"))
}

func generateToken(id string) (string, error) {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		panic("Env variable: JWT_SECRET does not exist")
	}

	exp := jwt.NewNumericDate(time.Now().Add(time.Hour))
	iat := jwt.NewNumericDate(time.Now())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": exp,
		"iat": iat,
	})
	signed, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return signed, nil
}

func verifyToken(tokenString string) bool {
	key := os.Getenv("JWT_SECRET")

	type CustomClaims struct {
		Exp int `json:"exp"`
		Iat int `json:"iat"`
		jwt.RegisteredClaims
	}
	claims := &CustomClaims{}
	t, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return false
	}
	if !t.Valid {
		fmt.Printf("not valid\n")
		return false
	}

	return true
}
