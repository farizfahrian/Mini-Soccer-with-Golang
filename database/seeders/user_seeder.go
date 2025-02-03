package seeders

import (
	"user-service/constants"
	"user-service/domain/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RunUserSeeder(db *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	user := models.User{
		UUID:        uuid.New(),
		Name:        "Admin User",
		Username:    "Administrator",
		Password:    string(password),
		PhoneNumber: "+1234567890",
		Email:       "admin@example.com",
		RoleID:      constants.Admin,
	}

	err := db.FirstOrCreate(&user, models.User{Username: user.Username}).Error
	if err != nil {
		logrus.Errorf("failed to seed user: %v", err)
		panic(err)
	}
	logrus.Infof("user %s successfully seeded", user.Username)
}
