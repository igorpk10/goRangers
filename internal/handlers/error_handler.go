package handlers

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
	"igaopk.com/goPower/internal/domain/models"
	"igaopk.com/goPower/internal/logger"
)

func HandleError(err error) {
	// Register not found
	if errors.Is(err, gorm.ErrRecordNotFound) {

		panic(models.HTTPError{
			StatusCode: http.StatusNotFound,
			Message:    "Record now found",
			Error:      err.Error(),
		})
	}

	// Generic GORM error (SQL, conexão, constraint, etc)
	if errors.As(err, &gorm.ErrInvalidData) ||
		errors.As(err, &gorm.ErrInvalidTransaction) ||
		errors.As(err, &gorm.ErrMissingWhereClause) ||
		errors.As(err, &gorm.ErrPrimaryKeyRequired) ||
		errors.As(err, &gorm.ErrModelValueRequired) {

		logger.LogMessage("Database error: %v", logger.ERROR, err)

		panic(models.HTTPError{
			StatusCode: http.StatusBadRequest,
			Message:    "Internal server error",
			Error:      err.Error(),
		})
	}

	logger.LogMessage("Internal server error: %v", logger.ERROR, err)
	// Internal server error
	panic(models.HTTPError{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal server error",
		Error:      err.Error(),
	})
}
