package main

import (
	"fmt"
	"os"
)

// Teman struct untuk menyimpan informasi teman
type Friend struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// DataTeman menyimpan data teman-teman dalam kelas
var DataFriend = map[int]Friend{
	1: {"Alamin", "Muna Barat", "Fullstack Developer", "Ingin menjadi software engineer"},
	2: {"Rajab", "Kendari", "Developer", "Meningkatkan skill untuk pekerjaan"},
	3: {"Ferdi", "Konawe", "Mahasiswa", "Tertarik dengan scalable golang"},
	4: {"Agus", "Muna", "IT Support", "Membutuhkan skill programmer pada pekerjaan baru"},
	5: {"Wahyu", "Buton", "Teknisi Mesin", "Ingin Menjadi Developer Andal"},
}

// TampilkanDataTeman menampilkan data teman berdasarkan nomor absen
func ViewDataFriend(absen int) {
	friends, ok := DataFriend[absen]
	if !ok {
		fmt.Println("Teman dengan nomor absen tersebut tidak ditemukan")
		return
	}

	fmt.Println("Nama:", friends.Nama)
	fmt.Println("Alamat:", friends.Alamat)
	fmt.Println("Pekerjaan:", friends.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", friends.Alasan)
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Command: go run biodata.go <nomor_absen>")
		return
	}

	absen := args[1]
	var absenInt int
	_, err := fmt.Sscanf(absen, "%d", &absenInt)
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka")
		return
	}

	ViewDataFriend(absenInt)
}
