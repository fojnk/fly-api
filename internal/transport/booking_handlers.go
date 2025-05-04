package transport

import (
	"flyAPI/internal/dto/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create booking
// @Tags Booking
// @Description create booking
// @ID create-booking
// @Accept json
// @Produce  json
// @Param input body request.BookingRaceRequest true "Book data"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/book [post]
func (h *Handler) Book(c *gin.Context) {
	var input request.BookingRaceRequest

	if err := c.BindJSON(&input); err != nil {
		NewTransportErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	responses, err := h.services.IBookingService.CreateBooking(input)
	if err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": responses,
	})
}

// @Summary CheckIn
// @Tags Booking
// @Description check in
// @ID check-jn
// @Accept json
// @Produce  json
// @Param input body request.CheckInRequest true "CheckIn data"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/check-in [post]
func (h *Handler) CheckIn(c *gin.Context) {
	var input request.CheckInRequest

	if err := c.BindJSON(&input); err != nil {
		NewTransportErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.IBookingService.CheckIn(input)
	if err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"result": "done",
	})
}
