package main

import (
	"flag"
	"fmt"
)

func main() {

	username := flag.String("username", "root", "database username")
	password := flag.String("password", "root", "Password")
	hostname := flag.String("hostname", "localhost", "host name")
	port := flag.Int("port", 8080, "port")

	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Hostname", *hostname)
	fmt.Println("Port", *port)
}
