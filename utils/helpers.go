package utils

import (
	"net/url"
	"encoding/json"
)

// queryToJSON converts a query string to a JSON object
func queryToJSON(queryString string) (map[string]interface{}, error) {
    parsedQuery, err := url.ParseQuery(queryString)
    if err != nil {
        return nil, err
    }

    jsonData := make(map[string]interface{})
    for key, values := range parsedQuery {
        if len(values) > 0 {
            jsonData[key] = values[0]
        }
    }

    return jsonData, nil
}

func convertMapToString(dataMap map[string]interface{}) (string, error) {
	jsonBytes, err := json.Marshal(dataMap)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}