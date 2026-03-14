package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	ServicesHandler struct {
		BaseHandler
		manager domain.ServicesUseCase
	}

	serviceRequest struct {
		Name           string `json:"name" validate:"required"`
		Value          string `json:"value" validate:"required"`
		Duration       string `json:"duration" validate:"required"`
		ProfessionalId string `json:"professionalId" validate:"required"`
	}
)

func NewServices(manager domain.ServicesUseCase) *ServicesHandler {
	return &ServicesHandler{manager: manager}
}

// AddService godoc
// @Summary      Add a new service
// @Description  Create a new service with the provided data
// @Tags         Services
// @Accept       json
// @Produce      json
// @Param        service  body      domain.Services  true  "Service data"
// @Success      201      {object}  domain.HttpResponse{data=domain.Services}  "Service created successfully"
// @Failure      400      {object}  domain.HttpResponse  "Bad Request"
// @Failure      500      {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /services [post]
func (h *ServicesHandler) Add(ctx *gin.Context) {
	var req serviceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.RespondWithError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	serviceCreated := domain.Services{
		Name:           req.Name,
		Value:          req.Value,
		Duration:       req.Duration,
		ProfessionalId: req.ProfessionalId,
	}

	service, err := h.manager.Add(ctx.Request.Context(), &serviceCreated)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusCreated, "service created successfully", service)
}

// FindServiceById godoc
// @Summary      Find service by ID
// @Description  Retrieve a service using its unique ID
// @Tags         Services
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Service ID"
// @Success      200  {object}  domain.HttpResponse{data=domain.Services}  "Service found successfully"
// @Failure      404  {object}  domain.HttpResponse  "Service not found"
// @Failure      500  {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /services/{id} [get]
func (h *ServicesHandler) FindServiceById(ctx *gin.Context) {
	id := ctx.Param("id")
	service, err := h.manager.FindServiceById(ctx.Request.Context(), id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "service found successfully", service)
}

// FindServiceById godoc
// @Summary      GetAllServicesByProfessionalId
// @Description  GetAllServicesByProfessionalId
// @Tags         Services
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Service ID"
// @Success      200  {object}  domain.HttpResponse{data=domain.Services}  "all services found successfully"
// @Failure      404  {object}  domain.HttpResponse  "Service not found"
// @Failure      500  {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /services/{id}/all [get]
func (h *ServicesHandler) GetAllServicesByProfessionalId(ctx *gin.Context) {
	professionalId := ctx.Param("id")
	services, err := h.manager.GetAllServicesByProfessionalId(ctx.Request.Context(), professionalId)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "all services found successfully", services)
}
