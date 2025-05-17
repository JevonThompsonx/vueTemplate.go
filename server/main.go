package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Message is a simple struct to send JSON responses
type Message struct {
	Text string `json:"text"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())  // Logs requests
	e.Use(middleware.Recover()) // Recovers from panics

	// CORS Middleware Configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// IMPORTANT: Double-check that this URL matches exactly where your Vue dev server is running.
		// If Vite chose a different port (e.g., 5174), update this.
		AllowOrigins: []string{"http://localhost:5173"},

		// You can also use a wildcard for development, BUT BE CAREFUL and
		// restrict it for production:
		// AllowOrigins: []string{"*"}, // Less secure, for debugging only

		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization}, // Added Authorization
		// AllowCredentials: true, // Uncomment if your frontend sends credentials (cookies, authorization headers)
		// ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin"}, // Optional: if frontend needs to read these
		// MaxAge: 86400, // Optional: how long the result of a preflight request can be cached
	}))

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from Go Echo Server!")
	})

	e.GET("/api/hello", func(c echo.Context) error {
		msg := &Message{
			Text: "Hello from Go Echo API!",
		}
		return c.JSON(http.StatusOK, msg)
	})

	// Start server
	// You can change the port if needed, e.g., ":8080"
	e.Logger.Fatal(e.Start(":8080"))
}
