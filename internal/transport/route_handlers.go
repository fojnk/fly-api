package transport

import (
	"flyAPI/internal/dto/request"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Get all routes by params
// @Tags routes
// @Description get all routes
// @ID get-routes
// @Produce  json
// @Param src query string true "Departure airport/city"
// @Param dest query string true "Arrival airport/city"
// @Param date query string false "Date for start searching"
// @Param limit query int false "Flight limit"
// @Param conditions query string true "Fare conditions"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/routes [get]
func (h *Handler) AllRoutes(c *gin.Context) {
	src := c.Query("src")
	if src == "" || !h.services.IAirService.IsOriginExists(src) {
		NewTransportErrorResponse(c, http.StatusBadRequest, "source not found")
		return
	}

	dest := c.Query("dest")
	if dest == "" || !h.services.IAirService.IsOriginExists(dest) {
		NewTransportErrorResponse(c, http.StatusBadRequest, "destination not found")
		return
	}

	date := c.Query("date")
	parsedTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		parsedTime = time.Now().AddDate(-8, 4, 0)
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 0
	}
	fareConditions := c.Query("conditions")

	input := request.FlightParams{
		Src:            src,
		Dest:           dest,
		DepartureDate:  parsedTime.Format(time.RFC3339),
		LenghtLimit:    limit,
		FareConditions: fareConditions,
	}

	logrus.Info("get routes input", input)

	flights, err := h.services.IRouteService.GetRoutes(request.FlightParams{
		Src:            src,
		Dest:           dest,
		DepartureDate:  parsedTime.Format(time.RFC3339),
		LenghtLimit:    limit,
		FareConditions: fareConditions,
	})

	if err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"path": flights,
	})
}
