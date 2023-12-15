package main

import "fmt"

func main() {
	// Membuat slice daro array yang sudah ada

	// Ada variable name yang berisikan array (names) bertype string
	// yang tidak ditentukan panjang nya di awal
	// tetapi ditentukan ketika mengisikan data
	names := [...]string{"Jhon", "Smith", "Lance", "Abram", "Stif", "Kevin"}

	// Variable slice yang berisikan slice dari array names
	// array[low:high] 0,1,2
	slice := names[0:3]
	fmt.Println(slice)

	// array[low:]
	slice1 := names[4:]
	fmt.Println("array[low:]", slice1)

	// array[:high]
	slice2 := names[:4]
	fmt.Println("array[:high]", slice2)

	// array[:]
	slice3 := names[:]
	fmt.Print("array[:]", slice3)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	sliceDays := days[5:]
	fmt.Println(sliceDays)
	sliceDays[0] = "Sabtu Baru"
	sliceDays[1] = "Minggu Baru"
	fmt.Println("Day", days)

	sliceDays1 := append(sliceDays, "Libur Baru")
	fmt.Println("Day", sliceDays1)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Jhon"
	newSlice[1] = "Jhon"

	fmt.Println(newSlice)

}
