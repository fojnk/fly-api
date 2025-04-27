package transport

import (
	"flyAPI/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware(CORSOptions{""}))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/")
	{
		api.GET("cities", h.AllCities)
		api.GET("airports", h.AllAirports)
		api.GET("airports/:city", h.AirportByCity)
		api.GET("inbound-schedult/:airport", h.InboundSchedule)
		api.GET("outboundFlights", h.OutboundSchedule)
		api.GET("routes/", h.AllRoutes)
		api.POST("book", h.Book)
		api.POST("checkIn", h.CheckIn)
	}

	return router
}
