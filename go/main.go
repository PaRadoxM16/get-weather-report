package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Jeffail/gabs"
)

func getMetar(station string) []byte {
	url := fmt.Sprintf("https://api.paradox.ovh/metar?station=%s", station)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return bodyText
}

func main() {
	// Check if there is a user-supplied argument
	if len(os.Args) != 2 {
		panic("First argument should be an ICAO station.")
	}

	// Parse Json from metar API
	station := os.Args[1]
	result, err := gabs.ParseJSON(getMetar(station))

	if err != nil {
		panic(err)
	}

	// Search for the full description
	// If nil, the msg is printed
	path := "details.descriptions.full_description"
	full_description := result.Path(path).Data()

	if full_description == nil {
		log.Fatal(result.Search("msg"))
	} else {
		fmt.Println(full_description)
	}
}
