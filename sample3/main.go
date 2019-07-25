package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost:6379" // default host, maybe it works
	}

	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to my third website! \nPOST to /:key?value=value to save \nGET to /:key to retrieve")
	})

	e.POST("/:key", func(c echo.Context) error {
		key := c.Param("key")
		value := c.QueryParam("value")
		err := client.Set(key, value, 0).Err()

		if err != nil {
			return &echo.HTTPError{
				Message:  "Not able to save",
				Internal: err,
				Code:     http.StatusInternalServerError,
			}
		}

		msg := fmt.Sprintf("%s=%s saved", key, value)
		return c.String(http.StatusOK, msg)
	})

	e.GET("/:key", func(c echo.Context) error {
		key := c.Param("key")
		value, err := client.Get(key).Result()

		if err != nil {
			return &echo.HTTPError{
				Message:  "Not able to fetch",
				Internal: err,
				Code:     http.StatusInternalServerError,
			}
		}

		msg := fmt.Sprintf("%s=%s", key, value)
		return c.String(http.StatusOK, msg)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
