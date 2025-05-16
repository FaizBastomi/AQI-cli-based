package main

import (
	"fmt"
	"github.com/FaizBastomi/AQI-cli-based/interactive"
	"github.com/FaizBastomi/AQI-cli-based/utils"
	"os"
)

func main() {
	var opsi int

	// Read data from JSON file
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	data, err := utils.ReadFromJSON(path + "/data.json")
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	interactive.ClearConsole()
	for opsi != 8 {
		fmt.Println("Select Menu:")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Edit Data")
		fmt.Println("3. Hapus Data")
		fmt.Println("4. Tampilkan Data")
		fmt.Println("5. Daftar Polusi Tertinggi")
		fmt.Println("6. Daftar Polusi Terendah")
		fmt.Println("7. Cari Data")
		fmt.Println("8. Exit")
		fmt.Print("Masukan opsi: ")
		fmt.Scanln(&opsi)

		switch opsi {
		case 1:
			interactive.TambahData(&data)
		case 2:
			interactive.UbahData(&data)
		case 3:
			interactive.HapusData(&data)
		case 4:
			interactive.ShowData(&data)
		case 5:
			interactive.UrutPolusiTertinggi(&data)
		case 6:
			interactive.UrutPolusiTerendah(&data)
		case 7:
			interactive.CariData(&data)
		}

		interactive.ClearConsole()
	}

	// Write data to JSON file
	_ = utils.WriteToJSON(data, path+"/data.json")
}
