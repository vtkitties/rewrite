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
		http.Error(w, `{"error":"not logged in"}`, http.StatusUnauthorized)
		return
	}
	admin_id := claims["user_id"]
	var admin orm.User
	db.First(&admin, admin_id)

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"you":   admin,
	}) // if theres an error here u probsters cant encode the event struct

	// req := struct { // all fields are required
	// 	Name    string `json:"name"`
	// 	Surname string `json:"surname"`
	// 	Email   string `json:"email"`
	// }{}

	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	http.Error(w, `{"error":"invalid request body"}`, http.StatusBadRequest)
	// 	// log.Println(err)
	// 	return
	// }

	// user := orm.User{ // WHEN_EMAIL: send a notification to the user
	// 	Email:    req.Email,
	// 	Password: "12341234", // WHEN_EMAIL: send a random password to the user

	// 	Name:    req.Name,
	// 	Surname: req.Surname,
	// 	Role:    orm.RoleChair,
	// }

	// if err := db.Save(&user).Error; err != nil {
	// 	http.Error(w, `{"error":"couldn't save for some reason idk"}`, http.StatusBadRequest)
	// 	return
	// }
	// w.WriteHeader(http.StatusCreated)
	// _ = json.NewEncoder(w).Encode(map[string]any{
	// 	"message": "user created",
	// 	"email":   user.Email,
	// }) // if theres an error here u probsters cant encode the event struct

}
