package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/domain/ports/inbound"
)

type (
	ProfessionalAvailabilityHandler struct {
		BaseHandler
		manager inbound.ProfessionalsAvailabilityUseCase
	}

	professionalAvailabilityRequest struct {
		ProfessionalId uuid.UUID `json:"professionalId" validate:"required"`
		DayOfWeek      string    `json:"dayOfWeek" validate:"required"`
		StartTime      string    `json:"startTime" validate:"required"`
		EndTime        string    `json:"endTime" validate:"required"`
	}
)

func NewProfessionalAvailability(manager inbound.ProfessionalsAvailabilityUseCase) *ProfessionalAvailabilityHandler {
	return &ProfessionalAvailabilityHandler{manager: manager}
}

// Add godoc
// @Summary      Add
// @Description  Add
// @Tags         ProfessionalAvailability
// @Accept       json
// @Produce      json
// @Param        establishment  body      domain.Establishment  true  "ProfessionalAvailabilityHandler data"
// @Success      201            {object}  domain.HttpResponse{data=domain.ProfessionalAvailability}  "professional availability created successfully by id successfully retrieved"
// @Failure      400            {object}  domain.HttpResponse  "Bad Request"
// @Failure      500            {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /availability/ [post]
func (h *ProfessionalAvailabilityHandler) Add(ctx *gin.Context) {
	var req professionalAvailabilityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.RespondWithError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	professionalAvailabilityCreated := domain.ProfessionalAvailability{
		ProfessionalId: req.ProfessionalId,
		DayOfWeek:      req.DayOfWeek,
		StartTime:      req.StartTime,
		EndTime:        req.EndTime,
	}

	availability, err := h.manager.Add(ctx.Request.Context(), &professionalAvailabilityCreated)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusCreated, "pprofessional availability created successfully", availability)
}

// Add godoc
// @Summary      GetProfessionalAvailabilityById
// @Description  GetProfessionalAvailabilityById
// @Tags         ProfessionalAvailability
// @Accept       json
// @Produce      json
// @Param        establishment  body      domain.Establishment  true  "ProfessionalAvailabilityHandler data"
// @Success      201            {object}  domain.HttpResponse{data=domain.ProfessionalAvailability}  "professional availability retrieved successfullys availability created successfully by id successfully retrieved"
// @Failure      400            {object}  domain.HttpResponse  "Bad Request"
// @Failure      500            {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /availability/:id/professional [get]
func (h *ProfessionalAvailabilityHandler) GetProfessionalAvailabilityById(ctx *gin.Context) {
	professionalId := ctx.Param("id")
	availability, err := h.manager.GetProfessionalAvailabilityById(ctx.Request.Context(), professionalId)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "pprofessional availability retrieved successfully", availability)
}
