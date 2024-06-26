package app

import (
	"database/sql"
	"fmt"
)

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) InsertTask(t *Task) (int, error) {
	stmt, err := r.db.Prepare("INSERT INTO task (title, description, status, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	res, err := stmt.Exec(t.Title, t.Description, t.Status, t.CreatedAt)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *taskRepository) GetTasks() ([]Task, error) {
	rows, err := r.db.Query("SELECT * FROM task")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *taskRepository) UpdateTask(t *Task) error {
	// TODO: idが存在しないときの処理を検討する
	stmt, err := r.db.Prepare("UPDATE task SET title = ?, description = ?, status = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(t.Title, t.Description, t.Status, t.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *taskRepository) DeleteTask(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM task WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
