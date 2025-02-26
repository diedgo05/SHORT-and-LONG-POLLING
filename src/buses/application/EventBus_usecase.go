package application

import "polling/src/buses/domain"

type EventBusUseCase struct {
	db domain.IBusesRepository
}

func NewEventBusUseCase(db domain.IBusesRepository) *EventBusUseCase {
	return &EventBusUseCase{db: db}
}

func (u *EventBusUseCase) Run() []domain.Buses {
	allBuses, err := u.db.FindAllBuses()
	if err != nil {
		return []domain.Buses{}
	}
	var buses []domain.Buses
	buses = append(buses, allBuses...)
	return buses
}