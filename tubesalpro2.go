package main
import "fmt"
	const nmax int = 9999
	type Tagihan struct{
		nama, kategori, status string
		id, nominal, tanggal, bulan, tahun int
	}
	type arrtagihan [nmax]Tagihan
func main() {
	var pilihan, n int
	var data arrtagihan
	pilihan  = -1 	// bingung sebenarnya pilihan ini mau di deklarasi di main atau di prosedur jalankanpilihan, karna kan cuma di pake di prosedur itu doang
	jalankanpilihan(&data, &n) 	// bingung di main itu mau manggil prosedur yg mana antara menu utama ama jalankan pilihan
}
func menuutama(pilihanmenu *int) {
	var 
	fmt.Println("===========================")
	fmt.Println("		SELAMAT DATANG  	")
	fmt.Println("===========================")
	fmt.Println()
	fmt.Println("1. Tambah Tagihan")
	fmt.Println("2. Ubah Tagihan")
	fmt.Println("3. Hapus Tagihan")
	fmt.Println("4. Lihat Semua Tagihan")
	fmt.Println("5. cari Tagihan") // ini buat sequential dan binary search
	fmt.Println("6. Urutkan Tagihan") // ini buat selection dan insertion sort
	fmt.Println("0. keluar")
	fmt.Println()
	fmt.Println("===========================")
	fmt.Println("pilihan anda : ")
	fmt.Scan(pilihanmenu) 	
}
func jalankanpilihan(x *arrtagihan, n *int) {
	for pilihan != 0{ 	// disini nih mksudku dari comment line 11
		menuutama(&pilihan)
		switch pilihan{
		default:
			fmt.Println("Pilihan anda tidak tersedia :(")
		case 1:
			tambahtagihan(x, n)
		case 2: 
			ubahtagihan()
		case 3:
			hapustagihan()
		case 4:
			lihatsemuatagihan()
		case 5:
			caritagihan()	
		case 6:
			urutkantagihan()
		case 0:
			fmt.Println("Terima kasih")
		}
	}
}
func tambahtagihan(data *arrtagihan, jumlahtagihanuser *int) {
	
}
// bagian kaleb ke atas
//habis ganti akun github abangku ke akun ku