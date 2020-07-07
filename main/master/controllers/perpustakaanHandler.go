package controllers

import (
	"encoding/json"
	"gomux/main/master/middleware"
	"gomux/main/master/models"
	"gomux/main/master/response"
	"gomux/main/master/usecases"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type PerpustakaanHandler struct {
	PerpustakaanUseCase usecases.PerpustakaanUsecase
}

func PerpustakaanController(r *mux.Router, service usecases.PerpustakaanUsecase) {
	perpustakaanHandler := PerpustakaanHandler{service}

	r.Use(middleware.ActivityLogMiddleware)
	r.HandleFunc("/book", perpustakaanHandler.AllBook).Methods(http.MethodGet)
	r.HandleFunc("/book", perpustakaanHandler.AddBook).Methods(http.MethodPost)
	r.HandleFunc("/book", perpustakaanHandler.UpdateBook).Methods(http.MethodPut)
	r.HandleFunc("/book/{id_buku}", perpustakaanHandler.DeleteBook).Methods(http.MethodDelete)
	r.HandleFunc("/findBookByCategory/{nama_kategori}", perpustakaanHandler.FindBookByCategory).Methods(http.MethodGet)
	r.HandleFunc("/findBookByTitle/{judul_buku}", perpustakaanHandler.FindBookByTitle).Methods(http.MethodGet)
	r.HandleFunc("/findBookByAuthor/{nama_pengarang}", perpustakaanHandler.FindBookByAuthor).Methods(http.MethodGet)
	r.HandleFunc("/findBookByPublisher/{nama_penerbit}", perpustakaanHandler.FindBookByPublisher).Methods(http.MethodGet)
	r.HandleFunc("/ReportTotalBook", perpustakaanHandler.TotalBook).Methods(http.MethodGet)
	r.HandleFunc("/ReportTotalBookCategory", perpustakaanHandler.TotalBookCategory).Methods(http.MethodGet)
}

func (s PerpustakaanHandler) AllBook(w http.ResponseWriter, r *http.Request) {
	buku, err := s.PerpustakaanUseCase.GetAllBook()
	if err != nil {
		w.Write([]byte("Data Not Found"))
		log.Print(err)
		return
	}
	var response response.MyResponse = response.MyResponse{"Data Sukses di Tampikan", http.StatusOK, buku}
	byteOfBuku, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfBuku)
}

func (s PerpustakaanHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	buku := new(models.Perpustakaan)
	err := json.NewDecoder(r.Body).Decode(&buku)
	if err != nil {
		w.Write([]byte("Data Not Found"))
		log.Print(err)
		return
	}
	err = s.PerpustakaanUseCase.GetAddBook(buku)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Print(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data Sukses di Tambah"))
}

func (s PerpustakaanHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	buku := new(models.Perpustakaan)
	err := json.NewDecoder(r.Body).Decode(&buku)
	bukuResulth, err := s.PerpustakaanUseCase.GetUpdateBook(buku)
	if err != nil {
		w.Write([]byte("Data Not Found"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data Sukses di Ubah"))
	json.NewEncoder(w).Encode(bukuResulth)
}

func (s PerpustakaanHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := s.PerpustakaanUseCase.GetDeleteBook(vars["id_buku"])
	if err != nil {
		w.Write([]byte("Data tidak berhasil di hapus"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data Sukses di Hapus"))
}

func (s PerpustakaanHandler) FindBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := s.PerpustakaanUseCase.GetFindBookById(vars["id_buku"])
	if err != nil {
		w.Write([]byte("Data tidak di temukan"))
		log.Print(err)
		return
	}
	var response response.MyResponse = response.MyResponse{"Data Sukses di Tampikan", http.StatusOK, id}
	byteOfPerpustakaan, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfPerpustakaan)
}

func (s PerpustakaanHandler) FindBookByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	kategori, err := s.PerpustakaanUseCase.GetFindBookByCategory(vars["nama_kategori"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
		log.Print(err)
		return
	}
	var response response.MyResponse = response.MyResponse{"Data Sukses di Tampikan", http.StatusOK, kategori}
	byteOfPerpustakaan, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfPerpustakaan)
}

func (s PerpustakaanHandler) FindBookByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	judulBuku, err := s.PerpustakaanUseCase.GetFindBookByTitle(vars["judul_buku"])
	if err != nil {
		w.Write([]byte("Data tidak di temukan"))
		return
	}
	var response response.MyResponse = response.MyResponse{"Data Sukses di Tampikan", http.StatusOK, judulBuku}
	byteOfPerpustakaan, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfPerpustakaan)
}

func (s PerpustakaanHandler) FindBookByAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pengarang, err := s.PerpustakaanUseCase.GetFindBookByAuthor(vars["nama_pengarang"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
		return
	}
	var response response.MyResponse = response.MyResponse{"Data Sukses di Tampikan", http.StatusOK, pengarang}
	byteOfPerpustakaan, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfPerpustakaan)
}

func (s PerpustakaanHandler) FindBookByPublisher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	penerbit, err := s.PerpustakaanUseCase.GetFindBookByPublisher(vars["nama_penerbit"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
		return
	}
	var response response.MyResponse = response.MyResponse{"Data Sukses di Tampikan", http.StatusOK, penerbit}
	byteOfPerpustakaan, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfPerpustakaan)
}

func (s PerpustakaanHandler) TotalBook(w http.ResponseWriter, r *http.Request) {
	judulBuku, err := s.PerpustakaanUseCase.GetTotalBook()
	if err != nil {
		w.Write([]byte("Data tidak di temukan"))
		return
	}
	var response response.MyResponse = response.MyResponse{"Data Sukses di Tampikan", http.StatusOK, judulBuku}
	byteOfPerpustakaan, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfPerpustakaan)
}

func (s PerpustakaanHandler) TotalBookCategory(w http.ResponseWriter, r *http.Request) {
	judulBuku, err := s.PerpustakaanUseCase.GetTotalBookCategory()
	if err != nil {
		w.Write([]byte("Data tidak di temukan"))
		return
	}
	var response response.MyResponse = response.MyResponse{"Data Sukses di Tampikan", http.StatusOK, judulBuku}
	byteOfPerpustakaan, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfPerpustakaan)
}
