package orm

import (
	"context"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ctxKeyDB struct{}

var DBContextKey = &ctxKeyDB{}

func Init(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = AutoMigrate(db); err != nil {
		return nil, err
	}
	if err = db.Exec(`
		CREATE UNIQUE INDEX IF NOT EXISTS idx_vote_user_unique 
		ON vote_results (vote_id, user_id) 
		WHERE user_id IS NOT NULL;
	`).Error; err != nil {
		return nil, err
	}

	return db, nil
}

func MiddlewareWithDB(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), DBContextKey, db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Vote{},
		&VoteOption{},
		&VoteResult{},
		&Event{},
	)
}
