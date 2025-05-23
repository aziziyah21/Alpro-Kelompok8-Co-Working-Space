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
	// Data dummy
	coworkList = append(coworkList, CoWorkingSpace{
		Nama:      "Work Happy",
		Lokasi:    "Purwokerto",
		Fasilitas: []string{"WiFi", "Meeting Room", "Coffee", "Printer"},
		Harga:     50000,
		Rating:    4.5,
		Ulasan:    "Tempat nyaman dan tenang.",
	})

	coworkList = append(coworkList, CoWorkingSpace{
		Nama:      "Happy Life",
		Lokasi:    "Bandung",
		Fasilitas: []string{"WiFi", "Snack Bar", "AC", "Ruang Diskusi"},
		Harga:     60000,
		Rating:    4.7,
		Ulasan:    "Sangat cocok untuk kerja tim.",
	})

	coworkList = append(coworkList, CoWorkingSpace{
		Nama:      "Work band",
		Lokasi:    "Bali",
		Fasilitas: []string{"Coffe", "Snack Bar", "AC", "Ruang Diskusi"},
		Harga:     65000,
		Rating:    4.9,
		Ulasan:    "Bagus dan sangat nyaman.",
	})

	coworkList = append(coworkList, CoWorkingSpace{
		Nama:      "Work good",
		Lokasi:    "Tegal",
		Fasilitas: []string{"WiFi", "Snack Bar", "AC", "coffe"},
		Harga:     40000,
		Rating:    4.6,
		Ulasan:    "Sangat cocok untuk kerja tim.",
	})

	for {
		fmt.Println("MENU COWORKING SPACE")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Ubah Data")
		fmt.Println("3. Ubah Fasilitas")
		fmt.Println("4. Ubah Ulasan")
		fmt.Println("5. Ubah Rating")
		fmt.Println("6. Hapus Data")
		fmt.Println("7. Hapus Ulasan")
		fmt.Println("8. Cari Lokasi")
		fmt.Println("9. Urutan Data(bedasarkan rating)")
		fmt.Println("10. Tampilkan Fasilitas")
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
			cariLokasi()
		case 9:
			urutanData()
		case 10:
			Fasilitas()
		case 11:
			tampilkanSemua()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
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
	fmt.Print("Pilih nomer data yang ingin diubah:  ")
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
	for i := n; i < len(coworkList)-1; i++ {
		coworkList[i] = coworkList[i+1]
	}
	coworkList = coworkList[:len(coworkList)-1]
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
	coworkList[n-1].Ulasan = " "
	fmt.Println("Ulasan berhasil dihapus.")
}
func cariLokasi() {
	fmt.Print("Cari lokasi: ")
	var lokasi string
	fmt.Scan(&lokasi)

	ada := false

	for i := 0; i < len(coworkList); i++ {
		c := coworkList[i]
		if c.Lokasi == lokasi {
			if !ada {
				fmt.Println("Pencarian lokasi:")
				ada = true
			}
			fmt.Printf("Data %d\n", i+1)
			fmt.Printf("Nama    : %s\n", c.Nama)
			fmt.Printf("Lokasi  : %s\n", c.Lokasi)
			fmt.Printf("Harga   : Rp %d\n", c.Harga)
			fmt.Printf("Rating  : %.1f\n", c.Rating)
			fmt.Println("Fasilitas:")
			for j := 0; j < len(c.Fasilitas); j++ {
				fmt.Printf("  %d. %s\n", j+1, c.Fasilitas[j])
			}
			fmt.Printf("Ulasan  : %s\n", c.Ulasan)
		}
	}

	if !ada {
		fmt.Println("Nama tidak ditemukan")
	}
}
func urutanData() {
	for i := 0; i < len(coworkList)-1; i++ {
		max := i
		for j := i + 1; j < len(coworkList); j++ {
			if coworkList[j].Rating > coworkList[max].Rating {
				max = j
			}
		}
		if max != i {
			coworkList[i], coworkList[max] = coworkList[max], coworkList[i]
		}
	}

	fmt.Println("Data sudah diurutkan berdasarkan rating:")
	tampilkanSemua()
}
func tampilkanSemua() {
	for i := 0; i < len(coworkList); i++ {
		c := coworkList[i]

		fmt.Printf("Data ke  %d\n", i+1)
		fmt.Println("+----------------+--------------------------------------+")
		fmt.Printf("| Nama           | %-36s |\n", c.Nama)
		fmt.Println("+----------------+--------------------------------------+")
		fmt.Printf("| Lokasi         | %-36s |\n", c.Lokasi)
		fmt.Println("+----------------+--------------------------------------+")
		fmt.Printf("| Harga          | Rp %-33d |\n", c.Harga)
		fmt.Println("+----------------+--------------------------------------+")
		fmt.Printf("| Rating         | %-36.1f |\n", c.Rating)
		fmt.Println("+----------------+--------------------------------------+")

		fmt.Print("| Fasilitas ")
		if len(c.Fasilitas) > 0 {
			fmt.Printf("%d. %-40s |\n", 1, c.Fasilitas[0])
		} else {
			fmt.Printf("%-40s |\n", "")
		}
		for j := 1; j < len(c.Fasilitas); j++ {
			fmt.Printf("|           %d. %-40s |\n", j+1, c.Fasilitas[j])
		}
		fmt.Println("+----------------+--------------------------------------+")

		fmt.Printf("| Ulasan         | %-36s |\n", c.Ulasan)
		fmt.Println("+----------------+--------------------------------------+")
	}
}

func Fasilitas() {
	for i := 0; i < len(coworkList); i++ {
		c := coworkList[i]
		for j := 0; j < len(c.Fasilitas); j++ {
			fmt.Printf("  %d. %s\n", j+1, c.Fasilitas[j])
		}
	}
}
