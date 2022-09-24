package main

import "fmt"

type Siswa struct {
	nama, pekerjaan, alamat, alasan string
}

func (siswa Siswa)  Kalimat()  {
	fmt.Println("Halo saya", siswa.nama)
	fmt.Println("saya", siswa.pekerjaan)
	fmt.Println("saya tinggal di", siswa.alamat)
	fmt.Println("alasan saya mengikuti course adalah", siswa.alasan)
}

func main() {
	pilih := 1

	switch pilih {
	case 1:
		udin := Siswa{
			nama:"udin", 
			pekerjaan:"pelajar", 
			alamat:"bogor", 
			alasan:"gabut"}
		udin.Kalimat()
		break
	case 2:
		rara := Siswa{
			nama : "rara",
			pekerjaan : "karyawan",
			alamat : "jakarta",
			alasan : "iseng",
		}
		rara.Kalimat()
		break
	case 3:
		dudung := Siswa{
			nama : "dudung",
			pekerjaan : "pengamen",
			alamat : "bandung",
			alasan : "mengadu nasib",
		}
		dudung.Kalimat()
		break
	}

}
