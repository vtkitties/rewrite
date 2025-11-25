package handlers

import (
	"encoding/json"
	"kitties/orm"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

// POST /api/auth/login
func Login(tokenAuth *jwtauth.JWTAuth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := r.Context().Value(orm.DBContextKey).(*gorm.DB)

		var req AuthRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
			return
		}

		var user orm.User
		if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
			http.Error(w, `{"error":"invalid credentials"}`, http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			http.Error(w, `{"error":"invalid credentials"}`, http.StatusUnauthorized)
			return
		}

		_, accessToken, _ := tokenAuth.Encode(map[string]any{
			"user_id": user.ID,
			"exp":     time.Now().Add(7 * (24 * time.Hour)).Unix(),
		})

		resp := AuthResponse{
			AccessToken: accessToken,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}


// POST /api/auth/refresh
func Refresh(tokenAuth *jwtauth.JWTAuth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil || claims == nil {
			http.Error(w, `{"error":"invalid token"}`, http.StatusUnauthorized)
			return
		}

		userID := claims["user_id"]

		_, newAccessToken, _ := tokenAuth.Encode(map[string]any{
			"user_id": userID,
			"exp":     time.Now().Add(7 * (24 * time.Hour)).Unix(),
		})

		resp := AuthResponse{
			AccessToken: newAccessToken,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
