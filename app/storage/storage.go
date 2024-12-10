package storage

import (
	"fmt"
)

var Storage = make(map[string]*Cell, 10000)

// struct
/*func init() {

	Storage = make(map[string]Cell)

}*/

func Lookup() {

	for k, v := range Storage {
		fmt.Println(k, v)
	}
}

func Info() {

	fmt.Printf("Capacity: %v\n", len(Storage))
}
