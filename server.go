package main

import (
	_ "embed"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

//go:embed index.html
var indexHTML string

func NewServer(am *AMQPManager, queueName string) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	// Serve index.html
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, indexHTML)
	})

	// Log endpoint
	e.GET("/log/gps", func(c echo.Context) error {
		req := new(LogRequest)
		if err := c.Bind(req); err != nil {
			return c.String(http.StatusBadRequest, "Invalid request")
		}

		// Validate required fields
		if req.ID == "" || req.Key == "" || req.Lon == 0 || req.Lat == 0 {
			return c.String(http.StatusBadRequest, "All fields are required")
		}

		// set timestamp to now
		req.Timestamp = time.Now().UTC().Unix()
		// set event to "gps"
		req.Event = "gps"

		err := am.QueueMessage(queueName, req)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to queue message")
		}

		return c.String(http.StatusOK, "ok")
	})

	// Ignition endpoint
	e.GET("/log/ignition", func(c echo.Context) error {
		req := new(IgnitionLog)
		if err := c.Bind(req); err != nil {
			return c.String(http.StatusBadRequest, "Invalid request")
		}

		// Validate required fields
		if req.ID == "" || req.Key == "" {
			return c.String(http.StatusBadRequest, "All fields are required")
		}

		isEnabled :=  c.QueryParams().Get(("enabled")) == "1"
		req.Enabled = isEnabled

		// set timestamp to now
		req.Timestamp = time.Now().UTC().Unix()
		// set event to "ignition"
		req.Event = "ignition"

		err := am.QueueMessage(queueName, req)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to queue message")
		}

		return c.String(http.StatusOK, "ok")
	})

	// Ping endpoint
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	// Handle 404 for all other routes
	e.Any("/*", func(c echo.Context) error {
		return c.String(http.StatusNotFound, "Not Found")
	})

	return e
}
