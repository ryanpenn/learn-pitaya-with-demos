package serializer

import "github.com/vmihailenco/msgpack"

type Serializer struct{}

// NewMsgPackSerializer returns a new Serializer.
func NewMsgPackSerializer() *Serializer {
	return &Serializer{}
}

// Marshal returns the protobuf encoding of v.
func (s *Serializer) Marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

// Unmarshal parses the protobuf-encoded data and stores the result
// in the value pointed to by v.
func (s *Serializer) Unmarshal(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}

// GetName returns the name of the serializer.
func (s *Serializer) GetName() string {
	return "msgpack"
}
