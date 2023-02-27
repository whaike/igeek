package pkg

import "encoding/json"

type Any interface{}

func JsonFormat(data Any) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
