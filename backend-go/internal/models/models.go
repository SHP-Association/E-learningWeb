package models

import (
	"time"

	"gorm.io/gorm"
)

// Category represents the course category
type Category struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"unique;not null" json:"name"`
	Slug        string         `gorm:"unique;not null" json:"slug"`
	Description string         `json:"description"`
	Icon        *string        `json:"icon"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Category) TableName() string {
	return "Category_category"
}

// Course represents a course
type Course struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	Title            string         `gorm:"not null" json:"title"`
	Slug             string         `gorm:"unique;not null" json:"slug"`
	ShortDescription string         `json:"short_description"`
	Description      string         `gorm:"type:text" json:"description"`
	WhatYouWillLearn *string        `gorm:"type:text" json:"what_you_will_learn"`
	Requirements     *string        `gorm:"type:text" json:"requirements"`
	TargetAudience   *string        `gorm:"type:text" json:"target_audience"`
	InstructorID     *uint          `json:"instructor_id"`
	Instructor       *User          `gorm:"foreignKey:InstructorID" json:"instructor,omitempty"`
	CategoryID       *uint          `json:"category_id"`
	Category         *Category      `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Thumbnail        *string        `json:"thumbnail"`
	PromoVideoURL    *string        `json:"promo_video_url"`
	Price            float64        `gorm:"default:0.00" json:"price"`
	IsFree           bool           `gorm:"default:false" json:"is_free"`
	IsPublished      bool           `gorm:"default:false" json:"is_published"`
	Level            string         `gorm:"default:'beginner'" json:"level"` // beginner, intermediate, advanced
	Duration         *string        `json:"duration"`
	TotalLectures    uint           `gorm:"default:0" json:"total_lectures"`
	AverageRating    float64        `gorm:"default:0.00" json:"average_rating"`
	NumberOfReviews  uint           `gorm:"default:0" json:"number_of_reviews"`
	CreatedAt        time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Course) TableName() string {
	return "courses_course"
}

// Lesson represents a course lesson
type Lesson struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CourseID  uint           `gorm:"not null" json:"course_id"`
	Course    *Course        `gorm:"foreignKey:CourseID" json:"course,omitempty"`
	Title     string         `gorm:"not null" json:"title"`
	Slug      string         `gorm:"not null" json:"slug"`
	Content   *string        `gorm:"type:text" json:"content"`
	VideoURL  *string        `json:"video_url"`
	Order     uint           `gorm:"not null" json:"order"`
	IsPreview bool           `gorm:"default:false" json:"is_preview"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Lesson) TableName() string {
	return "Lesson_lesson"
}

// Enrollment represents a student's enrollment in a course
type Enrollment struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	StudentID   uint           `gorm:"not null" json:"student_id"`
	Student     *User          `gorm:"foreignKey:StudentID" json:"student,omitempty"`
	CourseID    uint           `gorm:"not null" json:"course_id"`
	Course      *Course        `gorm:"foreignKey:CourseID" json:"course,omitempty"`
	EnrolledAt  time.Time      `gorm:"autoCreateTime" json:"enrolled_at"`
	CompletedAt *time.Time     `json:"completed_at"`
	Progress    float64        `gorm:"default:0.0" json:"progress"`
	IsCompleted bool           `gorm:"default:false" json:"is_completed"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Enrollment) TableName() string {
	return "Enrollment_enrollment"
}

// Quiz represents a quiz
type Quiz struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CourseID     uint           `gorm:"not null" json:"course_id"`
	Course       *Course        `gorm:"foreignKey:CourseID" json:"course,omitempty"`
	LessonID     *uint          `json:"lesson_id"`
	Lesson       *Lesson        `gorm:"foreignKey:LessonID" json:"lesson,omitempty"`
	Title        string         `gorm:"not null" json:"title"`
	Description  *string        `gorm:"type:text" json:"description"`
	PassingScore float64        `gorm:"default:70.0" json:"passing_score"`
	TimeLimit    *int           `json:"time_limit"` // in minutes
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Quiz) TableName() string {
	return "Quiz_quiz"
}

// FAQ represents frequently asked questions
type FAQ struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Question  string         `gorm:"not null" json:"question"`
	Answer    string         `gorm:"type:text;not null" json:"answer"`
	Order     uint           `gorm:"default:0" json:"order"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (FAQ) TableName() string {
	return "FAQ_faq"
}

// Review represents a course review
type Review struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CourseID  uint           `gorm:"not null" json:"course_id"`
	Course    *Course        `gorm:"foreignKey:CourseID" json:"course,omitempty"`
	StudentID uint           `gorm:"not null" json:"student_id"`
	Student   *User          `gorm:"foreignKey:StudentID" json:"student,omitempty"`
	Rating    float64        `gorm:"not null" json:"rating"`
	Comment   *string        `gorm:"type:text" json:"comment"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Review) TableName() string {
	return "Review_review"
}

// Certificate represents a course completion certificate
type Certificate struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	StudentID     uint           `gorm:"not null" json:"student_id"`
	Student       *User         `gorm:"foreignKey:StudentID" json:"student,omitempty"`
	CourseID      uint           `gorm:"not null" json:"course_id"`
	Course        *Course        `gorm:"foreignKey:CourseID" json:"course,omitempty"`
	CertificateID string         `gorm:"unique;not null" json:"certificate_id"`
	IssuedAt      time.Time      `gorm:"autoCreateTime" json:"issued_at"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Certificate) TableName() string {
	return "Certificate_certificate"
}
