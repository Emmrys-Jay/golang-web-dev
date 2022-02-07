package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	lstner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer lstner.Close()

	for {
		conn, err := lstner.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
			//remember to put continue so the current iteration is skipped
			//therefore the line where a string is being wriiten to conn wont run
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
			if ln == "" {
				break
			}
		}

		fmt.Println("Code got here.")
		io.WriteString(conn, "I see you connected.")

		conn.Close()
	}
}
