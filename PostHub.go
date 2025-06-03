package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const consoleWidth int = 133
const NMAX int = 100

type konten struct {
	id, no             int
	nama               string
	judulPart          [NMAX]string
	long               int
	kategori, platform string
	waktu              waktu
	engagement         engagement
}
type waktu struct {
	jam, hari, bulan, tahun, converted int
}
type engagement struct {
	like, comment, share int
}
type data [NMAX]konten

// Merapikan visual aplikasi dengan menghapus riwayat instruksi yang sudah dijalankan
func cls() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

// Visual header selama aplikasi berjalaln
func header() {
	fmt.Printf("%-35s┌─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐\n", "")
	fmt.Printf("%-35s│%-133s│\n", "", "")
	fmt.Printf("%-35s│%-56s▶▶  P o s t H u b  ◀◀%-56s│\n", "", "", "")
	fmt.Printf("%-35s│%-53s💡 Scroll less, Post more%-55s│\n", "", "", "")
	fmt.Printf("%-35s│%-133s│\n", "", "")
	fmt.Printf("%-35s└─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘\n", "")
}

// Main menu paling awal aplikasi
func menu(pilihan *int, counter *int) {
	cls()
	fmt.Println()
	fmt.Printf("%-50s┌──────────────────────────────────────────────────────────────────────────────────────────────────┐\n", "")
	fmt.Printf("%-50s│                                                                                                  │\n", "")
	fmt.Printf("%-50s│                                   ▶▶ Selamat Datang di PostHub! ◀◀                               │\n", "")
	fmt.Printf("%-50s│                                      💡 Scroll less, Post more                                   │\n", "")
	fmt.Printf("%-50s│                                                                                                  │\n", "")
	fmt.Printf("%-50s├──────────────────────────────────────────────────────────────────────────────────────────────────┤\n", "")
	fmt.Printf("%-50s│                    °❀⋆.ೃ࿔*:･°❀⋆.ೃ࿔*:･        MAIN MENU         °❀⋆.ೃ࿔*:･°❀⋆.ೃ࿔*           	     │\n", "")
	fmt.Printf("%-50s└──────────────────────────────────────────────────────────────────────────────────────────────────┘\n", "")
	if (*pilihan < 1 || *pilihan > 2) && *counter != 0 {
		fmt.Printf("%-50sPilihan invalid.\n", "")
	} else {
		fmt.Println()
	}
	fmt.Printf("%-50s[1] 🔧 Manage PostHub\n", "")
	fmt.Printf("%-50s[2] ↩️ Exit app\n", "")
	fmt.Println()
	fmt.Printf("%-50s╰┈➤ ", "")
	fmt.Scan(pilihan)
	*counter++
}

// Output data konten
func outputData(A data, n int, status *string) {
	var i int
	cls()
	fmt.Println()
	header()
	if n != 0 {
		fmt.Printf("%-35s┌───┬───────────────────────────────────┬────────────────┬────────────────┬──────┬────────────────┬───────────┬───────────┬───────────┐\n", "")
		fmt.Printf("%-35s│%-3s│ %-34s│ %-15s│ %-15s│ %-5v│ %-15s│ %-10s│ %-10s│ %-10s│\n", "", "No", "Judul", "Kategori", "Platform", "Jam", "Tanggal Post", "Like", "Comment", "Share")
		fmt.Printf("%-35s├───┼───────────────────────────────────┼────────────────┼────────────────┼──────┼────────────────┼───────────┼───────────┼───────────┤\n", "")
		for i = 0; i < n; i++ {
			A[i].no = i + 1
			fmt.Printf("%-35s│%-3v│ %-34v│ %-15v│ %-15v│ %-5v│ %02d/%02d/%-9v│ %-10v│ %-10v│ %-10v│\n", "", A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam, A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like, A[i].engagement.comment, A[i].engagement.share)
		}
		fmt.Printf("%-35s└───┴───────────────────────────────────┴────────────────┴────────────────┴──────┴────────────────┴───────────┴───────────┴───────────┘\n", "")

	} else {
		fmt.Printf("%-35s┌─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐\n", "")
		fmt.Printf("%-35s│%-60s⌕  Data Kosong%-59s│\n", "", "", "")
		fmt.Printf("%-35s└─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘\n", "")
	}
	fmt.Println()
	fmt.Printf("%-35s%v\n", "", *status)
	*status = ""
}

// Menu inti aplikasi, berisi data konten dan pilihan fitur untuk dipakai
func daftarKonten(pilihan *int, A data, n *int, status *string) {
	outputData(A, *n, *&status)
	fmt.Printf("%-35s[1] ➕ Tambahkan konten\n", "")
	fmt.Printf("%-35s[2] 🔎 Cari judul konten\n", "")
	fmt.Printf("%-35s[3] 📝 Edit data konten\n", "")
	fmt.Printf("%-35s[4] ☰  Filter konten\n", "")
	fmt.Printf("%-35s[5] 📶 Sortir konten\n", "")
	fmt.Printf("%-35s[6] ⬅️ Kembali ke menu\n", "")
	fmt.Printf("%-35s╰┈➤ ", "")
	fmt.Scan(pilihan)

	if *n == 0 && (*pilihan == 2 || *pilihan == 3 || *pilihan == 4 || *pilihan == 5) {
		*status = "Belum ada data untuk melakukan perintah."
	} else if *pilihan > 6 || *pilihan < 1 {
		*status = "Pilihan invalid."
	}
}

// Input data konten
func inputKonten(A *data, n *int, status *string) {
	var i int = 0
	var j int
	var namaKonten string
	outputData(*A, *n, *&status)
	fmt.Printf("%-35sFormat inputan:\n", "")
	fmt.Printf("%-35s[ Judul Kategori Platform Jam DD/MM/YYYY Like Comment Share ] (Back -)\n", "")
	fmt.Println()
	if A[0].nama == "" {
		fmt.Printf("%-35s", "")
		fmt.Scan(&namaKonten)
	} else if A[0].nama != "" {
		for A[i].nama != "" {
			i++
		}
		fmt.Printf("%-35s", "")
		fmt.Scan(&namaKonten)
	}
	for namaKonten != "-" && namaKonten != "." {
		j = 0
		A[i].judulPart[j] = namaKonten
		if strings.Contains(A[i].judulPart[j], ".") {
			A[i].nama = A[i].judulPart[j]
		} else {
			isiJudul(*&A, i, 1)
		}
		fmt.Scanf("%v %v %v %v/%v/%v %v %v %v", &A[i].kategori, &A[i].platform, &A[i].waktu.jam, &A[i].waktu.hari, &A[i].waktu.bulan, &A[i].waktu.tahun, &A[i].engagement.like, &A[i].engagement.comment, &A[i].engagement.share)
		A[i].id = i + 1
		A[i].no = A[i].id
		A[i].waktu.converted = timeConvert(*A, i)
		*n++
		i++
		*status = "Konten berhasil ditambahkan."
		fmt.Printf("%-35s", "")
		fmt.Scan(&namaKonten)
	}
}

// Input judul dengan spasi dan diakhiri dengan titik (.)
func isiJudul(A *data, i int, j int) {
	fmt.Scan(&A[i].judulPart[j])
	for !strings.Contains(A[i].judulPart[j], ".") {
		j++
		fmt.Scan(&A[i].judulPart[j])
	}
	A[i].long = j + 1
	A[i].nama = gabungJudul(*A, i)
}

// Menggabungkan array judul menjadi judul utuh
func gabungJudul(A data, i int) string {
	var j int
	var hasil string
	for j = 0; j < A[i].long; j++ {
		if j > 0 {
			hasil += " "
		}
		hasil += A[i].judulPart[j]
	}
	return hasil
}

// Mencari judul berdasarkan keyword yang diinputkan pengguna dengan function strings.Contains()
func findJudul(pilihan *int, A *data, n *int, F data, nFilter int, status *string) {
	var key string
	var search string
	var i int
	var back bool
	var showStatus bool = true
	outputData(*A, *n, *&status)
	fmt.Println(*status)
	fmt.Printf("%-35sMasukkan kata kunci (Back - ): ", "")
	fmt.Scan(&key)
	if key != "-" {
		for !back {
			nFilter = 0
			key = strings.ToLower(key)
			cls()
			fmt.Println()
			header()
			fmt.Printf("%-35s┌───┬───────────────────────────────────┬────────────────┬────────────────┬──────┬────────────────┬───────────┬───────────┬───────────┐\n", "")
			fmt.Printf("%-35s│%-3s│ %-34s│ %-15s│ %-15s│ %-5v│ %-15s│ %-10s│ %-10s│ %-10s│\n", "", "No", "Judul", "Kategori", "Platform", "Jam", "Tanggal Post", "Like", "Comment", "Share")
			fmt.Printf("%-35s├───┼───────────────────────────────────┼────────────────┼────────────────┼──────┼────────────────┼───────────┼───────────┼───────────┤\n", "")

			for i = 0; i < *n; i++ {
				search = strings.ToLower(A[i].nama)
				if strings.Contains(search, key) {
					F[nFilter] = A[i]
					F[nFilter].no = nFilter + 1
					fmt.Printf("%-35s│%-3v│ %-34v│ %-15v│ %-15v│ %-5v│ %02d/%02d/%-9v│ %-10v│ %-10v│ %-10v│\n", "", F[nFilter].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam, A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like, A[i].engagement.comment, A[i].engagement.share)
					nFilter++
				}
			}
			if nFilter != 0 && showStatus {
				*status = "Judul dengan kata kunci ditemukan."
				showStatus = false
			}
			filterOption(*&pilihan, *&A, *&n, F, nFilter, *&status)
			*status = ""
			if *pilihan == 4 || *pilihan == 1 {
				back = true
			}

		}
	}
}

// Menu opsi edit data yang menyediakan perubahan dan penghapusan
func editData(pilihan *int, A *data, n *int, status *string) {
	var back bool
	var idx int
	for !back {
		outputData(*A, *n, *&status)
		fmt.Printf("%-35s[1] 📝 Edit data\n", "")
		fmt.Printf("%-35s[2] 🗑️ Hapus data\n", "")
		fmt.Printf("%-35s[3] ⬅️ Kembali\n", "")
		fmt.Printf("%-35s╰┈➤ ", "")
		fmt.Scan(pilihan)
		switch *pilihan {
		case 1:
			replaceData(A, *n, status)
		case 2:
			deleteData(A, n, status, &idx)
		case 3:
			back = true
			*status = ""
		default:
			*status = "Pilihan tidak valid."
		}
	}
}

// Mengubah data berdasarkan kolom yang diinginkan dengan mengimplementasikan binary search
func replaceData(A *data, n int, status *string) {
	var noData, idx int
	var col string
	outputData(*A, n, *&status)
	fmt.Printf("%-35sFormat inputan:\n", "")
	fmt.Printf("%-35s[ Nomor_data Kolom ] (Back: 0 0)\n", "")
	fmt.Printf("%-35s", "")
	fmt.Scan(&noData, &col)
	col = strings.ToLower(col)
	idx = dataSearch(*A, n, noData)
	if idx != -1 {
		outputData(*A, n, *&status)
		fmt.Printf("%-35s%v%v%v%v", "", col, " baru untuk data nomor ", noData, ": ")
		*status = "Konten berhasil diedit."
		switch col {
		case "judul":
			isiJudul(*&A, idx, 0)
		case "kategori":
			fmt.Scan(&A[idx].kategori)
		case "platform":
			fmt.Scan(&A[idx].platform)
		case "waktu":
			fmt.Scanf("\n%d %d/%d/%d\n", &A[idx].waktu.jam, &A[idx].waktu.hari, &A[idx].waktu.bulan, &A[idx].waktu.tahun)
		case "jam":
			fmt.Scan(&A[idx].waktu.jam)
		case "hari":
			fmt.Scan(&A[idx].waktu.hari)
		case "bulan":
			fmt.Scan(&A[idx].waktu.bulan)
		case "tahun":
			fmt.Scan(&A[idx].waktu.tahun)
		case "engagement":
			fmt.Scan(&A[idx].engagement.like, &A[idx].engagement.comment, &A[idx].engagement.share)
		case "like":
			fmt.Scan(&A[idx].engagement.like)
		case "comment":
			fmt.Scan(&A[idx].engagement.comment)
		case "share":
			fmt.Scan(&A[idx].engagement.share)
		default:
			*status = "Kolom invalid."
		}
	} else if noData == 0 && col == "0" {
	} else {
		if noData == 0 {
		} else {
			*status = "Data tidak ditemukan."
		}
	}
}

// Menghapus data sesuai nomor data yang mengimplementasikan binary search
func deleteData(A *data, n *int, status *string, idx *int) {
	var noData int
	var i int
	outputData(*A, *n, *&status)
	fmt.Printf("%-35sNomor data (Back 0): ", "")
	fmt.Scan(&noData)
	*idx = dataSearch(*A, *n, noData)
	if noData != 0 && *idx != -1 {
		for i = *idx; i < *n; i++ {
			A[i] = A[i+1]
			A[i].no--
			A[i].id--
		}
		*n--
		A[i].nama = ""
		*status = "Konten berhasil dihapus."
	} else if noData != 0 && *idx == -1 {
		*status = "Data tidak ditemukan."
	} else {
	}
}

// Menyaring data konten berdasarkan kategori dan/atau platform menggunakan sequential search
func filter(pilihan *int, A *data, n *int, F data, nFilter int, status *string) {
	var pilKat, pilPlat string
	var lowPlat, lowKat string
	var i int
	var back bool
	var showStatus bool = true
	outputData(*A, *n, *&status)
	fmt.Printf("%-35sMasukkan kategori atau platform (Back - -): ", "")
	fmt.Scan(&pilKat, &pilPlat)
	if pilKat == "-" && pilPlat == "-" {
	} else if (pilKat == "-" && pilPlat != "-" && len(pilPlat) < 4) ||
		(pilPlat == "-" && pilKat != "-" && len(pilKat) < 4) ||
		(pilKat != "-" && pilPlat != "-" && (len(pilKat) < 4 || len(pilPlat) < 4)) {
		*status = "Masukan tidak valid."
	} else {
		for !back {
			cls()
			fmt.Println()
			header()
			fmt.Printf("%-35s┌───┬───────────────────────────────────┬────────────────┬────────────────┬───────────────────────┬───────────┬───────────┬───────────┐\n", "")
			fmt.Printf("%-35s│%-3s│ %-34s│ %-15s│ %-15s│ %-5v│ %-15s│ %-10s│ %-10s│ %-10s│\n", "", "No", "Judul", "Kategori", "Platform", "Jam", "Tanggal Post", "Like", "Comment", "Share")
			fmt.Printf("%-35s├───┼───────────────────────────────────┼────────────────┼────────────────┼───────────────────────┼───────────┼───────────┼───────────┤\n", "")
			nFilter = 0
			pilKat = strings.ToLower(pilKat)
			pilPlat = strings.ToLower(pilPlat)
			for i = 0; i < *n; i++ {
				lowKat = strings.ToLower(A[i].kategori)
				lowPlat = strings.ToLower(A[i].platform)
				if pilPlat == "-" {
					if lowKat == pilKat {
						F[nFilter] = A[i]
						F[nFilter].no = nFilter + 1
						fmt.Printf("%-35s│%-3v│ %-34v│ %-15v│ %-15v│ %-5v│ %02d/%02d/%-9v│ %-10v│ %-10v│ %-10v│\n", "", F[nFilter].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam, A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like, A[i].engagement.comment, A[i].engagement.share)
						nFilter++
					}
				} else if pilKat == "-" {
					if lowPlat == pilPlat {
						F[nFilter] = A[i]
						F[nFilter].no = nFilter + 1
						fmt.Printf("%-35s│%-3v│ %-34v│ %-15v│ %-15v│ %-5v│ %02d/%02d/%-9v│ %-10v│ %-10v│ %-10v│\n", "", F[nFilter].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam, A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like, A[i].engagement.comment, A[i].engagement.share)
						nFilter++
					}
				} else {
					if lowPlat == pilPlat && lowKat == pilKat {
						F[nFilter] = A[i]
						F[nFilter].no = nFilter + 1
						fmt.Printf("%-35s│%-3v│ %-34v│ %-15v│ %-15v│ %-5v│ %02d/%02d/%-9v│ %-10v│ %-10v│ %-10v│\n", "", F[nFilter].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam, A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like, A[i].engagement.comment, A[i].engagement.share)
						nFilter++
					}
				}
			}
			if nFilter != 0 && showStatus {
				*status = "Konten berhasil difilter."
				showStatus = false
			}
			filterOption(*&pilihan, *&A, *&n, F, nFilter, *&status)
			if *pilihan == 4 || *pilihan == 1 {
				back = true
			}
		}
	}
}

// Menu setelah data difilter untuk pengelolaan konten lebih lanjut
func filterOption(pilihan *int, A *data, n *int, F data, nFilter int, status *string) {
	var i, j int
	if nFilter == 0 {

		fmt.Printf("%-35s│%-60s⌕  Data Kosong%-59s│\n", "", "", "")
		fmt.Println()
		fmt.Printf("%-35s[1] ⬅️ Kembali\n", "")
		fmt.Printf("%-35s╰┈➤ ", "")
		fmt.Scan(pilihan)
	} else {
		fmt.Printf("%-35s└───┴───────────────────────────────────┴────────────────┴────────────────┴──────┴────────────────┴───────────┴───────────┴───────────┘\n", "")
		fmt.Println()

		fmt.Printf("%-35s%v\n", "", *status)
		*status = ""

		fmt.Printf("%-35s[1] 📶 Sortir konten\n", "")
		fmt.Printf("%-35s[2] 📝 Edit konten\n", "")
		fmt.Printf("%-35s[3] 🗑️ Hapus konten\n", "")
		fmt.Printf("%-35s[4] ⬅️ Kembali\n", "")
		fmt.Printf("%-35s╰┈➤ ", "")
		fmt.Scan(pilihan)
		if *pilihan == 1 && nFilter > 0 {
			sort(*&pilihan, *&A, *&n, F, nFilter, *&status, true)
		} else if *pilihan != 4 && nFilter == 0 {
			*status = "Data tidak ditemukan."
		} else if *pilihan == 2 && nFilter > 0 {
			replaceData(&F, nFilter, *&status)
			for i = 0; i < *n; i++ {
				for j = 0; j < nFilter; j++ {
					if A[i].id == F[j].id {
						A[i].nama = F[j].nama
						A[i].kategori = F[j].kategori
						A[i].platform = F[j].platform
						A[i].waktu = F[j].waktu
						A[i].engagement = F[j].engagement
					}
				}
			}
		} else if *pilihan == 3 && nFilter > 0 {
			filteredDelete(*&A, *&n, &F, &nFilter, status)
		}
	}
}

// Penghapusan data konten
func filteredDelete(A *data, n *int, F *data, nFilter *int, status *string) {
	var noData int
	var i, j, idx int
	var cek bool
	outputData(*F, *nFilter, *&status)
	fmt.Printf("%-35sNomor data (Back 0): ", "")
	fmt.Scan(&noData)
	idx = dataSearch(*F, *nFilter, noData)
	if noData != 0 && idx != -1 {
		for i = 0; i < *n && !cek; i++ {
			cek = A[i].id == F[idx].id
			if cek {
				for j = i; j < *n; j++ {
					A[j] = A[j+1]
					A[j].no--
					A[j].id--
				}
				*n--
				A[*n].nama = ""
			}
		}
		for i = idx; i < *nFilter-1; i++ {
			F[i] = F[i+1]
			F[i].no--
			F[i].id--
		}
		*nFilter--
		F[*nFilter].nama = ""
		*status = "Konten berhasil dihapus."
	} else if noData != 0 && idx == -1 {
		*status = "Data tidak ditemukan."
	} else {
	}
}

// Pencarian data menggunakan binary search
func dataSearch(A data, n int, find int) int {
	var left, right, mid, idx int
	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if find < A[mid].no {
			right = mid - 1
		} else if find > A[mid].no {
			left = mid + 1
		} else {
			idx = mid
		}
	}
	return idx
}

// Sortir data konten
func sort(pilihan *int, A *data, n *int, F data, nFilter int, status *string, isFiltered bool) {
	var whatSort string
	var i int
	if !isFiltered {
		nFilter = 0
		for i = 0; i < *n; i++ {
			F[i] = A[i]
			nFilter++
		}
	}
	outputData(F, nFilter, *&status)
	fmt.Printf("%-35s[1] ⬆ Ascending (Membesar)\n", "")
	fmt.Printf("%-35s[2] ⬇ Descending (Mengecil)\n", "")
	fmt.Printf("%-35s[ Pilihan_sort Berdasarkan ] (Back: 0 0)\n", "")
	fmt.Printf("%-35s", "")
	fmt.Scan(pilihan, &whatSort)
	whatSort = strings.ToLower(whatSort)
	if *pilihan == 0 && whatSort == "0" {
		*status = ""
	} else {
		sortProcess(&F, nFilter, *&pilihan, whatSort)
		sortOption(*&pilihan, *&A, *&n, &F, &nFilter, *&status)
	}
}

// Proses sortir data konten
func sortProcess(F *data, n int, pilihan *int, x string) {
	var pass, idx, i int
	var temp konten
	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for i = pass + 1; i < n; i++ {
			switch x {
			case "waktu":
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
			case "like":
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
			case "comment":
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
			case "share":
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
		temp.nama = F[pass].nama
		temp.kategori = F[pass].kategori
		temp.platform = F[pass].platform
		temp.waktu = F[pass].waktu
		temp.engagement = F[pass].engagement
		temp.id = F[pass].id

		F[pass].nama = F[idx].nama
		F[pass].kategori = F[idx].kategori
		F[pass].platform = F[idx].platform
		F[pass].waktu = F[idx].waktu
		F[pass].engagement = F[idx].engagement
		F[pass].id = F[idx].id

		F[idx].nama = temp.nama
		F[idx].kategori = temp.kategori
		F[idx].platform = temp.platform
		F[idx].waktu = temp.waktu
		F[idx].engagement = temp.engagement
		F[idx].id = temp.id

	}
}

// Menu opsi sortir
func sortOption(pilihan *int, A *data, n *int, F *data, nFilter *int, status *string) {
	var back bool
	var showStatus bool = true
	for !back {
		var i, j int
		if *nFilter != 0 && showStatus {
			*status = "Konten berhasil disortir."
			showStatus = false
		}
		outputData(*F, *nFilter, *&status)
		fmt.Printf("%-35s[1] 📝 Edit data\n", "")
		fmt.Printf("%-35s[2] 🗑️ Hapus data\n", "")
		fmt.Printf("%-35s[3] ⬅️ Kembali\n", "")
		fmt.Printf("%-35s╰┈➤ ", "")
		fmt.Scan(pilihan)
		*status = ""
		if *pilihan == 1 && *nFilter > 0 {
			replaceData(*&F, *nFilter, *&status)
			for i = 0; i < *n; i++ {
				for j = 0; j < *nFilter; j++ {
					if A[i].id == F[j].id {
						A[i] = F[j]
					}
				}
			}

		} else if *pilihan == 2 && *nFilter > 0 {
			filteredDelete(*&A, *&n, *&F, *&nFilter, status)
		} else if *pilihan == 3 {
			back = true
		} else {
			*status = "Pilihan invalid."
		}
	}
}

// Konversi waktu untuk sortir
func timeConvert(A data, i int) int {
	var tahun, bulan, hari, jam int
	tahun = A[i].waktu.tahun * 360 * 24
	bulan = A[i].waktu.bulan * 30 * 24
	hari = A[i].waktu.hari * 24
	jam = A[i].waktu.jam
	return tahun + bulan + hari + jam
}

// Program Utama
func main() {
	var pilihan int = 1
	var dataKonten data
	var Filtered data
	var nFiltered int = 0
	var nData int = 0
	var status string = ""
	var counter int = 0

	dummyNew(&dataKonten, &nData)
	for pilihan != 2 || pilihan == 999 {
		menu(&pilihan, &counter)
		if pilihan == 1 {
			counter = 0
			for pilihan != 6 {
				daftarKonten(&pilihan, dataKonten, &nData, &status)
				if nData != 0 || pilihan == 1 {
					switch pilihan {
					case 1:
						inputKonten(&dataKonten, &nData, &status)
					case 2:
						findJudul(&pilihan, &dataKonten, &nData, Filtered, nFiltered, &status)
					case 3:
						editData(&pilihan, &dataKonten, &nData, &status)
					case 4:
						filter(&pilihan, &dataKonten, &nData, Filtered, nFiltered, &status)
					case 5:
						sort(&pilihan, &dataKonten, &nData, Filtered, nFiltered, &status, false)
					}
				}
			}
		}
	}
}

// Input data dummy
func dummyNew(A *data, n *int) {
	var i int = 0

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Misteri Rumah Tua.", "Horror", "YouTube", 20, 11, 5, 2025, 2489, 322, 157
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Tips Lolos CPNS.", "Edukasi", "Instagram", 9, 2, 10, 2025, 781, 89, 54
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Review HP Terbaru.", "Teknologi", "TikTok", 18, 6, 12, 2025, 3056, 410, 230
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Vlog Ke Jepang.", "Vlog", "YouTube", 16, 7, 12, 2025, 1298, 120, 98
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Resep Sambal Pedas.", "Kuliner", "Facebook", 11, 3, 12, 2025, 1523, 145, 67
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Fakta Ular Paling Berbisa.", "Edukasi", "TikTok", 14, 1, 12, 2025, 2201, 290, 101
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Konser Musik Jepang.", "Hiburan", "Instagram", 17, 4, 12, 2025, 3324, 410, 185
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Eksperimen CocaCola.", "Edukasi", "YouTube", 15, 8, 11, 2025, 1984, 250, 120
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Hantu Penjaga Sekolah.", "Horror", "TikTok", 22, 6, 12, 2025, 2876, 315, 204
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Jalan Jalan Bandung.", "Vlog", "Instagram", 10, 5, 12, 2025, 1133, 104, 75
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Nasi Goreng KakiLima.", "Kuliner", "Facebook", 8, 4, 11, 2025, 1760, 166, 82
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Penampakan Lembah Angker.", "Horror", "YouTube", 23, 3, 10, 2025, 2545, 337, 199
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Rumus Cepat Matematika.", "Edukasi", "TikTok", 12, 1, 12, 2025, 2430, 280, 111
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Ganti Layar HP Sendiri.", "Teknologi", "YouTube", 19, 2, 11, 2025, 1989, 210, 95
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Cemilan Sehat 5Menit.", "Kuliner", "Instagram", 13, 7, 12, 2025, 1430, 130, 77
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Kost Murah Dekat Kampus.", "Vlog", "YouTube", 21, 2, 10, 2025, 1032, 88, 49
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Misteri Jalan Berbisik.", "Horror", "TikTok", 20, 4, 12, 2025, 2620, 309, 167
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Alat Unik Jepang.", "Teknologi", "Facebook", 7, 3, 12, 2025, 1887, 195, 92
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Tempat Wisata Tersembunyi.", "Vlog", "Instagram", 14, 5, 12, 2025, 1344, 110, 66
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	A[i].id, A[i].no, A[i].nama, A[i].kategori, A[i].platform, A[i].waktu.jam,
		A[i].waktu.hari, A[i].waktu.bulan, A[i].waktu.tahun, A[i].engagement.like,
		A[i].engagement.comment, A[i].engagement.share = i+1, i+1, "Belajar Bahasa Korea.", "Edukasi", "YouTube", 16, 6, 12, 2025, 2120, 240, 99
	A[i].waktu.converted = timeConvert(*A, i)
	i++

	*n = i
}
