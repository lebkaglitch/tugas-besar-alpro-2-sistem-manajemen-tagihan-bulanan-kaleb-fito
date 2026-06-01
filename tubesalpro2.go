package main

import "fmt"

const nmax int = 9999

type Tagihan struct {
	nama, kategori, status             string
	id, nominal, tanggal, bulan, tahun int
}
type arrtagihan [nmax]Tagihan

func main() {
	var jumlahtagihan int
	var data arrtagihan
	jumlahtagihan = 0
	// bingung sebenarnya pilihan ini mau di deklarasi di main atau di prosedur jalankanpilihan, karna kan cuma di pake di prosedur itu doang (solusinya udah ku pindah)
	jalankanpilihan(&data, &jumlahtagihan) // bingung di main itu mau manggil prosedur yg mana antara menu utama ama jalankan pilihan
}
func menuutama(pilihanmenu *int) {
	fmt.Println("===========================")
	fmt.Println("		SELAMAT DATANG  	")
	fmt.Println("===========================")
	fmt.Println()
	fmt.Println("1. Tambah Tagihan")
	fmt.Println("2. Ubah Tagihan")
	fmt.Println("3. Hapus Tagihan")
	fmt.Println("4. Lihat Semua Tagihan")
	fmt.Println("5. cari Tagihan")    // ini buat sequential dan binary search
	fmt.Println("6. Urutkan Tagihan") // ini buat selection dan insertion sort
	fmt.Println("0. keluar")
	fmt.Println()
	fmt.Println("===========================")
	fmt.Print("pilihan anda : ")
	fmt.Scan(pilihanmenu)
}
func jalankanpilihan(x *arrtagihan, n *int) {
	var pilihan int
	pilihan = -1
	for pilihan != 0 { // disini nih mksudku dari comment line 11
		menuutama(&pilihan)
		switch pilihan {
		default:
			fmt.Println("Pilihan anda tidak tersedia :(")
		case 1:
			tambahtagihan(x, n)
		case 2:
			ubahtagihan()
		case 3:
			hapustagihan()
		case 4:
			lihatsemuatagihan(x, *n)
		case 5:
			caritagihan()
		case 6:
			urutkantagihan()
		case 0:
			fmt.Println("Terima kasih")
		}
	}
}
func tambahtagihan(data *arrtagihan, jumlahtagihan *int) {
	var t Tagihan
	if *jumlahtagihan < nmax {
		fmt.Println("======TAMBAH TAGIHAN======")
		fmt.Println()
		fmt.Print("ID Tagihan: ")
		fmt.Scan(&t.id)

		fmt.Print("Nama Tagihan: ")
		fmt.Scan(&t.nama)

		fmt.Print("Kategori: ")
		fmt.Scan(&t.kategori)

		fmt.Print("Nominal: ")
		fmt.Scan(&t.nominal)

		fmt.Print("Tanggal Jatuh Tempo: ")
		fmt.Scan(&t.tanggal)

		fmt.Print("Bulan Jatuh Tempo: ")
		fmt.Scan(&t.bulan)

		fmt.Print("Tahun Jatuh Tempo: ")
		fmt.Scan(&t.tahun)

		t.status = "Belum Lunas"
		(*data)[*jumlahtagihan] = t
		*jumlahtagihan++

		fmt.Println("Tagihan Berhasil Ditambahkan!")
	} else {
		fmt.Println("Maaf tagihan anda sudah penuh")
	}
}

func ubahtagihan() {
	fmt.Println("menu ubah tagihan")
}

func hapustagihan() {
	fmt.Println("menu hapus tagihan")
}

func lihatsemuatagihan(x *arrtagihan, jumlahtagihan int) {
	var i int
	if jumlahtagihan == 0 {
		fmt.Println("Maaf anda belum memiliki tagihan, silahkan tambahkan tagihan anda!")
	} else {
		for i = 0; i < jumlahtagihan; i++ {
			fmt.Println(x[i].id)
			fmt.Println(x[i].nama)
			fmt.Println(x[i].kategori)
			fmt.Println(x[i].nominal)
			fmt.Println(x[i].tanggal)
			fmt.Println(x[i].bulan)
			fmt.Println(x[i].tahun)
			fmt.Println(x[i].status)
		}
	}
}

func caritagihan() {
	fmt.Println("menu cari tagihan")
}

func urutkantagihan() {
	fmt.Println("menu urutkan tagihan")
	// aku ada rencana  di bagian lihat semua tagihan and sorting, mau nge limit list nya only 10 (jika data nya lebih dari 10), nanti akan tekan enter untuk ganti halamaan
}

// bagian kaleb ke atas

// aku mau nambahin nanti agar ID itu otomatis nambah tanpa si user perlu ngisi sendiri, nanti akan nambahin var idterkahir buat ngatasin solusi di hapustagihan
