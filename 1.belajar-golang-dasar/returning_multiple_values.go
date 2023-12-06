package main

func getFullName() (string, string) {

	return "Jhon", "Smith"
}
func main() {
	firstName, lastName := getFullName()
	println(firstName, lastName)
}
