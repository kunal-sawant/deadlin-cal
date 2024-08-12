package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"os"
	"time"

	"github.com/kunal-sawant/deadlin-cal/genscripts"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed scripts\schema.sql
var ddl string

func main() {

	fileName := "deadline_cal.sqlite3"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating or opening the file:", err)
		return
	}
	defer file.Close()

	ctx := context.Background()

	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		panic(err)
	}

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		panic(err)
	}

	queries := genscripts.New(db)

	task, err := queries.CreateTask(ctx, genscripts.CreateTaskParams{
		TaskName:  "sell tata stock",
		StartDate: time.Now().Format("02/01/2006"),
		EndDate:   "31/01/2025",
	})

	task, err = queries.CreateTask(ctx, genscripts.CreateTaskParams{
		TaskName:  "renew insurance",
		StartDate: time.Now().Format("02/01/2006"),
		EndDate:   "10/11/2024",
	})

	fmt.Print(task)

	fetchedTask, err := queries.GetAllTasks(ctx)
	fmt.Println(fetchedTask)
}
