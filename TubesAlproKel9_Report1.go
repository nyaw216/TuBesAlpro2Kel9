package main

import "fmt"

const NMAX int = 100

type konten struct {
	id                       int
	nama, kategori, platform string
	waktu                    waktu
	engagement               engagement
}
type waktu struct {
	jam, hari, bulan, tahun, converted int
}
type engagement struct {
	like, comment, share int
}
type data [NMAX]konten

func menu(pilihan *int) {
	fmt.Println()
	fmt.Println("               || Selamat Datang di PostHub! ||")
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("          °❀⋆.ೃ࿔*:･°❀⋆.ೃ࿔*:･ MENU °❀⋆.ೃ࿔*:･°❀⋆.ೃ࿔*:･")
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("1. Manage PostHub")
	fmt.Println("2. Exit app")
	fmt.Println("---------------------------------------------------------------")
	loopBack(*&pilihan, 1, 2)
}
func outputData(A data, n int) {
	var i int
	fmt.Println()
	if n != 0 {
		fmt.Println("---------------------------------------------------------------------------------------------------")
		fmt.Println("No |            Judul         | Kategori | Platform | Jam | Tanggal Post | Like | Comment | Share |")
		fmt.Println("---------------------------------------------------------------------------------------------------")
		for i = 0; i < n; i++ {
			fmt.Printf("%-3v %-26v %-10v %-10v %-5v %v/%v/%-9v %-7v %-8v %-7v\n", A[i].id, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam, A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like, A[i].engagement.comment, A[i].engagement.share)
		}
		fmt.Println("---------------------------------------------------------------------------------------------------")

	} else {
		fmt.Println("---------------------------------------------------------------")
		fmt.Println("                       ⌕  Data Kosong")
		fmt.Println("---------------------------------------------------------------")
	}
	fmt.Println()
}
func daftarKonten(pilihan *int, A data, n *int) {
	outputData(A, *n)
	fmt.Println("1. Tambahkan konten")
	fmt.Println("2. Edit data konten")
	fmt.Println("3. Filter konten")
	fmt.Println("4. Sortir konten")
	fmt.Println("5. Kembali ke menu")
	fmt.Print("╰┈➤ ")
	fmt.Scan(pilihan)

	for (*pilihan > 5 || *pilihan < 1) || (*n == 0 && (*pilihan == 2 || *pilihan == 3 || *pilihan == 4)) {
		if *n == 0 && (*pilihan == 2 || *pilihan == 3 || *pilihan == 4) {
			fmt.Println("Belum ada data untuk melakukan perintah.")
			fmt.Print("╰┈➤ ")
			fmt.Scan(pilihan)
		} else if *pilihan > 5 || *pilihan < 1 {
			fmt.Println("Pilihan invalid.")
			fmt.Print("╰┈➤ ")
			fmt.Scan(pilihan)
		}
	}
}
func inputKonten(A *data, n *int) {
	var i int = 0
	var namaKonten string
	fmt.Println()
	fmt.Println("Format inputan:")
	fmt.Println("[ Judul Kategori Platform Jam DD/MM/YYYY Like Comment Share ] (-)")
	fmt.Println()
	if A[0].nama == "" {
		fmt.Scan(&namaKonten)
	} else if A[0].nama != "" {
		for A[i].nama != "" {
			i++
		}
		fmt.Scan(&namaKonten)
	}
	for namaKonten != "-" {
		A[i].nama = namaKonten
		fmt.Scanf("%v %v %v %v/%v/%v %v %v %v", &A[i].kategori, &A[i].platform, &A[i].waktu.jam, &A[i].waktu.hari, &A[i].waktu.bulan, &A[i].waktu.tahun, &A[i].engagement.like, &A[i].engagement.comment, &A[i].engagement.share)
		A[i].id = i + 1
		A[i].waktu.converted = timeConvert(*A, i)
		*n++
		i++
		fmt.Scan(&namaKonten)
	}
}
func editData(pilihan *int, A *data, n *int) {
	outputData(*A, *n)
	fmt.Println("1. Edit data")
	fmt.Println("2. Hapus data")
	fmt.Println("3. Kembali")
	loopBack(*&pilihan, 1, 3)
}
func replaceData(A *data, n int) {
	var noData, idx int
	var col string
	outputData(*A, n)
	fmt.Println("Format inputan:")
	fmt.Println("[ Nomor_data Kolom ] (Back: 0 0)")
	fmt.Scan(&noData, &col)
	idx = dataSearch(*A, n, noData)
	if idx != -1 {
		fmt.Print(col, " baru: ")
		switch col {
		case "Nama", "nama":
			fmt.Scan(&A[idx].nama)
		case "Kategori", "kategori":
			fmt.Scan(&A[idx].kategori)
		case "Platform", "platform":
			fmt.Scan(&A[idx].platform)
		case "Waktu", "waktu":
			fmt.Scanf("%v %v/%v/%v", &A[idx].waktu.jam, &A[idx].waktu.hari, &A[idx].waktu.bulan, &A[idx].waktu.tahun)
		case "Engagement", "engagement":
			fmt.Scan(&A[idx].engagement.like, &A[idx].engagement.comment, &A[idx].engagement.share)

		default:
			fmt.Println("Kolom invalid")
		}
	} else if noData == 0 && col == "0" {
	} else {
		if noData == 0 {
		} else {
			fmt.Println("Data tidak ditemukan")
		}
	}

}
func filter(pilihan *int, A *data, n int, F data, nFilter int) {
	var pilKat, pilPlat string
	fmt.Print("Masukkan kategori atau platform ( - ): ")
	fmt.Scan(&pilKat, &pilPlat)
	if pilKat == "-" && pilPlat == "-" {
	} else {
		var i int
		fmt.Println()
		fmt.Println("---------------------------------------------------------------------------------------------------")
		fmt.Println("No |            Judul         | Kategori | Platform | Jam | Tanggal Post | Like | Comment | Share |")
		fmt.Println("---------------------------------------------------------------------------------------------------")

		for i = 0; i < n; i++ {
			if pilPlat == "-" {
				if A[i].kategori == pilKat {
					fmt.Printf("%-3v %-26v %-10v %-10v %-5v %v/%v/%-9v %-7v %-8v %-7v\n", A[i].id, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam, A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like, A[i].engagement.comment, A[i].engagement.share)
					F[nFilter] = A[i]
					nFilter++
				}
			} else if pilKat == "-" {
				if A[i].platform == pilPlat {
					fmt.Printf("%-3v %-26v %-10v %-10v %-5v %v/%v/%-9v %-7v %-8v %-7v\n", A[i].id, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam, A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like, A[i].engagement.comment, A[i].engagement.share)
					F[nFilter] = A[i]
					nFilter++
				}
			} else {
				if A[i].platform == pilPlat && A[i].kategori == pilKat {
					fmt.Printf("%-3v %-26v %-10v %-10v %-5v %v/%v/%-9v %-7v %-8v %-7v\n", A[i].id, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam, A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like, A[i].engagement.comment, A[i].engagement.share)
					F[nFilter] = A[i]
					nFilter++
				}
			}
		}

		if nFilter == 0 {
			fmt.Println("                                      ⌕  Data tidak ditemukan")
		} else {
			fmt.Println("---------------------------------------------------------------------------------------------------")
			fmt.Println()
			fmt.Println("1. Sortir konten")
			fmt.Println("2. Kembali")
			loopBack(*&pilihan, 1, 2)
			if *pilihan == 1 {
				sort(F, nFilter, pilihan)
			}
		}
	}
}
func deleteData(A *data, n *int) {
	var idx, noData int
	var i int
	fmt.Print("Nomor data (Back 0): ")
	fmt.Scan(&noData)
	idx = dataSearch(*A, *n, noData)
	if noData != 0 && idx != -1 {
		for i = idx; i < *n; i++ {
			A[i] = A[i+1]
			A[i].id--
		}
		*n--
		A[i].nama = ""
	} else if noData != 0 && idx == -1 {
		fmt.Println("Data tidak ditemukan.")
	} else {

	}
}

func dataSearch(A data, n int, find int) int {
	var left, right, mid, idx int
	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if find < A[mid].id {
			right = mid - 1
		} else if find > A[mid].id {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}
func loopBack(pilihan *int, left, right int) {
	fmt.Print("╰┈➤ ")
	fmt.Scan(pilihan)
	for *pilihan < left || *pilihan > right {
		fmt.Println("Pilihan invalid.")
		fmt.Print("Pilihan kamu: ")
		fmt.Print("╰┈➤ ")
		fmt.Scan(pilihan)
	}
}
func sort(F data, n int, pilihan *int) {
	var whatSort string
	fmt.Println("1. Ascending (Membesar)")
	fmt.Println("2. Descending (Mengecil)")
	fmt.Println("[ Pilihan_sort Berdasarkan ] (Back: 0 0)")
	fmt.Scan(pilihan, &whatSort)
	if *pilihan == 0 && whatSort == "0" {
	} else {
		sortProcess(&F, n, *&pilihan, whatSort)
		outputData(F, n)
		fmt.Println("1. Kembali")
		fmt.Print("╰┈➤ ")
		fmt.Scan(pilihan)
	}
}
func sortProcess(F *data, n int, pilihan *int, x string) {
	var pass, idx, i int
	var temp konten
	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for i = pass + 1; i < n; i++ {
			switch x {
			case "Waktu", "waktu":
				switch *pilihan {
				case 1:
					if F[idx].waktu.converted > F[i].waktu.converted {
						idx = i
					}
				case 2:
					if F[idx].waktu.converted < F[i].waktu.converted {
						idx = i
					}
				}
			case "Like", "like":
				switch *pilihan {
				case 1:
					if F[idx].engagement.like > F[i].engagement.like {
						idx = i
					}
				case 2:
					if F[idx].engagement.like < F[i].engagement.like {
						idx = i
					}
				}
			case "Comment", "comment":
				switch *pilihan {
				case 1:
					if F[idx].engagement.comment > F[i].engagement.comment {
						idx = i
					}
				case 2:
					if F[idx].engagement.comment < F[i].engagement.comment {
						idx = i
					}
				}
			case "Share", "share":
				switch *pilihan {
				case 1:
					if F[idx].engagement.share > F[i].engagement.share {
						idx = i
					}
				case 2:
					if F[idx].engagement.share < F[i].engagement.share {
						idx = i
					}
				}
			}
		}
		temp = F[pass]
		F[pass] = F[idx]
		F[idx] = temp
	}
}
func timeConvert(A data, i int) int {
	var tahun, bulan, hari, jam int
	tahun = A[i].waktu.tahun * 360 * 24
	bulan = A[i].waktu.bulan * 30 * 24
	hari = A[i].waktu.hari * 24
	jam = A[i].waktu.jam
	return tahun + bulan + hari + jam
}
func dummy(A *data) {
	
}
func main() {
	var pilihan int = 1
	var dataKonten data
	var Filtered data
	var nFiltered int = 0
	var nData int = 0

	for pilihan != 2 {
		menu(&pilihan)
		if pilihan == 1 {
			for pilihan != 5 {
				daftarKonten(&pilihan, dataKonten, &nData)
				switch pilihan {
				case 1:
					inputKonten(&dataKonten, &nData)
				case 2:
					for pilihan != 3 {
						editData(&pilihan, &dataKonten, &nData)
						switch pilihan {
						case 1:
							replaceData(&dataKonten, nData)
						case 2:
							deleteData(&dataKonten, &nData)
						}
					}
				case 3:
					filter(&pilihan, &dataKonten, nData, Filtered, nFiltered)
				case 4:
					for pilihan != 0 && pilihan != 1 {
						sort(dataKonten, nData, &pilihan)
					}
				}
			}
		}
	}
}
