package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hebertzin/scheduler/internal/domain"
	"github.com/hebertzin/scheduler/internal/domain/ports/inbound"
)

type (
	AccountHandler struct {
		BaseHandler
		manager inbound.AccountUseCase
	}

	accountRequest struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
)

func NewAccount(manager inbound.AccountUseCase) *AccountHandler {
	return &AccountHandler{manager: manager}
}

// Add godoc
// @Summary      Add a new account
// @Description  Create a new account with the provided data
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        account  body      domain.account  true  "account data"
// @Success      201   {object}  domain.HttpResponse{data=dto.accountDTO}  "account created successfully"
// @Failure      400   {object}  domain.HttpResponse  "Bad Request"
// @Failure      500   {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /accounts [post]
func (h *AccountHandler) Add(ctx *gin.Context) {
	var req accountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.RespondWithError(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	accountCreated := domain.Account{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	account, err := h.manager.Add(ctx.Request.Context(), &accountCreated)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusCreated, "account created successfully", account)
}

// FindaccountById godoc
// @Summary      Find a account by ID
// @Description  Retrieve a account by their unique ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id  path      string  true  "account ID"
// @Success      200  {object}  domain.HttpResponse{data=dto.accountDTO}  "account found successfully"
// @Failure      400  {object}  domain.HttpResponse  "Bad Request"
// @Failure      404  {object}  domain.HttpResponse  "account not found"
// @Failure      500  {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /accounts/{id} [get]
func (h *AccountHandler) FindAccountById(ctx *gin.Context) {
	id := ctx.Param("id")
	account, err := h.manager.FindAccountById(ctx.Request.Context(), id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "account found successfully", account)
}

// we realy need this route ???
// FindAllaccounts godoc
// @Summary      Get all accounts
// @Description  Retrieve a list of all accounts in the system
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {object}  domain.HttpResponse{data=[]dto.accountDTO}  "accounts retrieved successfully"
// @Failure      500  {object}  domain.HttpResponse  "Internal Server Error"
// @Router       /accounts [get]
func (h *AccountHandler) FindAllAccounts(ctx *gin.Context) {
	accounts, err := h.manager.FindAllAccounts(ctx.Request.Context())
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
		return
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "accounts retrieved", accounts)
}

func (h *AccountHandler) FindAllEstablishmentsByAccountId(ctx *gin.Context) {
	id := ctx.Param("id")
	establishments, err := h.manager.FindAllEstablishmentsByAccountId(ctx.Request.Context(), id)
	if err != nil {
		h.RespondWithError(ctx, err.Code, err.Message, err)
	}

	h.RespondWithSuccess(ctx, http.StatusOK, "accounts establishments retrieved", establishments)
}
