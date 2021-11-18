package main

import "log"
import "os"
import "fmt"
import "bitbucket.org/ix-specops/golangeventserver/src/truevieweventserver"

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Must call this with the location of a config.json file i.e.  trueview.exe myconifg.json")
		os.Exit(0)
	}

	var configFile string = os.Args[1]
	log.Println("Using config file:", configFile)

	// Get full path of test.json
	targetDestConfig := []truevieweventserver.Configuration{}
	errJsonFileToStruct := truevieweventserver.JsonFileToStruct(configFile, &targetDestConfig)

	if errJsonFileToStruct != nil {
		log.Fatal(errJsonFileToStruct)
	}

	targetDestConfigArrayLength := len(targetDestConfig)

	// Sping it all off using goroutines
	for i := 0; i < targetDestConfigArrayLength; i += 1 {

		target := targetDestConfig[i]

		go func(value truevieweventserver.Configuration) {
			truevieweventserver.RunConfiguration(value)
		}(target)

	}

	log.Println("*****> Running ....")

	truevieweventserver.WaitForSystemSignals()

}
