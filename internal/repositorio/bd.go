package repositorio

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}

// salvando os dados
func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "")

	if err != nil {
		return err
	}

	return os.WriteFile(s.FileName, fileData, 0644)
}

// carregando os dados
func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.FileName)

	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}
