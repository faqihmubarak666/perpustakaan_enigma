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

	//Category
	categoryRepo := repositories.InitCategoryRepoImpl(db)
	categoryUseCase := usecases.InitCategoryUseCase(categoryRepo)
	controllers.CategoryController(r, categoryUseCase)

	//Author
	authorRepo := repositories.InitAuthorRepoImpl(db)
	authorUseCase := usecases.InitAuthorUseCase(authorRepo)
	controllers.AuthorController(r, authorUseCase)

	//Publisher
	publisherRepo := repositories.InitPublisherRepoImpl(db)
	publisherUseCase := usecases.InitPublisherUseCase(publisherRepo)
	controllers.PublisherController(r, publisherUseCase)

}
