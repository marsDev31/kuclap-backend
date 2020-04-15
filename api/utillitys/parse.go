package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"
	"github.com/marsDev31/kuclap-backend/api/models"
)


func main(){
	var classesOld []models.OldClass
	// var classesNew models.Classes
	var result []models.Class

	// This I/O for reading from old data (KUnit version)
	data, err := ioutil.ReadFile("../classes.json")
	if err != nil {
      fmt.Print("error (read file)",err)
	}
	if err := json.Unmarshal(data, &classesOld); err != nil {
        fmt.Println("error (old-classes unmarshal):", err)
	}

	// Make slice of Classes (KUclap version)
	for _, class := range classesOld {
		new := models.Class{
			class.Value,
			strings.Fields(class.Label)[1],
			strings.ReplaceAll(strings.Split(class.Label, "(")[1], ")", ""),
			class.Label}
		result = append(result, new)
	}
	file, _ := json.MarshalIndent(result, "", "	")
	_ = ioutil.WriteFile("../classesParsed.json", file, 0644)
	
}