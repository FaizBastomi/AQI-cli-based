package interactive

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/FaizBastomi/AQI-cli-based/utils"
)

func ClearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func TambahData(A *utils.AirPolutions) {
	var lokasi, sumberPolusi string
	var IdxUdara int
	var scanner *bufio.Scanner

	ClearConsole()

	scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("Silahkan masukkan data baru")

	lokasi = utils.GetNonEmptyInput(scanner, "Lokasi: ")
	sumberPolusi = utils.GetNonEmptyInput(scanner, "Sumber Polusi: ")
	IdxUdara = utils.GetIntInput(scanner, "Index Udara: ")

	utils.AddData(A, lokasi, sumberPolusi, IdxUdara)
}

func UbahData(A *utils.AirPolutions) {
	var i, index, idxUdaraBaru, totalPages int
	var choice, lokasiBaru, sumberBaru string
	var item utils.AirPolution
	var dataPage []utils.AirPolution
	var scanner *bufio.Scanner
	var err error

	var currentPage int = 1

	ClearConsole()
	scanner = bufio.NewScanner(os.Stdin)

	fmt.Println("Silahkan pilih data yang ingin diubah:")
	for {
		dataPage = utils.PaginateData(*A, currentPage)
		totalPages = len(dataPage) / 5
		if len(dataPage)%5 != 0 {
			totalPages++
		}

		if len(dataPage) == 0 {
			fmt.Println("Tidak ada data untuk ditampilkan.")
		} else {
			fmt.Printf("Data halaman %d dari %d:\n", currentPage, totalPages)
			for i, item = range dataPage {
				fmt.Printf("%d. Lokasi: %s\nSumber: %s\nIndex: %d\nTingkat: %s\nWaktu: %v\n",
					i+1, item.Lokasi, item.SumberPolusi, item.IdxUdara, item.TingkatBahaya, item.Waktu.Format("02-January-2006 15:04"))
				fmt.Println(strings.Repeat("-", 50))
			}
		}
		fmt.Println("[n] Next page\n[p] Previous page\n[q] Main Menu, atau masukan nomor data")
		fmt.Print("Select: ")

		scanner.Scan()
		choice = scanner.Text()

		ClearConsole()
		switch choice {
		case "q":
			return
		case "n":
			if currentPage < totalPages {
				currentPage++
			} else {
				fmt.Println("Sudah di halaman terakhir.")
			}
		case "p":
			if currentPage > 1 {
				currentPage--
			} else {
				fmt.Println("Sudah di halaman pertama.")
			}
		default:
			index, err = strconv.Atoi(choice)
			if err != nil || index < 1 || index > len(dataPage) {
				fmt.Println("Pilihan tidak valid.")
			} else {
				ClearConsole()
				item = dataPage[index-1]
				fmt.Println("Ubah Data untuk:")
				fmt.Printf("Lokasi: %s\nSumber: %s\nIndex: %d\n", item.Lokasi, item.SumberPolusi, item.IdxUdara)

				fmt.Print("Lokasi (tekan enter untuk skip): ")
				scanner.Scan()
				lokasiBaru = scanner.Text()
				if lokasiBaru != "" {
					item.Lokasi = lokasiBaru
				}

				fmt.Print("Sumber polusi (tekan enter untuk skip): ")
				scanner.Scan()
				sumberBaru = scanner.Text()
				if sumberBaru != "" {
					item.SumberPolusi = sumberBaru
				}

				fmt.Print("Index udara (tekan enter untuk skip): ")
				scanner.Scan()
				choice = scanner.Text()
				if choice != "" {
					idxUdaraBaru, err = strconv.Atoi(choice)
					if err != nil {
						fmt.Println("Index udara tidak valid.")
					} else {
						item.IdxUdara = idxUdaraBaru
						if idxUdaraBaru >= 0 && idxUdaraBaru <= 50 {
							item.TingkatBahaya = "Baik"
						} else if idxUdaraBaru >= 51 && idxUdaraBaru <= 100 {
							item.TingkatBahaya = "Sedang"
						} else if idxUdaraBaru >= 101 && idxUdaraBaru <= 150 {
							item.TingkatBahaya = "Tidak Sehat"
						} else {
							item.TingkatBahaya = "Berbahaya"
						}
					}
				}
				utils.EditData(A, item.Lokasi, item.SumberPolusi, item.IdxUdara, item.AqiID)
			}
		}
	}
}

func HapusData(A *utils.AirPolutions) {
	var i, index, totalPages int
	var choice, confirm string
	var item utils.AirPolution
	var dataPage []utils.AirPolution
	var scanner *bufio.Scanner
	var err error

	var currentPage int = 1

	ClearConsole()
	scanner = bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Silahkan pilih data yang ingin dihapus:")

		dataPage = utils.PaginateData(*A, currentPage)
		totalPages = len(dataPage) / 5
		if len(dataPage)%5 != 0 {
			totalPages++
		}

		if len(dataPage) == 0 {
			fmt.Println("Tidak ada data untuk ditampilkan.")
		} else {
			fmt.Printf("Data halaman %d dari %d:\n", currentPage, totalPages)
			for i, item = range dataPage {
				fmt.Printf("%d. Lokasi: %s\nSumber: %s\nIndex: %d\nTingkat: %s\nWaktu: %v\n",
					i+1, item.Lokasi, item.SumberPolusi, item.IdxUdara, item.TingkatBahaya, item.Waktu.Format("02-January-2006 15:04"))
				fmt.Println(strings.Repeat("-", 50))
			}
		}
		fmt.Println("[n] Next page\n[p] Previous page\n[q] Main Menu, atau masukan nomor data")
		fmt.Print("Select: ")

		scanner.Scan()
		choice = scanner.Text()

		ClearConsole()
		switch choice {
		case "q":
			return
		case "n":
			if currentPage < totalPages {
				currentPage++
			} else {
				fmt.Println("Sudah di halaman terakhir.")
			}
		case "p":
			if currentPage > 1 {
				currentPage--
			} else {
				fmt.Println("Sudah di halaman pertama.")
			}
		default:
			index, err = strconv.Atoi(choice)
			if err != nil || index < 1 || index > len(dataPage) {
				fmt.Println("Pilihan tidak valid.")
			} else {
				ClearConsole()
				item = dataPage[index-1]

				fmt.Print("Apakah kamu yakin (y/n): ")
				scanner.Scan()
				confirm = scanner.Text()

				if confirm == "y" {
					utils.DeleteData(A, item.AqiID)
				}
			}
		}
	}
}

func ShowData(A *utils.AirPolutions) {
	var currentPage, totalPages, page, i int
	var item utils.AirPolution
	var dataPage []utils.AirPolution
	var scanner *bufio.Scanner
	var choice string
	var err error

	ClearConsole()
	scanner = bufio.NewScanner(os.Stdin)
	currentPage = 1

	for {
		dataPage = utils.PaginateData(*A, currentPage)
		totalPages = len(dataPage) / 5
		if len(dataPage)%5 != 0 {
			totalPages++
		}

		if len(dataPage) == 0 {
			fmt.Println("Tidak ada data untuk ditampilkan.")
		} else {
			fmt.Printf("Data halaman %d dari %d:\n", currentPage, totalPages)
			for i, item = range dataPage {
				fmt.Printf("%d Lokasi: %s\nSumber: %s\nIndex: %d\nTingkat: %s\nWaktu: %v\n",
					i+1, item.Lokasi, item.SumberPolusi, item.IdxUdara, item.TingkatBahaya, item.Waktu.Format("02-January-2006 15:04"))
				fmt.Println(strings.Repeat("-", 50))
			}
		}
		fmt.Println("[n] Next page\n[p] Previous page\n[q] Main Menu, atau masukan nomor halaman")
		fmt.Print("Select: ")
		scanner.Scan()
		choice = scanner.Text()

		ClearConsole()
		switch choice {
		case "q":
			return
		case "n":
			if currentPage < totalPages {
				currentPage++
			} else {
				fmt.Println("Sudah di halaman terakhir.")
			}
		case "p":
			if currentPage > 1 {
				currentPage--
			} else {
				fmt.Println("Sudah di halaman pertama.")
			}
		default:
			page, err = strconv.Atoi(choice)
			if err != nil || page < 1 {
				fmt.Println("Halaman tidak valid.")
			} else {
				currentPage = page
			}
		}
	}
}

func CariData(A *utils.AirPolutions) {
	var lokasi string
	var scanner *bufio.Scanner

	scanner = bufio.NewScanner(os.Stdin)
	lokasi = utils.GetNonEmptyInput(scanner, "Masukkan nama lokasi yang dicari: ")

	result := utils.SequentialSearch(*A, lokasi)
	if result != nil {
		fmt.Println("Data ditemukan:")
		fmt.Printf("Lokasi: %s\nSumber: %s\nIndex: %d\nTingkat: %s\nWaktu: %v\n",
			result.Lokasi, result.SumberPolusi, result.IdxUdara, result.TingkatBahaya, result.Waktu.Format("02-January-2006 15:04"))
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
	fmt.Println("Tekan Enter untuk kembali...")
	scanner.Scan()
}

func UrutPolusiTerendah(A *utils.AirPolutions) {
	utils.SortAscendingByIdxUdara(A)
	ShowData(A)
}

func UrutPolusiTertinggi(A *utils.AirPolutions) {
	utils.SortDescendingByIdxUdara(A)
	ShowData(A)
}
