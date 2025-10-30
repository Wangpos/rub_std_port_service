package models

import "time"

type College struct {
    ID          uint `gorm:"primaryKey"`
    Name        string `gorm:"unique;not null"`
    Departments []Department
    Programs    []Program
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

type Department struct {
    ID        uint `gorm:"primaryKey"`
    Name      string `gorm:"not null"`
    CollegeID uint
    College   College
    Programs  []Program
    CreatedAt time.Time
    UpdatedAt time.Time
}

type Program struct {
    ID           uint `gorm:"primaryKey"`
    Name         string `gorm:"not null"`
    DepartmentID uint
    Department   Department
    CollegeID    uint
    College      College
    Students     []Student
    CreatedAt    time.Time
    UpdatedAt    time.Time
}

type Student struct {
    ID             uint `gorm:"primaryKey"`
    StudentNumber  string `gorm:"unique;not null"`
    FirstName      string `gorm:"not null"`
    LastName       string `gorm:"not null"`
    Email          string `gorm:"unique;not null"`
    DOB            *time.Time
    EnrollmentYear int
    ProgramID      uint
    Program        Program
    CreatedAt      time.Time
    UpdatedAt      time.Time
}
