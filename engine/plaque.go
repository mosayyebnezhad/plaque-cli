/*
this is engine of plaque cli powerd by plates.json
*/
package engine

import (
	_ "embed"
	"encoding/json"
)

//go:embed plates.json
var Plates []byte

type PlateInfo struct {
	Province string `json:"province"`
	City     string `json:"city"`
}

func Engine(y, w string) PlateInfo {

	var newJsonPlates map[string]map[string]PlateInfo

	json.Unmarshal(Plates, &newJsonPlates)

	return newJsonPlates[w][y]
}
