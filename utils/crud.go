package utils

import (
	"time"
)

const NMAX int = 100

type AirPolution struct {
	AqiID         string    `json:"aqiID"`
	Lokasi        string    `json:"lokasi"`
	SumberPolusi  string    `json:"sumberPolusi"`
	TingkatBahaya string    `json:"tingkatBahaya"`
	IdxUdara      int       `json:"IdxUdara"`
	Waktu         time.Time `json:"waktu"`
}

type AirPolutions [NMAX]AirPolution

func AddData(data *AirPolutions, lokasi, sumberPolusi string, IdxUdara int) {
	var tingkat string
	var nonEmptyData []AirPolution

	nonEmptyData = filterNonEmpty(*data)
	var lastIdx int = len(nonEmptyData)

	if IdxUdara >= 0 && IdxUdara <= 50 {
		tingkat = "Baik"
	} else if IdxUdara >= 51 && IdxUdara <= 100 {
		tingkat = "Sedang"
	} else if IdxUdara >= 101 && IdxUdara <= 150 {
		tingkat = "Tidak Sehat"
	} else {
		tingkat = "Berbahaya"
	}

	(*data)[lastIdx].AqiID = randomID(5)
	(*data)[lastIdx].Lokasi = lokasi
	(*data)[lastIdx].SumberPolusi = sumberPolusi
	(*data)[lastIdx].IdxUdara = IdxUdara
	(*data)[lastIdx].Waktu = time.Now()
	(*data)[lastIdx].TingkatBahaya = tingkat
}

func EditData(data *AirPolutions, lokasi, sumberPolusi string, IdxUdara int, aqiID string) {
	var i int
	var user AirPolution

	for i, user = range *data {
		if user.AqiID == aqiID {
			(*data)[i].Lokasi = lokasi
			(*data)[i].SumberPolusi = sumberPolusi
			(*data)[i].IdxUdara = IdxUdara
			(*data)[i].Waktu = time.Now()

			if IdxUdara >= 0 && IdxUdara <= 50 {
				(*data)[i].TingkatBahaya = "Baik"
			} else if IdxUdara >= 51 && IdxUdara <= 100 {
				(*data)[i].TingkatBahaya = "Sedang"
			} else if IdxUdara >= 101 && IdxUdara <= 150 {
				(*data)[i].TingkatBahaya = "Tidak Sehat"
			} else {
				(*data)[i].TingkatBahaya = "Berbahaya"
			}
		}
	}
}

func DeleteData(data *AirPolutions, aqiID string) {
	var i, j int
	var user AirPolution

	for i, user = range *data {
		if user.AqiID == aqiID {
			for j = i; j < len(*data)-1; j++ {
				(*data)[j] = (*data)[j+1]
			}
			(*data)[len(*data)-1] = AirPolution{}
		}
	}
}
