package usecases

import (
	"gomux/main/master/models"
	"gomux/main/master/repositories"
	"gomux/utils"
)

type PerpustakaanUsecaseImpl struct {
	PerpustakaanRepo repositories.PerpustakaanRepository
}

func (s PerpustakaanUsecaseImpl) GetAllBook(page, limit, orderBy, sort, keyword string) ([]*models.AllBook, error) {
	book, err := s.PerpustakaanRepo.GetAllBook(page, limit, orderBy, sort, keyword)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s PerpustakaanUsecaseImpl) GetAddBook(buku *models.Perpustakaan) error {
	err := utils.ValidateInputNotNil(buku.JudulBuku, buku.IdKategori, buku.IdPengarang, buku.IdPenerbit)
	if err != nil {
		return err
	}
	err = utils.ValidateInputNotSymbol(buku.JudulBuku)
	if err != nil {
		return err
	}
	err = s.PerpustakaanRepo.GetAddBook(buku)
	if err != nil {
		return err
	}
	return nil
}

func (s PerpustakaanUsecaseImpl) GetUpdateBook(buku *models.Perpustakaan) (*models.Perpustakaan, error) {
	data, err := s.PerpustakaanRepo.GetUpdateBook(buku)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s PerpustakaanUsecaseImpl) GetDeleteBook(idBuku string) error {
	err := s.PerpustakaanRepo.GetDeleteBook(idBuku)
	if err != nil {
		return nil
	}
	return nil
}

func (s PerpustakaanUsecaseImpl) GetFindBookById(idBuku string) (*models.Perpustakaan, error) {
	buku, err := s.PerpustakaanRepo.GetFindBookById(idBuku)
	if err != nil {
		return nil, err
	}
	return buku, nil
}

func (s PerpustakaanUsecaseImpl) GetFindBookByCategory(namaKategori string) ([]*models.AllBook, error) {
	kategori, err := s.PerpustakaanRepo.GetFindBookByCategory(namaKategori)
	if err != nil {
		return nil, err
	}
	return kategori, nil
}

func (s PerpustakaanUsecaseImpl) GetFindBookByTitle(judulBuku string) (*models.AllBook, error) {
	buku, err := s.PerpustakaanRepo.GetFindBookByTitle(judulBuku)
	if err != nil {
		return nil, err
	}
	return buku, nil
}

func (s PerpustakaanUsecaseImpl) GetFindBookByAuthor(namaPengarang string) ([]*models.AllBook, error) {
	pengarang, err := s.PerpustakaanRepo.GetFindBookByAuthor(namaPengarang)
	if err != nil {
		return nil, err
	}
	return pengarang, nil
}

func (s PerpustakaanUsecaseImpl) GetFindBookByPublisher(namaPenerbit string) ([]*models.AllBook, error) {
	penerbit, err := s.PerpustakaanRepo.GetFindBookByPublisher(namaPenerbit)
	if err != nil {
		return nil, err
	}
	return penerbit, nil
}

func (s PerpustakaanUsecaseImpl) GetTotalBook() (*models.ReportBook, error) {
	totalBuku, err := s.PerpustakaanRepo.GetTotalBook()
	if err != nil {
		return nil, err
	}
	return totalBuku, nil
}

func (s PerpustakaanUsecaseImpl) GetTotalBookCategory() ([]*models.ReportBookCategory, error) {
	buku, err := s.PerpustakaanRepo.GetTotalBookCategory()
	if err != nil {
		return nil, err
	}
	return buku, nil
}

func InitPerpustakaanUseCase(PerpustakaanRepo repositories.PerpustakaanRepository) PerpustakaanUsecase {
	return &PerpustakaanUsecaseImpl{PerpustakaanRepo}
}
