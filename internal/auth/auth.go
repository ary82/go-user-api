package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(email string) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   email,
	}

	key := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	return ss, err
}

func ParseJWT(ss string) (*string, error) {
	token, err := jwt.ParseWithClaims(
		ss,
		&jwt.RegisteredClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
	if err != nil {
		fmt.Println("HERE")
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("claims corrupted, ok: %v", token.Claims)
	}

	email, err := claims.GetSubject()
	if err != nil {
		return nil, err
	}
	return &email, nil
}

func Middleware(c *fiber.Ctx) error {
	s := c.Cookies("jwt")
	if s == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "no cookie supplied",
		})
	}

	email, err := ParseJWT(s)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	c.Locals("email", *email)
	return c.Next()
}
