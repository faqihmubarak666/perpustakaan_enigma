package constantaQuery

const (
	//Perpustakaan
	GETALLBOOK             = "select buku.id_buku, buku.judul_buku, kategori.id_kategori, kategori.nama_kategori, pengarang.id_pengarang, pengarang.nama_pengarang, penerbit.id_penerbit, penerbit.nama_penerbit from buku join kategori on buku.id_kategori=kategori.id_kategori join pengarang on buku.id_pengarang=pengarang.id_pengarang join penerbit on buku.id_penerbit=penerbit.id_penerbit where buku.judul_buku like ? ORDER BY  ?, %s LIMIT  %s %s " 
	GETALLCATEGORY         = "select * from  kategori"
	GETALLAUTHOR           = "select * from pengarang"
	GETALLPUBLISHER        = "select * from penerbit"
	GETALLUSER             = "select * from user;"
	GETADDBOOK             = "insert into buku values(?, ?, ?, ?, ?)"
	GETUPDATEBOOK          = "update buku set judul_buku=?, id_kategori=?, id_pengarang=? where id_buku=?"
	GETDELETEBOOK          = "delete from buku where id_buku=?"
	GETFINDBOOKBYID        = "select buku.id_buku, buku.judul_buku, kategori.nama_kategori, pengarang.nama_pengarang, penerbit nama_penerbit from buku join kategori on buku.id_kategori=kategori.id_kategori join pengarang on buku id_pengarang=pengarang.id_pengarang join penerbit on buku.id_penerbit=penerbit.id_penerbit where buku.id_buku=?"
	GETFINDBOOKBYCATEGORY  = "select buku.id_buku, buku.judul_buku, kategori.nama_kategori, pengarang.nama_pengarang, penerbit.nama_penerbit from buku join kategori on buku.id_kategori=kategori.id_kategori join pengarang on buku.id_pengarang=pengarang.id_pengarang join penerbit on buku.id_penerbit=penerbit.id_penerbit where kategori.nama_kategori=?"
	GETFINDBOOKBYTITLE     = "select buku.id_buku, buku.judul_buku, kategori.nama_kategori, pengarang.nama_pengarang, penerbit.nama_penerbit from buku join kategori on buku.id_kategori=kategori.id_kategori join pengarang on buku.id_pengarang=pengarang.id_pengarang join penerbit on buku.id_penerbit=penerbit.id_penerbit where buku.judul_buku=?"
	GETFINDBOOKBYAUTHOR    = "select buku.id_buku, buku.judul_buku, kategori.nama_kategori, pengarang.nama_pengarang, penerbit.nama_penerbit from buku join kategori on buku.id_kategori=kategori.id_kategori join pengarang on buku.id_pengarang=pengarang.id_pengarang join penerbit on buku.id_penerbit=penerbit.id_penerbit where pengarang.nama_pengarang=?"
	GETFINDBOOKBYPUBLISHER = "select buku.id_buku, buku.judul_buku, kategori.nama_kategori, pengarang.nama_pengarang, penerbit.nama_penerbit from buku join kategori on buku.id_kategori=kategori.id_kategori join pengarang on buku.id_pengarang=pengarang.id_pengarang join penerbit on buku.id_penerbit=penerbit.id_penerbit where penerbit.nama_penerbit=?"
	GETTOTALBOOK           = "select count(buku.judul_buku) as total_buku from buku"
	GETTOTALBOOKCATEGORY   = "select count(buku.judul_buku) as jumlah_buku, kategori.nama_kategori from buku join kategori on buku.id_kategori=kategori.id_kategori group by nama_kategori;"

	//USER
	GETUSER = "select username, password from user where username=? and password =?"
)
