package main

import (
	"fmt"
)

const NMAX int = 100

type ID_Petugas struct {
	username [NMAX]string
	password [NMAX]string
}

type ID_Pabrik struct {
	nama    string
	alamat  string
	wilayah string
	telepon string
}

type ID_Motor struct {
	merek string
	warna string
	jenis string
}

type Tab_Petugas struct {
	data  [NMAX]ID_Petugas
	count int
}

type Tab_Pabrik struct {
	data  [NMAX]ID_Pabrik
	count int
}

type Tab_Motor struct {
	data  [NMAX]ID_Motor
	count int
}

var users Tab_Petugas
var pabrik Tab_Pabrik
var motor Tab_Motor

func main() {
	var opsi, pilih string

	for {
		tampilan_home()
		fmt.Scan(&opsi)
		switch opsi {
		case "1":
			register()
		case "2":
			if login() {
				for {
					tampilan_menu()
					fmt.Scan(&pilih)
					switch pilih {
					case "1":
						for {
							menu_pabrik()
							fmt.Scan(&pilih)
							switch pilih {
							case "1":
								daftar_pabrik()
							case "2":
								tambah_pabrik()
							case "3":
								ubah_pabrik()
							case "4":
								hapus_pabrik()
							case "5":
								break
							default:
								fmt.Println("Masukan Anda Tidak Valid")
							}
						}
					case "2":
						for {
							menu_motor()
							fmt.Scan(&pilih)
							switch pilih {
							case "1":
								daftar_motor()
							case "2":
								tambah_motor()
							case "3":
								ubah_motor()
							case "4":
								hapus_motor()
							case "5":
								break
							default:
								fmt.Println("Masukan Anda Tidak Valid")
							}
						}
					case "3":
						menu_cari()
					case "4":
						menu_tertinggi()
					case "5":
						fmt.Println("Terimakasih Sudah Memakai Aplikasi Kami")
						return
					default:
						fmt.Println("Masukan Anda Tidak Valid")
					}
				}
			}
		case "3":
			fmt.Println("Terimakasih Sudah Memakai Aplikasi Kami")
			return
		default:
			fmt.Println("Masukan Anda Tidak Valid")
		}
	}
}

func tampilan_home() {
	fmt.Println("+--------------------------------+")
	fmt.Println("|           DEALER MOTOR         |")
	fmt.Println("+--------------------------------+")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")
	fmt.Print("Pilih (1/2/3): ")
}

func tampilan_menu() {
	fmt.Println("+--------------------------------+")
	fmt.Println("|           DAFTAR MENU          |")
	fmt.Println("+--------------------------------+")
	fmt.Println("1. Daftar Pabrik")
	fmt.Println("2. Daftar Motor")
	fmt.Println("3. Cari Data Motor")
	fmt.Println("4. Daftar Penjualan Tertinggi")
	fmt.Println("5. Exit")
	fmt.Print("Pilih (1/2/3/4/5): ")
}

func register() {
	var UserName, pass string

	fmt.Println("+--------------------------------+")
	fmt.Println("|         Registrasi Akun        |")
	fmt.Println("+--------------------------------+")
	fmt.Print("Username: ")
	fmt.Scan(&UserName)

	for i := 0; i < users.count; i++ {
		if users.data[i].username[0] == UserName {
			fmt.Println("Username sudah digunakan. Coba lagi.")
			return
		}
	}

	fmt.Print("Password: ")
	fmt.Scan(&pass)

	if users.count < NMAX {
		users.data[users.count].username[0] = UserName
		users.data[users.count].password[0] = pass
		users.count++
		fmt.Println("Registrasi berhasil")
	} else {
		fmt.Println("Kapasitas pengguna penuh")
	}
}

func login() bool {
	var UserName, pass string

	fmt.Println("+--------------------------------+")
	fmt.Println("|           Login Akun           |")
	fmt.Println("+--------------------------------+")
	fmt.Print("Username: ")
	fmt.Scan(&UserName)
	fmt.Print("Password: ")
	fmt.Scan(&pass)

	for i := 0; i < users.count; i++ {
		if users.data[i].username[0] == UserName && users.data[i].password[0] == pass {
			fmt.Println("Login berhasil")
			return true
		}
	}
	fmt.Println("Login tidak berhasil")
	return false
}

func menu_pabrik() {
	fmt.Println("+--------------------------------+")
	fmt.Println("|           Menu Pabrik          |")
	fmt.Println("+--------------------------------+")
	fmt.Println("1. Daftar Pabrik")
	fmt.Println("2. Tambah Pabrik")
	fmt.Println("3. Ubah Pabrik")
	fmt.Println("4. Hapus Pabrik")
	fmt.Println("5. Kembali")
	fmt.Print("Pilih (1/2/3/4/5): ")
}

func daftar_pabrik() {
	fmt.Println("+--------------------------------+")
	fmt.Println("|       Daftar Data Pabrik       |")
	fmt.Println("+--------------------------------+")
	for i := 0; i < pabrik.count; i++ {
		fmt.Printf("%d. %s, %s, %s, %s\n", i+1, pabrik.data[i].nama, pabrik.data[i].alamat, pabrik.data[i].wilayah, pabrik.data[i].telepon)
	}
}

func tambah_pabrik() {
	var nama, alamat, wilayah, telepon string
	fmt.Println("+--------------------------------+")
	fmt.Println("|        Tambah Data Pabrik      |")
	fmt.Println("+--------------------------------+")
	fmt.Print("Nama Pabrik: ")
	fmt.Scan(&nama)
	fmt.Print("Alamat Pabrik: ")
	fmt.Scan(&alamat)
	fmt.Print("Wilayah Pabrik: ")
	fmt.Scan(&wilayah)
	fmt.Print("Telepon Pabrik: ")
	fmt.Scan(&telepon)

	if pabrik.count < NMAX {
		pabrik.data[pabrik.count] = ID_Pabrik{nama: nama, alamat: alamat, wilayah: wilayah, telepon: telepon}
		pabrik.count++
		fmt.Println("Data pabrik berhasil ditambahkan")
	} else {
		fmt.Println("Kapasitas pabrik penuh")
	}
}

func ubah_pabrik() {
	var index int
	daftar_pabrik()
	fmt.Print("Masukkan nomor pabrik yang ingin diubah: ")
	fmt.Scan(&index)
	index--
	if index < 0 || index >= pabrik.count {
		fmt.Println("Nomor pabrik tidak valid")
		return
	}
	fmt.Println("Masukkan data baru untuk pabrik")
	fmt.Print("Nama Pabrik: ")
	fmt.Scan(&pabrik.data[index].nama)
	fmt.Print("Alamat Pabrik: ")
	fmt.Scan(&pabrik.data[index].alamat)
	fmt.Print("Wilayah Pabrik: ")
	fmt.Scan(&pabrik.data[index].wilayah)
	fmt.Print("Telepon Pabrik: ")
	fmt.Scan(&pabrik.data[index].telepon)
	fmt.Println("Data pabrik berhasil diubah")
}

func hapus_pabrik() {
	var index int
	daftar_pabrik()
	fmt.Print("Masukkan nomor pabrik yang ingin dihapus: ")
	fmt.Scan(&index)
	index--
	if index < 0 || index >= pabrik.count {
		fmt.Println("Nomor pabrik tidak valid")
		return
	}
	for i := index; i < pabrik.count-1; i++ {
		pabrik.data[i] = pabrik.data[i+1]
	}
	pabrik.count--
	fmt.Println("Data pabrik berhasil dihapus")
}

func menu_motor() {
	fmt.Println("+--------------------------------+")
	fmt.Println("|           Menu Motor           |")
	fmt.Println("+--------------------------------+")
	fmt.Println("1. Daftar Motor")
	fmt.Println("2. Tambah Motor")
	fmt.Println("3. Ubah Motor")
	fmt.Println("4. Hapus Motor")
	fmt.Println("5. Kembali")
	fmt.Print("Pilih (1/2/3/4/5): ")
}

func daftar_motor() {
	fmt.Println("+--------------------------------+")
	fmt.Println("|        Daftar Data Motor       |")
	fmt.Println("+--------------------------------+")
	for i := 0; i < motor.count; i++ {
		fmt.Printf("%d. %s, %s, %s\n", i+1, motor.data[i].merek, motor.data[i].warna, motor.data[i].jenis)
	}
}

func tambah_motor() {
	var merek, warna, jenis string
	fmt.Println("+--------------------------------+")
	fmt.Println("|        Tambah Data Motor       |")
	fmt.Println("+--------------------------------+")
	fmt.Print("Merek Motor: ")
	fmt.Scan(&merek)
	fmt.Print("Warna Motor: ")
	fmt.Scan(&warna)
	fmt.Print("Jenis Motor: ")
	fmt.Scan(&jenis)

	if motor.count < NMAX {
		motor.data[motor.count] = ID_Motor{merek: merek, warna: warna, jenis: jenis}
		motor.count++
		fmt.Println("Data motor berhasil ditambahkan")
	} else {
		fmt.Println("Kapasitas motor penuh")
	}
}

func ubah_motor() {
	var index int
	daftar_motor()
	fmt.Print("Masukkan nomor motor yang ingin diubah: ")
	fmt.Scan(&index)
	index--
	if index < 0 || index >= motor.count {
		fmt.Println("Nomor motor tidak valid")
		return
	}
	fmt.Println("Masukkan data baru untuk motor")
	fmt.Print("Merek Motor: ")
	fmt.Scan(&motor.data[index].merek)
	fmt.Print("Warna Motor: ")
	fmt.Scan(&motor.data[index].warna)
	fmt.Print("Jenis Motor: ")
	fmt.Scan(&motor.data[index].jenis)
	fmt.Println("Data motor berhasil diubah")
}

func hapus_motor() {
	var index int
	daftar_motor()
	fmt.Print("Masukkan nomor motor yang ingin dihapus: ")
	fmt.Scan(&index)
	index--
	if index < 0 || index >= motor.count {
		fmt.Println("Nomor motor tidak valid")
		return
	}
	for i := index; i < motor.count-1; i++ {
		motor.data[i] = motor.data[i+1]
	}
	motor.count--
	fmt.Println("Data motor berhasil dihapus")
}

func menu_cari() {
	fmt.Println("+--------------------------------+")
	fmt.Println("|       Menu Cari Motor          |")
	fmt.Println("+--------------------------------+")
	fmt.Println("1. Cari Motor")
	fmt.Print("Pilih (1): ")
}

func menu_tertinggi() {
	fmt.Println("+--------------------------------+")
	fmt.Println("|      Penjualan Tertinggi       |")
	fmt.Println("+--------------------------------+")
}
