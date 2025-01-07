package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/prateeksrivastav2/UrlShortner/api/database"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")

	// Create the Redis client for data retrieval (db 0)
	r := database.CreateClient(0)
	defer r.Close()

	// Corrected line: Only pass the key (url), no need to pass context
	value, err := r.Get(url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Short URL not found in the Database",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	// Track usage - Create another Redis client for incrementing the counter (db 1)
	rInr := database.CreateClient(1)
	defer rInr.Close()

	// Increment the counter for tracking the usage of short URLs
	_ = rInr.Incr("counter")

	// Redirect to the original URL
	return c.Redirect(value, 301)
}
