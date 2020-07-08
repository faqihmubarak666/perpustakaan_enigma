package master

import (
	"database/sql"
	"gomux/main/master/controllers"
	"gomux/main/master/repositories"
	"gomux/main/master/usecases"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *sql.DB) {
	//perpustakaan
	perpustakaanRepo := repositories.InitPerpustakaanRepoImpl(db)
	perpustakaanUseCase := usecases.InitPerpustakaanUseCase(perpustakaanRepo)
	controllers.PerpustakaanController(r, perpustakaanUseCase)

	//User
	userRepo := repositories.InitUserRepoImpl(db)
	userUseCase := usecases.InitUserUsecase(userRepo)
	controllers.UserController(r, userUseCase)

}
