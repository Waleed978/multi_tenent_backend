package handlers

import (
	"net/http"
	"strconv"

	"github.com/Waleed978/multi_tenent_backend/models"
	"github.com/Waleed978/multi_tenent_backend/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// StudentHandler handles HTTP requests related to students.
type StudentHandler struct {
	studentService services.StudentService
}

// NewStudentHandler creates a new instance of StudentHandler.
func NewStudentHandler(service services.StudentService) *StudentHandler {
	return &StudentHandler{studentService: service}
}

// CreateStudent handles the creation of a new student.
// @Summary Create a new student
// @Description Creates a new student record in the database
// @Accept json
// @Produce json
// @Param student body models.Student true "Student object to be created"
// @Success 201 {object} models.Student
// @Failure 400 {object} gin.H "Bad Request"
// @Failure 500 {object} gin.H "Internal Server Error"
// @Router /students [post]
func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var student models.Student
	// Bind JSON request body to the Student struct and perform validation.
	if err := c.ShouldBindJSON(&student); err != nil {
		// Handle validation errors
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errors := make(map[string]string)
			for _, fieldErr := range validationErrors {
				errors[fieldErr.Field()] = fieldErr.Tag()
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": errors})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service layer to create the student.
	if err := h.studentService.CreateStudent(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, student)
}

// GetStudent handles retrieving a student by ID.
// @Summary Get a student by ID
// @Description Retrieves a single student record by their ID
// @Produce json
// @Param id path int true "Student ID"
// @Success 200 {object} models.Student
// @Failure 400 {object} gin.H "Bad Request"
// @Failure 404 {object} gin.H "Not Found"
// @Failure 500 {object} gin.H "Internal Server Error"
// @Router /students/{id} [get]
func (h *StudentHandler) GetStudent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	student, err := h.studentService.GetStudentByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve student", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}

// GetStudents handles retrieving all students.
// @Summary Get all students
// @Description Retrieves all student records
// @Produce json
// @Success 200 {array} models.Student
// @Failure 500 {object} gin.H "Internal Server Error"
// @Router /students [get]
func (h *StudentHandler) GetStudents(c *gin.Context) {
	students, err := h.studentService.GetAllStudents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve students", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

// UpdateStudent handles updating an existing student.
// @Summary Update an existing student
// @Description Updates an existing student record by ID
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Param student body models.Student true "Updated student object"
// @Success 200 {object} models.Student
// @Failure 400 {object} gin.H "Bad Request"
// @Failure 404 {object} gin.H "Not Found"
// @Failure 500 {object} gin.H "Internal Server Error"
// @Router /students/{id} [put]
func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		// Handle validation errors
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errors := make(map[string]string)
			for _, fieldErr := range validationErrors {
				errors[fieldErr.Field()] = fieldErr.Tag()
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": errors})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the ID from the URL parameter to ensure the correct student is updated.
	student.ID = uint(id)

	// Check if the student exists before attempting to update
	existingStudent, err := h.studentService.GetStudentByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing student", "details": err.Error()})
		return
	}

	// Preserve CreatedAt from existing student if not provided in update payload
	if student.CreatedAt.IsZero() {
		student.CreatedAt = existingStudent.CreatedAt
	}
	// Preserve EnrolledAt from existing student if not provided in update payload
	if student.EnrolledAt.IsZero() {
		student.EnrolledAt = existingStudent.EnrolledAt
	}

	if err := h.studentService.UpdateStudent(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}

// DeleteStudent handles deleting a student by ID.
// @Summary Delete a student by ID
// @Description Deletes a student record by their ID
// @Produce json
// @Param id path int true "Student ID"
// @Success 204 "No Content"
// @Failure 400 {object} gin.H "Bad Request"
// @Failure 404 {object} gin.H "Not Found"
// @Failure 500 {object} gin.H "Internal Server Error"
// @Router /students/{id} [delete]
func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	// Check if the student exists before attempting to delete
	_, err = h.studentService.GetStudentByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check existing student", "details": err.Error()})
		return
	}

	if err := h.studentService.DeleteStudent(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student", "details": err.Error()})
		return
	}

	c.Status(http.StatusNoContent) // 204 No Content for successful deletion
}
