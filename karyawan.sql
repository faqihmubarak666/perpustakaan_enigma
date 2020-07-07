create database karyawan;
drop table karyawan_permanen;
create table karyawan_permanen (id varchar(40) primary key, nama varchar(40), alamat varchar(40),
ttl varchar(20), gapok int, tunjangan int, total_gaji int, status varchar(20));

create table karyawan_kontrak (id varchar(40) primary key, nama varchar(40), alamat varchar(40),
ttl varchar(20), gapok int, total_gaji int, status varchar(20));


insert into karyawan_permanen values
("KP001", "andika", "jl. camar", "jakarta, 12-02-1996", 1000, 500, 1500, "aktif"),
("KP002", "budi", "jl. bangau", "bali, 05-05-1996", 1000, 500, 1500,"aktif"),
("KP003", "annisa", "jl. perkutut", "cirebon, 20-10-1996", 1000, 500, 1500,"tidak aktif");
select * from karyawan_permanen;
select id, nama, alamat, ttl, total_gaji from karyawan_permanen;
select * from karyawan_permanen where status="aktif";



insert into karyawan_kontrak values
("KK001", "ryan", "jl. robusta", "jakarta, 12-02-1996", 1000, 1000,"aktif"),
("KK002", "lestari", "jl. arabika", "bali, 05-05-1996", 1000, 1000,"aktif"),
("KK003", "sarah", "jl. gayo", "cirebon, 20-10-1996", 1000,  1000,"tidak aktif");
select * from karyawan_kontrak;