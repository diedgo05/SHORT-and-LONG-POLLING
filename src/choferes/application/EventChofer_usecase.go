package application

import  "polling/src/choferes/domain"

type EventChoferUseCase struct {
	db domain.IChoferesRepository
}

func NewEventChoferUseCase(db domain.IChoferesRepository) *EventChoferUseCase {
	return &EventChoferUseCase{db: db}
}

func (u *EventChoferUseCase) Run()  []domain.Chofer {
	allChoferes, err := u.db.FindAll()
	if err != nil {
		return []domain.Chofer{}
	}
	var choferes []domain.Chofer
	choferes = append(choferes, allChoferes...)
	return choferes
}