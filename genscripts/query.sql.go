// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package genscripts

import (
	"context"
)

const createTask = `-- name: CreateTask :one
INSERT INTO Tasks (TASK_NAME, START_DATE, END_DATE) VALUES
(?, ?, ?) RETURNING id, task_name, start_date, end_date
`

type CreateTaskParams struct {
	TaskName  string
	StartDate string
	EndDate   string
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask, arg.TaskName, arg.StartDate, arg.EndDate)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.TaskName,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const deleteTaskName = `-- name: DeleteTaskName :exec
DELETE FROM Tasks WHERE ID = ?
`

func (q *Queries) DeleteTaskName(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTaskName, id)
	return err
}

const getAllTasks = `-- name: GetAllTasks :many
SELECT id, task_name, start_date, end_date FROM Tasks
`

func (q *Queries) GetAllTasks(ctx context.Context) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getAllTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.TaskName,
			&i.StartDate,
			&i.EndDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTask = `-- name: GetTask :one
SELECT id, task_name, start_date, end_date FROM Tasks WHERE ID = ? LIMIT 1
`

func (q *Queries) GetTask(ctx context.Context, id int64) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTask, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.TaskName,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const getTaskByName = `-- name: GetTaskByName :one
SELECT id, task_name, start_date, end_date FROM Tasks WHERE TASK_NAME = ? LIMIT 1
`

func (q *Queries) GetTaskByName(ctx context.Context, taskName string) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTaskByName, taskName)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.TaskName,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const updateTaskName = `-- name: UpdateTaskName :exec
UPDATE Tasks SET TASK_NAME = ? WHERE ID = ?
`

type UpdateTaskNameParams struct {
	TaskName string
	ID       int64
}

func (q *Queries) UpdateTaskName(ctx context.Context, arg UpdateTaskNameParams) error {
	_, err := q.db.ExecContext(ctx, updateTaskName, arg.TaskName, arg.ID)
	return err
}
