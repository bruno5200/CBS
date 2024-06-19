package middleware

import (
	"encoding/base64"
	"errors"
	"log"
	"strings"

	db "github.com/bruno5200/CSM/database"
	r "github.com/bruno5200/CSM/service/infrastructure/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils/v2"
)

const (
	headerKey = "X-API-KEY"
	preFix    = "Key "
)

var ErrInvalidApiKey = errors.New("invalid Key")

func ApiKey() fiber.Handler {
	return func(c *fiber.Ctx) error {

		// decode base64 string
		if err := apiKey(utils.CopyString(c.Get(headerKey))); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error(), "success": false})
		}

		return c.Next()
	}
}

func apiKey(auth string) error {

	if !strings.Contains(auth, preFix) {
		return ErrInvalidApiKey
	}

	// Split the base64 encoded string
	decoded, err := base64.StdEncoding.DecodeString(auth[len(preFix):])

	if err != nil {
		log.Printf("Error decoding key: %s", err)
		return ErrInvalidApiKey
	}

	if _, err = r.NewServiceRepository(db.PostgresDB()).ReadServiceByKey(string(decoded)); err != nil {
		log.Printf("Error reading service with Header Key: %s", err)
		return ErrInvalidApiKey
	}

	return nil
}
