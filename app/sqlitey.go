package app

import "fmt"

type Sqlitey struct{}

var log = fmt.Println

func New() *Sqlitey {
	return &Sqlitey{}
}

func (s *Sqlitey) Run() {
	log("run")
}
