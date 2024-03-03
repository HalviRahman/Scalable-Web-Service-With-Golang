package biodata

import "fmt"

type Biodata struct {
	Nama    string
	Alamat string
	Pekerjaan    []string
	Alasan  string
}

func (a Biodata) GetBiodata(number int) {
	fmt.Println("No Absen:", number)
	fmt.Println("**************************")
	fmt.Println("Nama:", a.Nama)
	fmt.Println("Alamat:", a.Alamat)
	fmt.Println("Pekerjaan:", a.Pekerjaan)
	fmt.Println("Alasan:", a.Alasan)
}

func (a Biodata) GetListBiodata() []Biodata {
	return ListStudents
}
