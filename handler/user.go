package handler

import (
	"MaintenanceSystem/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 解析 JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "参数错误"})
		return
	}

	// 调用 service
	token, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	// 返回 token
	c.JSON(200, gin.H{
		"token": token,
	})
}

func (h *UserHandler) Profile(c *gin.Context) {
	userID, _ := c.Get("user_id")

	c.JSON(200, gin.H{
		"msg":     "你已登录",
		"user_id": userID,
	})
}
