package api

import (
	"encoding/json"
	"net/http"

	"github.com/SHP-Association/E-learningWeb/backend-go/internal/models"
	"github.com/SHP-Association/E-learningWeb/backend-go/internal/store"
)

// listCourses handles GET /api/courses
func (api *API) listCourses(w http.ResponseWriter, r *http.Request) {
	page, pageSize := getPaginationParams(r)
	search := r.URL.Query().Get("search")

	opts := store.ListOptions{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
		Filters:  make(map[string]interface{}),
	}

	// Add filters
	if isPublished := r.URL.Query().Get("is_published"); isPublished != "" {
		opts.Filters["is_published"] = isPublished == "true"
	}
	if level := r.URL.Query().Get("level"); level != "" {
		opts.Filters["level"] = level
	}
	if categoryID := r.URL.Query().Get("category_id"); categoryID != "" {
		opts.Filters["category_id"] = categoryID
	}

	courses, total, err := api.store.ListCourses(r.Context(), opts)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch courses")
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"results":   courses,
		"count":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// getCourse handles GET /api/courses/{id}
func (api *API) getCourse(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid course ID")
		return
	}

	course, err := api.store.GetCourseByID(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, "Course not found")
		return
	}

	respondJSON(w, http.StatusOK, course)
}

// createCourse handles POST /api/courses
func (api *API) createCourse(w http.ResponseWriter, r *http.Request) {
	var course models.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate required fields
	if course.Title == "" {
		respondError(w, http.StatusBadRequest, "Title is required")
		return
	}

	if err := api.store.CreateCourse(r.Context(), &course); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create course")
		return
	}

	respondJSON(w, http.StatusCreated, course)
}

// updateCourse handles PUT /api/courses/{id}
func (api *API) updateCourse(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid course ID")
		return
	}

	// Get existing course
	course, err := api.store.GetCourseByID(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, "Course not found")
		return
	}

	// Decode update data
	if err := json.NewDecoder(r.Body).Decode(course); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Ensure ID doesn't change
	course.ID = id

	if err := api.store.UpdateCourse(r.Context(), course); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to update course")
		return
	}

	respondJSON(w, http.StatusOK, course)
}

// deleteCourse handles DELETE /api/courses/{id}
func (api *API) deleteCourse(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid course ID")
		return
	}

	if err := api.store.DeleteCourse(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to delete course")
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Course deleted successfully"})
}
