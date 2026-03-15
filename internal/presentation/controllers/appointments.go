package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/domain/ports/inbound"
)

type (
	AppointmentHandler struct {
		BaseHandler
		manager inbound.AppointmentUseCase
	}

	appointmentRequest struct {
		ProfessionalID string `json:"professionalId" validate:"required"`
		ServiceID      string `json:"serviceId" validate:"required"`
		StartTime      string `json:"startTime" validate:"required"`
		EndTime        string `json:"endTime" validate:"required"`
		Email          string `json:"email" validate:"required"`
		Phone          string `json:"phone" validate:"required"`
		Notes          string `json:"notes"`
	}
)

func NewAppointment(manager inbound.AppointmentUseCase) *AppointmentHandler {
	return &AppointmentHandler{manager: manager}
}

// Add godoc
// @Summary      Create an Appointment
// @Description  Create a new Appointment
// @Tags         Appointment
// @Accept       json
// @Produce      json
// @Param        establishment  body      domain.Establishment  true  "Appointment data"
// @Success      201            {object}  domain.HttpResponse{data=domain.Appointment}  "Appointment created successfully"
// @Failure      400            {object}  domain.HttpResponse  "Bad Request"
// @Failure      500            {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /appointments [post]
func (h *AppointmentHandler) Add(ctx *gin.Context) {
	var req appointmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	appointment := domain.Appointment{
		ProfessionalId: req.ProfessionalID,
		ServiceId:      req.ServiceID,
		StartTime:      req.StartTime,
		EndTime:        req.EndTime,
		Email:          req.Email,
		Phone:          req.Phone,
		Notes:          req.Notes,
	}

	appointmentCreated, err := h.manager.Add(ctx.Request.Context(), &appointment)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusCreated, "Appointment created successfully", appointmentCreated)
}

// Add godoc
// @Summary      Get all appointments by professional id
// @Description  Get all appointments by professional ids
// @Tags         Appointment
// @Accept       json
// @Produce      json
// @Param        establishment  body      domain.Establishment  true  "Appointment data"
// @Success      201            {object}  domain.HttpResponse{data=domain.Appointment}  "appointment by id successfully retrieved"
// @Failure      400            {object}  domain.HttpResponse  "Bad Request"
// @Failure      500            {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /appointments/:id/professional [get]
func (h *AppointmentHandler) GetAllAppointmentsByProfessionalId(ctx *gin.Context) {
	id := ctx.Param("id")
	appointments, err := h.manager.GetAllAppointmentsByProfessionalId(ctx.Request.Context(), id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "appointments by professional id successfully retrieved", appointments)
}

// Add godoc
// @Summary      GetAppointmentById
// @Description  GetAppointmentById
// @Tags         Appointment
// @Accept       json
// @Produce      json
// @Param        establishment  body      domain.Establishment  true  "Appointment data"
// @Success      201            {object}  domain.HttpResponse{data=domain.Appointment}  "appointment by id successfully retrieved"
// @Failure      400            {object}  domain.HttpResponse  "Bad Request"
// @Failure      500            {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /appointments/:id [get]
func (h *AppointmentHandler) GetAppointmentById(ctx *gin.Context) {
	id := ctx.Param("id")
	appointment, err := h.manager.GetAppointmentById(ctx.Request.Context(), id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "appointment by id successfully retrieved", appointment)
}

// @Summary      DeleteAppointment
// @Description  Delete an appointment by ID
// @Tags         Appointment
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Appointment ID"
// @Success      204  {object}  domain.HttpResponse  "appointment deleted"
// @Failure      400  {object}  domain.HttpResponse  "Bad Request"
// @Failure      500  {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /appointments/{id} [delete]
func (h *AppointmentHandler) DeleteAppointment(ctx *gin.Context) {
	id := ctx.Param("id")
	err := h.manager.DeleteAppointment(ctx.Request.Context(), id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusNoContent, "appointment deleted", nil)
}
