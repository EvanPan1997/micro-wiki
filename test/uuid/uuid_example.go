package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {
	uid := uuid.NewV4().String()
	fmt.Println(uid)
}
