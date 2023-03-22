package modules

import (
	"encoding/json"
	"os"
	"strings"
	"time"

	"github.com/dodysat/gue-auth/database"
	"github.com/dodysat/gue-auth/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateResponseUser(user models.User) User {
	return User{
		ID:       user.ID,
		Username: user.Username,
	}
}

func UserRegister(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing user",
		})
	}

	username := user.Username
	hash, _ := HashPassword(user.Password)

	database.Database.Db.Create(&models.User{
		Username: username,
		Password: hash,
	})

	return c.Status(201).JSON(CreateResponseUser(user))
}

func UserLogin(c *fiber.Ctx) error {
	var user models.User

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing user",
		})
	}

	username := user.Username
	password := user.Password

	var dbUser models.User
	database.Database.Db.Where("username = ?", username).First(&dbUser)

	if CheckPasswordHash(password, dbUser.Password) {
		userId := dbUser.ID

		accessTokenExpires := time.Now().Add(time.Minute * 15).Unix()
		refreshTokenExpires := time.Now().Add(time.Hour * 24 * 7).Unix()

		secret := os.Getenv("JWT_SECRET")
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"userId": userId, "exp": accessTokenExpires})
		accessTokenString, err := accessToken.SignedString([]byte(secret))

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"message": "Error creating access token"})
		}

		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"userId": userId, "exp": refreshTokenExpires})
		refreshTokenString, err := refreshToken.SignedString([]byte(secret))

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"message": "Error creating refresh token"})
		}

		database.Database.Db.Create(&models.Session{
			UserID:       userId,
			AccessToken:  accessTokenString,
			RefreshToken: refreshTokenString,
		})

		database.Database.Db.Create(&models.Activity{
			UserID:    userId,
			Service:   "auth",
			Path:      c.Path(),
			Method:    c.Method(),
			Timestamp: time.Now(),
		})

		return c.Status(200).JSON(fiber.Map{
			"message":            "Login success",
			"access_token":       accessTokenString,
			"refresh_token":      refreshTokenString,
			"access_token_exp":   accessTokenExpires,
			"refresh_token_exp":  refreshTokenExpires,
			"access_token_type":  "Bearer",
			"refresh_token_type": "Bearer",
		})

	} else {
		return c.Status(400).JSON(fiber.Map{
			"message": "Wrong password",
		})
	}
}

func UserLogout(c *fiber.Ctx) error {
	accessToken := c.Get("Authorization")
	accessToken = accessToken[7:]

	var session models.Session

	database.Database.Db.Where("access_token = ?", accessToken).First(&session)

	database.Database.Db.Delete(&session)

	return c.Status(200).JSON(fiber.Map{
		"message": "Logout success",
	})
}

func UserRefresh(c *fiber.Ctx) error {
	refreshtoken := c.Get("Authorization")
	refreshtoken = refreshtoken[7:]

	var session models.Session

	database.Database.Db.Where("refresh_token = ?", refreshtoken).First(&session)

	userId := session.UserID

	accessTokenExpires := time.Now().Add(time.Minute * 15).Unix()
	refreshTokenExpires := time.Now().Add(time.Hour * 24 * 7).Unix()

	secret := os.Getenv("JWT_SECRET")
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"userId": userId, "exp": accessTokenExpires})
	accessTokenString, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error creating access token"})
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"userId": userId, "exp": refreshTokenExpires})
	refreshTokenString, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Error creating refresh token"})
	}

	database.Database.Db.Model(&session).Updates(models.Session{AccessToken: accessTokenString, RefreshToken: refreshTokenString})

	return c.Status(200).JSON(fiber.Map{
		"message":            "Refresh success",
		"access_token":       accessTokenString,
		"refresh_token":      refreshTokenString,
		"access_token_exp":   accessTokenExpires,
		"refresh_token_exp":  refreshTokenExpires,
		"access_token_type":  "Bearer",
		"refresh_token_type": "Bearer",
	})
}

func UserVerify(c *fiber.Ctx) error {
	header := c.Get("Authorization")

	token := strings.Split(header, " ")[1]
	secret := os.Getenv("JWT_SECRET")
	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) { return []byte(secret), nil })
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Invalid Token"})
	}

	claims := decoded.Claims.(jwt.MapClaims)
	userID := claims["userId"]

	bodyRequest := c.Body()

	body := make(map[string]interface{})
	json.Unmarshal(bodyRequest, &body)

	database.Database.Db.Create(&models.Activity{
		UserID:    uint(userID.(float64)),
		Service:   body["service"].(string),
		Path:      body["path"].(string),
		Method:    body["method"].(string),
		Timestamp: time.Now(),
	})

	return c.Status(200).JSON(fiber.Map{"message": "Verified", "userId": userID})
}
