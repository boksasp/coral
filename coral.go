package coral

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Filter(jsonString string, key string, value string, matchNull bool) (bool, error) {
	keys := strings.Split(key, ".")
	var data map[string]interface{}

	if err := json.Unmarshal([]byte(jsonString), &data); err != nil {
		return false, err
	}

	_, isMatch := filter(data, keys, value, matchNull)

	return isMatch, nil
}

func filter(data map[string]interface{}, keys []string, match string, matchNull bool) (map[string]interface{}, bool) {
	if len(keys) == 1 {
		value, exists := data[keys[0]]
		if exists {
			if matchNull {
				return nil, value == nil
			} else {
				return nil, fmt.Sprint(value) == match
			}
		}
		return make(map[string]interface{}), false
	}

	var nextNode map[string]interface{}
	value, exists := data[keys[0]]

	if exists {
		// Check for an array in the value.
		// The filter must run on each element in the array.
		if strings.HasPrefix(fmt.Sprint(value), "[") {
			arrayNode := value.([]interface{})
			for i := 0; i < len(arrayNode); i++ {
				arrayElement := arrayNode[i].(map[string]interface{})
				result, isMatch := filter(arrayElement, keys[1:], match, matchNull)
				if isMatch {
					return result, isMatch
				}
			}
		} else {
			nextNode = value.(map[string]interface{})
		}
	}

	return filter(nextNode, keys[1:], match, matchNull)
}
