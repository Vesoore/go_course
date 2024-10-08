package main

import (
	"fmt"
	"less_1/internal/pkg/storage"
	"log"
)

func main() {

	s, err := storage.NewStorage()
	if err != nil {
		log.Fatal(err)
	}
	s.Set("key_1", "value_1")
	s.Set("key_2", "value_2")
	s.Set("key_3", "12.3")
	s.Set("key_4", "123")
	st := s.GetKind("key_1")
	num_1 := s.GetKind("key_3")
	num_2 := s.GetKind("key_4")
	fmt.Println(st, num_1, num_2)
}
