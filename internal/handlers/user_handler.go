package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/darendsen/test-gin/internal/models"
	"github.com/darendsen/test-gin/internal/scopes"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	query := h.DB

	age, err := ParseAge(c.Query("age"))
	if err == nil {
		query = query.Scopes(scopes.Age(age))
	}

	var users []models.User
	query.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	userId := c.Param("id")

	var user models.User
	result := h.DB.Where("id = ?", userId).Find(&user)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"user": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	email := "darendsen@gamepoint.com"
	birthday := time.Now()
	user := models.User{Name: "Jinzhu", Email: email, Birthday: &birthday, Age: 18}
	h.DB.Select("Name", "Email", "Birthday", "Age").Create(&user)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	userId := c.Param("id")

	var user models.User
	result := h.DB.Where("id = ?", userId).Find(&user)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"user": nil,
		})
		return
	}

	age, err := ParseAge(c.PostForm("age"))
	if err == nil {
		user.Age = age
	}

	h.DB.Save(&user)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func ParseAge(ageStr string) (uint8, error) {
	if ageStr == "" {
		return 0, fmt.Errorf("age string is empty")
	}

	ageUint64, err := strconv.ParseUint(ageStr, 10, 8)
	if err != nil {
		return 0, fmt.Errorf("failed to parse age: %w", err)
	}

	return uint8(ageUint64), nil
}
