package repositories

import "gomux/main/master/models"

type PerpustakaanRepository interface {
	GetAllBook() ([]*models.AllBook, error)
	GetAddBook(buku *models.Perpustakaan) error
	GetUpdateBook(buku *models.Perpustakaan) (*models.Perpustakaan, error)
	GetDeleteBook(idBuku string) error
	GetFindBookById(idBuku string) (*models.Perpustakaan, error)
	GetFindBookByCategory(namaKategori string) ([]*models.AllBook, error)
	GetFindBookByTitle(judulBuku string) (*models.AllBook, error)
	GetFindBookByAuthor(namaPengarang string) ([]*models.AllBook, error)
	GetFindBookByPublisher(namaPenerbit string) ([]*models.AllBook, error)
	GetTotalBook() (*models.ReportBook, error)
	GetTotalBookCategory() ([]*models.ReportBookCategory, error)
}
