package main

import (
	"fmt"
	"math/rand"
	"time"
)

var users = []string{
	"",
}

func main() {
	sorted := randomSort(users)
	for _, u := range sorted {
		fmt.Println(u)
	}
}

func randomSort(users []string) []string {
	sorted := make([]string, 0, len(users))
	rand.Seed(time.Now().UnixNano())
	for _ = range users {
		index := rand.Intn(len(users))
		sorted = append(sorted, users[index])
		if index < len(users) {
			users = append(users[:index], users[index+1:]...)
		}
	}
	return sorted
}
