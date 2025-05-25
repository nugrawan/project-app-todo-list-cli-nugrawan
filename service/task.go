package service

import (
	"errors"
	"fmt"
	"projectgo/model"
	"projectgo/repository"
	"strings"
	"time"
)

var tasks []model.Task

// fungsi untuk mengambil daftar tugas
func Load() error {
	var err error
	tasks, err = repository.LoadTasks()
	return err
}

// menambahkan tugas
func Add(title, desc string) error {
	if strings.TrimSpace(title) == "" {
		return errors.New("judul tidak boleh kosong")
	}
	for _, t := range tasks {
		if t.Title == title {
			return errors.New("judul sudah digunakan")
		}
	}
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}
	t := model.Task{
		ID:          id,
		Title:       title,
		Description: desc,
		CreatedAt:   time.Now(),
	}
	tasks = append(tasks, t)
	return repository.SaveTasks(tasks)
}

// menampilkan daftar tugas
func List(keyword string) []model.Task {
	var result []model.Task
	for _, t := range tasks {
		if keyword == "" || strings.Contains(strings.ToLower(t.Title+t.Description), strings.ToLower(keyword)) {
			result = append(result, t)
		}
	}
	return result
}

// menandai tugas sebagai selesai
func MarkDone(id int) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Completed = true
			return repository.SaveTasks(tasks)
		}
	}
	return fmt.Errorf("tugas dengan ID %d tidak ditemukan", id)
}

// menghapus tugas
func Delete(id int) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return repository.SaveTasks(tasks)
		}
	}
	return fmt.Errorf("tugas dengan ID %d tidak ditemukan", id)
}
