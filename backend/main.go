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
	orm.InitSuperuser("pebis", gormDB)

	// routing setup
	r := chi.NewRouter()
	r.Use(orm.MiddlewareWithDB(gormDB))

	// public auth
	r.Post("/api/auth/login", handlers.Login(tokenAuth))

	// protected routes
	r.Route("/api", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))

		r.Route("/events", func(r chi.Router) {
			r.Post("/", handlers.NewEvent)
		})

		r.Route("/auth", func(r chi.Router) {
			r.Post("/refresh", handlers.Refresh(tokenAuth))
			// r.Post("/register", handlers.Register(tokenAuth))
		})

		r.Route("/admin", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				_, claims, _ := jwtauth.FromContext(r.Context())
				fmt.Fprintf(w, "hello, %v\n%v", claims["user_id"], claims)
			})
			r.Post("/new_user", handlers.NewUser)
			// r.Post("/register", handlers.Register(tokenAuth))
		})
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
