package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHelloWorldRoute(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Error during test: %v", err)
	}

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Could not read response body: %v", err)
	}
	resp.Body.Close()

	assert.Equal(t, "Hello, World!", string(body))
}

func TestGetValueRoute(t *testing.T) {
	app := setupApp()

	req := httptest.NewRequest(http.MethodGet, "/testvalue", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body := httptest.NewRecorder()
	body.Body.ReadFrom(resp.Body)

	assert.Equal(t, "value: testvalue", body.Body.String())
}
