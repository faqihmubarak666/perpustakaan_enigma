package repositories

import (
	"database/sql"
	"fmt"
	"gomux/main/master/constantaQuery"
	"gomux/main/master/models"
	"log"
)

type PerpustakaanRepoImpl struct {
	db *sql.DB
}

func (s PerpustakaanRepoImpl) GetAllBook() ([]*models.AllBook, error) {
	dataBuku := []*models.AllBook{}
	query := constantaQuery.GETALLBOOK
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		buku := models.AllBook{}
		var err = data.Scan(&buku.IdBuku, &buku.JudulBuku, &buku.NamaKategori, &buku.NamaPengarang, &buku.NamaPenerbit)
		if err != nil {
			return nil, err
		}
		dataBuku = append(dataBuku, &buku)
	}
	fmt.Println("Data Sukses di Tampikan")
	return dataBuku, nil
}

func (s PerpustakaanRepoImpl) GetAddBook(buku *models.Perpustakaan) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	query := constantaQuery.GETADDBOOK

	stmt, err := s.db.Prepare(query)
	if err != nil {
		tx.Rollback()
		log.Print(err)
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(buku.IdBuku, buku.JudulBuku, buku.IdKategori, buku.IdPengarang, buku.IdPenerbit); err != nil {
		tx.Rollback()
		log.Printf("%v", err)
		return err
	}
	fmt.Println("Data sukses di tambahkan")
	return tx.Commit()
}

func (s PerpustakaanRepoImpl) GetUpdateBook(buku *models.Perpustakaan) (*models.Perpustakaan, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	query := constantaQuery.GETUPDATEBOOK
	_, err = tx.Exec(query, buku.JudulBuku, buku.IdKategori, buku.IdPengarang, buku.IdBuku)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	fmt.Println("Data Karyawan Sukses di Ubah")
	return s.GetFindBookById(buku.IdBuku)
}

func (s PerpustakaanRepoImpl) GetDeleteBook(IdBuku string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	query := constantaQuery.GETDELETEBOOK
	_, err = tx.Exec(query, IdBuku)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	if err != nil {
		return err
	}
	fmt.Println("Data Sukses di Hapus")
	return nil
}

func (s PerpustakaanRepoImpl) GetFindBookById(idBuku string) (*models.Perpustakaan, error) {
	var buku models.Perpustakaan
	query := constantaQuery.GETFINDBOOKBYID
	err := s.db.QueryRow(query, idBuku).Scan(&buku.IdBuku, &buku.JudulBuku, &buku.NamaKategori, &buku.NamaPengarang, &buku.NamaPenerbit)
	if err != nil {
		return nil, err
	}
	fmt.Println("Data Id Sukses di Tampikan")
	return &buku, nil
}

func (s PerpustakaanRepoImpl) GetFindBookByCategory(namaKategori string) ([]*models.AllBook, error) {
	dataKategori := []*models.AllBook{}
	query := constantaQuery.GETFINDBOOKBYCATEGORY
	data, err := s.db.Query(query, namaKategori)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		kategori := models.AllBook{}
		var err = data.Scan(&kategori.IdBuku, &kategori.JudulBuku, &kategori.NamaKategori, &kategori.NamaPengarang, &kategori.NamaPenerbit)
		if err != nil {
			return nil, err
		}
		dataKategori = append(dataKategori, &kategori)
	}
	fmt.Println("Data Sukses di Tampikan")
	return dataKategori, nil
}

func (s PerpustakaanRepoImpl) GetFindBookByTitle(judulBuku string) (*models.AllBook, error) {
	var buku models.AllBook
	query := constantaQuery.GETFINDBOOKBYTITLE
	err := s.db.QueryRow(query, judulBuku).Scan(&buku.IdBuku, &buku.JudulBuku, &buku.NamaKategori, &buku.NamaPengarang, &buku.NamaPenerbit)
	if err != nil {
		return nil, err
	}
	fmt.Println("Data Id Sukses di Tampikan")
	return &buku, nil
}

func (s PerpustakaanRepoImpl) GetFindBookByAuthor(namaPengarang string) ([]*models.AllBook, error) {
	dataPengarang := []*models.AllBook{}
	query := constantaQuery.GETFINDBOOKBYAUTHOR
	data, err := s.db.Query(query, namaPengarang)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		pengarang := models.AllBook{}
		var err = data.Scan(&pengarang.IdBuku, &pengarang.JudulBuku, &pengarang.NamaKategori, &pengarang.NamaPengarang, &pengarang.NamaPenerbit)
		if err != nil {
			return nil, err
		}
		dataPengarang = append(dataPengarang, &pengarang)
	}
	fmt.Println("Data Sukses di Tampikan")
	return dataPengarang, nil
}

func (s PerpustakaanRepoImpl) GetFindBookByPublisher(namaPenerbit string) ([]*models.AllBook, error) {
	dataPenerbit := []*models.AllBook{}
	query := constantaQuery.GETFINDBOOKBYPUBLISHER
	data, err := s.db.Query(query, namaPenerbit)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		penerbit := models.AllBook{}
		var err = data.Scan(&penerbit.IdBuku, &penerbit.JudulBuku, &penerbit.NamaKategori, &penerbit.NamaPengarang, &penerbit.NamaPenerbit)
		if err != nil {
			return nil, err
		}
		dataPenerbit = append(dataPenerbit, &penerbit)
	}
	fmt.Println("Data Sukses di Tampikan")
	return dataPenerbit, nil
}

func InitPerpustakaanRepoImpl(db *sql.DB) PerpustakaanRepository {
	return &PerpustakaanRepoImpl{db}
}
