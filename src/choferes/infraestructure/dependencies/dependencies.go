package dependencies

import (
	"fmt"
	"polling/src/choferes/application"
	"polling/src/choferes/infraestructure"
	"polling/src/choferes/infraestructure/http/controllers"
	"polling/src/core"
)

var (
	mySQL infraestructure.MySQL
)

func Init() {
	db, err := core.InitMySQL()
	if err != nil {
		fmt.Println("Error de servidor")
		return
	}

	//defer db.Close()
	mySQL = *infraestructure.NewMySQL(db)
}

func AddChoferController() *controllers.AddChoferController {
	ucAddChofer := application.NewAddChoferUseCase(&mySQL)

	return controllers.NewAddChoferController(ucAddChofer)
}

func GetAllChoferesController() *controllers.GetAllChoferesController {
	ucGetAllChoferes := application.NewGetAllChoferesUseCase(&mySQL)

	return controllers.NewGetAllChoferesController(ucGetAllChoferes)
}

func UpdateChoferController() *controllers.UpdateByIDChoferController {
	ucUpdateChofer := application.NewUpdateByIDChoferUseCase(&mySQL)

	return controllers.NewUpdateByIDChoferController(ucUpdateChofer)
}

func DeleteChoferController() *controllers.DeleteByIDChoferController {
	ucDeleteChofer := application.NewDeleteByIDChoferUseCase(&mySQL)

	return controllers.NewDeleteByIDChoferController(ucDeleteChofer)
}