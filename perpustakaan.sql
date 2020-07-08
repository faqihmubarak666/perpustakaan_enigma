create database perpustakaan_enigma;
drop table fk;
create table kategori (id_kategori varchar(40) primary key, nama_kategori varchar(40));
insert into kategori values("Kat01", "ekonomi"), ("Kat02", "telekomunikasi"), ("Kat03", "geografi");

create table buku (id_buku varchar(40) primary key, judul_buku varchar(40), id_kategori varchar(40),
id_pengarang varchar(40), id_penerbit varchar(40),
FOREIGN KEY(id_kategori) REFERENCES kategori(id_kategori),
FOREIGN KEY(id_pengarang) REFERENCES pengarang(id_pengarang),
FOREIGN KEY(id_penerbit) REFERENCES penerbit(id_penerbit));
insert into buku values
("Buk01", "analisa keuangan", "Kat01", "Peng01", "Bit01"),
("Buk02", "pertumbuhan ekonomi", "Kat01", "Peng01", "Bit01"), 
("Buk03", "mengatur keuangan", "Kat01", "Peng01", "Bit01"),
("Buk04", "teknologi komputer", "Kat02", "Peng02", "Bit02"),
("Buk05", "jaringan komputer", "Kat02", "Peng02", "Bit02"),
("Buk06", "Software komputer", "Kat02", "Peng02", "Bit02"),
("Buk07", "lempengan bumi", "Kat03", "Peng03", "Bit03"),
("Buk08", "analisa cuacah", "Kat03", "Peng03", "Bit03"),
("Buk09", "pergerakan angin", "Kat03", "Peng03", "Bit03");
select * from buku;
select buku.id_buku, buku.judul_buku, kategori.nama_kategori, pengarang.nama_pengarang, penerbit.nama_penerbit 
from buku join kategori on 
buku.id_kategori=kategori.id_kategori join pengarang on buku.id_pengarang=pengarang.id_pengarang
join penerbit on buku.id_penerbit=penerbit.id_penerbit;

select buku.id_buku, buku.judul_buku, kategori.nama_kategori, pengarang.nama_pengarang, penerbit.nama_penerbit 
from buku join kategori on 
buku.id_kategori=kategori.id_kategori join pengarang on buku.id_pengarang=pengarang.id_pengarang
join penerbit on buku.id_penerbit=penerbit.id_penerbit where buku.id_buku="Buk01";

select buku.id_buku, buku.judul_buku, kategori.nama_kategori, pengarang.nama_pengarang, penerbit.nama_penerbit from buku join kategori on 
buku.id_kategori=kategori.id_kategori join pengarang on buku.id_pengarang=pengarang.id_pengarang
join penerbit on buku.id_penerbit=penerbit.id_penerbit where kategori.nama_kategori="ekonomi";

select buku.id_buku, buku.judul_buku, kategori.nama_kategori, pengarang.nama_pengarang, penerbit.nama_penerbit from buku join kategori on 
buku.id_kategori=kategori.id_kategori join pengarang on buku.id_pengarang=pengarang.id_pengarang
join penerbit on buku.id_penerbit=penerbit.id_penerbit where buku.judul_buku="analisa keuangan";

select buku.id_buku, buku.judul_buku, kategori.nama_kategori, pengarang.nama_pengarang, penerbit.nama_penerbit from buku join kategori on 
buku.id_kategori=kategori.id_kategori join pengarang on buku.id_pengarang=pengarang.id_pengarang
join penerbit on buku.id_penerbit=penerbit.id_penerbit where pengarang.nama_pengarang="aldi fauzi";

select buku.id_buku, buku.judul_buku, kategori.nama_kategori, pengarang.nama_pengarang, penerbit.nama_penerbit from buku join kategori on 
buku.id_kategori=kategori.id_kategori join pengarang on buku.id_pengarang=pengarang.id_pengarang
join penerbit on buku.id_penerbit=penerbit.id_penerbit where penerbit.nama_penerbit="sinar gemilang";



create table pengarang(id_pengarang varchar(40) primary key, nama_pengarang varchar(40));
insert into pengarang values("Peng01", "aldi fauzi"), ("Peng02", "ryan setiadi"), ("Peng03", "nadia ulfa");

create table penerbit (id_penerbit varchar(40) primary key, nama_penerbit varchar(40));
insert into penerbit values("Bit01", "sinar gemilang"), ("Bit02", "cipta agung"), ("Bit03", "sinar cahaya");

select count(buku.judul_buku) as jumlah_buku, kategori.nama_kategori from buku join kategori on 
buku.id_kategori=kategori.id_kategori group by nama_kategori;

select count(buku.judul_buku) as total_buku from buku;