package handler

import (
	"net/http"

	"go-profile-service-magangku/internal/domain"
	"go-profile-service-magangku/internal/repository"
	"go-profile-service-magangku/internal/response"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	Repo *repository.ProfileRepository
}

func NewProfileHandler(r *repository.ProfileRepository) *ProfileHandler {
	return &ProfileHandler{Repo: r}
}

func (h *ProfileHandler) GetMyProfile(c *gin.Context) {
	userID := c.GetString("user_id")

	profile, err := h.Repo.GetByUserID(c, userID)
	if err != nil {
		c.JSON(500, response.APIResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	if profile == nil {
		c.JSON(200, response.APIResponse{
			Message: "Profile not found",
			Data:    nil,
		})
		return
	}

	c.JSON(200, response.APIResponse{
		Message: "Success",
		Data:    profile,
	})
}

func (h *ProfileHandler) CreateMyProfile(c *gin.Context) {
	userID := c.GetString("user_id")

	existing, _ := h.Repo.GetByUserID(c, userID)
	if existing != nil {
		c.JSON(409, response.APIResponse{
			Message: "Profile already exists",
			Data:    existing,
		})
		return
	}


	var req domain.Profile
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.APIResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	req.UserID = userID

	if err := h.Repo.Create(c, &req); err != nil {
		c.JSON(500, response.APIResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	profile, _ := h.Repo.GetByUserID(c, userID)

	c.JSON(http.StatusCreated, response.APIResponse{
		Message: "Profile created",
		Data:    profile,
	})
}

func (h *ProfileHandler) UpdateMyProfile(c *gin.Context) {
	userID := c.GetString("user_id")

	var req domain.Profile
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, response.APIResponse{
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	req.UserID = userID

	if err := h.Repo.Update(c, &req); err != nil {
		c.JSON(404, response.APIResponse{
			Message: "Profile not found",
			Data:    nil,
		})
		return
	}

	profile, _ := h.Repo.GetByUserID(c, userID)

	c.JSON(200, response.APIResponse{
		Message: "Profile updated",
		Data:    profile,
	})
}
