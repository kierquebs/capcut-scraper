package util

import (
	"encoding/json"
	"net/url"
)

// DataDecoder DataDecoder converts Unescaped Query into json object
func DataDecoder(data string) (interface{}, error) {
	decodedStr, err := url.QueryUnescape(data)
	if err != nil {
		return nil, err
	}


	// Parse JSON
	var jsonData map[string]interface{}
	err = json.Unmarshal([]byte(decodedStr), &jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}