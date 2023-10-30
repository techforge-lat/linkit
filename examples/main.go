package main

import (
	"fmt"
	"log"

	"github.com/techforge-lat/dependor"
)

func main() {
	foo := &Foo{}
	dependor.Set[*Foo](dependor.Config{
		Value: foo,
		AuxiliaryDependencies: map[string]string{
			"Bar": dependor.Name(Bar{}),
		},
	})

	dependor.Set[*Bar](dependor.Config{})

	if err := dependor.SetAuxiliarDependencies(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(dependor.Name(Foo{}))
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
