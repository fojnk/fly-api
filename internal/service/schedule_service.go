package service

import (
	"flyAPI/internal/repository"
)

type ScheduleService struct {
	scheduleRepo repository.IScheduleRepository
}

func NewScheduleService(scheduleRepo repository.IScheduleRepository) *ScheduleService {
	return &ScheduleService{
		scheduleRepo: scheduleRepo,
	}
}

func (s *ScheduleService) GetInboundSchedule(airport string, time string, offset int, limit int) ([]repository.InboundSchedule, error) {
	return s.scheduleRepo.GetInboundScheduleForAirport(airport, time, offset, limit)
}

func (s *ScheduleService) GetOutboundSchedule(airport string, time string, offset int, limit int) ([]repository.OutboundSchedule, error) {
	return s.scheduleRepo.GetOutboundScheduleForAirport(airport, time, offset, limit)
}
