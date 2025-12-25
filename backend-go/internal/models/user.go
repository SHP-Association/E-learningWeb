package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents the CustomUser model from Django
type User struct {
	ID                   uint           `gorm:"primaryKey" json:"id"`
	Username             string         `gorm:"uniqueIndex;not null" json:"username"`
	Email                string         `gorm:"uniqueIndex;not null" json:"email"`
	Password             string         `gorm:"not null" json:"-"` // Never expose password in JSON
	FirstName            string         `json:"first_name"`
	LastName             string         `json:"last_name"`
	Role                 string         `gorm:"default:'student'" json:"role"` // student, instructor, admin
	Bio                  string         `json:"bio"`
	ProfilePicture       *string        `json:"profile_picture"`
	DateOfBirth          *time.Time     `json:"date_of_birth"`
	Gender               string         `json:"gender"`
	ContactNumber        string         `json:"contact_number"`
	Address              string         `json:"address"`
	Country              string         `json:"country"`
	IsEmailVerified      bool           `gorm:"default:false" json:"is_email_verified"`
	HighestQualification string         `json:"highest_qualification"`
	Institution          string         `json:"institution"`
	Skills               string         `json:"skills"`
	LinkedinProfile      string         `json:"linkedin_profile"`
	GithubProfile        string         `json:"github_profile"`
	QuizScores           *string        `gorm:"type:jsonb" json:"quiz_scores"` // JSONB field
	InstructorRating     float64        `gorm:"default:0.0" json:"instructor_rating"`
	TotalReviews         uint           `gorm:"default:0" json:"total_reviews"`
	IsActive             bool           `gorm:"default:true" json:"is_active"`
	IsStaff              bool           `gorm:"default:false" json:"is_staff"`
	IsSuperuser          bool           `gorm:"default:false" json:"is_superuser"`
	LastActivity         *time.Time     `json:"last_activity"`
	LoginIP              *string        `json:"login_ip"`
	TwoFactorEnabled     bool           `gorm:"default:false" json:"two_factor_enabled"`
	DateJoined           time.Time      `gorm:"autoCreateTime" json:"date_joined"`
	LastLogin            *time.Time     `json:"last_login"`
	CreatedAt            time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for GORM (matches Django's table)
func (User) TableName() string {
	return "Account_customuser"
}

// SetPassword hashes and sets the user's password
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword verifies if the provided password matches the user's password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// BeforeSave GORM hook to auto-assign staff status based on role
func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.Role == "instructor" || u.Role == "admin" {
		u.IsStaff = true
	} else {
		u.IsStaff = false
	}
	return nil
}

// UserResponse is the response structure (excludes sensitive fields)
type UserResponse struct {
	ID                   uint       `json:"id"`
	Username             string     `json:"username"`
	Email                string     `json:"email"`
	FirstName            string     `json:"first_name"`
	LastName             string     `json:"last_name"`
	Role                 string     `json:"role"`
	Bio                  string     `json:"bio"`
	ProfilePicture       *string    `json:"profile_picture"`
	DateOfBirth          *time.Time `json:"date_of_birth"`
	Gender               string     `json:"gender"`
	ContactNumber        string     `json:"contact_number"`
	Address              string     `json:"address"`
	Country              string     `json:"country"`
	IsEmailVerified      bool       `json:"is_email_verified"`
	HighestQualification string     `json:"highest_qualification"`
	Institution          string     `json:"institution"`
	Skills               string     `json:"skills"`
	LinkedinProfile      string     `json:"linkedin_profile"`
	GithubProfile        string     `json:"github_profile"`
	InstructorRating     float64    `json:"instructor_rating"`
	TotalReviews         uint       `json:"total_reviews"`
	IsActive             bool       `json:"is_active"`
	DateJoined           time.Time  `json:"date_joined"`
	LastLogin            *time.Time `json:"last_login"`
}

// ToResponse converts User to UserResponse
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:                   u.ID,
		Username:             u.Username,
		Email:                u.Email,
		FirstName:            u.FirstName,
		LastName:             u.LastName,
		Role:                 u.Role,
		Bio:                  u.Bio,
		ProfilePicture:       u.ProfilePicture,
		DateOfBirth:          u.DateOfBirth,
		Gender:               u.Gender,
		ContactNumber:        u.ContactNumber,
		Address:              u.Address,
		Country:              u.Country,
		IsEmailVerified:      u.IsEmailVerified,
		HighestQualification: u.HighestQualification,
		Institution:          u.Institution,
		Skills:               u.Skills,
		LinkedinProfile:      u.LinkedinProfile,
		GithubProfile:        u.GithubProfile,
		InstructorRating:     u.InstructorRating,
		TotalReviews:         u.TotalReviews,
		IsActive:             u.IsActive,
		DateJoined:           u.DateJoined,
		LastLogin:            u.LastLogin,
	}
}
