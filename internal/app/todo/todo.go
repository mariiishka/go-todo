package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(id int) error {
	list := *t
	index := id - 1

	if id <= 0 || id > len(list) {
		return errors.New("invalid index")
	}

	if list[index].Done {
		return errors.New("todo is already completed")
	}

	list[index].CompletedAt = time.Now()
	list[index].Done = true

	return nil
}

func (t *Todos) Delete(id int) error {
	list := *t
	index := id - 1
	lastIndex := len(list) - 1

	if index < 0 || index > lastIndex {
		return errors.New("invalid index")
	}

	*t = append(list[:index], list[index+1:]...)

	return nil
}

func (t *Todos) Load(filename string) error {
	const op = "internal.todo.Load"

	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return fmt.Errorf("%s: %w", op, err)
	}

	if len(file) == 0 {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Store(filename string) error {
	const op = "internal.todo.Store"

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
