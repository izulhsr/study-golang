package main

import "fmt"

type ValidationError struct {
	message string
}

func (v *ValidationError) Error() string {
	return v.message
}

type NotFoundError struct {
	message string
}

func (v *NotFoundError) Error() string {
	return v.message
}
func SaveData(id string, data any) error {
	if id == "" {
		return &ValidationError{"validation error"}
	}

	if id != "jhon" {
		return &NotFoundError{"data not found"}
	}

	// ok

	return nil
}
func main() {
	err := SaveData("jhon", nil)
	if err != nil {
		// terjadi error
		//if validationErr, ok := err.(*validationError); ok {
		//	fmt.Println("validation error:", validationErr.Error())
		//} else if notFoundErr, ok := err.(*notFoundError); ok {
		//	fmt.Println("not found error:", notFoundErr.Error())
		//} else {
		//	fmt.Println("unknown error:", err.Error())
		//}

		switch finalError := err.(type) {
		case *ValidationError:
			fmt.Println("validation error:", finalError.Error())
		case *NotFoundError:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("unknown error:", finalError.Error())
		}

	} else {
		// sukse
		fmt.Println("Sukses")
	}
}
