package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	EstablishmentHandler struct {
		BaseHandler
		uc domain.EstablishmentUseCase
	}

	establishmentRequest struct {
		Name       string    `json:"name" validate:"required"`
		City       string    `json:"city" validate:"required"`
		State      string    `json:"state" validate:"required"`
		PostalCode string    `json:"postal_code" validate:"required"`
		Number     string    `json:"number" validate:"required"`
		UserId     string    `json:"user_id" validate:"required"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}
)

func NewEstablishmentController(uc domain.EstablishmentUseCase) *EstablishmentHandler {
	return &EstablishmentHandler{uc: uc}
}

// Add godoc
// @Summary      Add a new establishment
// @Description  Create a new establishment with the provided data
// @Tags         Establishments
// @Accept       json
// @Produce      json
// @Param        establishment  body      domain.Establishment  true  "Establishment data"
// @Success      201            {object}  domain.HttpResponse{data=domain.Establishment}  "Establishment created successfully"
// @Failure      400            {object}  domain.HttpResponse  "Bad Request"
// @Failure      500            {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /establishments [post]
func (h *EstablishmentHandler) Add(ctx *gin.Context) {
	var req establishmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.RespondWithError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	estblishmentCreated := domain.Establishment{
		Name:       req.Name,
		City:       req.City,
		PostalCode: req.PostalCode,
		State:      req.State,
		Number:     req.Number,
		UserId:     req.UserId,
	}

	establishment, err := h.uc.Add(ctx.Request.Context(), &estblishmentCreated)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusCreated, "establishment created successfully", establishment)
}

// FindEstablishmentById godoc
// @Summary      Find establishment by id
// @Description  The enpoint retrieve an establishment using its unique ID
// @Tags         Establishments
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Establishment ID"
// @Success      200  {object}  domain.HttpResponse{data=domain.Establishment}  "Establishment found successfully"
// @Failure      404  {object}  domain.HttpResponse  "Establishment not found"
// @Failure      500  {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /establishment_id/{id} [get]
func (h *EstablishmentHandler) FindEstablishmentById(ctx *gin.Context) {
	id := ctx.Param("id")
	establishment, err := h.uc.FindEstablishmentById(ctx.Request.Context(), id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "establishment found successfully", establishment)
}

// GetAllProfessinalsByEstablishmentId godoc
// @Summary      Get All professionals by establishment ID
// @Description  The endpoint retrieve professionals using establishment id
// @Tags         Establishments
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "id"
// @Success      200  {object}  domain.HttpResponse{data=domain.Establishment}  "Establishment found successfully"
// @Failure      404  {object}  domain.HttpResponse  "Professionals found successfully"
// @Failure      500  {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /establishment/:id/professionals [get]
func (h *EstablishmentHandler) GetAllProfessinalsByEstablishmentId(ctx *gin.Context) {
	establishment_id := ctx.Param("id")
	professionals, err := h.uc.GetAllProfessionalsByEstablishmentId(ctx.Request.Context(), establishment_id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}
	h.RespondWithSuccess(ctx, http.StatusOK, "professionals found successfully", professionals)
}

// UpdateEstablishmentById godoc
// @Summary      Update establishment by id
// @Description  The endpoint update establishment using id
// @Tags         Establishments
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "id"
// @Failure      200  {object}  domain.HttpResponse  "Establishment update successfully"
// @Failure      500  {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /establishments/:id/update [put]
func (h *EstablishmentHandler) UpdateEstablishmentById(ctx *gin.Context) {
	establishment_id := ctx.Param("id")
	var input domain.Establishment
	establishments, err := h.uc.UpdateEstablishmentById(ctx.Request.Context(), establishment_id, &input)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "establishment update successfully", establishments)
}

// GetEstablishmentReport godoc
// @Summary      Get establishmentReport
// @Description  The endpoint get establishment report using id
// @Tags         Establishments
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "id"
// @Failure      200  {object}  domain.HttpResponse  "Establishment report"
// @Failure      400  {object}  domain.HttpResponse  "Establishment id not found"
// @Failure      500  {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /establishments/:id/report [get]
func (h *EstablishmentHandler) GetEstablishmentReport(ctx *gin.Context) {
	establishmentId := ctx.Param("id")
	establishmentReport, err := h.uc.GetEstablishmentReport(ctx.Request.Context(), establishmentId)
	if err != nil {
		h.RespondWithError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "establishment report", establishmentReport)
}
