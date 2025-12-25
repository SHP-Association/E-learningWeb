package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SHP-Association/E-learningWeb/backend-go/internal/config"
	"github.com/SHP-Association/E-learningWeb/backend-go/internal/store"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// API holds the dependencies for the API handlers
type API struct {
	store        store.Store
	config       *config.Config
	sessionStore sessions.Store
}

// Opts contains options for creating a new API instance
type Opts struct {
	Store         store.Store
	Config        *config.Config
	SessionSecret string
}

// New creates a new API instance
func New(opts Opts) *API {
	return &API{
		store:        opts.Store,
		config:       opts.Config,
		sessionStore: sessions.NewCookieStore([]byte(opts.SessionSecret)),
	}
}

// Routes sets up all API routes
func (api *API) Routes() http.Handler {
	r := mux.NewRouter()

	// Middleware
	r.Use(api.corsMiddleware)
	r.Use(api.loggingMiddleware)

	// Create /api subrouter
	apiRouter := r.PathPrefix("/api").Subrouter()

	// Health check
	apiRouter.HandleFunc("/health", api.healthCheck).Methods(http.MethodGet)

	// Authentication routes
	apiRouter.HandleFunc("/login", api.login).Methods(http.MethodPost)
	apiRouter.HandleFunc("/register", api.register).Methods(http.MethodPost)
	apiRouter.HandleFunc("/logout", api.logout).Methods(http.MethodPost)

	// User routes
	apiRouter.HandleFunc("/users", api.listUsers).Methods(http.MethodGet)
	apiRouter.HandleFunc("/users", api.createUser).Methods(http.MethodPost)
	apiRouter.HandleFunc("/users/{id}", api.getUser).Methods(http.MethodGet)
	apiRouter.HandleFunc("/users/{id}", api.updateUser).Methods(http.MethodPut)
	apiRouter.HandleFunc("/users/{id}", api.deleteUser).Methods(http.MethodDelete)

	// Category routes
	apiRouter.HandleFunc("/categories", api.listCategories).Methods(http.MethodGet)
	apiRouter.HandleFunc("/categories", api.createCategory).Methods(http.MethodPost)
	apiRouter.HandleFunc("/categories/{id}", api.getCategory).Methods(http.MethodGet)
	apiRouter.HandleFunc("/categories/{id}", api.updateCategory).Methods(http.MethodPut)
	apiRouter.HandleFunc("/categories/{id}", api.deleteCategory).Methods(http.MethodDelete)

	// Course routes
	apiRouter.HandleFunc("/courses", api.listCourses).Methods(http.MethodGet)
	apiRouter.HandleFunc("/courses", api.createCourse).Methods(http.MethodPost)
	apiRouter.HandleFunc("/courses/{id}", api.getCourse).Methods(http.MethodGet)
	apiRouter.HandleFunc("/courses/{id}", api.updateCourse).Methods(http.MethodPut)
	apiRouter.HandleFunc("/courses/{id}", api.deleteCourse).Methods(http.MethodDelete)

	// Lesson routes
	apiRouter.HandleFunc("/lessons", api.listLessons).Methods(http.MethodGet)
	apiRouter.HandleFunc("/lessons", api.createLesson).Methods(http.MethodPost)
	apiRouter.HandleFunc("/lessons/{id}", api.getLesson).Methods(http.MethodGet)
	apiRouter.HandleFunc("/lessons/{id}", api.updateLesson).Methods(http.MethodPut)
	apiRouter.HandleFunc("/lessons/{id}", api.deleteLesson).Methods(http.MethodDelete)

	// Enrollment routes
	apiRouter.HandleFunc("/enrollments", api.listEnrollments).Methods(http.MethodGet)
	apiRouter.HandleFunc("/enrollments", api.createEnrollment).Methods(http.MethodPost)
	apiRouter.HandleFunc("/enrollments/{id}", api.getEnrollment).Methods(http.MethodGet)
	apiRouter.HandleFunc("/enrollments/{id}", api.updateEnrollment).Methods(http.MethodPut)
	apiRouter.HandleFunc("/enrollments/{id}", api.deleteEnrollment).Methods(http.MethodDelete)

	// Quiz routes
	apiRouter.HandleFunc("/quizzes", api.listQuizzes).Methods(http.MethodGet)
	apiRouter.HandleFunc("/quizzes", api.createQuiz).Methods(http.MethodPost)
	apiRouter.HandleFunc("/quizzes/{id}", api.getQuiz).Methods(http.MethodGet)
	apiRouter.HandleFunc("/quizzes/{id}", api.updateQuiz).Methods(http.MethodPut)
	apiRouter.HandleFunc("/quizzes/{id}", api.deleteQuiz).Methods(http.MethodDelete)

	// FAQ routes
	apiRouter.HandleFunc("/faqs", api.listFAQs).Methods(http.MethodGet)
	apiRouter.HandleFunc("/faqs", api.createFAQ).Methods(http.MethodPost)
	apiRouter.HandleFunc("/faqs/{id}", api.getFAQ).Methods(http.MethodGet)
	apiRouter.HandleFunc("/faqs/{id}", api.updateFAQ).Methods(http.MethodPut)
	apiRouter.HandleFunc("/faqs/{id}", api.deleteFAQ).Methods(http.MethodDelete)

	return r
}

// Helper functions

// respondJSON sends a JSON response
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// respondError sends an error response
func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

// getIDFromPath extracts the ID parameter from the URL path
func getIDFromPath(r *http.Request) (uint, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

// getPaginationParams extracts pagination parameters from query string
func getPaginationParams(r *http.Request) (page, pageSize int) {
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("page_size")

	page, _ = strconv.Atoi(pageStr)
	pageSize, _ = strconv.Atoi(pageSizeStr)

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	if pageSize > 100 {
		pageSize = 100
	}

	return page, pageSize
}

// healthCheck handler
func (api *API) healthCheck(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{
		"status":  "healthy",
		"service": "e-learning-backend-go",
	})
}
