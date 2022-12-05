package user

import (
	"fmt"
	"net/http/httptest"
	"test/db"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetHandler(t *testing.T) {
	database, err := db.Connect()
	assert.Nil(t, err)

	repo := NewRepository(database)
	service := NewService(repo)
	handler := NewHandler(service)

	app := fiber.New()
	app.Get("/users/:id", handler.Get)

	id, err := repo.Create(Model{Name: "test", Email: "test@testmail.com"})
	assert.Nil(t, err)

	req := httptest.NewRequest("GET", fmt.Sprintf("/users/%d", id), nil)
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

}
