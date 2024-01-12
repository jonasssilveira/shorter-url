package handlers

import (
	"encoding/json"
	"github.com/btcsuite/btcutil/base58"
	"github.com/gofiber/fiber/v2"
	"urlShorter/domain/entity"
)

func Handler(c *fiber.Ctx) error {
	body := c.Body()

	var input entity.Url
	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	encodedLink := base58.Encode([]byte(input.FullUrl))

	return c.JSON(map[string]string{"short_url": encodedLink[:6]})
}
