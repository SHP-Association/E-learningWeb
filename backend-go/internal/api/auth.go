package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/SHP-Association/E-learningWeb/backend-go/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

// JWT Claims
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// generateJWT creates a new JWT token for a user
func (api *API) generateJWT(user *models.User) (string, error) {
	expirationTime := time.Now().Add(api.config.JWT.Expiry)
	claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(api.config.JWT.Secret))
}

// verifyJWT verifies a JWT token and returns the user ID
func (api *API) verifyJWT(tokenString string) (uint, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(api.config.JWT.Secret), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	return claims.UserID, nil
}

// Context helpers
type contextKey string

const userContextKey contextKey = "user"

func setUserInContext(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func getUserFromContext(ctx context.Context) *models.User {
	user, ok := ctx.Value(userContextKey).(*models.User)
	if !ok {
		return nil
	}
	return user
}

// Auth handlers

// LoginRequest represents the login request payload
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	Token string              `json:"token"`
	User  models.UserResponse `json:"user"`
}

// login handles user login
func (api *API) login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Get user by username or email
	var user *models.User
	var err error

	user, err = api.store.GetUserByUsername(r.Context(), req.Username)
	if err != nil {
		// Try by email
		user, err = api.store.GetUserByEmail(r.Context(), req.Username)
		if err != nil {
			respondError(w, http.StatusUnauthorized, "Invalid credentials")
			return
		}
	}

	// Check password
	if !user.CheckPassword(req.Password) {
		respondError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Update last login
	now := time.Now()
	user.LastLogin = &now
	api.store.UpdateUser(r.Context(), user)

	// Generate JWT token
	token, err := api.generateJWT(user)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	respondJSON(w, http.StatusOK, LoginResponse{
		Token: token,
		User:  user.ToResponse(),
	})
}

// RegisterRequest represents the registration request payload
type RegisterRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// register handles user registration
func (api *API) register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate required fields
	if req.Username == "" || req.Email == "" || req.Password == "" {
		respondError(w, http.StatusBadRequest, "Username, email, and password are required")
		return
	}

	// Check if user already exists
	_, err := api.store.GetUserByUsername(r.Context(), req.Username)
	if err == nil {
		respondError(w, http.StatusConflict, "Username already exists")
		return
	}

	_, err = api.store.GetUserByEmail(r.Context(), req.Email)
	if err == nil {
		respondError(w, http.StatusConflict, "Email already exists")
		return
	}

	// Create new user
	user := &models.User{
		Username:  req.Username,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      "student",
		IsActive:  true,
	}

	// Hash password
	if err := user.SetPassword(req.Password); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	// Save user
	if err := api.store.CreateUser(r.Context(), user); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	// Generate JWT token
	token, err := api.generateJWT(user)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	respondJSON(w, http.StatusCreated, LoginResponse{
		Token: token,
		User:  user.ToResponse(),
	})
}

// logout handles user logout
func (api *API) logout(w http.ResponseWriter, r *http.Request) {
	// For JWT-based auth, logout is handled client-side by removing the token
	respondJSON(w, http.StatusOK, map[string]string{"message": "Logged out successfully"})
}
