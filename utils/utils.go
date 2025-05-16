package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func ReadFromJSON(filename string) (AirPolutions, error) {
	var data AirPolutions
	var dataByte []byte
	var err error

	// Empty data to return if file does not exist
	var emptyData AirPolutions

	dataByte, err = os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			initialData, errM := json.Marshal(emptyData)
			if errM != nil {
				return emptyData, errM
			}
			err = os.WriteFile(filename, initialData, 0644)
			if err != nil {
				return emptyData, err
			}
			return emptyData, nil
		}
		return data, err
	}

	err = json.Unmarshal(dataByte, &data)
	if err != nil {
		return emptyData, err
	}

	return data, nil
}

func WriteToJSON(data AirPolutions, filename string) error {
	var dataByte []byte
	var err error
	var nonEmptyData []AirPolution

	nonEmptyData = filterNonEmpty(data)

	dataByte, err = json.Marshal(nonEmptyData)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, dataByte, 0644)
	if err != nil {
		return err
	}

	return nil
}

func filterNonEmpty(data AirPolutions) []AirPolution {
	var entry AirPolution
	var n int = 0

	for _, entry = range data {
		if entry.AqiID != "" {
			data[n] = entry
			n++
		}
	}

	return data[:n]
}

func PaginateData(data AirPolutions, page int) []AirPolution {
	var perPage, start, end int
	var filteredData []AirPolution

	perPage = 5
	start = (page - 1) * perPage
	filteredData = filterNonEmpty(data)

	if start >= len(filteredData) {
		return []AirPolution{}
	}
	end = start + perPage
	if end > len(filteredData) {
		end = len(filteredData)
	}
	return filteredData[start:end]
}

func GetNonEmptyInput(scanner *bufio.Scanner, prompt string) string {
	var input string
	for strings.TrimSpace(input) == "" {
		fmt.Print(prompt)
		scanner.Scan()
		input = scanner.Text()

		if strings.TrimSpace(input) == "" {
			fmt.Println("Data tidak boleh kosong.")
		}
	}
	return input
}

func GetIntInput(scanner *bufio.Scanner, prompt string) int {
	var input int
	var inputTrim string
	var err error

	for strings.TrimSpace(inputTrim) == "" {
		fmt.Print(prompt)
		scanner.Scan()
		inputTrim = scanner.Text()

		if strings.TrimSpace(inputTrim) == "" {
			fmt.Println("Data tidak boleh kosong.")
		}

		if input, err = strconv.Atoi(inputTrim); err != nil {
			fmt.Println("Input tidak valid. Harap masukkan angka.")
			inputTrim = ""
		}
	}
	return input
}

func randomID(length int) string {
	var i int

	var charset = "abcdefghijklmnopqrstuvwxyz"
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	var result = make([]byte, length)

	for i = range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

func SortDescendingByIdxUdara(A *AirPolutions) {
	sort.Slice(*A, func(i, j int) bool {
		return (*A)[i].IdxUdara > (*A)[j].IdxUdara
	})
}

func SortAscendingByIdxUdara(A *AirPolutions) {
	sort.Slice(*A, func(i, j int) bool {
		return (*A)[i].IdxUdara < (*A)[j].IdxUdara
	})
}

func SequentialSearch(data AirPolutions, target string) *AirPolution {
	for _, item := range data {
		if strings.EqualFold(item.Lokasi, target) {
			return &item
		}
	}
	return nil
}
