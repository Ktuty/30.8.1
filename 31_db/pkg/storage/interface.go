package storage

import "modules/pkg/storage/postgres"

type Interface interface {
	Tasks(int, int) ([]postgres.Task, error)
	NewTask(postgres.Task) (int, error)
	DeleteTask(postgres.Task) (int, error)
	EditTask(postgres.Task) (int, error)
	TasksByLabel(postgres.Label) ([]postgres.Task, error)
}