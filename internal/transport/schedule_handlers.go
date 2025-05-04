package transport

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Get inbound flights by airport
// @Tags schedule
// @Description get all inbound flights
// @ID get-inbound-flights
// @Param airport path string true "Airport"
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/inbound-schedule/{airport} [get]
func (h *Handler) InboundSchedule(c *gin.Context) {
	airport := c.Param("airport")

	currTime := time.Now().AddDate(-8, 4, 0)

	schedule, err := h.services.IScheduleService.GetInboundSchedule(airport, currTime.Format(time.RFC3339))

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

	currTime := time.Now().AddDate(-8, 4, 0)

	schedule, err := h.services.IScheduleService.GetOutboundSchedule(airport, currTime.Format(time.RFC3339))

	if err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"schedule": schedule,
	})
}
