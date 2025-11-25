package handlers

import (
	"encoding/json"
	"kitties/orm"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"gorm.io/gorm"
)

// POST /api/admin/new_user
func NewUser(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(orm.DBContextKey).(*gorm.DB)
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, ErrorResponse("not logged in"), http.StatusUnauthorized)
		return
	}
	admin_id := claims["user_id"]
	var admin orm.User
	db.First(&admin, admin_id)

	if admin.Role != orm.RoleAdmin {
		http.Error(w, ErrorResponse("you're not an admin"), http.StatusUnauthorized)
		return
	}

	req := struct {
		Name     string `json:"name"`
		Surname  string `json:"surname"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, ErrorResponse("invalid body"), http.StatusBadRequest)
		return
	}

	user := orm.User{ // WHEN_EMAIL: send a notification to the user
		Email:    req.Email,
		Password: req.Password, // WHEN_EMAIL: send a random password to the user

		Name:    req.Name,
		Surname: req.Surname,
		Role:    orm.RoleChair,
	}
	if err := user.Register(db); err != nil {
		http.Error(w, ErrorResponse("couldn't save for some reason idk"), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"message": "user created",
		"email":   user.Email,
	}) // if theres an error here u probsters cant encode the event struct

}
