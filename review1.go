package main

import "fmt"

type CoWorkingSpace struct {
	Nama      string
	Lokasi    string
	Fasilitas []string
	Harga     int
	Rating    float64
	Ulasan    string
}

var coworkList []CoWorkingSpace

func main() {
	for {
		fmt.Println("\n==== MENU COWORKING SPACE ====")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Ubah Data")
		fmt.Println("3. Ubah Fasilitas")
		fmt.Println("4. Ubah Ulasan")
		fmt.Println("5. Ubah Rating")
		fmt.Println("6. Hapus Data")
		fmt.Println("7. Hapus Ulasan")
		fmt.Println("8. Cari Nama")
		fmt.Println("9. Cari Lokasi")
		fmt.Println("10. Urutan Data")
		fmt.Println("11. Tampilkan Semua")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")
		var pilih int
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahData()
		case 2:
			ubahData()
		case 3:
			ubahFasilitas()
		case 4:
			ubahUlasan()
		case 5:
			ubahRating()
		case 6:
			hapusData()
		case 7:
			hapusUlasan()
		case 8:
			cariNama()
		case 9:
			cariLokasi()
		case 10:
			urutanData()
		case 11:
			tampilkanSemua()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tampilkanSemua() {
	if len(coworkList) == 0 {
		fmt.Println("Belum ada data")
		return
	}
	for i := 0; i < len(coworkList); i++ {
		s := coworkList[i]
		fmt.Printf("Data %d\n", i+1)
		fmt.Printf("Nama    : %s\n", s.Nama)
		fmt.Printf("Lokasi  : %s\n", s.Lokasi)
		fmt.Printf("Harga   : Rp %d\n", s.Harga)
		fmt.Printf("Rating  :  %.1f\n", s.Rating)
		fmt.Println("Fasilitas:")
		for j, f := range s.Fasilitas {
			fmt.Printf("  %d. %s\n", j+1, f)
		}
		fmt.Printf("Ulasan  : %s\n", s.Ulasan)
	}
}

func tambahData() {
	var nama, lokasi, ulasan string
	var harga int
	var rating float64
	var namaFasilitas []string

	fmt.Print("Nama: ")
	fmt.Scan(&nama)

	fmt.Print("Lokasi: ")
	fmt.Scan(&lokasi)

	fmt.Println("Masukkan fasilitas (ketik 'STOP' untuk berhenti):")
	for {
		var fasilitas string
		fmt.Print("Fasilitas: ")
		fmt.Scan(&fasilitas)
		if fasilitas == "STOP" {
			break
		}
		namaFasilitas = append(namaFasilitas, fasilitas)
	}

	fmt.Print("Harga: ")
	fmt.Scan(&harga)

	fmt.Print("Rating: ")
	fmt.Scan(&rating)

	fmt.Print("Ulasan: ")
	fmt.Scan(&ulasan)

	coworkList = append(coworkList, CoWorkingSpace{
		Nama:      nama,
		Lokasi:    lokasi,
		Fasilitas: namaFasilitas,
		Harga:     harga,
		Rating:    rating,
		Ulasan:    ulasan,
	})

	fmt.Println("Data berhasil ditambahkan.")
	tampilkanSemua()
}

func ubahData() {
	tampilkanSemua()
	fmt.Print("Pilih nomer data yang ingin diubah: ")
	var n int
	fmt.Scan(&n)
	if n < 1 || n > len(coworkList) {
		fmt.Println("Data tidak ditemukan")
		return
	}
	n--

	var nama, lokasi string
	var harga int

	fmt.Print("Nama baru: ")
	fmt.Scan(&nama)
	fmt.Print("Lokasi baru: ")
	fmt.Scan(&lokasi)
	fmt.Print("Harga baru: ")
	fmt.Scan(&harga)

	coworkList[n].Nama = nama
	coworkList[n].Lokasi = lokasi
	coworkList[n].Harga = harga

	fmt.Println("Data berhasil diubah")
	tampilkanSemua()
}

func ubahFasilitas() {
	tampilkanSemua()
	fmt.Print("Pilih nomer data yang ingin diubah: ")
	var n int
	fmt.Scan(&n)
	if n < 1 || n > len(coworkList) {
		fmt.Println("Data tidak ditemukan")
		return
	}
	n--

	var fasilitas string
	var namaFasilitas [10]string
	var jumlah int

	for i := 0; i < 10; i++ {
		fmt.Printf("Fasilitas %d (kosongkan untuk berhenti): ", i+1)
		fmt.Scan(&fasilitas)
		if fasilitas == "STOP" {
			break
		}
		namaFasilitas[i] = fasilitas
		jumlah++
	}

	coworkList[n].Fasilitas = namaFasilitas[:jumlah]
	fmt.Println("Fasilitas berhasil diubah")
}

func ubahRating() {
	tampilkanSemua()
	fmt.Print("Pilih nomer data yang ingin diubah: ")
	var n int
	fmt.Scan(&n)
	if n < 1 || n > len(coworkList) {
		fmt.Println("Data tidak ditemukan")
		return
	}
	n--

	fmt.Print("Masukkan rating baru: ")
	var rating float64
	fmt.Scan(&rating)

	coworkList[n].Rating = rating
	fmt.Println("Rating berhasil diubah")
}

func ubahUlasan() {
	tampilkanSemua()
	fmt.Print("Pilih nomer data yang ingin diubah: ")
	var n int
	fmt.Scan(&n)
	if n < 1 || n > len(coworkList) {
		fmt.Println("Data tidak ditemukan")
		return
	}
	n--

	fmt.Print("Masukkan ulasan baru: ")
	var ulasan string
	fmt.Scan(&ulasan)

	coworkList[n].Ulasan = ulasan
	fmt.Println("Ulasan berhasil diubah.")
}

func hapusData() {
	tampilkanSemua()
	fmt.Print("Pilih nomer data yang ingin diubah: ")
	var n int
	fmt.Scan(&n)
	if n < 1 || n > len(coworkList) {
		fmt.Println("Data tidak ditemukan")
		return
	}
	n--
	coworkList = append(coworkList[:n], coworkList[n+1:]...)
	fmt.Println("Data berhasil dihapus")
	tampilkanSemua()
}

func hapusUlasan() {
	tampilkanSemua()
	fmt.Print("Pilih nomer data yang ingin diubah: ")
	var n int
	fmt.Scan(&n)
	if n < 1 || n > len(coworkList) {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	coworkList[n-1].Ulasan = ""
	fmt.Println("Ulasan berhasil dihapus.")
}

func cariNama() {
	if len(coworkList) == 0 {
		fmt.Println("Data tidak ada")
		return
	}
	fmt.Print("Nama yang dicari: ")
	var nama string
	fmt.Scan(&nama)

	ada := false

	for i, s := range coworkList {
		if s.Nama == nama {
			if !ada {
				fmt.Println("Pencarian nama:")
				ada = true
			}
			fmt.Printf("Data %d\n", i+1)
			fmt.Printf("Nama    : %s\n", s.Nama)
			fmt.Printf("Lokasi  : %s\n", s.Lokasi)
			fmt.Printf("Harga   : Rp %d\n", s.Harga)
			fmt.Printf("Rating  : %.1f\n", s.Rating)
			fmt.Println("Fasilitas:")
			for j, f := range s.Fasilitas {
				fmt.Printf("  %d. %s\n", j+1, f)
			}
			fmt.Printf("Ulasan  : %s\n\n", s.Ulasan)
		}
	}

	if !ada {
		fmt.Println("Nama tidak ditemukan")
	}
}

func cariLokasi() {
	if len(coworkList) == 0 {
		fmt.Println("Tidak ada data yang dicari.")
		return
	}
	fmt.Print("Lokasi yang dicari: ")
	var lokasiCari string
	fmt.Scan(&lokasiCari)

	ada := false

	for i, s := range coworkList {
		if s.Lokasi == lokasiCari {
			if !ada {
				fmt.Printf("Pencarian lokasi:\n")
				ada = true
			}
			fmt.Printf("Data %d\n", i+1)
			fmt.Printf("Nama    : %s\n", s.Nama)
			fmt.Printf("Lokasi  : %s\n", s.Lokasi)
			fmt.Printf("Harga   : Rp %d\n", s.Harga)
			fmt.Printf("Rating  : %.1f\n", s.Rating)
			fmt.Printf("Fasilitas:\n")
			for j, f := range s.Fasilitas {
				fmt.Printf("  %d. %s\n", j+1, f)
			}
			fmt.Printf("Ulasan  : %s\n\n", s.Ulasan)
		}
	}

	if !ada {
		fmt.Printf("Lokasi tidak ditemukan.\n")
	}
}

func urutanData() {
	if len(coworkList) == 0 {
		fmt.Println("Tidak ada data yang diurutkan")
		return
	}

	fmt.Println("\n==== MENU COWORKING SPACE ====")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Ubah Data")
	fmt.Println("3. Ubah Fasilitas")
	fmt.Println("4. Ubah Ulasan")
	fmt.Println("5. Ubah Rating")
	fmt.Println("6. Hapus Data")
	fmt.Println("7. Hapus Ulasan")
	fmt.Println("8. Cari Nama")
	fmt.Println("9. Cari Lokasi")
	fmt.Println("10. Urutan Data")
	fmt.Println("11. Tampilkan Semua")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih: ")

	var pilih int
	fmt.Scan(&pilih)

	fmt.Println("Data sesudah diurutkan:")
	tampilkanSemua()
}
