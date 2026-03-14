package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	ClientHandler struct {
		BaseHandler
		uc domain.ClientUseCase
	}

	clientRequest struct {
		Phone string `json:"phone" validate:"required"`
		Email string `json:"email" validate:"required"`
	}
)

func NewClientsController(uc domain.ClientUseCase) *ClientHandler {
	return &ClientHandler{uc: uc}
}

// Add godoc
// @Summary      Create an Account
// @Description  Create a new Account
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param        establishment  body      domain.Establishment  true  "Account data"
// @Success      201            {object}  domain.HttpResponse{data=domain.Account}  "Account created successfully"
// @Failure      400            {object}  domain.HttpResponse  "Bad Request"
// @Failure      500            {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /clients [post]
func (h *ClientHandler) Add(ctx *gin.Context) {
	var req clientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.RespondWithError(ctx, http.StatusBadRequest, "invalid request payload", err)
		return
	}

	client := domain.Client{
		Email: req.Email,
		Phone: req.Phone,
	}

	c, err := h.uc.Add(ctx.Request.Context(), &client)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusCreated, "account created successfully", c)
}
