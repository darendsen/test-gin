package handlers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/darendsen/test-gin/internal/models"
	"github.com/darendsen/test-gin/internal/utils"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{DB: db}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	query := h.DB

	ageStr := c.Query("age")
	if ageStr != "" {
		ageUint64, err := strconv.ParseUint(ageStr, 10, 8)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid age parameter",
			})
			return
		}
		age := uint8(ageUint64)
		query = query.Scopes(utils.Age(age))
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
	user := models.User{Name: "Jinzhu", Email: &email, Birthday: &birthday, Age: 18}
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

	user.Age = 100
	h.DB.Save(&user)

	c.JSON(200, gin.H{
		"user": user,
	})
}
