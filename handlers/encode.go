package handlers

import (
	"encoding/json"
	"github.com/btcsuite/btcutil/base58"
	"github.com/gofiber/fiber/v2"
	"time"
	db "urlShorter/db/sqlc"
	"urlShorter/domain/entity"
	"urlShorter/usecase"
)

type EncodeURL struct {
	useCaseURL usecase.URL
}

func NewEncodeURL(useCaseURL usecase.URL) EncodeURL {
	return EncodeURL{useCaseURL: useCaseURL}
}

func (e *EncodeURL) Handler(c *fiber.Ctx) error {

	body := c.Body()
	var input entity.URL
	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	encodedLink := base58.Encode([]byte(input.FullURL))[:6]

	url, err := e.useCaseURL.CreateURL(c.Context(), db.CreateURLParams{
		UrlEncoded:     encodedLink,
		UrlOriginal:    input.FullURL,
		ExpirationDate: time.Time{},
	})

	if err != nil {
		return err
	}
	return c.JSON(url)
}
