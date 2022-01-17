// Joe Allen Butarbutar (GLNG020ONL003)

package main

import (
	"fmt"
	"os"
	"strconv"
)

type Classmate struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var classmate []Classmate

func init() {
	classmate = []Classmate{
		{Absen: 1, Nama: "Joe Allen Butarbutar", Alamat: "abc", Pekerjaan: "Backend Engineer 1", Alasan: "Lorem ipsum dolor sit amet"},
		{Absen: 2, Nama: "Akhyar Adin", Alamat: "def", Pekerjaan: "Backend Engineer 1", Alasan: "Lacus sed turpis tincidunt id"},
		{Absen: 3, Nama: "Rizki Akbari", Alamat: "ghi", Pekerjaan: "Backend Engineer 1", Alasan: "Pellentesque elit ullamcorper dignissim cras tincidunt"},
		{Absen: 4, Nama: "Fernando", Alamat: "jkl", Pekerjaan: "Backend Engineer 1", Alasan: "Duis at consectetur lorem donec"},
		{Absen: 5, Nama: "Felix Otto Trihardjo", Alamat: "mno", Pekerjaan: "Backend Engineer 1", Alasan: "Cras sed felis eget velit aliquet"},
		{Absen: 6, Nama: "Panji Ragil Wibisono", Alamat: "pqr", Pekerjaan: "Backend Engineer 1", Alasan: "Purus sit amet luctus venenatis lectus magna"},
	}
}
func main() {
	allArgs := os.Args[1:]
	if len(os.Args) > 1 {
		for i := 0; i < len(allArgs); i++ {
			absen, err := strconv.Atoi(allArgs[i])
			if err != nil {
				fmt.Println("Something went wrong. Please re-check your input")
			} else {
				search := 0
				for i := 0; i < len(classmate); i++ {
					if classmate[i].Absen == absen {
						ShowData(classmate[i])
						search++
						break
					}
				}
				if search == 0 {
					fmt.Println("Data Not Found")
				}
			}
		}
	} else {
		fmt.Println("Please input the student's absent number in argument")
	}
}

func ShowData(classmate Classmate) {
	fmt.Println("Nama :", classmate.Nama)
	fmt.Println("Alamat :", classmate.Alamat)
	fmt.Println("Pekerjaan :", classmate.Pekerjaan)
	fmt.Println("Alasan :", classmate.Alasan)
}
