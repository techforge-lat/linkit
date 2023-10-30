package main

import (
	"fmt"
	"log"

	"github.com/techforge-lat/linkit"
)

func main() {
	foo := &Foo{}
	linkit.Set[*Foo](
		linkit.WithValue(foo),
		linkit.WithAuxiliaryDependencies(map[string]string{
			"Bar": linkit.Name(Bar{}),
		}),
	)

	linkit.Set[*Bar]()

	if err := linkit.SetAuxiliaryDependencies(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(linkit.Name(Foo{}))
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
