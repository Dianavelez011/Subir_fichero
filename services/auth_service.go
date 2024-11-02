package services

import (
    "errors"
    "github.com/gin/models"  // Reemplaza con tu nombre de módulo
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

type AuthService struct {
    db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
    return &AuthService{db: db}
}

func (s *AuthService) CreateUser(name, email, password string) error {
    var existingUser models.User
    if result := s.db.Where("email = ?", email).First(&existingUser); result.Error == nil {
        return errors.New("email already exists")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user := models.User{
        Name:     name,
        Email:    email,
        Password: string(hashedPassword),
    }

    if result := s.db.Create(&user); result.Error != nil {
        return result.Error
    }

    return nil
}

func (s *AuthService) ValidateUser(email, password string) error {
    var user models.User
    if result := s.db.Where("email = ?", email).First(&user); result.Error != nil {
        return errors.New("invalid credentials")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return errors.New("invalid credentials")
    }

    return nil
}