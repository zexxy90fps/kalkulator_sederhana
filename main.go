package main

import "fmt"

func main() {
	var num1, num2 int
	var operator string

	//meminta input dari pengguna
	fmt.Println("masukan angka pertamamu: ")
	_, err1 := fmt.Scanln(&num1)
	if err1 != nil {
		fmt.Println("input tidak valid: ", err1)
		return
	}
	fmt.Println("masukan operator (+, -, *, /):")
	_, err2 := fmt.Scanln(&operator)
	if err2 != nil || (operator != "+" && operator != "-" && operator != "*" && operator != "/") {
		fmt.Println("operator tidak valid: ", err2)
		return
	}

	fmt.Println("masukan angka kedua:")
	_, err3 := fmt.Scanln(&num2)
	if err3 != nil {
		fmt.Println("input tidak valid: ", err3)
		return
	}

	//melakukan operasi yang dipilih
	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("error : kesalahan:pembagian nol tidak diperbolehkan.")
			return
		}
		result = num1 / num2
	}

	//menampilkan hasil
	fmt.Printf("Ini lah hasilnya: %d", result)
	fmt.Printf("")

}
