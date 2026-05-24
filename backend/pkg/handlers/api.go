package handlers

import (
	"crypto/rand"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/SHP-Association/E-learningWeb/backend/ent"
	"github.com/SHP-Association/E-learningWeb/backend/ent/user"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/services"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/tasks"
	"github.com/labstack/echo/v4"
)

const (
	otpTTL            = 15 * time.Minute
	otpResendCooldown = 60 * time.Second
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
	auth.POST("/verify-otp", h.verifyOTP)
	auth.GET("/signup-config", h.signupConfig)
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

	// Check if email has been verified.
	if !u.Verified && !u.IsEmailVerified {
		return ctx.JSON(http.StatusForbidden, echo.Map{
			"error":                "email verification required",
			"verification_pending": true,
			"email":                u.Email,
		})
	}

	// Log the user in.
	err = h.auth.Login(ctx, u.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "unable to log in user"})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "logged in successfully",
		"user":    publicUserPayload(u),
	})
}

// generateOTP generates a secure, random 6-digit numeric string
func generateOTP() string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, 6)
	n, err := io.ReadAtLeast(rand.Reader, b, 6)
	if n != 6 || err != nil {
		return "582914" // Secure fallback
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func (h *API) register(ctx echo.Context) error {
	type registerInput struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
		Role     string `json:"role"`
	}

	var input registerInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	if strings.TrimSpace(input.Username) == "" || strings.TrimSpace(input.Email) == "" || strings.TrimSpace(input.Password) == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "username, email and password are required"})
	}

	if input.Role != "" && strings.ToLower(strings.TrimSpace(input.Role)) != "student" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "web signup is only available for students"})
	}

	normalizedEmail := strings.ToLower(strings.TrimSpace(input.Email))
	trimmedUsername := strings.TrimSpace(input.Username)

	existingUser, err := h.orm.ORM.User.
		Query().
		Where(user.Email(normalizedEmail)).
		Only(ctx.Request().Context())

	if err == nil {
		if existingUser.Verified || existingUser.IsEmailVerified {
			return ctx.JSON(http.StatusConflict, echo.Map{"error": "a user with this email or username already exists"})
		}

		if strings.ToLower(existingUser.Role) != "student" {
			return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "web signup is only available for students"})
		}

		if !existingUser.OtpExpiresAt.IsZero() {
			lastSentAt := existingUser.OtpExpiresAt.Add(-otpTTL)
			nextAllowedAt := lastSentAt.Add(otpResendCooldown)
			if time.Now().Before(nextAllowedAt) {
				retryAfter := int(time.Until(nextAllowedAt).Seconds())
				if retryAfter < 1 {
					retryAfter = 1
				}
				return ctx.JSON(http.StatusTooManyRequests, echo.Map{
					"error":               "please wait before requesting another code",
					"retry_after_seconds": retryAfter,
				})
			}
		}

		otpCode := generateOTP()
		otpExpiry := time.Now().Add(otpTTL)

		updatedUser, updateErr := existingUser.Update().
			SetUsername(trimmedUsername).
			SetPassword(input.Password).
			SetRole("student").
			SetVerified(false).
			SetIsEmailVerified(false).
			SetOtpCode(otpCode).
			SetOtpExpiresAt(otpExpiry).
			Save(ctx.Request().Context())
		if updateErr != nil {
			if ent.IsConstraintError(updateErr) {
				return ctx.JSON(http.StatusConflict, echo.Map{"error": "a user with this email or username already exists"})
			}
			return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "unable to update pending signup"})
		}

		_, _ = h.orm.Tasks.
			Add(tasks.SendOTPTask{
				Email:    updatedUser.Email,
				Username: updatedUser.Username,
				OTP:      otpCode,
			}).
			Save()

		return ctx.JSON(http.StatusOK, echo.Map{
			"message":               "verification code sent to your email",
			"verification_required": true,
			"email":                 updatedUser.Email,
			"role":                  "student",
		})
	}

	if err != nil && !ent.IsNotFound(err) {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "unable to check existing user"})
	}

	otpCode := generateOTP()
	otpExpiry := time.Now().Add(otpTTL)

	u, err := h.orm.ORM.User.
		Create().
		SetUsername(trimmedUsername).
		SetEmail(normalizedEmail).
		SetPassword(input.Password).
		SetRole("student").
		SetVerified(false).
		SetIsEmailVerified(false).
		SetOtpCode(otpCode).
		SetOtpExpiresAt(otpExpiry).
		Save(ctx.Request().Context())

	if err != nil {
		if ent.IsConstraintError(err) {
			return ctx.JSON(http.StatusConflict, echo.Map{"error": "a user with this email or username already exists"})
		}
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "unable to create user"})
	}

	_, _ = h.orm.Tasks.
		Add(tasks.SendOTPTask{
			Email:    u.Email,
			Username: u.Username,
			OTP:      otpCode,
		}).
		Save()

	return ctx.JSON(http.StatusCreated, echo.Map{
		"message":               "account created, verification code sent to your email",
		"verification_required": true,
		"email":                 u.Email,
		"role":                  "student",
	})
}

func (h *API) verifyOTP(ctx echo.Context) error {
	type verifyInput struct {
		Email string `json:"email" validate:"required,email"`
		OTP   string `json:"otp" validate:"required,len=6"`
	}

	var input verifyInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	// Find the user by email
	u, err := h.orm.ORM.User.
		Query().
		Where(user.Email(strings.ToLower(input.Email))).
		Only(ctx.Request().Context())

	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "user not found"})
	}

	if u.Verified || u.IsEmailVerified {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "email is already verified"})
	}

	// Validate OTP and expiration
	if u.OtpCode == "" || u.OtpExpiresAt.IsZero() {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "no verification request active"})
	}

	if u.OtpCode != input.OTP {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "incorrect verification code"})
	}

	if time.Now().After(u.OtpExpiresAt) {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "verification code has expired"})
	}

	// Update the user as verified and clear OTP fields
	u, err = u.Update().
		SetVerified(true).
		SetIsEmailVerified(true).
		ClearOtpCode().
		ClearOtpExpiresAt().
		Save(ctx.Request().Context())

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "unable to complete verification"})
	}

	// Log the user in
	err = h.auth.Login(ctx, u.ID)
	if err != nil {
		return ctx.JSON(http.StatusOK, echo.Map{
			"message": "email verified successfully, please log in manually",
			"user":    publicUserPayload(u),
		})
	}

	// Dispatch welcome email task as post-verification follow-up
	_, _ = h.orm.Tasks.
		Add(tasks.WelcomeEmailTask{
			UserID:   u.ID,
			Username: u.Username,
			Email:    u.Email,
		}).
		Save()

	return ctx.JSON(http.StatusOK, echo.Map{
		"message":            "verification successful, logged in",
		"user":               publicUserPayload(u),
		"onboarding_required": userNeedsOnboarding(u),
	})
}

func (h *API) signupConfig(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"instructor_contact_email": h.orm.Config.Mail.FromAddress,
	})
}

func (h *API) logout(ctx echo.Context) error {
	if err := h.auth.Logout(ctx); err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to logout"})
	}
	return ctx.JSON(http.StatusOK, echo.Map{"message": "logged out"})
}

func userNeedsOnboarding(u *ent.User) bool {
	return strings.TrimSpace(u.FirstName) == "" ||
		strings.TrimSpace(u.LastName) == "" ||
		strings.TrimSpace(u.ContactNumber) == "" ||
		strings.TrimSpace(u.Country) == ""
}

func publicUserPayload(u *ent.User) echo.Map {
	return echo.Map{
		"id":                  u.ID,
		"username":            u.Username,
		"email":               u.Email,
		"role":                u.Role,
		"first_name":          u.FirstName,
		"last_name":           u.LastName,
		"contact_number":      u.ContactNumber,
		"country":             u.Country,
		"onboarding_required": userNeedsOnboarding(u),
	}
}
