package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Username    string     `gorm:"uniqueIndex;not null" json:"username" binding:"required"`
	Email       string     `gorm:"uniqueIndex;not null" json:"email" binding:"required,email"`
	Password    string     `gorm:"not null" json:"-"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Avatar      string     `json:"avatar"`
	IsActive    bool       `gorm:"default:true" json:"is_active"`
	IsVerified  bool       `gorm:"default:false" json:"is_verified"`
	LastLoginAt *time.Time `json:"last_login_at"`

	// Relationships
	Roles    []Role    `gorm:"many2many:user_roles;" json:"roles"`
	Sessions []Session `gorm:"foreignKey:UserID" json:"-"`
}

// BeforeCreate generates UUID for new users
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

// Role represents a role in the system
type Role struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name        string `gorm:"uniqueIndex;not null" json:"name" binding:"required"`
	Description string `json:"description"`
	Permissions string `gorm:"type:text" json:"-"`
	IsActive    bool   `gorm:"default:true" json:"is_active"`

	// Relationships
	Users []User `gorm:"many2many:user_roles;" json:"-"`
}

// BeforeCreate generates UUID for new roles
func (r *Role) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}

// GetPermissions returns permissions as string slice
func (r *Role) GetPermissions() []string {
	if r.Permissions == "" {
		return []string{}
	}
	var permissions []string
	json.Unmarshal([]byte(r.Permissions), &permissions)
	return permissions
}

// SetPermissions sets permissions from string slice
func (r *Role) SetPermissions(permissions []string) {
	if permissions == nil {
		permissions = []string{}
	}
	data, _ := json.Marshal(permissions)
	r.Permissions = string(data)
}

// Policy represents an authorization policy
type Policy struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Name        string            `gorm:"uniqueIndex;not null" json:"name" binding:"required"`
	Description string            `json:"description"`
	Resource    string            `gorm:"not null" json:"resource" binding:"required"`
	Action      string            `gorm:"not null" json:"action" binding:"required"`
	Effect      string            `gorm:"not null;default:'allow'" json:"effect" binding:"required,oneof=allow deny"`
	Conditions  map[string]string `gorm:"type:jsonb" json:"conditions"`
	IsActive    bool              `gorm:"default:true" json:"is_active"`
}

// BeforeCreate generates UUID for new policies
func (p *Policy) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

// Session represents a user session
type Session struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Token     string    `gorm:"uniqueIndex;not null" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	IPAddress string    `json:"ip_address"`
	UserAgent string    `json:"user_agent"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`

	// Relationships
	User User `gorm:"foreignKey:UserID" json:"user"`
}

// BeforeCreate generates UUID for new sessions
func (s *Session) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}

// IsExpired checks if the session is expired
func (s *Session) IsExpired() bool {
	return time.Now().After(s.ExpiresAt)
}

// AuditLog represents an audit log entry
type AuditLog struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`

	UserID     *uuid.UUID             `gorm:"type:uuid" json:"user_id"`
	Action     string                 `gorm:"not null" json:"action"`
	Resource   string                 `gorm:"not null" json:"resource"`
	ResourceID *uuid.UUID             `gorm:"type:uuid" json:"resource_id"`
	Details    map[string]interface{} `gorm:"type:jsonb" json:"details"`
	IPAddress  string                 `json:"ip_address"`
	UserAgent  string                 `json:"user_agent"`
	Success    bool                   `gorm:"default:true" json:"success"`

	// Relationships
	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// BeforeCreate generates UUID for new audit logs
func (a *AuditLog) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
}

// CreateUserRequest represents the request to create a new user
type CreateUserRequest struct {
	Username  string      `json:"username" binding:"required,min=3,max=50"`
	Email     string      `json:"email" binding:"required,email"`
	Password  string      `json:"password" binding:"required,min=6"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	RoleIDs   []uuid.UUID `json:"role_ids"`
}

// UpdateUserRequest represents the request to update a user
type UpdateUserRequest struct {
	Username  *string     `json:"username,omitempty" binding:"omitempty,min=3,max=50"`
	Email     *string     `json:"email,omitempty" binding:"omitempty,email"`
	FirstName *string     `json:"first_name,omitempty"`
	LastName  *string     `json:"last_name,omitempty"`
	Avatar    *string     `json:"avatar,omitempty"`
	IsActive  *bool       `json:"is_active,omitempty"`
	RoleIDs   []uuid.UUID `json:"role_ids,omitempty"`
}

// LoginRequest represents the login request
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	User         *User  `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

// RefreshTokenRequest represents the refresh token request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// ChangePasswordRequest represents the change password request
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=6"`
}
