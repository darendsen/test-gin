package utils

import (
	"gorm.io/gorm"
)

// Age filter users by specific age
func Age(age uint8) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("age = ?", age)
	}
}

// Adults filter users by age >= 18
func Adults(db *gorm.DB) *gorm.DB {
	return db.Where("age >= ?", 18)
}

// Minors filter users by age < 18
func Minors(db *gorm.DB) *gorm.DB {
	return db.Where("age < ?", 18)
}
