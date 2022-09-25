package cmd

import (
	"fmt"
	"net/http"
	"os"
)

const BaseURL = "https://api.met.gov.my/v2.1"

func RunCLI() {
	key := os.Getenv("METMY_KEY")
	fmt.Println("KEY: ", key)
	// curl -H "Authorization: METToken <METMY_KEY>" \
	//		"https://api.met.gov.my/v2/data?datasetid=FORECAST&datacategoryid=GENERAL&locationid=LOCATION:237&start_date=2022-09-14&end_date=2022-09-14"
	rangeStr := "start_date=2022-09-14&end_date=2022-09-14"
	URL := fmt.Sprintf("%s/data?datasetid=FORECAST&datacategoryid=GENERAL&locationid=LOCATION:237&%s", BaseURL, rangeStr)
	fmt.Println("URL: ", URL)
	//http.Head()
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	fmt.Println("STATUS: ", resp.Status)
	if resp.StatusCode == http.StatusUnauthorized {
		fmt.Println("lets br tass")
	}
	//spew.Dump(resp.Body)
	defer resp.Body.Close()
}
