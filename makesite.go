package main

import (
	"html/template"
	"os"
)

type entry struct {
	Name string
	Done bool
}

type ToDo struct {
	User string
	List []entry
}

func main() {
	paths := []string{
		"todo.tmpl",
	}

	todos := []entry{
		entry{Name: "do this", Done: false},
	}
	user := ToDo{
		User: "Me myself and I",
		List: todos,
	}

	t := template.Must(template.New("todo.tmpl").ParseFiles(paths...))
	err := t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
