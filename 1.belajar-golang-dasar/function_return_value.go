package main

func getHello(firstName string) string {
	hello := "Hello " + firstName
	return hello
}
func main() {
	println(getHello("Jhon"))
}
