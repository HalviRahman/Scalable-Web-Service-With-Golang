package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/HalviRahman/Scalable-Web-Service-with-Golang/biodata"
)

func checkInput(input int, s []biodata.Biodata) (biodata.Biodata, error) {
	if input > len(s) || input < 1 {
		msg := fmt.Sprintf("INDEX OUT OF RANGE: Mohon Masukan Angka dari 1 - %d", len(s))
		return biodata.Biodata{}, errors.New(msg)
	} else {
		return s[input-1], nil
	}
}

func main() {
	biodataStruct := biodata.Biodata{}
	listStudent := biodataStruct.GetListBiodata()

	if len(os.Args) > 1 {
		input, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic("Harap masukan Angka terlebih dahulu...")
		}
		bio, err := checkInput(input, listStudent)
		if err != nil {
			panic(err)
		}
		bio.GetBiodata(input)
	} else {
		panic("Harap masukan angka terlebih dahulu...")
	}
	fmt.Println("\nProgram Selesai...")
}
