package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func NewStorage(connstr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), connstr)
  if err != nil {
    return nil, err
  }

	s := Storage{
		db: db,
	}
  return &s, nil
}

type Task struct {
	ID          int
	Opened      int64
	Closed 			int64
	AuthorId 		int
	AssignedId 	int
	Title     	string
	Content 		string
}

func (s *Storage) Tasks(taskID, authorId int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
    SELECT 
						id, 
						opened, 
						closed, 
						author_id, 
						assigned_id, 
						title, 
						content
    FROM 
					tasks
    WHERE
					($1 = 0 OR id = $1) AND
					($2 = 0 OR author_id = $2)
    ORDER BY id;
		`,
		 taskID, authorId)

  if err != nil {
    return nil, err
  }

  defer rows.Close()

  var tasks []Task
  for rows.Next() {
    var task Task
    err = rows.Scan(
				&task.ID, 
				&task.Opened, 
				&task.Closed, 
				&task.AuthorId, 
				&task.AssignedId, 
				&task.Title, 
				&task.Content,
			)

    if err != nil {
      return nil, err
    }
    tasks = append(tasks, task)
  }
  return tasks, rows.Err()
}

func (s *Storage) NewTask(task Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
    INSERT INTO tasks (title, content)
    VALUES ($1, $2)
    RETURNING id;
    `,
		 task.Title, task.Content,
		 ).Scan(&id)

  if err != nil {
    return 0, err
  }

  return id, nil
}