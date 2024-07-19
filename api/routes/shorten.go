package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/j0sal/go-url-shortner/helpers"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/gofiber/fiber.v2"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}
type response struct {
	URL            string        `json:"url"`
	CustomShort    string        `json:"short"`
	Expiry         time.Duration `json:"expiry"`
	XRateRemaining int           `json:"rate_limit"`
	XRateLimitRest time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	// Implement rate limiting

	// Check if input is an actual URL

	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	// Check for domain error

	if helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "Invalid URL - Nice try -_*"})
	}
	// enforce Https, SSL

	body.URL = helpers.enforceHTTP(body.URL)
}
