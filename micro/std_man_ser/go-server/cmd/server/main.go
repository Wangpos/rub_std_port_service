package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/yourorg/student-management-service-go/internal/handlers"
	"github.com/yourorg/student-management-service-go/internal/models"
)

func main() {
    // init DB
    db, err := gorm.Open(sqlite.Open("go_dev.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to open db: %v", err)
    }

    // Migrate schema
    if err := db.AutoMigrate(&models.College{}, &models.Department{}, &models.Program{}, &models.Student{}); err != nil {
        log.Fatalf("migrate failed: %v", err)
    }

    // Seed minimal data if empty
    seed(db)

    // Setup Gin
    r := gin.Default()

    api := r.Group("/api")
    {
        h := handlers.NewHandler(db)
        api.GET("/students", h.ListStudents)
        api.POST("/students", h.CreateStudent)
        api.GET("/students/:id", h.GetStudent)
        api.PUT("/students/:id", h.UpdateStudent)
        api.DELETE("/students/:id", h.DeleteStudent)

        api.GET("/colleges", h.ListColleges)
    }

    // simple health
    r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"ok": true}) })

    port := 4100
    log.Printf("Starting Go Student Management Service on :%d", port)
    if err := r.Run(":" + strconv.Itoa(port)); err != nil {
        log.Fatalf("server failed: %v", err)
    }
}

func seed(db *gorm.DB) {
    var count int64
    db.Model(&models.College{}).Count(&count)
    if count > 0 {
        return
    }

    c := models.College{Name: "Example College"}
    db.Create(&c)

    d := models.Department{Name: "Computer Science", CollegeID: c.ID}
    db.Create(&d)

    p := models.Program{Name: "BSc Computer Science", DepartmentID: d.ID, CollegeID: c.ID}
    db.Create(&p)

    students := []models.Student{
        {StudentNumber: "G2025001", FirstName: "Charlie", LastName: "Clark", Email: "charlie@example.edu", EnrollmentYear: 2025, ProgramID: p.ID},
        {StudentNumber: "G2025002", FirstName: "Dana", LastName: "Dawson", Email: "dana@example.edu", EnrollmentYear: 2025, ProgramID: p.ID},
    }
    for _, s := range students {
        db.Create(&s)
    }
}
