package services

import (
	"github.com/Waleed978/multi_tenent_backend/models"

	"gorm.io/gorm"
)

// StudentService defines the interface for student-related business logic.
type StudentService interface {
	CreateStudent(student *models.Student) error
	GetStudentByID(id uint) (*models.Student, error)
	GetAllStudents() ([]models.Student, error)
	UpdateStudent(student *models.Student) error
	DeleteStudent(id uint) error
}

// studentService implements the StudentService interface.
type studentService struct {
	db *gorm.DB
}

// NewStudentService creates a new instance of StudentService.
func NewStudentService(db *gorm.DB) StudentService {
	return &studentService{db: db}
}

// CreateStudent creates a new student record in the database.
func (s *studentService) CreateStudent(student *models.Student) error {
	return s.db.Create(student).Error
}

// GetStudentByID retrieves a student by their ID.
func (s *studentService) GetStudentByID(id uint) (*models.Student, error) {
	var student models.Student
	err := s.db.First(&student, id).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

// GetAllStudents retrieves all student records from the database.
func (s *studentService) GetAllStudents() ([]models.Student, error) {
	var students []models.Student
	err := s.db.Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}

// UpdateStudent updates an existing student record in the database.
func (s *studentService) UpdateStudent(student *models.Student) error {
	// Use Save for full update, or Updates for partial updates.
	// Save will update all fields, including zero values.
	return s.db.Save(student).Error
}

// DeleteStudent deletes a student record by their ID.
func (s *studentService) DeleteStudent(id uint) error {
	// Soft delete by default with gorm.Model
	return s.db.Delete(&models.Student{}, id).Error
}
