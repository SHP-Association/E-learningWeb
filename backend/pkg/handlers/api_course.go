package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/SHP-Association/E-learningWeb/backend/ent"
	"github.com/SHP-Association/E-learningWeb/backend/ent/category"
	"github.com/SHP-Association/E-learningWeb/backend/ent/course"
	"github.com/SHP-Association/E-learningWeb/backend/pkg/services"
)

type CourseAPI struct {
	orm *ent.Client
}

func init() {
	Register(new(CourseAPI))
}

func (h *CourseAPI) Init(c *services.Container) error {
	h.orm = c.ORM
	return nil
}

func (h *CourseAPI) Routes(g *echo.Group) {
	api := g.Group("/api")

	// Categories
	api.GET("/categories", h.listCategories)
	api.GET("/categories/:id", h.getCategory)

	// Courses
	api.GET("/courses", h.listCourses)
	api.GET("/courses/:id", h.getCourse)
}

func (h *CourseAPI) listCategories(ctx echo.Context) error {
	categories, err := h.orm.Category.
		Query().
		Order(ent.Asc(category.FieldName)).
		All(ctx.Request().Context())

	if err != nil {
		return fail(err, "unable to list categories")
	}

	return ctx.JSON(http.StatusOK, categories)
}

func (h *CourseAPI) getCategory(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	cat, err := h.orm.Category.
		Get(ctx.Request().Context(), id)

	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "category not found"})
	}

	return ctx.JSON(http.StatusOK, cat)
}

func (h *CourseAPI) listCourses(ctx echo.Context) error {
	courses, err := h.orm.Course.
		Query().
		WithInstructor().
		WithCategory().
		Order(ent.Desc(course.FieldCreatedAt)).
		All(ctx.Request().Context())

	if err != nil {
		return fail(err, "unable to list courses")
	}

	return ctx.JSON(http.StatusOK, courses)
}

func (h *CourseAPI) getCourse(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	c, err := h.orm.Course.
		Query().
		Where(course.ID(id)).
		WithInstructor().
		WithCategory().
		WithLessons().
		Only(ctx.Request().Context())

	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "course not found"})
	}

	return ctx.JSON(http.StatusOK, c)
}
