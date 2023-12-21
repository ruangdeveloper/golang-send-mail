package main

import (
	"fmt"
)

func main() {
	err := SendMail(
		"user@example.com",
		"Halo, ini subject",
		"Halo, ini isi pesan yang dikirim menggunakan GoMail ya!",
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Email sent!")
}
