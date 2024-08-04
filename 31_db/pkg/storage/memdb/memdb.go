package memdb

import "modules/pkg/storage/postgres"

type DB []postgres.Task

func (db DB) Tasks(int, int) ([]postgres.Task, error) {
	return db, nil
}

func (db DB) NewTask(postgres.Task) (int, error) {
	return 0, nil
}

func (db DB) DeleteTask(postgres.Task) (int, error) {
	return 0, nil
}

func (db DB) EditTaskTask(postgres.Task) (int, error) {
  return 0, nil
}

func (db DB) TasksByLabel(label string) ([]postgres.Task, error) {
	return db, nil
}