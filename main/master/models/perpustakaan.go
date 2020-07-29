package models

type Perpustakaan struct {
	IdKategori    string `json:"id_kategori"`
	NamaKategori  string `json:"nama_kategori"`
	IdPengarang   string `json:"id_pengarang"`
	NamaPengarang string `json:"nama_pengarang"`
	IdPenerbit    string `json:"id_penerbit"`
	NamaPenerbit  string `json:"nama_penerbit"`
	IdBuku        string `json:"id_buku"`
	JudulBuku     string `json:"judul_buku"`
}

type AllBook struct {
	IdBuku        string `json:"id_buku"`
	JudulBuku     string `json:"judul_buku"`
	IdKategori    string `json:"id_kategori"`
	NamaKategori  string `json:"nama_kategori"`
	IdPengarang   string `json:"id_pengarang"`
	NamaPengarang string `json:"nama_pengarang"`
	IdPenerbit    string `json:"id_penerbit"`
	NamaPenerbit  string `json:"nama_penerbit"`
}

type ReportBook struct {
	TotalBook string `json:"total_buku"`
}

type ReportBookCategory struct {
	TotalBook    string `json:"total_buku"`
	NamaKategori string `json:"nama_kategori"`
}
