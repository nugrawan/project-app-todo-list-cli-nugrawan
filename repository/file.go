package repository

import (
	"encoding/json"
	"os"
	"path/filepath"
	"projectgo/model"
)

var (
	dataDir  = "data"
	dataFile = filepath.Join(dataDir, "tasks.json")
)

// membuat direktori data jika belum ada
func ensureDataDir() error {
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		return os.Mkdir(dataDir, 0755)
	}
	return nil
}

// mengembalikan daftar tugas
func LoadTasks() ([]model.Task, error) {
	if err := ensureDataDir(); err != nil {
		return nil, err
	}
	file, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Task{}, nil
		}
		return nil, err
	}
	var tasks []model.Task
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

// menyimpan daftar tugas
func SaveTasks(tasks []model.Task) error {
	if err := ensureDataDir(); err != nil {
		return err
	}
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}
