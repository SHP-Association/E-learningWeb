package postgres

import (
	"context"

	"github.com/SHP-Association/E-learningWeb/backend-go/internal/models"
	"github.com/SHP-Association/E-learningWeb/backend-go/internal/store"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// PostgresStore implements the Store interface using PostgreSQL
type PostgresStore struct {
	db *gorm.DB
}

// New creates a new PostgreSQL store
func New(dsn string) (*PostgresStore, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

// GetDB returns the underlying GORM database instance
func (s *PostgresStore) GetDB() *gorm.DB {
	return s.db
}

// User operations
func (s *PostgresStore) CreateUser(ctx context.Context, user *models.User) error {
	return s.db.WithContext(ctx).Create(user).Error
}

func (s *PostgresStore) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := s.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *PostgresStore) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := s.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *PostgresStore) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := s.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *PostgresStore) UpdateUser(ctx context.Context, user *models.User) error {
	return s.db.WithContext(ctx).Save(user).Error
}

func (s *PostgresStore) DeleteUser(ctx context.Context, id uint) error {
	return s.db.WithContext(ctx).Delete(&models.User{}, id).Error
}

func (s *PostgresStore) ListUsers(ctx context.Context, opts store.ListOptions) ([]*models.User, int64, error) {
	var users []*models.User
	var total int64

	query := s.db.WithContext(ctx).Model(&models.User{})

	// Apply search
	if opts.Search != "" {
		query = query.Where("username ILIKE ? OR email ILIKE ? OR first_name ILIKE ? OR last_name ILIKE ?",
			"%"+opts.Search+"%", "%"+opts.Search+"%", "%"+opts.Search+"%", "%"+opts.Search+"%")
	}

	// Apply filters
	for key, value := range opts.Filters {
		query = query.Where(key+" = ?", value)
	}

	// Count total
	query.Count(&total)

	// Apply pagination
	if opts.PageSize > 0 {
		offset := (opts.Page - 1) * opts.PageSize
		query = query.Offset(offset).Limit(opts.PageSize)
	}

	// Apply ordering
	if opts.OrderBy != "" {
		query = query.Order(opts.OrderBy)
	} else {
		query = query.Order("created_at DESC")
	}

	err := query.Find(&users).Error
	return users, total, err
}

// Category operations
func (s *PostgresStore) CreateCategory(ctx context.Context, category *models.Category) error {
	return s.db.WithContext(ctx).Create(category).Error
}

func (s *PostgresStore) GetCategoryByID(ctx context.Context, id uint) (*models.Category, error) {
	var category models.Category
	err := s.db.WithContext(ctx).First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (s *PostgresStore) GetCategoryBySlug(ctx context.Context, slug string) (*models.Category, error) {
	var category models.Category
	err := s.db.WithContext(ctx).Where("slug = ?", slug).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (s *PostgresStore) UpdateCategory(ctx context.Context, category *models.Category) error {
	return s.db.WithContext(ctx).Save(category).Error
}

func (s *PostgresStore) DeleteCategory(ctx context.Context, id uint) error {
	return s.db.WithContext(ctx).Delete(&models.Category{}, id).Error
}

func (s *PostgresStore) ListCategories(ctx context.Context, opts store.ListOptions) ([]*models.Category, int64, error) {
	var categories []*models.Category
	var total int64

	query := s.db.WithContext(ctx).Model(&models.Category{})

	if opts.Search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+opts.Search+"%", "%"+opts.Search+"%")
	}

	query.Count(&total)

	if opts.PageSize > 0 {
		offset := (opts.Page - 1) * opts.PageSize
		query = query.Offset(offset).Limit(opts.PageSize)
	}

	if opts.OrderBy != "" {
		query = query.Order(opts.OrderBy)
	} else {
		query = query.Order("name ASC")
	}

	err := query.Find(&categories).Error
	return categories, total, err
}

// Course operations
func (s *PostgresStore) CreateCourse(ctx context.Context, course *models.Course) error {
	return s.db.WithContext(ctx).Create(course).Error
}

func (s *PostgresStore) GetCourseByID(ctx context.Context, id uint) (*models.Course, error) {
	var course models.Course
	err := s.db.WithContext(ctx).Preload("Instructor").Preload("Category").First(&course, id).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (s *PostgresStore) GetCourseBySlug(ctx context.Context, slug string) (*models.Course, error) {
	var course models.Course
	err := s.db.WithContext(ctx).Preload("Instructor").Preload("Category").Where("slug = ?", slug).First(&course).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (s *PostgresStore) UpdateCourse(ctx context.Context, course *models.Course) error {
	return s.db.WithContext(ctx).Save(course).Error
}

func (s *PostgresStore) DeleteCourse(ctx context.Context, id uint) error {
	return s.db.WithContext(ctx).Delete(&models.Course{}, id).Error
}

func (s *PostgresStore) ListCourses(ctx context.Context, opts store.ListOptions) ([]*models.Course, int64, error) {
	var courses []*models.Course
	var total int64

	query := s.db.WithContext(ctx).Model(&models.Course{}).Preload("Instructor").Preload("Category")

	if opts.Search != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+opts.Search+"%", "%"+opts.Search+"%")
	}

	for key, value := range opts.Filters {
		query = query.Where(key+" = ?", value)
	}

	query.Count(&total)

	if opts.PageSize > 0 {
		offset := (opts.Page - 1) * opts.PageSize
		query = query.Offset(offset).Limit(opts.PageSize)
	}

	if opts.OrderBy != "" {
		query = query.Order(opts.OrderBy)
	} else {
		query = query.Order("created_at DESC")
	}

	err := query.Find(&courses).Error
	return courses, total, err
}

// Lesson operations
func (s *PostgresStore) CreateLesson(ctx context.Context, lesson *models.Lesson) error {
	return s.db.WithContext(ctx).Create(lesson).Error
}

func (s *PostgresStore) GetLessonByID(ctx context.Context, id uint) (*models.Lesson, error) {
	var lesson models.Lesson
	err := s.db.WithContext(ctx).Preload("Course").First(&lesson, id).Error
	if err != nil {
		return nil, err
	}
	return &lesson, nil
}

func (s *PostgresStore) UpdateLesson(ctx context.Context, lesson *models.Lesson) error {
	return s.db.WithContext(ctx).Save(lesson).Error
}

func (s *PostgresStore) DeleteLesson(ctx context.Context, id uint) error {
	return s.db.WithContext(ctx).Delete(&models.Lesson{}, id).Error
}

func (s *PostgresStore) ListLessons(ctx context.Context, courseID uint, opts store.ListOptions) ([]*models.Lesson, int64, error) {
	var lessons []*models.Lesson
	var total int64

	query := s.db.WithContext(ctx).Model(&models.Lesson{}).Where("course_id = ?", courseID)

	query.Count(&total)

	if opts.PageSize > 0 {
		offset := (opts.Page - 1) * opts.PageSize
		query = query.Offset(offset).Limit(opts.PageSize)
	}

	query = query.Order("\"order\" ASC")

	err := query.Find(&lessons).Error
	return lessons, total, err
}

// Enrollment operations
func (s *PostgresStore) CreateEnrollment(ctx context.Context, enrollment *models.Enrollment) error {
	return s.db.WithContext(ctx).Create(enrollment).Error
}

func (s *PostgresStore) GetEnrollmentByID(ctx context.Context, id uint) (*models.Enrollment, error) {
	var enrollment models.Enrollment
	err := s.db.WithContext(ctx).Preload("Student").Preload("Course").First(&enrollment, id).Error
	if err != nil {
		return nil, err
	}
	return &enrollment, nil
}

func (s *PostgresStore) GetEnrollment(ctx context.Context, studentID, courseID uint) (*models.Enrollment, error) {
	var enrollment models.Enrollment
	err := s.db.WithContext(ctx).Where("student_id = ? AND course_id = ?", studentID, courseID).First(&enrollment).Error
	if err != nil {
		return nil, err
	}
	return &enrollment, nil
}

func (s *PostgresStore) UpdateEnrollment(ctx context.Context, enrollment *models.Enrollment) error {
	return s.db.WithContext(ctx).Save(enrollment).Error
}

func (s *PostgresStore) DeleteEnrollment(ctx context.Context, id uint) error {
	return s.db.WithContext(ctx).Delete(&models.Enrollment{}, id).Error
}

func (s *PostgresStore) ListEnrollments(ctx context.Context, studentID uint, opts store.ListOptions) ([]*models.Enrollment, int64, error) {
	var enrollments []*models.Enrollment
	var total int64

	query := s.db.WithContext(ctx).Model(&models.Enrollment{}).Preload("Course").Where("student_id = ?", studentID)

	query.Count(&total)

	if opts.PageSize > 0 {
		offset := (opts.Page - 1) * opts.PageSize
		query = query.Offset(offset).Limit(opts.PageSize)
	}

	query = query.Order("enrolled_at DESC")

	err := query.Find(&enrollments).Error
	return enrollments, total, err
}

// Quiz operations
func (s *PostgresStore) CreateQuiz(ctx context.Context, quiz *models.Quiz) error {
	return s.db.WithContext(ctx).Create(quiz).Error
}

func (s *PostgresStore) GetQuizByID(ctx context.Context, id uint) (*models.Quiz, error) {
	var quiz models.Quiz
	err := s.db.WithContext(ctx).Preload("Course").Preload("Lesson").First(&quiz, id).Error
	if err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (s *PostgresStore) UpdateQuiz(ctx context.Context, quiz *models.Quiz) error {
	return s.db.WithContext(ctx).Save(quiz).Error
}

func (s *PostgresStore) DeleteQuiz(ctx context.Context, id uint) error {
	return s.db.WithContext(ctx).Delete(&models.Quiz{}, id).Error
}

func (s *PostgresStore) ListQuizzes(ctx context.Context, courseID uint, opts store.ListOptions) ([]*models.Quiz, int64, error) {
	var quizzes []*models.Quiz
	var total int64

	query := s.db.WithContext(ctx).Model(&models.Quiz{}).Where("course_id = ?", courseID)

	query.Count(&total)

	if opts.PageSize > 0 {
		offset := (opts.Page - 1) * opts.PageSize
		query = query.Offset(offset).Limit(opts.PageSize)
	}

	err := query.Find(&quizzes).Error
	return quizzes, total, err
}

// FAQ operations
func (s *PostgresStore) CreateFAQ(ctx context.Context, faq *models.FAQ) error {
	return s.db.WithContext(ctx).Create(faq).Error
}

func (s *PostgresStore) GetFAQByID(ctx context.Context, id uint) (*models.FAQ, error) {
	var faq models.FAQ
	err := s.db.WithContext(ctx).First(&faq, id).Error
	if err != nil {
		return nil, err
	}
	return &faq, nil
}

func (s *PostgresStore) UpdateFAQ(ctx context.Context, faq *models.FAQ) error {
	return s.db.WithContext(ctx).Save(faq).Error
}

func (s *PostgresStore) DeleteFAQ(ctx context.Context, id uint) error {
	return s.db.WithContext(ctx).Delete(&models.FAQ{}, id).Error
}

func (s *PostgresStore) ListFAQs(ctx context.Context, opts store.ListOptions) ([]*models.FAQ, int64, error) {
	var faqs []*models.FAQ
	var total int64

	query := s.db.WithContext(ctx).Model(&models.FAQ{}).Where("is_active = ?", true)

	query.Count(&total)

	if opts.PageSize > 0 {
		offset := (opts.Page - 1) * opts.PageSize
		query = query.Offset(offset).Limit(opts.PageSize)
	}

	query = query.Order("\"order\" ASC")

	err := query.Find(&faqs).Error
	return faqs, total, err
}

// Review operations
func (s *PostgresStore) CreateReview(ctx context.Context, review *models.Review) error {
	return s.db.WithContext(ctx).Create(review).Error
}

func (s *PostgresStore) GetReviewByID(ctx context.Context, id uint) (*models.Review, error) {
	var review models.Review
	err := s.db.WithContext(ctx).Preload("Student").Preload("Course").First(&review, id).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (s *PostgresStore) UpdateReview(ctx context.Context, review *models.Review) error {
	return s.db.WithContext(ctx).Save(review).Error
}

func (s *PostgresStore) DeleteReview(ctx context.Context, id uint) error {
	return s.db.WithContext(ctx).Delete(&models.Review{}, id).Error
}

func (s *PostgresStore) ListReviews(ctx context.Context, courseID uint, opts store.ListOptions) ([]*models.Review, int64, error) {
	var reviews []*models.Review
	var total int64

	query := s.db.WithContext(ctx).Model(&models.Review{}).Preload("Student").Where("course_id = ?", courseID)

	query.Count(&total)

	if opts.PageSize > 0 {
		offset := (opts.Page - 1) * opts.PageSize
		query = query.Offset(offset).Limit(opts.PageSize)
	}

	query = query.Order("created_at DESC")

	err := query.Find(&reviews).Error
	return reviews, total, err
}

// Certificate operations
func (s *PostgresStore) CreateCertificate(ctx context.Context, certificate *models.Certificate) error {
	return s.db.WithContext(ctx).Create(certificate).Error
}

func (s *PostgresStore) GetCertificateByID(ctx context.Context, id uint) (*models.Certificate, error) {
	var certificate models.Certificate
	err := s.db.WithContext(ctx).Preload("Student").Preload("Course").First(&certificate, id).Error
	if err != nil {
		return nil, err
	}
	return &certificate, nil
}

func (s *PostgresStore) GetCertificateByCertificateID(ctx context.Context, certificateID string) (*models.Certificate, error) {
	var certificate models.Certificate
	err := s.db.WithContext(ctx).Preload("Student").Preload("Course").Where("certificate_id = ?", certificateID).First(&certificate).Error
	if err != nil {
		return nil, err
	}
	return &certificate, nil
}

func (s *PostgresStore) ListCertificates(ctx context.Context, studentID uint, opts store.ListOptions) ([]*models.Certificate, int64, error) {
	var certificates []*models.Certificate
	var total int64

	query := s.db.WithContext(ctx).Model(&models.Certificate{}).Preload("Course").Where("student_id = ?", studentID)

	query.Count(&total)

	if opts.PageSize > 0 {
		offset := (opts.Page - 1) * opts.PageSize
		query = query.Offset(offset).Limit(opts.PageSize)
	}

	query = query.Order("issued_at DESC")

	err := query.Find(&certificates).Error
	return certificates, total, err
}
