package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/SHP-Association/E-learningWeb/backend/ent"
	"github.com/SHP-Association/E-learningWeb/backend/ent/course"
	"github.com/SHP-Association/E-learningWeb/backend/ent/enrollment"
	"github.com/SHP-Association/E-learningWeb/backend/ent/faq"
	"github.com/SHP-Association/E-learningWeb/backend/ent/lesson"
	"github.com/SHP-Association/E-learningWeb/backend/ent/quiz"
	"github.com/SHP-Association/E-learningWeb/backend/ent/review"
	"github.com/SHP-Association/E-learningWeb/backend/ent/user"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/context"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/middleware"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/services"
	"github.com/labstack/echo/v4"
)

type LMSAPI struct {
	orm *ent.Client
}

func init() {
	Register(new(LMSAPI))
}

func (h *LMSAPI) Init(c *services.Container) error {
	h.orm = c.ORM
	return nil
}

func (h *LMSAPI) Routes(g *echo.Group) {
	api := g.Group("/api")

	// Lessons
	api.GET("/lessons/:id", h.getLesson)

	// Quizzes
	api.GET("/quizzes/:id", h.getQuiz)
	api.POST("/quizzes/:id/submit", h.submitQuiz, middleware.RequireAuthentication)

	// FAQs
	api.GET("/faqs", h.listFAQs)

	// Reviews
	api.GET("/courses/:id/reviews", h.listReviews)
	api.POST("/courses/:id/reviews", h.submitReview, middleware.RequireAuthentication)

	// Enrollment (Protected)
	auth := api.Group("", middleware.RequireAuthentication)
	auth.POST("/enroll/:slug", h.enroll)
	auth.GET("/profile", h.getProfile)
	auth.PATCH("/profile/onboarding", h.updateOnboarding)
}

func (h *LMSAPI) getLesson(ctx echo.Context) error {
	id, err := parseIntParam(ctx, "id")
	if err != nil {
		return err
	}
	l, err := h.orm.Lesson.
		Query().
		Where(lesson.ID(id)).
		WithQuizzes().
		Only(ctx.Request().Context())

	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "lesson not found"})
	}

	return ctx.JSON(http.StatusOK, l)
}

func (h *LMSAPI) getQuiz(ctx echo.Context) error {
	id, err := parseIntParam(ctx, "id")
	if err != nil {
		return err
	}
	q, err := h.orm.Quiz.
		Query().
		Where(quiz.ID(id)).
		WithQuestions(func(q *ent.QuestionQuery) {
			q.WithChoices()
		}).
		Only(ctx.Request().Context())

	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "quiz not found"})
	}

	return ctx.JSON(http.StatusOK, q)
}

func (h *LMSAPI) listFAQs(ctx echo.Context) error {
	faqs, err := h.orm.FAQ.
		Query().
		Where(faq.IsPublished(true)).
		Order(ent.Asc(faq.FieldOrder)).
		All(ctx.Request().Context())

	if err != nil {
		return jsonInternalError(ctx, "unable to list FAQs")
	}

	return ctx.JSON(http.StatusOK, faqs)
}

func (h *LMSAPI) enroll(ctx echo.Context) error {
	slug := ctx.Param("slug")
	u := ctx.Get(context.AuthenticatedUserKey).(*ent.User)

	// Get course
	c, err := h.orm.Course.
		Query().
		Where(course.Slug(slug)).
		Only(ctx.Request().Context())

	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "course not found"})
	}

	// Check if already enrolled
	exists, err := h.orm.Enrollment.
		Query().
		Where(
			enrollment.HasStudentWith(user.ID(u.ID)),
			enrollment.HasCourseWith(course.ID(c.ID)),
		).
		Exist(ctx.Request().Context())

	if err != nil {
		return jsonInternalError(ctx, "error checking enrollment")
	}

	if exists {
		return ctx.JSON(http.StatusConflict, echo.Map{"error": "already enrolled"})
	}

	// Create enrollment
	enroll, err := h.orm.Enrollment.
		Create().
		SetStudent(u).
		SetCourse(c).
		Save(ctx.Request().Context())

	if err != nil {
		return jsonInternalError(ctx, "failed to enroll")
	}

	return ctx.JSON(http.StatusCreated, echo.Map{
		"message":    fmt.Sprintf("successfully enrolled in %s", c.Title),
		"enrollment": enroll,
	})
}

func (h *LMSAPI) getProfile(ctx echo.Context) error {
	u := ctx.Get(context.AuthenticatedUserKey).(*ent.User)

	// Get user with enrollments
	userWithData, err := h.orm.User.
		Query().
		Where(user.ID(u.ID)).
		WithEnrollments(func(q *ent.EnrollmentQuery) {
			q.WithCourse()
		}).
		Only(ctx.Request().Context())

	if err != nil {
		return jsonInternalError(ctx, "unable to load profile")
	}

	return ctx.JSON(http.StatusOK, userWithData)
}

func (h *LMSAPI) updateOnboarding(ctx echo.Context) error {
	type onboardingInput struct {
		FirstName     string `json:"first_name"`
		LastName      string `json:"last_name"`
		ContactNumber string `json:"contact_number"`
		Country       string `json:"country"`
	}

	var input onboardingInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	input.FirstName = strings.TrimSpace(input.FirstName)
	input.LastName = strings.TrimSpace(input.LastName)
	input.ContactNumber = strings.TrimSpace(input.ContactNumber)
	input.Country = strings.TrimSpace(input.Country)

	validationErrors := map[string][]string{}
	if input.FirstName == "" {
		validationErrors["first_name"] = []string{"first name is required"}
	}
	if input.LastName == "" {
		validationErrors["last_name"] = []string{"last name is required"}
	}
	if input.ContactNumber == "" {
		validationErrors["contact_number"] = []string{"contact number is required"}
	}
	if input.Country == "" {
		validationErrors["country"] = []string{"country is required"}
	}

	if len(validationErrors) > 0 {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error":  "validation failed",
			"errors": validationErrors,
		})
	}

	authUser := ctx.Get(context.AuthenticatedUserKey).(*ent.User)
	updatedUser, err := h.orm.User.
		UpdateOneID(authUser.ID).
		SetFirstName(input.FirstName).
		SetLastName(input.LastName).
		SetContactNumber(input.ContactNumber).
		SetCountry(input.Country).
		Save(ctx.Request().Context())
	if err != nil {
		return jsonInternalError(ctx, "unable to update onboarding details")
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "onboarding details saved",
		"user":    updatedUser,
	})
}

func (h *LMSAPI) submitReview(ctx echo.Context) error {
	courseID, err := parseIntParam(ctx, "id")
	if err != nil {
		return err
	}
	u := ctx.Get(context.AuthenticatedUserKey).(*ent.User)

	type reviewInput struct {
		Rating  int    `json:"rating" validate:"required,min=1,max=5"`
		Comment string `json:"comment"`
	}

	var input reviewInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	// Create review
	rev, err := h.orm.Review.
		Create().
		SetCourseID(courseID).
		SetStudent(u).
		SetRating(input.Rating).
		SetComment(input.Comment).
		SetIsApproved(true). // Auto-approve for now
		Save(ctx.Request().Context())

	if err != nil {
		return jsonInternalError(ctx, "failed to submit review")
	}

	return ctx.JSON(http.StatusCreated, rev)
}

func (h *LMSAPI) listReviews(ctx echo.Context) error {
	courseID, err := parseIntParam(ctx, "id")
	if err != nil {
		return err
	}
	reviews, err := h.orm.Review.
		Query().
		Where(
			review.HasCourseWith(course.ID(courseID)),
			review.IsApproved(true),
		).
		WithStudent().
		Order(ent.Desc(review.FieldCreatedAt)).
		All(ctx.Request().Context())

	if err != nil {
		return jsonInternalError(ctx, "unable to list reviews")
	}

	return ctx.JSON(http.StatusOK, reviews)
}

func (h *LMSAPI) submitQuiz(ctx echo.Context) error {
	quizID, err := parseIntParam(ctx, "id")
	if err != nil {
		return err
	}
	u := ctx.Get(context.AuthenticatedUserKey).(*ent.User)

	type attemptInput struct {
		Answers map[int]int `json:"answers"` // questionID -> choiceID
	}

	var input attemptInput
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	// In a real app, you'd calculate the score here.
	// For now, we'll just record the attempt.
	attempt, err := h.orm.UserQuizAttempt.
		Create().
		SetQuizID(quizID).
		SetStudent(u).
		SetScore(100.0). // Simulated score
		SetPassed(true).
		Save(ctx.Request().Context())

	if err != nil {
		return jsonInternalError(ctx, "failed to record quiz attempt")
	}

	return ctx.JSON(http.StatusCreated, attempt)
}
