//package main

// import (
// 	"encoding/json"
// 	"net/http"
// 	"os"

// 	"github.com/go-chi/chi"
// 	"github.com/go-chi/chi/middleware"
// 	"github.com/go-playground/validator/v10"
// 	"github.com/joho/godotenv"
// )

// // Structs with validation tags
// type StudentCreateRequest struct {
// 	ID   string `json:"id" validate:"required"`
// 	Name string `json:"name" validate:"required,min=2"`
// }

// type StudentUpdateRequest struct {
// 	Name string `json:"name" validate:"required,min=2"`
// }

// var validate = validator.New()

// // Error response helper
// func respondWithError(w http.ResponseWriter, code int, message string) {
// 	w.WriteHeader(code)
// 	json.NewEncoder(w).Encode(map[string]string{"error": message})
// }

// func main() {
// 	godotenv.Load()
// 	port := os.Getenv("PORT")

// 	router := chi.NewRouter()
// 	router.Use(middleware.Logger)

// 	Students := map[string]string{
// 		"1": "waleed",
// 		"2": "Jane",
// 		"3": "Jim",
// 		"4": "Jill",
// 		"5": "Jack",
// 		"6": "Jill",
// 	}

// 	// Read student
// 	router.Get("/student/{id}", func(w http.ResponseWriter, r *http.Request) {
// 		id := chi.URLParam(r, "id")
// 		student, ok := Students[id]
// 		if !ok {
// 			respondWithError(w, http.StatusNotFound, "Student not found")
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(map[string]string{"id": id, "name": student})
// 	})

// 	// Create student with validation
// 	router.Post("/student", func(w http.ResponseWriter, r *http.Request) {
// 		var req StudentCreateRequest
// 		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 			respondWithError(w, http.StatusBadRequest, "Invalid JSON body")
// 			return
// 		}
// 		if err := validate.Struct(req); err != nil {
// 			respondWithError(w, http.StatusBadRequest, "Validation failed: "+err.Error())
// 			return
// 		}
// 		if _, exists := Students[req.ID]; exists {
// 			respondWithError(w, http.StatusConflict, "Student already exists")
// 			return
// 		}
// 		Students[req.ID] = req.Name
// 		w.WriteHeader(http.StatusCreated)
// 		json.NewEncoder(w).Encode(map[string]string{"message": "Student created"})
// 	})

// 	// Update student with validation
// 	router.Put("/student/{id}", func(w http.ResponseWriter, r *http.Request) {
// 		id := chi.URLParam(r, "id")
// 		var req StudentUpdateRequest
// 		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 			respondWithError(w, http.StatusBadRequest, "Invalid JSON body")
// 			return
// 		}
// 		if err := validate.Struct(req); err != nil {
// 			respondWithError(w, http.StatusBadRequest, "Validation failed: "+err.Error())
// 			return
// 		}
// 		if _, exists := Students[id]; !exists {
// 			respondWithError(w, http.StatusNotFound, "Student not found")
// 			return
// 		}
// 		Students[id] = req.Name
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(map[string]string{"message": "Student updated"})
// 	})

// 	// Delete student
// 	router.Delete("/student/{id}", func(w http.ResponseWriter, r *http.Request) {
// 		id := chi.URLParam(r, "id")
// 		if _, exists := Students[id]; !exists {
// 			respondWithError(w, http.StatusNotFound, "Student not found")
// 			return
// 		}
// 		delete(Students, id)
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(map[string]string{"message": "Student deleted"})
// 	})

//		http.ListenAndServe(":"+port, router)
//	}
package main

import (
	"fmt"
	"log"

	"github.com/Waleed978/multi_tenent_backend/config"
	"github.com/Waleed978/multi_tenent_backend/database"
	"github.com/Waleed978/multi_tenent_backend/handlers"
	"github.com/Waleed978/multi_tenent_backend/routes"
	"github.com/Waleed978/multi_tenent_backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	fmt.Println("Configuration loaded successfully.")

	// 2. Initialize Database Connection
	database.InitDB(cfg)
	fmt.Println("Database initialized.")

	// 3. Initialize Services
	studentService := services.NewStudentService(database.DB)
	fmt.Println("Student service initialized.")

	// 4. Initialize Handlers
	studentHandler := handlers.NewStudentHandler(studentService)
	fmt.Println("Student handler initialized.")

	// 5. Setup Gin Router
	router := gin.Default()
	fmt.Println("Gin router initialized.")

	// 6. Setup Routes
	routes.SetupStudentRoutes(router, studentHandler)
	fmt.Println("Routes set up.")

	// 7. Start the Server
	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
