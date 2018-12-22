package json

import (
	"encoding/json"
	"io"
)

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func MarshalWriter(writer io.Writer, v interface{}) error {
	return json.NewEncoder(writer).Encode(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func UnmarshalReader(reader io.Reader, v interface{}) error {
	return json.NewDecoder(reader).Decode(v)
}
