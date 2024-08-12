-- name: GetTask :one
SELECT * FROM Tasks WHERE ID = ? LIMIT 1;

-- name: GetTaskByName :one
SELECT * FROM Tasks WHERE TASK_NAME = ? LIMIT 1;

-- name: GetAllTasks :many
SELECT * FROM Tasks;

-- name: CreateTask :one
INSERT INTO Tasks (TASK_NAME, START_DATE, END_DATE) VALUES
(?, ?, ?) RETURNING *;

-- name: UpdateTaskName :exec
UPDATE Tasks SET TASK_NAME = ? WHERE ID = ?;

-- name: DeleteTaskName :exec
DELETE FROM Tasks WHERE ID = ?;