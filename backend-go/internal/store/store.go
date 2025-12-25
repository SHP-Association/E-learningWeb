package store

import (
	"context"

	"github.com/SHP-Association/E-learningWeb/backend-go/internal/models"
)

// Store defines the interface for database operations
type Store interface {
	// User operations
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id uint) error
	ListUsers(ctx context.Context, opts ListOptions) ([]*models.User, int64, error)

	// Category operations
	CreateCategory(ctx context.Context, category *models.Category) error
	GetCategoryByID(ctx context.Context, id uint) (*models.Category, error)
	GetCategoryBySlug(ctx context.Context, slug string) (*models.Category, error)
	UpdateCategory(ctx context.Context, category *models.Category) error
	DeleteCategory(ctx context.Context, id uint) error
	ListCategories(ctx context.Context, opts ListOptions) ([]*models.Category, int64, error)

	// Course operations
	CreateCourse(ctx context.Context, course *models.Course) error
	GetCourseByID(ctx context.Context, id uint) (*models.Course, error)
	GetCourseBySlug(ctx context.Context, slug string) (*models.Course, error)
	UpdateCourse(ctx context.Context, course *models.Course) error
	DeleteCourse(ctx context.Context, id uint) error
	ListCourses(ctx context.Context, opts ListOptions) ([]*models.Course, int64, error)

	// Lesson operations
	CreateLesson(ctx context.Context, lesson *models.Lesson) error
	GetLessonByID(ctx context.Context, id uint) (*models.Lesson, error)
	UpdateLesson(ctx context.Context, lesson *models.Lesson) error
	DeleteLesson(ctx context.Context, id uint) error
	ListLessons(ctx context.Context, courseID uint, opts ListOptions) ([]*models.Lesson, int64, error)

	// Enrollment operations
	CreateEnrollment(ctx context.Context, enrollment *models.Enrollment) error
	GetEnrollmentByID(ctx context.Context, id uint) (*models.Enrollment, error)
	GetEnrollment(ctx context.Context, studentID, courseID uint) (*models.Enrollment, error)
	UpdateEnrollment(ctx context.Context, enrollment *models.Enrollment) error
	DeleteEnrollment(ctx context.Context, id uint) error
	ListEnrollments(ctx context.Context, studentID uint, opts ListOptions) ([]*models.Enrollment, int64, error)

	// Quiz operations
	CreateQuiz(ctx context.Context, quiz *models.Quiz) error
	GetQuizByID(ctx context.Context, id uint) (*models.Quiz, error)
	UpdateQuiz(ctx context.Context, quiz *models.Quiz) error
	DeleteQuiz(ctx context.Context, id uint) error
	ListQuizzes(ctx context.Context, courseID uint, opts ListOptions) ([]*models.Quiz, int64, error)

	// FAQ operations
	CreateFAQ(ctx context.Context, faq *models.FAQ) error
	GetFAQByID(ctx context.Context, id uint) (*models.FAQ, error)
	UpdateFAQ(ctx context.Context, faq *models.FAQ) error
	DeleteFAQ(ctx context.Context, id uint) error
	ListFAQs(ctx context.Context, opts ListOptions) ([]*models.FAQ, int64, error)

	// Review operations
	CreateReview(ctx context.Context, review *models.Review) error
	GetReviewByID(ctx context.Context, id uint) (*models.Review, error)
	UpdateReview(ctx context.Context, review *models.Review) error
	DeleteReview(ctx context.Context, id uint) error
	ListReviews(ctx context.Context, courseID uint, opts ListOptions) ([]*models.Review, int64, error)

	// Certificate operations
	CreateCertificate(ctx context.Context, certificate *models.Certificate) error
	GetCertificateByID(ctx context.Context, id uint) (*models.Certificate, error)
	GetCertificateByCertificateID(ctx context.Context, certificateID string) (*models.Certificate, error)
	ListCertificates(ctx context.Context, studentID uint, opts ListOptions) ([]*models.Certificate, int64, error)
}

// ListOptions contains pagination and filtering options
type ListOptions struct {
	Page     int
	PageSize int
	OrderBy  string
	Search   string
	Filters  map[string]interface{}
}
