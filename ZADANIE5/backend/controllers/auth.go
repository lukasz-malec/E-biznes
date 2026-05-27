package controllers

import (
	"net/http"
	"time"

	"backend/db"
	"backend/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)


var jwtSecret = []byte("jakis_tajny_klycz")

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}


func generateToken(userID int, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(jwtSecret)
}


func Register(c echo.Context) error {
	req := new(RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nieprawidłowe dane"})
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Wszystkie pola są wymagane"})
	}

	result, err := db.DB.Exec(
		"INSERT INTO users (name, email, password) VALUES (?, ?, ?)",
		req.Name, req.Email, req.Password,
	)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Email już zajęty"})
	}

	userID, _ := result.LastInsertId()

	token, err := generateToken(int(userID), req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Błąd generowania tokenu"})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"user": models.User{
			ID:    int(userID),
			Name:  req.Name,
			Email: req.Email,
		},
		"token": token,
	})
}



func Login(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nieprawidłowe dane"})
	}

	var user models.User
	err := db.DB.QueryRow(
		"SELECT id, name, email, password FROM users WHERE email = ?",
		req.Email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Nieprawidłowy email lub hasło"})
	}

	if user.Password != req.Password {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Nieprawidłowy email lub hasło"})
	}

	token, err := generateToken(user.ID, user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Błąd generowania tokenu"})
	}

	user.Password = ""

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":  user,
		"token": token,
	})
}