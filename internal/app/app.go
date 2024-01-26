package app

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/mariiishka/go-todo/internal/app/todo"
	"github.com/mariiishka/go-todo/internal/lib/stringspretty"
	"github.com/mariiishka/go-todo/internal/lib/tablespretty"
)

const (
	todoFile = "todos.json"
)

func Run() {
	var (
		add      bool
		complete int
		delete   int
		list     bool
	)

	flag.BoolVar(&add, "add", false, "add a new todo")
	flag.IntVar(&complete, "complete", 0, "mark todo as completed")
	flag.IntVar(&delete, "delete", 0, "delete todo")
	flag.BoolVar(&list, "list", false, "list all todos")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case add:
		task, err := Input(os.Stdin, flag.Args()...)

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		todos.Add(task)
		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		fmt.Fprintln(os.Stdout, "task added successfully!")
	case complete > 0:
		err := todos.Complete(complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		fmt.Printf("todo number %d completed\n", complete)
	case delete > 0:
		err := todos.Delete(delete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		err = todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		fmt.Printf("todo number %d deleted\n", delete)
	case list:
		err := PrintTodos(todos)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(1)
	}
}

func PrintTodos(t *todo.Todos) error {
	if len(*t) == 0 {
		return errors.New("todo list is empty")
	}

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "TASK"},
			{Align: simpletable.AlignCenter, Text: "DONE"},
			{Align: simpletable.AlignCenter, Text: "CREATED AT"},
			{Align: simpletable.AlignCenter, Text: "COMPLETED AT"},
		},
	}

	for i, todo := range *t {
		var isDone string
		var completedAt string

		if todo.Done {
			isDone = stringspretty.Green(strconv.FormatBool(todo.Done))
			completedAt = stringspretty.Green(todo.CompletedAt.Format(time.ANSIC))
		} else {
			isDone = stringspretty.Red(strconv.FormatBool(todo.Done))
			completedAt = stringspretty.Red(todo.CompletedAt.Format(time.ANSIC))
		}

		row := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: stringspretty.Blue(fmt.Sprintf("%d", i+1))},
			{Align: simpletable.AlignLeft, Text: todo.Task},
			{Align: simpletable.AlignLeft, Text: isDone},
			{Align: simpletable.AlignLeft, Text: stringspretty.Blue(todo.CreatedAt.Format(time.ANSIC))},
			{Align: simpletable.AlignLeft, Text: completedAt},
		}

		table.Body.Cells = append(table.Body.Cells, row)
	}

	table.SetStyle(tablespretty.StyleDefaultColorful)
	fmt.Fprintln(os.Stdout, table.String())
	return nil
}

func Input(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("empty todo is not allowed")
	}

	return text, nil
}
