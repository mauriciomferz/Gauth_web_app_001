package database

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gauth-web-app/backend/internal/config"
	"gauth-web-app/backend/internal/models"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	if cfg.Environment == "production" {
		gormConfig.Logger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully")

	return db, nil
}

func Migrate(db *gorm.DB) error {
	log.Println("Starting database migration...")

	err := db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Policy{},
		&models.Session{},
		&models.AuditLog{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database migration completed successfully")
	return nil
}

func Seed(db *gorm.DB) error {
	log.Println("Seeding initial data...")

	// Create default admin role
	adminRole := &models.Role{
		Name:        "admin",
		Description: "System Administrator",
	}
	adminRole.SetPermissions([]string{
		"user:create", "user:read", "user:update", "user:delete",
		"role:create", "role:read", "role:update", "role:delete",
		"policy:create", "policy:read", "policy:update", "policy:delete",
		"audit:read",
	})

	// Create default user role
	userRole := &models.Role{
		Name:        "user",
		Description: "Regular User",
	}
	userRole.SetPermissions([]string{
		"profile:read", "profile:update",
	})

	// Check if roles already exist
	var existingAdminRole models.Role
	if err := db.Where("name = ?", "admin").First(&existingAdminRole).Error; err == gorm.ErrRecordNotFound {
		if err := db.Create(adminRole).Error; err != nil {
			return fmt.Errorf("failed to create admin role: %w", err)
		}
		log.Println("Admin role created successfully")
	}

	var existingUserRole models.Role
	if err := db.Where("name = ?", "user").First(&existingUserRole).Error; err == gorm.ErrRecordNotFound {
		if err := db.Create(userRole).Error; err != nil {
			return fmt.Errorf("failed to create user role: %w", err)
		}
		log.Println("User role created successfully")
	}

	// Create test admin user
	var existingAdmin models.User
	if err := db.Where("username = ?", "admin").First(&existingAdmin).Error; err == gorm.ErrRecordNotFound {
		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}

		// Get admin role
		var adminRoleFromDB models.Role
		if err := db.Where("name = ?", "admin").First(&adminRoleFromDB).Error; err != nil {
			return fmt.Errorf("failed to find admin role: %w", err)
		}

		adminUser := &models.User{
			Username:   "admin",
			Email:      "admin@gauth.local",
			Password:   string(hashedPassword),
			FirstName:  "Admin",
			LastName:   "User",
			IsActive:   true,
			IsVerified: true,
			Roles:      []models.Role{adminRoleFromDB},
		}

		if err := db.Create(adminUser).Error; err != nil {
			return fmt.Errorf("failed to create admin user: %w", err)
		}
		log.Println("Admin user created successfully (username: admin, password: password)")
	}

	log.Println("Database seeding completed successfully")
	return nil
}
