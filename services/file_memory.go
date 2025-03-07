package services

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type JsonRecord struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}

var OLIVER_PATH = ".oli"

func FindJsonMemoryRecord(path string, id string) (*ChatConversation, error) {
	filePath := filepath.Join(path, fmt.Sprintf("%v.json", id))

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var cc ChatConversation

	err = json.Unmarshal(byteValue, &cc)
	if err != nil {
		return nil, fmt.Errorf("error parsing json: %v", err)
	}

	return &cc, nil
}

func SaveJsonMemoryRecord(path string, cc *ChatConversation) error {
	jsonData, err := json.MarshalIndent(cc, "", "  ")
	if err != nil {
		return fmt.Errorf("error converting to json: %v", err)
	}

	filePath := filepath.Join(path, fmt.Sprintf("%v.json", cc.Id))

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}

	return nil
}
