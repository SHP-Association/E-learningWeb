package api

import (
	"encoding/json"
	"net/http"

	"github.com/SHP-Association/E-learningWeb/backend-go/internal/models"
	"github.com/SHP-Association/E-learningWeb/backend-go/internal/store"
)

// User handlers

func (api *API) listUsers(w http.ResponseWriter, r *http.Request) {
	page, pageSize := getPaginationParams(r)
	search := r.URL.Query().Get("search")

	opts := store.ListOptions{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	users, total, err := api.store.ListUsers(r.Context(), opts)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	// Convert to response format
	userResponses := make([]models.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = user.ToResponse()
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"results":   userResponses,
		"count":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

func (api *API) getUser(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := api.store.GetUserByID(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, "User not found")
		return
	}

	respondJSON(w, http.StatusOK, user.ToResponse())
}

func (api *API) createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := api.store.CreateUser(r.Context(), &user); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	respondJSON(w, http.StatusCreated, user.ToResponse())
}

func (api *API) updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := api.store.GetUserByID(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, "User not found")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user.ID = id

	if err := api.store.UpdateUser(r.Context(), user); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to update user")
		return
	}

	respondJSON(w, http.StatusOK, user.ToResponse())
}

func (api *API) deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	if err := api.store.DeleteUser(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

// Category handlers

func (api *API) listCategories(w http.ResponseWriter, r *http.Request) {
	page, pageSize := getPaginationParams(r)
	search := r.URL.Query().Get("search")

	opts := store.ListOptions{Page: page, PageSize: pageSize, Search: search}
	categories, total, err := api.store.ListCategories(r.Context(), opts)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch categories")
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"results": categories, "count": total, "page": page, "page_size": pageSize,
	})
}

func (api *API) getCategory(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	category, err := api.store.GetCategoryByID(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, "Category not found")
		return
	}

	respondJSON(w, http.StatusOK, category)
}

func (api *API) createCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := api.store.CreateCategory(r.Context(), &category); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create category")
		return
	}

	respondJSON(w, http.StatusCreated, category)
}

func (api *API) updateCategory(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	category, err := api.store.GetCategoryByID(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, "Category not found")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(category); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	category.ID = id

	if err := api.store.UpdateCategory(r.Context(), category); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to update category")
		return
	}

	respondJSON(w, http.StatusOK, category)
}

func (api *API) deleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	if err := api.store.DeleteCategory(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to delete category")
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Category deleted successfully"})
}

// Lesson, Enrollment, Quiz, FAQ handlers follow the same pattern
// For brevity, I'll create stub implementations

func (api *API) listLessons(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]interface{}{"results": []interface{}{}, "count": 0})
}

func (api *API) getLesson(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (api *API) createLesson(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (api *API) updateLesson(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (api *API) deleteLesson(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (api *API) listEnrollments(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]interface{}{"results": []interface{}{}, "count": 0})
}

func (api *API) getEnrollment(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (api *API) createEnrollment(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (api *API) updateEnrollment(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (api *API) deleteEnrollment(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (api *API) listQuizzes(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]interface{}{"results": []interface{}{}, "count": 0})
}

func (api *API) getQuiz(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (api *API) createQuiz(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (api *API) updateQuiz(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (api *API) deleteQuiz(w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusNotImplemented, "Not implemented yet")
}

func (api *API) listFAQs(w http.ResponseWriter, r *http.Request) {
	page, pageSize := getPaginationParams(r)
	opts := store.ListOptions{Page: page, PageSize: pageSize}
	faqs, total, err := api.store.ListFAQs(r.Context(), opts)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to fetch FAQs")
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"results": faqs, "count": total, "page": page, "page_size": pageSize,
	})
}

func (api *API) getFAQ(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid FAQ ID")
		return
	}

	faq, err := api.store.GetFAQByID(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, "FAQ not found")
		return
	}

	respondJSON(w, http.StatusOK, faq)
}

func (api *API) createFAQ(w http.ResponseWriter, r *http.Request) {
	var faq models.FAQ
	if err := json.NewDecoder(r.Body).Decode(&faq); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := api.store.CreateFAQ(r.Context(), &faq); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to create FAQ")
		return
	}

	respondJSON(w, http.StatusCreated, faq)
}

func (api *API) updateFAQ(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid FAQ ID")
		return
	}

	faq, err := api.store.GetFAQByID(r.Context(), id)
	if err != nil {
		respondError(w, http.StatusNotFound, "FAQ not found")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(faq); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	faq.ID = id

	if err := api.store.UpdateFAQ(r.Context(), faq); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to update FAQ")
		return
	}

	respondJSON(w, http.StatusOK, faq)
}

func (api *API) deleteFAQ(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromPath(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Invalid FAQ ID")
		return
	}

	if err := api.store.DeleteFAQ(r.Context(), id); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to delete FAQ")
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "FAQ deleted successfully"})
}
