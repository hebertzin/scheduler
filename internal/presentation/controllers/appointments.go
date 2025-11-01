package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domain"
)

type (
	AppointmentHandler struct {
		BaseHandler
		uc domain.AppointmentUseCase
	}

	appointmentRequest struct {
		ProfessionalID string    `json:"professional_id" validate:"required"`
		ServiceID      string    `json:"service_id" validate:"required"`
		ScheduledDate  time.Time `json:"schedule_date" validate:"required"`
		Email          string    `json:"user_email" validate:"required"`
		Phone          string    `json:"user_phone" validate:"required"`
		Notes          string    `json:"notes"`
	}
)

func NewAppointmentController(uc domain.AppointmentUseCase) *AppointmentHandler {
	return &AppointmentHandler{uc: uc}
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
		ScheduledDate:  req.ScheduledDate,
		Email:          req.Email,
		Phone:          req.Phone,
		Notes:          req.Notes,
	}

	appointmentCreated, err := h.uc.Add(ctx.Request.Context(), &appointment)
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
	appointments, err := h.uc.GetAllAppointmentsByProfessionalId(ctx.Request.Context(), id)
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
	appointment, err := h.uc.GetAppointmentById(ctx.Request.Context(), id)
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
	err := h.uc.DeleteAppointment(ctx.Request.Context(), id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusNoContent, "appointment deleted", nil)
}
