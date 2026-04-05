package storage

import (
	"encoding/json"
	"os"

	"github.com/danchenkov/todo-gin/internal/models"
)

const filePath = "data/todos.json"

func LoadTodos() ([]models.Todo, error) {
	var todos []models.Todo

	file, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Todo{}, nil
		}
		return nil, err
	}

	err = json.Unmarshal(file, &todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func SaveTodos(todos []models.Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}
