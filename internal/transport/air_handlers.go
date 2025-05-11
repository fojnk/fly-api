package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all cities
// @Tags cities
// @Description get all cities
// @ID get-cities
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/cities [get]
func (h *Handler) AllCities(c *gin.Context) {

	cities, err := h.services.IAirService.GetAllCities()

	if err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"src_sities":  cities.SrcCities,
		"dest_sities": cities.DestCities,
	})
}

// @Summary Get all airports
// @Tags aiports
// @Description get all aiports
// @ID get-aiports
// @Produce  json
// @Param lang query string true "Language"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/airports [get]
func (h *Handler) AllAirports(c *gin.Context) {
	lang := c.Query("lang")
	airports, err := h.services.IAirService.GetAllAirports(lang)

	if err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"src_airports":  airports.SrcAirports,
		"dest_airports": airports.DestAirports,
	})
}

// @Summary Get Aiports by city
// @Tags aiports
// @Description Get aiports by city
// @ID get-airports-by-city
// @Accept  json
// @Produce  json
// @Param city path string true "City"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/airports/{city} [get]
func (h *Handler) AirportByCity(c *gin.Context) {
	city := c.Param("city")

	airports, err := h.services.IAirService.GetAirportsByCity(city)

	if err != nil {
		NewTransportErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"airports": airports,
	})
}
