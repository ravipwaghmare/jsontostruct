package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jsontostruct/ironic"
	"os"
)

func main() {
	jsonFile, err := os.Open("introspectedData.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	jsonString, _ := ioutil.ReadAll(jsonFile)

	ironicData := ironic.Data{}
	json.Unmarshal([]byte(jsonString), &ironicData)
	fmt.Println(ironicData)
}
