package handlers
package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "github.com/yourorg/student-management-service-go/internal/models"
)

type Handler struct{
    db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler { return &Handler{db: db} }

func (h *Handler) CreateStudent(c *gin.Context) {
    var input models.Student
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": gin.H{"code": "validation.failed", "message": err.Error()}})
        return
    }

    // check conflicts
    var ex models.Student
    if err := h.db.Where("email = ? OR student_number = ?", input.Email, input.StudentNumber).First(&ex).Error; err == nil {
        c.JSON(http.StatusConflict, gin.H{"ok": false, "error": gin.H{"code": "student.duplicate", "message": "Student with same email or student number exists"}})
        return
    }

    if err := h.db.Create(&input).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": gin.H{"code": "internal_error", "message": err.Error()}})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"ok": true, "data": input})
}

func (h *Handler) ListStudents(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
    search := c.Query("search")
    programId := c.Query("programId")
    enrollmentYear := c.Query("enrollmentYear")

    var students []models.Student
    query := h.db.Preload("Program")

    if search != "" {
        like := "%" + search + "%"
        query = query.Where("first_name LIKE ? OR last_name LIKE ? OR email LIKE ? OR student_number LIKE ?", like, like, like, like)
    }
    if programId != "" {
        if pid, err := strconv.Atoi(programId); err == nil {
            query = query.Where("program_id = ?", pid)
        }
    }
    if enrollmentYear != "" {
        if ey, err := strconv.Atoi(enrollmentYear); err == nil {
            query = query.Where("enrollment_year = ?", ey)
        }
    }

    var total int64
    query.Model(&models.Student{}).Count(&total)

    offset := (page - 1) * limit
    if err := query.Offset(offset).Limit(limit).Find(&students).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": gin.H{"code": "internal_error", "message": err.Error()}})
        return
    }

    c.JSON(http.StatusOK, gin.H{"ok": true, "data": gin.H{"items": students}, "meta": gin.H{"total": total, "page": page, "limit": limit}})
}

func (h *Handler) GetStudent(c *gin.Context) {
    id := c.Param("id")
    var s models.Student
    if err := h.db.Preload("Program").First(&s, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"ok": false, "error": gin.H{"code": "student.not_found", "message": "Student not found"}})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": gin.H{"code": "internal_error", "message": err.Error()}})
        return
    package handlers

    import (
        "net/http"
        "strconv"

        "github.com/gin-gonic/gin"
        "gorm.io/gorm"

        "github.com/yourorg/student-management-service-go/internal/models"
    )

    type Handler struct{
        db *gorm.DB
    }

    func NewHandler(db *gorm.DB) *Handler { return &Handler{db: db} }

    func (h *Handler) CreateStudent(c *gin.Context) {
        var input models.Student
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": gin.H{"code": "validation.failed", "message": err.Error()}})
            return
        }

        // check conflicts
        var ex models.Student
        if err := h.db.Where("email = ? OR student_number = ?", input.Email, input.StudentNumber).First(&ex).Error; err == nil {
            c.JSON(http.StatusConflict, gin.H{"ok": false, "error": gin.H{"code": "student.duplicate", "message": "Student with same email or student number exists"}})
            return
        }

        if err := h.db.Create(&input).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": gin.H{"code": "internal_error", "message": err.Error()}})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"ok": true, "data": input})
    }

    func (h *Handler) ListStudents(c *gin.Context) {
        page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
        limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
        search := c.Query("search")
        programId := c.Query("programId")
        enrollmentYear := c.Query("enrollmentYear")

        var students []models.Student
        query := h.db.Preload("Program")

        if search != "" {
            like := "%" + search + "%"
            query = query.Where("first_name LIKE ? OR last_name LIKE ? OR email LIKE ? OR student_number LIKE ?", like, like, like, like)
        }
        if programId != "" {
            if pid, err := strconv.Atoi(programId); err == nil {
                query = query.Where("program_id = ?", pid)
            }
        }
        if enrollmentYear != "" {
            if ey, err := strconv.Atoi(enrollmentYear); err == nil {
                query = query.Where("enrollment_year = ?", ey)
            }
        }

        var total int64
        query.Model(&models.Student{}).Count(&total)

        offset := (page - 1) * limit
        if err := query.Offset(offset).Limit(limit).Find(&students).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": gin.H{"code": "internal_error", "message": err.Error()}})
            return
        }

        c.JSON(http.StatusOK, gin.H{"ok": true, "data": gin.H{"items": students}, "meta": gin.H{"total": total, "page": page, "limit": limit}})
    }

    func (h *Handler) GetStudent(c *gin.Context) {
        id := c.Param("id")
        var s models.Student
        if err := h.db.Preload("Program").First(&s, id).Error; err != nil {
            if err == gorm.ErrRecordNotFound {
                c.JSON(http.StatusNotFound, gin.H{"ok": false, "error": gin.H{"code": "student.not_found", "message": "Student not found"}})
                return
            }
            c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": gin.H{"code": "internal_error", "message": err.Error()}})
            return
        package handlers

        import (
            "net/http"
            "strconv"

            "github.com/gin-gonic/gin"
            "gorm.io/gorm"

            "github.com/yourorg/student-management-service-go/internal/models"
        )

        type Handler struct{
            db *gorm.DB
        }

        func NewHandler(db *gorm.DB) *Handler { return &Handler{db: db} }

        func (h *Handler) CreateStudent(c *gin.Context) {
            var input models.Student
            if err := c.ShouldBindJSON(&input); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": gin.H{"code": "validation.failed", "message": err.Error()}})
                return
            }

            // check conflicts
            var ex models.Student
            if err := h.db.Where("email = ? OR student_number = ?", input.Email, input.StudentNumber).First(&ex).Error; err == nil {
                c.JSON(http.StatusConflict, gin.H{"ok": false, "error": gin.H{"code": "student.duplicate", "message": "Student with same email or student number exists"}})
                return
            }

            if err := h.db.Create(&input).Error; err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": gin.H{"code": "internal_error", "message": err.Error()}})
                return
            }

            c.JSON(http.StatusCreated, gin.H{"ok": true, "data": input})
        }

        func (h *Handler) ListStudents(c *gin.Context) {
            page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
            limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
            search := c.Query("search")
            programId := c.Query("programId")
            enrollmentYear := c.Query("enrollmentYear")

            var students []models.Student
            query := h.db.Preload("Program")

            if search != "" {
                like := "%" + search + "%"
                query = query.Where("first_name LIKE ? OR last_name LIKE ? OR email LIKE ? OR student_number LIKE ?", like, like, like, like)
            }
            if programId != "" {
                if pid, err := strconv.Atoi(programId); err == nil {
                    query = query.Where("program_id = ?", pid)
                }
            }
            if enrollmentYear != "" {
                if ey, err := strconv.Atoi(enrollmentYear); err == nil {
                    query = query.Where("enrollment_year = ?", ey)
                }
            }

            var total int64
            query.Model(&models.Student{}).Count(&total)

            offset := (page - 1) * limit
            if err := query.Offset(offset).Limit(limit).Find(&students).Error; err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": gin.H{"code": "internal_error", "message": err.Error()}})
                return
            }

            c.JSON(http.StatusOK, gin.H{"ok": true, "data": gin.H{"items": students}, "meta": gin.H{"total": total, "page": page, "limit": limit}})
        }

        func (h *Handler) GetStudent(c *gin.Context) {
            id := c.Param("id")
            var s models.Student
            if err := h.db.Preload("Program").First(&s, id).Error; err != nil {
                if err == gorm.ErrRecordNotFound {
                    c.JSON(http.StatusNotFound, gin.H{"ok": false, "error": gin.H{"code": "student.not_found", "message": "Student not found"}})
                    return
                }
                c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": gin.H{"code": "internal_error", "message": err.Error()}})
                return
            package handlers

            import (
                "net/http"
                "strconv"

                "github.com/gin-gonic/gin"
                "gorm.io/gorm"

                "github.com/yourorg/student-management-service-go/internal/models"
            )

            type Handler struct{
                db *gorm.DB
            }

            func NewHandler(db *gorm.DB) *Handler { return &Handler{db: db} }

            func (h *Handler) CreateStudent(c *gin.Context) {
                var input models.Student
                if err := c.ShouldBindJSON(&input); err != nil {
                    c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": gin.H{"code": "validation.failed", "message": err.Error()}})
                    return
                }

                // check conflicts
                var ex models.Student
                if err := h.db.Where("email = ? OR student_number = ?", input.Email, input.StudentNumber).First(&ex).Error; err == nil {
                    c.JSON(http.StatusConflict, gin.H{"ok": false, "error": gin.H{"code": "student.duplicate", "message": "Student with same email or student number exists"}})
                    return
                }

                if err := h.db.Create(&input).Error; err != nil {
                    c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": gin.H{"code": "internal_error", "message": err.Error()}})
                    return
                }

                c.JSON(http.StatusCreated, gin.H{"ok": true, "data": input})
            }

            func (h *Handler) ListStudents(c *gin.Context) {
                page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
                limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
                search := c.Query("search")
                programId := c.Query("programId")
                enrollmentYear := c.Query("enrollmentYear")

                var students []models.Student
                query := h.db.Preload("Program")

                if search != "" {
                    like := "%" + search + "%"
                    query = query.Where("first_name LIKE ? OR last_name LIKE ? OR email LIKE ? OR student_number LIKE ?", like, like, like, like)
                }
                if programId != "" {
                    if pid, err := strconv.Atoi(programId); err == nil {
                        query = query.Where("program_id = ?", pid)
                    }
                }
                if enrollmentYear != "" {
                    if ey, err := strconv.Atoi(enrollmentYear); err == nil {
                        query = query.Where("enrollment_year = ?", ey)
                    }
                }

                var total int64
                query.Model(&models.Student{}).Count(&total)

                offset := (page - 1) * limit
                if err := query.Offset(offset).Limit(limit).Find(&students).Error; err != nil {
                    c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": gin.H{"code": "internal_error", "message": err.Error()}})
                    return
                }

                c.JSON(http.StatusOK, gin.H{"ok": true, "data": gin.H{"items": students}, "meta": gin.H{"total": total, "page": page, "limit": limit}})
            }

            func (h *Handler) GetStudent(c *gin.Context) {
                id := c.Param("id")
                var s models.Student
                if err := h.db.Preload("Program").First(&s, id).Error; err != nil {
                    if err == gorm.ErrRecordNotFound {
                        c.JSON(http.StatusNotFound, gin.H{"ok": false, "error": gin.H{"code": "student.not_found", "message": "Student not found"}})
                        return
                    }
                    c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": gin.H{"code": "internal_error", "message": err.Error()}})
                    return
                }
                c.JSON(http.StatusOK, gin.H{"ok": true, "data": s})
            }

            func (h *Handler) UpdateStudent(c *gin.Context) {
                id := c.Param("id")
                var s models.Student
                if err := h.db.First(&s, id).Error; err != nil {
                    c.JSON(http.StatusNotFound, gin.H{"ok": false, "error": gin.H{"code": "student.not_found", "message": "Student not found"}})
                    return
                }

                var input models.Student
                if err := c.ShouldBindJSON(&input); err != nil {
                    c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": gin.H{"code": "validation.failed", "message": err.Error()}})
                    return
                }

                // check uniqueness
                var conflict models.Student
                if err := h.db.Where("(email = ? OR student_number = ?) AND id <> ?", input.Email, input.StudentNumber, s.ID).First(&conflict).Error; err == nil {
                    c.JSON(http.StatusConflict, gin.H{"ok": false, "error": gin.H{"code": "student.duplicate", "message": "Conflict with existing student"}})
                    return
                }

                s.FirstName = input.FirstName
                s.LastName = input.LastName
                s.Email = input.Email
                s.ProgramID = input.ProgramID
                s.EnrollmentYear = input.EnrollmentYear

                if err := h.db.Save(&s).Error; err != nil {
                    c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": gin.H{"code": "internal_error", "message": err.Error()}})
                    return
                }
                c.JSON(http.StatusOK, gin.H{"ok": true, "data": s})
            }

            func (h *Handler) DeleteStudent(c *gin.Context) {
                id := c.Param("id")
                if err := h.db.Delete(&models.Student{}, id).Error; err != nil {
                    c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": gin.H{"code": "internal_error", "message": err.Error()}})
                    return
                }
                c.Status(http.StatusNoContent)
            }

            func (h *Handler) ListColleges(c *gin.Context) {
                var colleges []models.College
                if err := h.db.Preload("Programs").Preload("Departments").Find(&colleges).Error; err != nil {
                    c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": gin.H{"code": "internal_error", "message": err.Error()}})
                    return
                }
                c.JSON(http.StatusOK, gin.H{"ok": true, "data": colleges})
            }
