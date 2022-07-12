package main

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	Nama     string
	Email    string
	Password string
}

var (
	arrData []User
)

func main() {
	var isRunning = true
	var inputMenu int
	scanner := bufio.NewScanner(os.Stdin)

	for isRunning {
		fmt.Println("1. Register")
		fmt.Println("2. Lihat Semua Data")
		fmt.Println("99. Keluar")

		fmt.Print("Masukkan menu pilihan")
		fmt.Scanln(&inputMenu)

		switch inputMenu {
		case 1:
			var newUser User
			fmt.Println("Registrasi")
			fmt.Print("Masukkan Nama\t: ")
			scanner.Scan()
			newUser.Nama = scanner.Text()
			fmt.Print("Masukkan Email\t: ")
			scanner.Scan()
			newUser.Email = scanner.Text()
			fmt.Print("Masukkan Password\t: ")
			scanner.Scan()
			newUser.Password = scanner.Text()

			arrData = append(arrData, newUser)
		case 2:
			fmt.Println(arrData)
		case 99:
			isRunning = false
		default:
			continue
		}
	}

}
