package utils

import (
	"encoding/json"
	"os"
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

	nonEmptyData = FilterNonEmpty(data)

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
