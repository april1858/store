package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	store := NewKeyValueStore()
	scanner := bufio.NewScanner(os.Stdin)

	if err := scanner.Err(); err != nil {
		fmt.Println("scanner.Err() - ", err)
	}

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 0 {
			continue
		}

		switch strings.ToUpper(fields[0]) {
		case "PUT": // PUT key value
			if len(fields) != 3 {
				fmt.Println("usage: PUT key value")
				continue
			}
			if err := store.Put(fields[1], fields[2]); err != nil {
				fmt.Println("error:", err)
			}

		case "GET": // GET key
			if len(fields) != 2 {
				fmt.Println("usage: GET key")
				continue
			}
			value, err := store.Get(fields[1])
			if err != nil {
				fmt.Println("error:", err)
				continue
			}
			fmt.Println(value)

		case "DELETE": // DELETE key
			if len(fields) != 2 {
				fmt.Println("usage: DELETE key")
				continue
			}
			if err := store.Delete(fields[1]); err != nil {
				fmt.Println("error:", err)
			}

		default:
			fmt.Println("unknown command:", fields[0])
		}
	}
}
