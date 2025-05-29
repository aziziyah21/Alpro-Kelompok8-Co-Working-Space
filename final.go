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
		fmt.Println("MENU COWORKING SPACE")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Ubah Data")
		fmt.Println("3. Ubah Fasilitas")
		fmt.Println("4. Ubah Ulasan")
		fmt.Println("5. Ubah Rating")
		fmt.Println("6. Hapus Data")
		fmt.Println("7. Hapus Ulasan")
		fmt.Println("8. Cari Nama")
		fmt.Println("9. Cari Lokasi")
		fmt.Println("10. Urutan Data (berdasarkan rating)")
		fmt.Println("11. Urutan Data (berdasarkan harga)")
		fmt.Println("12. Tampilkan Fasilitas")
		fmt.Println("13. Tampilkan Semua")
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
			urutanRating()
		case 11:
			urutanHarga()
		case 12:
			Fasilitas()
		case 13:
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
	fmt.Scanln()
	input := ""
	for {
		var line string
		n, _ := fmt.Scanln(&line)
		if n == 0 {
			break
		}
		input += line + "  "
	}
	ulasan = input

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
		fmt.Printf("Fasilitas %d (ketik 'STOP' untuk berhenti): ", i+1)
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
	fmt.Print("Pilih nomor data yang ingin diubah ulasannya: ")
	var n int
	fmt.Scan(&n)

	if n < 1 || n > len(coworkList) {
		fmt.Println("Data tidak ditemukan.")
		return
	}
	n--

	fmt.Print("Masukkan ulasan baru : ")
	fmt.Scanln()
	var ulasanBaru string
	input := ""
	for {
		var line string
		n, _ := fmt.Scanln(&line)
		if n == 0 {
			break
		}
		input += line + "  "
	}
	ulasanBaru = input
	coworkList[n].Ulasan = ulasanBaru
	fmt.Println("Ulasan berhasil diubah.")
	tampilkanSemua()
}

func hapusData() {
	tampilkanSemua()
	fmt.Print("Pilih nomer data yang ingin dihapus: ")
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
	coworkList[n-1].Ulasan = ""
	fmt.Println("Ulasan berhasil dihapus.")
}
func cariNama() {
	fmt.Print("Masukkan nama yang ingin dicari: ")
	var nama string
	fmt.Scan(&nama)

	// baca satu baris lengkap (termasuk spasi)

	ditemukan := false
	for i := 0; i < len(coworkList); i++ {
		if coworkList[i].Nama == nama { // sequential
			fmt.Println("Data ditemukan:")
			fmt.Printf("Data ke-%d\n", i+1)
			fmt.Printf("Nama    : %s\n", coworkList[i].Nama)
			fmt.Printf("Lokasi  : %s\n", coworkList[i].Lokasi)
			fmt.Printf("Harga   : Rp %d\n", coworkList[i].Harga)
			fmt.Printf("Rating  : %.1f\n", coworkList[i].Rating)
			fmt.Printf("Ulasan  : %s\n", coworkList[i].Ulasan)
			fmt.Println("Fasilitas:")
			for j := 0; j < len(coworkList[i].Fasilitas); j++ {
				fmt.Printf("  %d. %s\n", j+1, coworkList[i].Fasilitas[j])
			}
			ditemukan = true
			break
		}
	}

	if !ditemukan {
		fmt.Println("Nama tidak ditemukan.")
	}
}
func cariLokasi() {
	if len(coworkList) == 0 {
		fmt.Println("Data kosong.")
		return
	}

	// Selection sort berdasarkan Lokasi
	for i := 0; i < len(coworkList)-1; i++ {
		min := i
		for j := i + 1; j < len(coworkList); j++ {
			if coworkList[j].Lokasi < coworkList[min].Lokasi {
				min = j
			}
		}
		if min != i {
			coworkList[i], coworkList[min] = coworkList[min], coworkList[i]
		}
	}

	fmt.Print("Masukkan lokasi yang ingin dicari: ")
	var lokasi string
	fmt.Scan(&lokasi)

	// Binary search lokasi
	low, high := 0, len(coworkList)-1
	ditemukan := false

	for low <= high {
		mid := (low + high) / 2
		if coworkList[mid].Lokasi == lokasi {
			fmt.Println("Data ditemukan:")
			fmt.Printf("Nama    : %s\n", coworkList[mid].Nama)
			fmt.Printf("Lokasi  : %s\n", coworkList[mid].Lokasi)
			fmt.Printf("Harga   : Rp %d\n", coworkList[mid].Harga)
			fmt.Printf("Rating  : %.1f\n", coworkList[mid].Rating)
			fmt.Printf("Ulasan  : %s\n", coworkList[mid].Ulasan)

			fmt.Println("Fasilitas:")
			for j := 0; j < len(coworkList[mid].Fasilitas); j++ {
				fmt.Printf("  %d. %s\n", j+1, coworkList[mid].Fasilitas[j])
			}
			ditemukan = true
			break
		} else if coworkList[mid].Lokasi < lokasi {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !ditemukan {
		fmt.Println("Lokasi tidak ditemukan.")
	}
}

func urutanRating() { //slection shor
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
func urutanHarga() { //inster shorting
	for i := 1; i < len(coworkList); i++ {
		key := coworkList[i]
		j := i - 1
		for j >= 0 && coworkList[j].Harga > key.Harga {
			coworkList[j+1] = coworkList[j]
			j--
		}
		coworkList[j+1] = key
	}
	fmt.Println("Data sudah diurutkan berdasarkan harga (insertion sort):")
	tampilkanSemua()
}

func tampilkanSemua() {
	for i, c := range coworkList {
		fmt.Printf("Data ke %d\n", i+1)
		fmt.Println("---------------------------------------------------------")
		fmt.Printf("| Nama           | %-36s |\n", c.Nama)
		fmt.Println("---------------------------------------------------------")
		fmt.Printf("| Lokasi         | %-36s |\n", c.Lokasi)
		fmt.Println("---------------------------------------------------------")
		fmt.Printf("| Harga          | Rp %-33d |\n", c.Harga)
		fmt.Println("---------------------------------------------------------")
		fmt.Printf("| Rating         | %-36.1f |\n", c.Rating)
		fmt.Println("---------------------------------------------------------")
		fmt.Print("| Fasilitas ")
		if len(c.Fasilitas) > 0 {
			fmt.Printf("%d. %-40s |\n", 1, c.Fasilitas[0])
			for j := 1; j < len(c.Fasilitas); j++ {
				fmt.Printf("|           %d. %-40s |\n", j+1, c.Fasilitas[j])
			}
		} else {
			fmt.Printf("%-40s |\n", "")
		}
		fmt.Println("---------------------------------------------------------")
		text := c.Ulasan
		width := 36
		start := 0
		line := 0
		for start < len(text) {
			end := start + width
			if end > len(text) {
				end = len(text)
			}
			sub := text[start:end]
			if line == 0 {
				fmt.Printf("| Ulasan         | %-36s |\n", sub)
			} else {
				fmt.Printf("| %-15s | %-36s |\n", "", sub)
			}
			start = end
			line++
		}
		fmt.Println("---------------------------------------------------------")
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
