package handlers

import (
	"net/http"
	"strings"

	"github.com/SHP-Association/E-learningWeb/backend/ent"
	"github.com/SHP-Association/E-learningWeb/backend/ent/user"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/services"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/tasks"
	"github.com/labstack/echo/v4"
)

type API struct {
	orm  *services.Container
	auth *services.AuthClient
}

func init() {
	Register(new(API))
}

func (h *API) Init(c *services.Container) error {
	h.orm = c
	h.auth = c.Auth
	return nil
}

func (h *API) Routes(g *echo.Group) {
	// API group
	api := g.Group("/api")

	// Health check
	api.GET("/health", h.health)

	// Auth routes
	auth := api.Group("/auth")
	auth.POST("/login", h.login)
	auth.POST("/register", h.register)
	auth.POST("/logout", h.logout)

	// Backward-compatible auth aliases under /api/*
	api.POST("/login", h.login)
	api.POST("/register", h.register)
	api.POST("/logout", h.logout)
}

func (h *API) health(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"status": "up",
	})
}

func (h *API) login(ctx echo.Context) error {
	type loginInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	var input loginInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	// Attempt to load the user.
	u, err := h.orm.ORM.User.
		Query().
		Where(user.Email(strings.ToLower(input.Email))).
		Only(ctx.Request().Context())

	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid credentials"})
	}

	// Check if the password is correct.
	err = h.auth.CheckPassword(input.Password, u.Password)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid credentials"})
	}

	// Log the user in.
	err = h.auth.Login(ctx, u.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "unable to log in user"})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "logged in successfully",
		"user": echo.Map{
			"id":       u.ID,
			"username": u.Username,
			"email":    u.Email,
			"role":     u.Role,
		},
	})
}

func (h *API) register(ctx echo.Context) error {
	type registerInput struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	var input registerInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	// Attempt creating the user.
	u, err := h.orm.ORM.User.
		Create().
		SetUsername(input.Username).
		SetEmail(input.Email).
		SetPassword(input.Password).
		Save(ctx.Request().Context())

	if err != nil {
		if ent.IsConstraintError(err) {
			return ctx.JSON(http.StatusConflict, echo.Map{"error": "a user with this email or username already exists"})
		}
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "unable to create user"})
	}

	// Log the user in.
	err = h.auth.Login(ctx, u.ID)
	if err != nil {
		return ctx.JSON(http.StatusOK, echo.Map{
			"message": "account created, but failed to log in automatically",
			"user":    u,
		})
	}

	// Dispatch welcome email task
	_, _ = h.orm.Tasks.
		Add(tasks.WelcomeEmailTask{
			UserID:   u.ID,
			Username: u.Username,
			Email:    u.Email,
		}).
		Save()

	return ctx.JSON(http.StatusCreated, echo.Map{
		"message": "account created and logged in",
		"user": echo.Map{
			"id":       u.ID,
			"username": u.Username,
			"email":    u.Email,
			"role":     u.Role,
		},
	})
}

func (h *API) logout(ctx echo.Context) error {
	if err := h.auth.Logout(ctx); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to logout"})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "logged out"})
}
