package orm

import "gorm.io/gorm"

// a superuser is the ultramegaadmin who has all permissions. there can only be one
func InitSuperuser(password string, db *gorm.DB) error {
	user := User{
		Email:    "admin",
		Password: password,
		Role:     RoleAdmin,
		Name:     "Robert",
		Surname:  "Tables",
	}
	return user.NewUser(db)
}
