package user

import (
	"net/http"
	"strings"

	"github.com/isd-sgcu/johnjud-gateway/src/app/constant"
	"github.com/isd-sgcu/johnjud-gateway/src/app/dto"
	"github.com/isd-sgcu/johnjud-gateway/src/app/router"
	"github.com/isd-sgcu/johnjud-gateway/src/app/validator"
	userconst "github.com/isd-sgcu/johnjud-gateway/src/constant/user"
	"github.com/isd-sgcu/johnjud-gateway/src/pkg/service/user"
)

type Handler struct {
	service  user.Service
	validate validator.IDtoValidator
}

func NewHandler(service user.Service, validate validator.IDtoValidator) *Handler {
	return &Handler{service, validate}
}

// FindOne is a function that returns a user by id from database
// @Summary finds one user
// @Description Returns the data of user if successful
// @Param id path string true "user id"
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} dto.User
// @Failure 400 {object} dto.ResponseBadRequestErr "Invalid request body"
// @Failure 500 {object} dto.ResponseInternalErr "Internal service error"
// @Failure 503 {object} dto.ResponseServiceDownErr "Service is down"
// @Router /v1/users/{id} [get]
func (h *Handler) FindOne(c *router.FiberCtx) {
	id, err := c.ID()
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	user, errRes := h.service.FindOne(id)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusCreated, dto.ResponseSuccess{
		StatusCode: http.StatusOK,
		Message:    userconst.FindOneUserSuccessMessage,
		Data:       user,
	})
	return
}

// Update is a function that updates user in database
// @Summary updates user
// @Description Returns the data of user if successfully
// @Param update body dto.UpdateUserRequest true "update user dto"
// @Tags auth
// @Accept json
// @Produce json
// @Success 201 {object} dto.User
// @Failure 400 {object} dto.ResponseBadRequestErr "Invalid request body"
// @Failure 500 {object} dto.ResponseInternalErr "Internal service error"
// @Failure 503 {object} dto.ResponseServiceDownErr "Service is down"
// @Router /v1/users [put]
func (h *Handler) Update(c *router.FiberCtx) {
	usrId := c.UserID()

	request := &dto.UpdateUserRequest{}

	err := c.Bind(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    constant.BindingRequestErrorMessage + err.Error(),
			Data:       nil,
		})
		return
	}

	if err := h.validate.Validate(request); err != nil {
		var errorMessage []string
		for _, reqErr := range err {
			errorMessage = append(errorMessage, reqErr.Message)
		}
		c.JSON(http.StatusBadRequest, dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    constant.InvalidRequestBodyMessage + strings.Join(errorMessage, ", "),
			Data:       nil,
		})
		return
	}

	user, errRes := h.service.Update(usrId, request)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, dto.ResponseSuccess{
		StatusCode: http.StatusOK,
		Message:    userconst.UpdateUserSuccessMessage,
		Data:       user,
	})
	return
}
