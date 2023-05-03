package routes

import (
	models "myapp/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func NewUserRoute(uGroup *echo.Group) {
	uGroup.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		u := models.GetUser(id)
		if u == nil {
			return c.JSON(http.StatusBadRequest, models.MessageResponse{
				Message: "User not found",
			})
		}
		return c.JSON(http.StatusOK, models.GetUserResponse{
			ID:   u.GetID(),
			Name: u.GetName(),
			Date: time.Now().UTC(),
		})
	})

	uGroup.POST("/register", func(c echo.Context) error {
		req := new(models.RegisterRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, models.MessageResponse{
				Message: err.Error(),
			})
		}

		id := req.ID
		name := req.Name

		if id == "" || name == "" {
			return c.JSON(http.StatusBadRequest, models.MessageResponse{
				Message: "Invaild id or name",
			})
		}

		if models.GetUser(id) != nil {
			return c.JSON(http.StatusBadRequest, models.MessageResponse{
				Message: "This id register",
			})
		}

		models.NewUser(id, name)
		return c.JSON(http.StatusOK, models.MessageResponse{
			Message: "Register Success",
		})
	})

	uGroup.GET("/:id/name", func(c echo.Context) error {
		id := c.Param("id")
		u := models.GetUser(id)
		if u == nil {
			return c.JSON(http.StatusBadRequest, models.MessageResponse{
				Message: "Name not found",
			})
		}

		return c.JSON(http.StatusOK, u.GetName())
	})
}
