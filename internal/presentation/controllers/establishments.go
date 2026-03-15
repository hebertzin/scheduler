package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/domain/ports/inbound"
)

type (
	EstablishmentHandler struct {
		BaseHandler
		manager inbound.EstablishmentUseCase
	}

	establishmentRequest struct {
		Name       string `json:"name" validate:"required"`
		City       string `json:"city" validate:"required"`
		State      string `json:"state" validate:"required"`
		PostalCode string `json:"postalCode" validate:"required"`
		Number     string `json:"number" validate:"required"`
		UserId     string `json:"userId" validate:"required"`
	}
)

func NewEstablishment(manager inbound.EstablishmentUseCase) *EstablishmentHandler {
	return &EstablishmentHandler{manager: manager}
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

	establishment, err := h.manager.Add(ctx.Request.Context(), &estblishmentCreated)
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
	establishment, err := h.manager.FindEstablishmentById(ctx.Request.Context(), id)
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
	id := getIdParam(ctx, "id")
	professionals, err := h.manager.GetAllProfessionalsByEstablishmentId(ctx.Request.Context(), id)
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
	id := getIdParam(ctx, "id")

	var req establishmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.RespondWithError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	dto := domain.Establishment{
		Name:       req.Name,
		City:       req.City,
		PostalCode: req.PostalCode,
		State:      req.State,
		Number:     req.Number,
		UserId:     req.UserId,
	}

	establishments, err := h.manager.UpdateEstablishmentById(ctx.Request.Context(), id, &dto)
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
	id := getIdParam(ctx, "id")
	metrics, err := h.manager.GetEstablishmentReport(ctx.Request.Context(), id)
	if err != nil {
		h.RespondWithError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "establishment metrics", metrics)
}

func getIdParam(ctx *gin.Context, paramName string) string {
	return ctx.Param(paramName)
}
