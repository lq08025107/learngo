package search

import (
	"encoding/json"
	"os"
)

const dataFile = "H:\\workspace\\go\\learngo\\ch2\\sample\\data\\data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	file, err = os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	defer file.close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
