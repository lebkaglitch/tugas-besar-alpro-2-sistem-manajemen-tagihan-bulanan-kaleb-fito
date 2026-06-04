package main

import "fmt"

const nmax int = 9999

type Tagihan struct {
	nama, kategori, status             string
	id, nominal, tanggal, bulan, tahun int
}
type arrtagihan [nmax]Tagihan

func main() {
	var jumlahtagihan, idterkahir int
	var data arrtagihan
	jumlahtagihan = 0
	idterkahir = 0
	// bingung sebenarnya pilihan ini mau di deklarasi di main atau di prosedur jalankanpilihan, karna kan cuma di pake di prosedur itu doang (solusinya udah ku pindah)
	jalankanpilihan(&data, &jumlahtagihan, &idterkahir) // bingung di main itu mau manggil prosedur yg mana antara menu utama ama jalankan pilihan
}
func menuutama(pilihanmenu *int) {
	fmt.Println("==============================")
	fmt.Println("	SELAMAT DATANG  	")
	fmt.Println("==============================")
	fmt.Println()
	fmt.Println("1. Tambah Tagihan")
	fmt.Println("2. Ubah Tagihan")
	fmt.Println("3. Hapus Tagihan")
	fmt.Println("4. Lihat Semua Tagihan")
	fmt.Println("5. cari Tagihan")    // ini buat sequential dan binary search
	fmt.Println("6. Urutkan Tagihan") // ini buat selection dan insertion sort
	fmt.Println("0. keluar")
	fmt.Println()
	fmt.Println("==============================")
	fmt.Print("Masukkan Pilihan Anda: ")
	fmt.Scan(pilihanmenu)
}
func jalankanpilihan(x *arrtagihan, n *int, idterakhir *int) {
	var pilihan int
	pilihan = -1
	for pilihan != 0 { // disini nih mksudku dari comment line 11
		menuutama(&pilihan)
		switch pilihan {
		default:
			fmt.Println("Pilihan anda tidak tersedia :(")
		case 1:
			tambahtagihan(x, n, idterakhir)
		case 2:
			ubahtagihan()
		case 3:
			hapustagihan(x, n)
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
func tambahtagihan(data *arrtagihan, jumlahtagihan *int, idterakhir *int) {
	var t Tagihan
	if *jumlahtagihan < nmax {
		fmt.Println()
		fmt.Println("======TAMBAH TAGIHAN======")
		fmt.Println()
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

		*idterakhir = *idterakhir + 1
		t.id = *idterakhir
		(*data)[*jumlahtagihan] = t
		*jumlahtagihan = *jumlahtagihan + 1
		fmt.Println()
		fmt.Println("TAGIHAN ANDA BERHASIL DITAMBAHKAN :') ")
		fmt.Println()
	} else {
		fmt.Println("MAAF TAGIHAN ANDA SUDAH PENUH T_T")
	}
}

func ubahtagihan() {
	fmt.Println("menu ubah tagihan")
}

func hapustagihan(x *arrtagihan, jumlahtagihan *int) {
	var idfordelete, indexposisi, i int

	if *jumlahtagihan == 0 {
		fmt.Println("Belum ada tagihan")
		return
	}
	fmt.Println()
	fmt.Println("Masukkan ID tagihan yang akan dihaous: ")
	fmt.Scan(&idfordelete)

	indexposisi = cariposisiID(*x, *jumlahtagihan, idfordelete)
	if indexposisi == -1 {
		fmt.Println("Maaf, ID yang dicari tidak ditemukan")
	} else {
		for i = indexposisi; i < *jumlahtagihan-1; i++ {
			(*x)[i] = (*x)[i+1]
		}
		*jumlahtagihan = *jumlahtagihan - 1

		fmt.Println("Selamat, Tagihan anda berhasil dihapus")
	}
}
func cariposisiID(x arrtagihan, jumlahtagihan int, targetid int) int { // ini function yg digunakan fitur hapus dan ubah dalam mencari ID nya, krna utk hapus dan ubah itu aku buat dengan mencari ID yg dicari
	var i int
	for i = 0; i < jumlahtagihan; i++ {
		if x[i].id == targetid {
			return i
		}
	}
	return -1
}
func lihatsemuatagihan(x *arrtagihan, jumlahtagihan int) {
	var i int
	if jumlahtagihan == 0 {
		fmt.Println("Maaf anda belum memiliki tagihan, silahkan tambahkan tagihan anda!")
	} else {
		fmt.Println()
		fmt.Println("========================================================================================================================")
		fmt.Println("								LIST DAFTAR TAGIHAN ANDA")
		fmt.Println("========================================================================================================================")
		for i = 0; i < jumlahtagihan; i++ {
			fmt.Printf("| ID: %d\t | NAMA: %-10s\t | KATEGORI: %-10s\t | NOMINAL: %d\t | TANGGAL: %d\t | BULAN: %d\t | TAHUN: %d\t | Status Sekarang: %s\t |\n",
				x[i].id, x[i].nama, x[i].kategori, x[i].nominal, x[i].tanggal, x[i].bulan, x[i].tahun, x[i].status)
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
ac

//tanda lunas
func tandailunas(x *arrtagihan, jumlahtagihan int) {
	var idcari, idx int
 
	if jumlahtagihan == 0 {
		fmt.Println("Belum ada tagihan.")
		return
	}
 
	fmt.Println()
	fmt.Println("======TANDAI LUNAS======")
	fmt.Print("Masukkan ID tagihan yang sudah dibayar: ")
	fmt.Scan(&idcari)
 
	idx = cariposisiID(*x, jumlahtagihan, idcari)
	if idx == -1 {
		fmt.Println("Maaf, ID yang dicari tidak ditemukan")
	} else if (*x)[idx].status == "Lunas" {
		fmt.Printf("Tagihan '%s' sudah berstatus Lunas.\n", (*x)[idx].nama)
	} else {
		(*x)[idx].status = "Lunas"
		fmt.Printf("Tagihan '%s' berhasil ditandai Lunas!\n", (*x)[idx].nama)
	}
}

//
