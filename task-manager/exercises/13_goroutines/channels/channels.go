package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "Mesaj 1"
        messages <- "Mesaj 2"
        messages <- "Mesaj 3"
		close(messages)
	}()
	for msg := range messages {
		fmt.Println(msg)
	}
}