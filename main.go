package main

import (
	"fmt"
	"kitties/db"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"gorm.io/gorm"
)

func check_die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	gormDB, err := db.Init(":memory:")

	addr := "localhost:3333"
	log.Printf("serving on %v\n", addr)
	check_die(err)

	// routing setup
	initJWT()
	r := chi.NewRouter()
	r.Use(db.MiddlewareWithDB(gormDB))
	r.Use(jwtauth.Verifier(tokenAuth))
	r.Use(jwtauth.Authenticator(tokenAuth))
	r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		userID := claims["user_id"]
		fmt.Fprintf(w, "Hello protected user %v", userID)
	})
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		db := r.Context().Value(db.DBContextKey).(*gorm.DB)

		var count int64
		db.Raw("SELECT count(*) FROM sqlite_master").Scan(&count)

		fmt.Fprintf(w, "db has %d tables", count)
	})

	http.ListenAndServe(addr, r)
}

var tokenAuth *jwtauth.JWTAuth

// jwt
func initJWT() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("the JWT_SECRET env var is not set")
	}
	tokenAuth = jwtauth.New("HS256", []byte(secret), nil)

	_, tokenString, _ := tokenAuth.Encode(map[string]any{"user": 123})
	fmt.Printf("DEBUG: a sample JWT is %s\n", tokenString)
}
