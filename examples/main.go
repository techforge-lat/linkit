package main

import (
	"fmt"
	"log"

	"github.com/techforge-lat/dependor"
)

func main() {
	foo := &Foo{}
	dependor.Set[*Foo]("foo", foo, map[string]string{
		"Bar": "bar_storage",
	})

	bar := Bar{}
	dependor.Set[Bar]("bar_storage", bar, make(map[string]string, 0))

	if err := dependor.Populate(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(foo.execute())
}

type Foo struct {
	Bar Storage
}

func (f *Foo) SetStorage(s Storage) {
	f.Bar = s
}

func (f Foo) execute() string {
	return f.Bar.GetHello()
}

type Storage interface {
	GetHello() string
}

type Bar struct{}

func (b Bar) GetHello() string {
	return "Hello World!"
}
