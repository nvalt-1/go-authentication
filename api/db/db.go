package db

import (
	"authentication-test/api/models"
	"golang.org/x/crypto/bcrypt"
)

func LoadTestUser() *models.User {
	// Just for demonstration purpose, we create a user with the encrypted "test" password.
	// In real-world applications, you might load the user from the database by specific parameters (email, username, etc.)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("test"), 8)
	return &models.User{Password: string(hashedPassword), Username: "Test user"}
}
