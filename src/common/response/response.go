package response

import (
	"net/http"

	"banc-api/src/common/dto"
)

type ResponseContext interface {
	JSON(code int, obj any) error
}

// Success responde con éxito
func Success(c ResponseContext, data interface{}, message string) error {
	return c.JSON(http.StatusOK, dto.BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Created responde con recurso creado
func Created(c ResponseContext, data interface{}, message string) error {
	return c.JSON(http.StatusCreated, dto.BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error responde con error
func Error(c ResponseContext, statusCode int, message string, err string) error {
	return c.JSON(statusCode, dto.ErrorResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}

// NotFound responde con recurso no encontrado
func NotFound(c ResponseContext, message string) error {
	return Error(c, http.StatusNotFound, message, "")
}

// BadRequest responde con error de solicitud
func BadRequest(c ResponseContext, message string) error {
	return Error(c, http.StatusBadRequest, message, "")
}

// InternalError responde con error interno
func InternalError(c ResponseContext, message string) error {
	return Error(c, http.StatusInternalServerError, message, "")
}
