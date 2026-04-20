package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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
}

func (h *LMSAPI) getLesson(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
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
	id, _ := strconv.Atoi(ctx.Param("id"))
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
		return fail(err, "unable to list FAQs")
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
		return fail(err, "error checking enrollment")
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
		return fail(err, "failed to enroll")
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
		return fail(err, "unable to load profile")
	}

	return ctx.JSON(http.StatusOK, userWithData)
}

func (h *LMSAPI) submitReview(ctx echo.Context) error {
	courseID, _ := strconv.Atoi(ctx.Param("id"))
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
		return fail(err, "failed to submit review")
	}

	return ctx.JSON(http.StatusCreated, rev)
}

func (h *LMSAPI) listReviews(ctx echo.Context) error {
	courseID, _ := strconv.Atoi(ctx.Param("id"))
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
		return fail(err, "unable to list reviews")
	}

	return ctx.JSON(http.StatusOK, reviews)
}

func (h *LMSAPI) submitQuiz(ctx echo.Context) error {
	quizID, _ := strconv.Atoi(ctx.Param("id"))
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
		return fail(err, "failed to record quiz attempt")
	}

	return ctx.JSON(http.StatusCreated, attempt)
}
