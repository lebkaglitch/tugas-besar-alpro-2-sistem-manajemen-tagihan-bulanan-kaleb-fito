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
	fmt.Println("7. Lunaskan Tagihan")
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
			ubahtagihan(x, *n)
		case 3:
			hapustagihan(x, n)
		case 4:
			lihatsemuatagihan(x, *n)
		case 5:
			caritagihan()
		case 6:
			urutkantagihan()
		case 7:
			tandailunas(x, *n)
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
		fmt.Print("Nama Tagihan: ") // nama ama kategori ada masalah jika ada spasi nya
		fmt.Scan(&t.nama)
		fmt.Print("Kategori: ")
		fmt.Scan(&t.kategori)
		fmt.Print("Nominal: ")
		fmt.Scan(&t.nominal)
		fmt.Print("Tanggal Jatuh Tempo: ") // nanti mau nambahin tgl/bln/thn ada syaratnya
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
		fmt.Println()
	}
}

func ubahtagihan(x *arrtagihan, jumlahtagihan int) {
	var IDubah, pilihanUbah, indexposisi int
	var nominalBaru, tanggalBaru, bulanBaru, tahunBaru int
	var namaBaru, kategoriBaru string

	if jumlahtagihan == 0 {
		fmt.Println("Anda belum mempunyai tagihan")
	} else {
		fmt.Println()
		fmt.Println("Masukan ID tagihan: ")
		fmt.Scan(&IDubah)

		indexposisi = cariposisiID(*x, jumlahtagihan, IDubah)
		if indexposisi == -1 {
			fmt.Println("Maaf, ID yang dicari tidak ditemukan")
		} else {
			fmt.Println()
			fmt.Println("Tagihan Ditemukan!")
			fmt.Println()
			fmt.Printf("ID\t\t: %d\n", (*x)[indexposisi].id)
			fmt.Printf("Nama\t\t: %s\n", (*x)[indexposisi].nama)
			fmt.Printf("Kategori\t: %s\n", (*x)[indexposisi].kategori)
			fmt.Printf("Nominal\t\t: %d\n", (*x)[indexposisi].nominal)
			fmt.Printf("Jatuh Tempo\t: %d/%d/%d\n", (*x)[indexposisi].tanggal, (*x)[indexposisi].bulan, (*x)[indexposisi].tahun)
			fmt.Println()
			fmt.Println("Apa yang ingin diubah?")
			fmt.Println()
			fmt.Println("1. Nama")
			fmt.Println("2. Kategori")
			fmt.Println("3. Nominal")
			fmt.Println("4. Tanggal Jatuh Tempo")
			fmt.Println("5. Kembali")
			fmt.Println()
			fmt.Println("Pilihan Anda: ")
			fmt.Scan(&pilihanUbah)

			switch pilihanUbah {
			default:
				fmt.Println("Pilihan anda tidak tersedia :(")
			case 1:
				fmt.Println("Masukkan Nama Baru: ")
				fmt.Scan(&namaBaru)
				(*x)[indexposisi].nama = namaBaru
				fmt.Printf("Nama Berhasil Diubah menjadi '%s\n'", namaBaru)
				fmt.Println()
			case 2:
				fmt.Println("Masukkan Kategori Baru: ")
				fmt.Scan(&kategoriBaru)
				(*x)[indexposisi].kategori = kategoriBaru
				fmt.Printf("Kategori Berhasil Diubah menjadi '%s\n'", kategoriBaru)
				fmt.Println()
			case 3:
				fmt.Println("Masukkan Nominal Baru: ") // biasanya nominal ada Rp nya, tapi nanti aja, inget aja asal
				fmt.Scan(&nominalBaru)
				(*x)[indexposisi].nominal = nominalBaru
				fmt.Printf("Nominal Berhasil Diubah menjadi '%d\n'", nominalBaru)
				fmt.Println()
			case 4:
				fmt.Println("Masukkan Tanggal Baru: ")
				fmt.Scan(&tanggalBaru)
				fmt.Println("Masukkan Bulan Baru: ")
				fmt.Scan(&bulanBaru)
				fmt.Println("Masukkan Tahun Baru: ")
				fmt.Scan(&tahunBaru)

				(*x)[indexposisi].tanggal = tanggalBaru
				(*x)[indexposisi].bulan = bulanBaru
				(*x)[indexposisi].tahun = tahunBaru
				fmt.Println("Tanggal Jatuh Tempo Berhasil Diubah menjadi '%d/%d/%d'\n", tanggalBaru, bulanBaru, tahunBaru)
				fmt.Println()
			case 5:
				fmt.Println("Kembali ke menu utama")
				fmt.Println()
			}
		}
	}
	//aku ada kepikiran untuk setiap hapus dan ubah, nanti akan muncul lagi list nya, biar si user bisa langsung cek list nya tanpa perlu ke fitur liat tagihan
	fmt.Println()
}

func hapustagihan(x *arrtagihan, jumlahtagihan *int) {
	var idfordelete, indexposisi, i int

	if *jumlahtagihan == 0 {
		fmt.Println("Belum ada tagihan")
		return
	}
	//aku ada kepikiran untuk setiap hapus dan ubah, nanti akan muncul lagi list nya, biar si user bisa langsung cek list nya tanpa perlu ke fitur liat tagihan
	fmt.Println()
	fmt.Println("Masukkan ID tagihan yang akan dihapus: ")
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
		fmt.Println("==================================================================================================================================================================")
		fmt.Println("|								LIST DAFTAR TAGIHAN ANDA                                                                         |")
		fmt.Println("==================================================================================================================================================================")
		for i = 0; i < jumlahtagihan; i++ {
			fmt.Printf("| ID: %d\t | NAMA: %-10s\t | KATEGORI: %-10s\t | NOMINAL: %d\t | TANGGAL: %d\t | BULAN: %d\t | TAHUN: %d\t | Status Sekarang: %s\t |\n",
				x[i].id, x[i].nama, x[i].kategori, x[i].nominal, x[i].tanggal, x[i].bulan, x[i].tahun, x[i].status)
		}
		fmt.Println("==================================================================================================================================================================")
		fmt.Println()
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

// gw mau nanya to, kira kira kalau mau lunas syarat nya apa to?
//bagian fito
