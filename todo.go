package main

import (
	"errors"
	"github.com/aquasecurity/table"
	"os"
	"strconv"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) valueAtIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		return errors.New("Invalid Index")
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.valueAtIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos

	if err := t.valueAtIndex(index); err != nil {
		return err
	}
	t[index].Completed = !t[index].Completed
	ct := time.Now()
	t[index].CompletedAt = &ct
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos

	if err := t.valueAtIndex(index); err != nil {
		return err
	}
	t[index].Title = title
	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "CreatedAt", "CompletedAt")

	for i, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed == true {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(i), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
