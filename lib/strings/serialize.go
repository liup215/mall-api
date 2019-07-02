package strings

import (
	"encoding/json"

	"github.com/techleeone/gophp/serialize"
)

func PhpSerialize(d interface{}) ([]byte, error) {
	dataByte, err := json.Marshal(d)
	if err != nil {
		return []byte{}, err
	}

	var m map[string]interface{}
	err = json.Unmarshal(dataByte, &m)
	if err != nil {
		return []byte{}, err
	}
	jsonByte, err := serialize.Marshal(m)
	return jsonByte, err

}

func PhpUnserialize(data []byte, v interface{}) error {
	i, err := serialize.UnMarshal(data)
	if err != nil {
		return err
	}

	jsonByte, err := json.Marshal(i)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonByte, &v)
}
