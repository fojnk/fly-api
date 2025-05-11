package transport

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Get inbound flights by airport
// @Tags schedule
// @Description get all inbound flights
// @ID get-inbound-flights
// @Param airport path string true "Airport"
// @Param offset query string true "Offset"
// @Param limit query string true "Limit"
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/inbound-schedule/{airport} [get]
func (h *Handler) InboundSchedule(c *gin.Context) {
	airport := c.Param("airport")
	paramOff := c.Query("offset")
	paramLim := c.Query("limit")

	offset, _ := strconv.Atoi(paramOff)
	limit, _ := strconv.Atoi(paramLim)

	currTime := time.Now().AddDate(-8, 4, 0)

	schedule, err := h.services.IScheduleService.GetInboundSchedule(airport, currTime.Format(time.RFC3339), offset, limit)

	if err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"schedule": schedule,
	})
}

// @Summary Get outbound flights by airport
// @Tags schedule
// @Description get all outbound flights
// @Param offset query string true "Offset"
// @Param limit query string true "Limit"
// @Param airport path string true "Airport"
// @ID get-outbound-flights
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/outbound-schedule/{airport} [get]
func (h *Handler) OutboundSchedule(c *gin.Context) {
	airport := c.Param("airport")
	paramOff := c.Query("offset")
	paramLim := c.Query("limit")

	offset, _ := strconv.Atoi(paramOff)
	limit, _ := strconv.Atoi(paramLim)

	currTime := time.Now().AddDate(-8, 4, 0)

	schedule, err := h.services.IScheduleService.GetOutboundSchedule(airport, currTime.Format(time.RFC3339), offset, limit)

	if err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"schedule": schedule,
	})
}
