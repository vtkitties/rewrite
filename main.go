package main

import (
	"fmt"
	"kitties/handlers"
	"kitties/orm"
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
	gormDB, err := orm.Init(":memory:")

	addr := "localhost:3333"
	log.Printf("serving on %v\n", addr)
	check_die(err)
	initJWT()

	// routing setup
	r := chi.NewRouter()
	r.Use(orm.MiddlewareWithDB(gormDB))
	// public auth
	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/login", handlers.Login(tokenAuth))
		r.Post("/register", handlers.Register(tokenAuth))
		r.Post("/refresh", handlers.Refresh(tokenAuth))
	})
	// protected routes
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			fmt.Fprintf(w, "Hello user %v", claims["user_id"])
		})
	})

	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		db := r.Context().Value(orm.DBContextKey).(*gorm.DB)

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
		// panic("the JWT_SECRET env var is not set")
		secret = "asdf" // TEMP
	}
	tokenAuth = jwtauth.New("HS256", []byte(secret), nil)

	// _, tokenString, _ := tokenAuth.Encode(map[string]any{"user": 123})
	// fmt.Printf("DEBUG: a sample JWT is %s\n", tokenString)
}
