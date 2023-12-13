package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"jhon", "smith", "lance"})
	_ = writer.Write([]string{"celyn", "adios", "dug"})
	_ = writer.Write([]string{"joko", "morro", "diah"})

	writer.Flush()
}
