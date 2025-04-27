package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all routes by params
// @Tags routes
// @Description get all routes
// @ID get-routes
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} transort_error
// @Failure 500 {object} transort_error
// @Failure default {object} transort_error
// @Router /api/routes [get]
func (h *Handler) AllRoutes(c *gin.Context) {

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
