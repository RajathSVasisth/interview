package main

import (
	"log"

	models "github.com/RajathSVasisth/interview/models"
	flag "github.com/spf13/pflag"
)

func main() {
	endpoint := flag.String("endpoint", "", "Logging endpoint")
	debug := flag.Bool("debug", false, "Debug flag used to run locally")
	days := flag.Int("days", 7, "Number of days of data to parse")
	flag.Parse()

	var data []models.IndexInfo
	var err error

	if *debug {
		data, err = getDataFromFile("example-in.json")
		if err != nil {
			log.Fatalln("Error reading data from file. Error: ", err.Error())
		}
	} else {
		data, err = getDataFromServer(*endpoint, *days)
		if err != nil {
			log.Fatalln("Error reading data from API endpoint. Error: ", err.Error())
		}
	}

	printLargestIndexes(data)
	printMostShards(data)
	printLeastBalanced(data)
}
