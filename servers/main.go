/**********
* Servers - Basic
**********/
package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Important Go pattern when caching a method
	// var name, err then if err != nil

	//We are establishing a tcp server with the net package
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nHello from TCP Server\n")
		fmt.Println(conn, "How was your day?")
		fmt.Fprintf(conn, "%v", "well I hope!")

		conn.Close()
	}

}
