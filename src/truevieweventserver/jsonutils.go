package truevieweventserver

import "encoding/json"
import "log"
import "path/filepath"
import "io/ioutil"

func JsonFileToStruct(jsonFilePath string, targetStruct interface{}) error {

	jsonFilePath, errDirectory := filepath.Abs(jsonFilePath)
	if errDirectory != nil {
		log.Fatal(errDirectory)
	} else {
		log.Println("Reading in jsonFile:", jsonFilePath)
	}

	byteArray, errFile := ioutil.ReadFile(jsonFilePath)
	if errFile != nil {
		log.Fatal(errFile)
		return errFile
	}

	errUnmarshal := json.Unmarshal(byteArray, &targetStruct)

	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
		return errUnmarshal
	}

	return nil

}

func JsonStringStruct(jsonStr string, targetStruct interface{}) error {

	var byteArray []byte = []byte(jsonStr)
	errUnmarshal := json.Unmarshal(byteArray, &targetStruct)

	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
		return errUnmarshal
	}

	return nil

}

func StructToJson(config interface{}) (string, error) {

	configJson, errorMarshal := json.Marshal(config)

	if errorMarshal != nil {
		log.Fatal(errorMarshal)
		return "", errorMarshal
	}

	configJsonString := string(configJson[:])

	return configJsonString, nil
}
