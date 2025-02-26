package dependencies

import (
	"polling/src/buses/application"
	"polling/src/buses/infraestructure"
	"polling/src/buses/infraestructure/http/controllers"
	"polling/src/core"
)

var (
	mySQL infraestructure.MySQL
)

func InitBus() {
	db, err := core.InitMySQL()
	if err != nil {
		return
	}
	mySQL = *infraestructure.NewMySQL(db)
}

func AddBusController() *controllers.AddBusController {
	ucAddBus := application.NewAddBusUseCase(&mySQL)

	return controllers.NewAddBusController(ucAddBus)
}

func GetAllBusesController() *controllers.GetAllBusesController {
	ucGetAllBuses := application.NewGetAllBusesUseCase(&mySQL)

	return controllers.NewGetAllBusesController(ucGetAllBuses)
}

func UpdateBusController() *controllers.UpdateBusByIDController {
	ucUpdateBus := application.NewUpdateBusByIDUseCase(&mySQL)

	return controllers.NewUpdateBusByIDController(ucUpdateBus)
}

func GetBusByIdChoferController() *controllers.GetBusByIdChoferController {
	ucGetBusByChofer := application.NewFindBusByIdChoferUseCase(&mySQL)

	return controllers.NewGetBusByIdChoferController(ucGetBusByChofer)
}

func DeleteBusController() *controllers.DeleteBusByIDController {
	ucDeleteBus := application.NewDeleteBusByIDUseCase(&mySQL)

	return controllers.NewDeleteBusByIDController(ucDeleteBus)
}

func PollingBusController() *controllers.EventBusController {
	ucEventBus := application.NewEventBusUseCase(&mySQL)

	return controllers.NewEventBusController(ucEventBus)
}
