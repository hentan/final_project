package models

import (
	jsi "github.com/json-iterator/go"
)

type Deserializable interface {
	Deserialize(data []byte) error
}

func (b *Book) Deserialize(data []byte) error {
	return jsi.Unmarshal(data, b)
}

func (a *Author) Deserialize(data []byte) error {
	return jsi.Unmarshal(data, a)
}

type Serializable interface {
	Serialize() ([]byte, error)
}

func (b *Book) Serialize() ([]byte, error) {
	return jsi.Marshal(b)
}

func (a *Author) Serialize() ([]byte, error) {
	return jsi.Marshal(a)
}
