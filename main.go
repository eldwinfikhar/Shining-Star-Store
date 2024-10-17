package main

import "fmt"

const nMAX int = 1000

type barang struct {
	nama, kategori string
	jumlah         int
}
type Inventory [nMAX]barang
type gudang struct {
	list Inventory
	n    int
}
type bayar struct {
	kode  int
	orang string
	kargo barang
}
type Payment [nMAX]bayar
type akun struct {
	masuk, keluar   Payment
	nMasuk, nKeluar int
}

func login() {
	/* I.S. -
	F.S. Menampilkan menu login dan input password.*/
	var password string
	fmt.Println("======== WELCOME ========")
	fmt.Println("     Store Inventory     ")
	fmt.Println("    Shining Star Store   ")
	fmt.Println("-------------------------")
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&password)
	for password != "sh1n3" {
		fmt.Print("Masukkan Password: ")
		fmt.Scanln(&password)
	}
}

func mainHeader() {
	fmt.Print("\n")
	fmt.Println("=========================")
	fmt.Println("     Store Inventory     ")
	fmt.Println("    Shining Star Store   ")
	fmt.Println("        MAIN MENU        ")
	fmt.Println("-------------------------")
	fmt.Println("1. Akses Data Barang")
	fmt.Println("2. Akses Data Transaksi")
	fmt.Println("3. Tampilkan Inventory")
	fmt.Println("0. Exit")
	fmt.Println("=========================")
	fmt.Print("Pilihan anda (1/2/3/0): ")
}

func mainMenu(P *gudang, Q *akun) {
	/* I.S. Terdefinisi struct gudang P berisi P.n data barang dan struct akun Q
	   berisi data dan jumlah dari transaksi masuk dan keluar.
	F.S. Menampilkan header MAIN MENU dan input opsi secara berulang. Pengguna akan
	diarahkan ke menu sesuai dengan opsi. Perulangan berakhir ketika opsi bernilai 0. */
	var opsi int
	for {
		mainHeader()
		fmt.Scan(&opsi)
		for opsi != 1 && opsi != 2 && opsi != 3 && opsi != 0 {
			fmt.Print("Pilihan anda (1/2/3/0): ")
			fmt.Scan(&opsi)
		}
		switch opsi {
		case 1:
			itemMenu(P)
		case 2:
			transactionMenu(P, Q)
		case 3:
			displayMenu(P)
		}
		if opsi == 0 {
			break
		}
	}
	fmt.Println("Terima kasih sudah menggunakan aplikasi.")
}

func itemHeader() {
	fmt.Print("\n")
	fmt.Println("=========================")
	fmt.Println("     Store Inventory     ")
	fmt.Println("    Shining Star Store   ")
	fmt.Println("        ITEM MENU        ")
	fmt.Println("-------------------------")
	fmt.Println("1. Tambah Barang")
	fmt.Println("2. Hapus Barang")
	fmt.Println("3. Edit Data Barang")
	fmt.Println("4. Cari Stok Barang")
	fmt.Println("0. Back")
	fmt.Println("=========================")
	fmt.Print("Pilihan anda (1/2/3/4/0): ")
}

func itemMenu(P *gudang) {
	/* I.S. Terdefinisi struct gudang P berisi P.n data barang.
	F.S. Menampilkan header ITEM MENU dan input opsi secara berulang. Pengguna akan
	diarahkan ke menu sesuai dengan opsi. Perulangan berakhir ketika opsi bernilai 0. */
	var opsi int
	for {
		itemHeader()
		fmt.Scan(&opsi)
		for opsi != 1 && opsi != 2 && opsi != 3 && opsi != 4 && opsi != 0 {
			fmt.Print("Pilihan anda (1/2/3/4/0): ")
			fmt.Scan(&opsi)
		}
		switch opsi {
		case 1:
			addItem(P)
		case 2:
			deleteItem(P)
		case 3:
			editItem(P)
		case 4:
			searchItem_Stock(*P)
		}
		if opsi == 0 {
			break
		}
	}
}

func checkItem(P gudang, x string) int {
	/* I.S. Terdefinisi struct gudang P berisi P.n data barang dan string x yang mewakili nama barang.
	F.S. Mengembalikan index lokasi barang dengan nama x pada array Inventory di struct gudang P jika
	barang ada pada array dan -1 jika barang tidak ada pada array. */
	var j int
	j = 0
	for j < P.n && P.list[j].nama != x {
		j++
	}
	if j == P.n {
		return -1
	} else {
		return j
	}
}

func addItem(P *gudang) {
	/* I.S. Terdefinisi struct gudang P berisi P.n data barang.
	Proses: Menambahkan jumlah data yang akan dimasukkan. Jika barang tambahan sudah
	        ada di array, maka hanya dilakukan penambahan stok pada barang tersebut.
	F.S. Stok barang tertentu di gudang bertambah. */
	var Nama, Kategori string
	var Jumlah, k, idx, nData int
	fmt.Print("Jumlah data masukan: ")
	fmt.Scan(&nData)
	fmt.Println("Masukkan data barang tambahan:")
	if nData > 0 {
		for {
			fmt.Scan(&Nama, &Kategori, &Jumlah)
			idx = checkItem(*P, Nama)
			if idx == -1 {
				P.list[P.n].nama = Nama
				P.list[P.n].kategori = Kategori
				P.list[P.n].jumlah = Jumlah
				P.n++
			} else {
				P.list[idx].jumlah += Jumlah
			}
			k++
			if k >= nData || P.n == nMAX {
				break
			}
		}
	}
	fmt.Println("-------------------------")
	fmt.Println("Jumlah barang di gudang:", P.n)
	if P.n == nMAX {
		fmt.Println("Inventory PENUH.")
	}
}

func deleteItem(P *gudang) {
	/* I.S. Terdefinisi struct gudang P berisi P.n data barang.
	Proses: Memasukkan NAMA dan KATEGORI barang yang akan dihapus.
	F.S. Jumlah barang di gudang berkurang 1 jika penghapusan BERHASIL. */
	var Nama, Kategori string
	var idx, i int
	var temp barang
	if P.n > 0 {
		fmt.Println("Masukkan data barang yang dihapus:")
		fmt.Print("Nama     : ")
		fmt.Scan(&Nama)
		fmt.Print("Kategori : ")
		fmt.Scan(&Kategori)
		idx = -1
		for i < P.n && idx == -1 {
			if P.list[i].nama == Nama && P.list[i].kategori == Kategori {
				idx = i
			}
			i++
		}
		if idx != -1 {
			for i = idx; i < P.n-1; i++ {
				temp = P.list[i]
				P.list[i] = P.list[i+1]
				P.list[i+1] = temp
			}
			P.n--
			fmt.Println("Penghapusan barang BERHASIL.")
		} else {
			fmt.Println("Barang TIDAK DITEMUKAN.")
		}

	}
	fmt.Println("Jumlah barang di gudang:", P.n)
	if P.n == 0 {
		fmt.Println("Inventory KOSONG.")
	}
}

func searchItem_Edit(P gudang) int {
	/* I.S. Terdefinisi struct P gudang berisi P.n data barang.
	Proses: Memasukkan NAMA dan KATEGORI barang yang ingin diubah. Lalu, melakukan
	        pencarian data dengan SEQUENTIAL SEARCH
	F.S. Mengembalikan indeks lokasi barang jika barang ada pada array dan -1 jika
	barang tidak ditemukan pada array. */
	var idx, i int
	var nama, kat string
	fmt.Scan(&nama, &kat)
	idx = -1
	i = 0
	for idx == -1 && i < P.n {
		if P.list[i].nama == nama && P.list[i].kategori == kat {
			idx = i
		}
		i++
	}
	return idx
}

func editItem(P *gudang) {
	/* I.S. Mencari indeks lokasi data yang ingin diubah.
	Proses: Memasukkan NAMA dan KATEGORI barang yang ingin diubah.
	F.S. Mengembalikan indeks lokasi barang jika barang ada pada array dan -1
	jika barang tidak ditemukan pada array. */
	var idx, newInt int
	var change, newString string
	fmt.Println("Masukkan NAMA dan KATEGORI barang:")
	idx = searchItem_Edit(*P)
	fmt.Print("Ingin mengganti (nama/kat/jum): ")
	fmt.Scan(&change)
	if idx != -1 {
		if change == "jum" {
			fmt.Print("Masukkan JUMLAH terbaru: ")
			fmt.Scan(&newInt)
			P.list[idx].jumlah = newInt
		} else if change == "nama" {
			fmt.Print("Masukkan NAMA terbaru: ")
			fmt.Scan(&newString)
			P.list[idx].nama = newString
		} else if change == "kat" {
			fmt.Print("Masukkan KATEGORI terbaru: ")
			fmt.Scan(&newString)
			P.list[idx].kategori = newString
		}
		fmt.Println("Edit data barang BERHASIL.")
	} else {
		fmt.Println("Barang TIDAK DITEMUKAN.")
	}
}

func searchItem_Stock(P gudang) {
	/* I.S. Terdefinisi struct P gudang berisi P.n data barang.
	Proses: Memasukkan NAMA dan KATEGORI barang yang ingin dilihat stoknya.
	F.S. Menampilkan stok barang yang dicari jika barang ada pada array dan string
	"Barang TIDAK DITEMUKAN" jika barang tidak ada pada array */
	var idx, i int
	var nama, kat string
	fmt.Println("Masukkan NAMA dan KATEGORI barang:")
	fmt.Scan(&nama, &kat)
	idx = -1
	i = 0
	for idx == -1 && i < P.n {
		if P.list[i].nama == nama && P.list[i].kategori == kat {
			idx = i
		}
		i++
	}
	if idx == -1 {
		fmt.Println("Barang TIDAK DITEMUKAN.")
	} else {
		fmt.Println("Stok", P.list[idx].nama, "di gudang:", P.list[idx].jumlah)
	}
}

func transactionHeader() {
	fmt.Print("\n")
	fmt.Println("===========================")
	fmt.Println("     Store Inventory     ")
	fmt.Println("    Shining Star Store   ")
	fmt.Println("     TRANSACTION MENU    ")
	fmt.Println("---------------------------")
	fmt.Println("1. Tambah Transaksi")
	fmt.Println("2. Hapus Transaksi")
	fmt.Println("3. Edit Transaksi")
	fmt.Println("4. Cari Transaksi")
	fmt.Println("5. Tampilkan Transaksi")
	fmt.Println("0. Back")
	fmt.Println("---------------------------")
	fmt.Println("NOTES:")
	fmt.Println("Kode transaksi bersifat UNIK.")
	fmt.Println("===========================")
	fmt.Print("Pilihan anda (1/2/3/4/5/0): ")
}

func transactionMenu(P *gudang, Q *akun) {
	/* I.S. Terdefinisi struct gudang P berisi n data barang dan struct akun Q berisi
	   data dan jumlah dari transaksi masuk dan keluar
	F.S. Menampilkan header TRANSACTION MENU dan input opsi secara berulang. Pengguna
	akan diarahkan ke menu sesuai dengan opsi. Perulangan berakhir ketika opsi bernilai 0 */
	var opsi int
	for {
		transactionHeader()
		fmt.Scan(&opsi)
		for opsi != 1 && opsi != 2 && opsi != 3 && opsi != 4 && opsi != 5 && opsi != 0 {
			fmt.Print("Pilihan anda (1/2/3/4/5/0): ")
			fmt.Scan(&opsi)
		}
		switch opsi {
		case 1:
			addTransaction(P, Q)
		case 2:
			deleteTransaction(Q)
		case 3:
			editTransaction(Q)
		case 4:
			searchTransactionByCode(Q)
		case 5:
			printTransaction(*Q)
		}
		if opsi == 0 {
			break
		}
	}
}

func addTransaction(P *gudang, Q *akun) {
	var pilih string
	fmt.Print("Transaksi MASUK/KELUAR (M/K)? ")
	fmt.Scan(&pilih)
	for pilih != "M" && pilih != "K" {
		fmt.Print("Transaksi MASUK/KELUAR (M/K)? ")
		fmt.Scan(&pilih)
	}
	if pilih == "M" {
		inTransaction(P, Q)
	} else {
		outTransaction(P, Q)
	}
}

func checkTransCode(Q akun, y int, trans string) bool {
	/* I.S. Terdefinisi struct gudang P berisi n data barang dan string x yang mewakili nama barang
	F.S. Mengembalikan index lokasi barang dengan nama x pada array Inventory di struct gudang P
	dan -1 jika barang tidak ada pada array */
	var j int
	j = 0
	if trans == "Masuk" {
		for j < Q.nMasuk && Q.masuk[j].kode != y {
			j++
		}
		if j == Q.nMasuk {
			return true
		} else {
			return false
		}
	} else {
		for j < Q.nKeluar && Q.keluar[j].kode != y {
			j++
		}
		if j == Q.nKeluar {
			return true
		} else {
			return false
		}
	}
}

func inTransaction(P *gudang, Q *akun) {
	var Nama, NamaBrg, KatBrg string
	var Kode, JumBrg, k, idx, nData int
	fmt.Println("FORMAT INPUT (tanpa koma):")
	fmt.Println("[KodeTransaksi, Nama, NamaBarang, KategoriBarang, JumlahBarang]")
	fmt.Print("Jumlah data masukan: ")
	fmt.Scan(&nData)
	fmt.Println("Masukkan data transaksi MASUK:")
	if nData > 0 {
		k = 0
		for {
			fmt.Scan(&Kode, &Nama, &NamaBrg, &KatBrg, &JumBrg)
			if checkTransCode(*Q, Kode, "Masuk") {
				Q.masuk[Q.nMasuk].kode = Kode
				Q.masuk[Q.nMasuk].orang = Nama
				Q.masuk[Q.nMasuk].kargo.nama = NamaBrg
				Q.masuk[Q.nMasuk].kargo.kategori = KatBrg
				Q.masuk[Q.nMasuk].kargo.jumlah = JumBrg
				Q.nMasuk++
				idx = checkItem(*P, NamaBrg)
				addItemFromTransaction(P, NamaBrg, KatBrg, JumBrg, idx)
			}
			k++
			if k >= nData || P.n == nMAX {
				break
			}
		}
	}
	fmt.Println("-------------------------")
	fmt.Println("Jumlah barang di gudang :", P.n)
	fmt.Println("Jumlah transaksi MASUK  :", Q.nMasuk)
}

func addItemFromTransaction(P *gudang, nama, kategori string, jumlah, idx int) {
	if idx == -1 {
		P.list[P.n].nama = nama
		P.list[P.n].kategori = kategori
		P.list[P.n].jumlah = jumlah
		P.n++
	} else {
		P.list[idx].jumlah += jumlah
	}
}

func outTransaction(P *gudang, Q *akun) {
	var Nama, NamaBrg, KatBrg, ulang string
	var Kode, JumBrg, idx, k int
	fmt.Println("FORMAT INPUT (tanpa koma):")
	fmt.Println("[KodeTransaksi, Nama, NamaBarang, KategoriBarang, JumlahBarang]")
	for {
		if P.n > 0 {
			fmt.Println("Masukkan data transaksi KELUAR:")
			fmt.Scan(&Kode, &Nama, &NamaBrg, &KatBrg, &JumBrg)
			for !checkTransCode(*Q, Kode, "Keluar") {
				fmt.Println("Kode transaksi INVALID.")
				fmt.Println("Masukkan data transaksi KELUAR:")
				fmt.Scan(&Kode, &Nama, &NamaBrg, &KatBrg, &JumBrg)
			}
			idx = checkItem(*P, NamaBrg)
			if idx != -1 {
				fmt.Println("-------------------------")
				fmt.Println("Stok", NamaBrg, "di gudang:", P.list[idx].jumlah)
				if JumBrg <= P.list[idx].jumlah {
					P.list[idx].jumlah -= JumBrg
					fmt.Println("Stok", NamaBrg, "terbaru  :", P.list[idx].jumlah)
					Q.keluar[k].kode = Kode
					Q.keluar[k].orang = Nama
					Q.keluar[k].kargo.nama = NamaBrg
					Q.keluar[k].kargo.kategori = KatBrg
					Q.keluar[k].kargo.jumlah = JumBrg
					k++
					if P.list[idx].jumlah == 0 {
						deleteItemFromTransaction(P, idx)
					}
					fmt.Println("Transaksi BERHASIL.")
				} else {
					fmt.Println("Pengambilan terlalu banyak, transaksi GAGAL.")
				}
			} else {
				fmt.Println("Barang TIDAK DITEMUKAN.")
			}
		} else {
			fmt.Println("Inventory KOSONG.")
		}
		fmt.Print("Ingin melakukan transaksi lagi (Ya/Tidak)? ")
		fmt.Scan(&ulang)
		if ulang == "Tidak" || P.n == 0 {
			break
		}
		fmt.Println("-------------------------")
	}
	Q.nKeluar += k
	fmt.Println("-------------------------")
	fmt.Println("Jumlah barang di gudang :", P.n)
	fmt.Println("Jumlah transaksi KELUAR :", Q.nKeluar)
}

func deleteItemFromTransaction(P *gudang, idx int) {
	var i int
	var temp barang
	for i = idx; i < P.n-1; i++ {
		temp = P.list[i]
		P.list[i] = P.list[i+1]
		P.list[i+1] = temp
	}
	P.n--
	fmt.Println("Jumlah barang di gudang:", P.n)
}

func deleteTransaction(Q *akun) {
	var Nama, pilih string
	var Kode, idx, i int
	var temp bayar
	fmt.Print("Ingin menghapus transaksi MASUK/KELUAR (M/K)? ")
	fmt.Scan(&pilih)
	for pilih != "M" && pilih != "K" {
		fmt.Print("Ingin menghapus transaksi MASUK/KELUAR (M/K)? ")
		fmt.Scan(&pilih)
	}
	fmt.Println("Masukkan data transaksi yang dihapus:")
	fmt.Print("Kode   : ")
	fmt.Scan(&Kode)
	fmt.Print("Nama   : ")
	fmt.Scan(&Nama)
	idx = -1
	if pilih == "M" {
		for i < Q.nMasuk && idx == -1 {
			if Q.masuk[i].kode == Kode && Q.masuk[i].orang == Nama {
				idx = i
			}
			i++
		}
	} else {
		for i < Q.nKeluar && idx == -1 {
			if Q.keluar[i].kode == Kode && Q.keluar[i].orang == Nama {
				idx = i
			}
			i++
		}
	}
	if idx != -1 {
		if pilih == "M" {
			for i = idx; i < Q.nMasuk-1; i++ {
				temp = Q.masuk[i]
				Q.masuk[i] = Q.masuk[i+1]
				Q.masuk[i+1] = temp
			}
			Q.nMasuk--
		} else {
			for i = idx; i < Q.nKeluar-1; i++ {
				temp = Q.keluar[i]
				Q.keluar[i] = Q.keluar[i+1]
				Q.keluar[i+1] = temp
			}
			Q.nKeluar--
		}
		fmt.Println("-------------------------")
		fmt.Println("Penghapusan transaksi BERHASIL.")
		fmt.Println("Jumlah transaksi MASUK :", Q.nMasuk)
		fmt.Println("Jumlah transaksi KELUAR:", Q.nKeluar)
	} else {
		fmt.Println("-------------------------")
		fmt.Println("Transaksi TIDAK DITEMUKAN.")
	}
}

func searchTrans_Edit(Q akun, pilih string) int {
	var idx, i, kode int
	fmt.Scan(&kode)
	idx = -1
	i = 0
	if pilih == "M" {
		for idx == -1 && i < Q.nMasuk {
			if Q.masuk[i].kode == kode {
				idx = i
			}
			i++
		}
	} else {
		for idx == -1 && i < Q.nKeluar {
			if Q.keluar[i].kode == kode {
				idx = i
			}
			i++
		}
	}
	return idx
}

func editTransaction(Q *akun) {
	var idx, newInt int
	var pilih, change, newString string
	fmt.Print("Mengganti transaksi MASUK/KELUAR (M/K)? ")
	fmt.Scan(&pilih)
	for pilih != "M" && pilih != "K" {
		fmt.Print("Mengganti transaksi MASUK/KELUAR (M/K)? ")
		fmt.Scan(&pilih)
	}
	fmt.Print("Masukkan KODE transaksi: ")
	idx = searchTrans_Edit(*Q, pilih)
	fmt.Print("Ingin mengganti (kode/nama): ")
	fmt.Scan(&change)
	if idx != -1 {
		if pilih == "M" {
			if change == "kode" {
				for {
					fmt.Print("Masukkan KODE terbaru: ")
					fmt.Scan(&newInt)
					if checkTransCode(*Q, newInt, pilih) {
						break
					}
					fmt.Println("Kode INVALID.")
				}
				Q.masuk[idx].kode = newInt
			} else if change == "nama" {
				fmt.Print("Masukkan NAMA terbaru: ")
				fmt.Scan(&newString)
				Q.masuk[idx].orang = newString
			}
		} else {
			if change == "kode" {
				for {
					fmt.Print("Masukkan KODE terbaru: ")
					fmt.Scan(&newInt)
					if checkTransCode(*Q, newInt, pilih) {
						break
					}
					fmt.Println("Kode INVALID.")
				}
				Q.keluar[idx].kode = newInt
			} else if change == "nama" {
				fmt.Print("Masukkan NAMA terbaru: ")
				fmt.Scan(&newString)
				Q.keluar[idx].orang = newString
			}
		}
		fmt.Println("Edit data transaksi BERHASIL.")
	} else {
		fmt.Println("Transaksi TIDAK DITEMUKAN.")
	}
}

func sortTransactionByCode(Q *akun, pilih string) {
	var pass, i int
	var temp bayar
	pass = 1
	if pilih == "M" {
		for pass < Q.nMasuk {
			i = pass
			temp = Q.masuk[pass]
			for i > 0 && temp.kode < Q.masuk[i-1].kode {
				Q.masuk[i] = Q.masuk[i-1]
				i--
			}
			Q.masuk[i] = temp
			pass++
		}
	} else {
		for pass < Q.nKeluar {
			i = pass
			temp = Q.keluar[pass]
			for i > 0 && temp.kode < Q.keluar[i-1].kode {
				Q.keluar[i] = Q.keluar[i-1]
				i--
			}
			Q.keluar[i] = temp
			pass++
		}
	}
}

func codeBinarySearch(Q *akun, Kode int, pilih string) int {
	var left, right, mid, idx int
	idx = -1
	if pilih == "M" {
		sortTransactionByCode(Q, pilih)
		left = 0
		right = Q.nMasuk - 1
		for left <= right && idx == -1 {
			mid = (left + right) / 2
			if Q.masuk[mid].kode < Kode {
				left = mid + 1
			} else if Q.masuk[mid].kode > Kode {
				right = left - 1
			} else {
				idx = mid
			}
		}
	} else {
		sortTransactionByCode(Q, pilih)
		left = 0
		right = Q.nKeluar - 1
		for left <= right && idx == -1 {
			mid = (left + right) / 2
			if Q.keluar[mid].kode < Kode {
				left = mid + 1
			} else if Q.keluar[mid].kode > Kode {
				right = left - 1
			} else {
				idx = mid
			}
		}
	}
	return idx
}

func searchTransactionByCode(Q *akun) {
	var Kode, k int
	var pilih string
	fmt.Print("Mencari transaksi MASUK/KELUAR (M/K)? ")
	fmt.Scan(&pilih)
	for pilih != "M" && pilih != "K" {
		fmt.Print("Mencari transaksi MASUK/KELUAR (M/K)? ")
		fmt.Scan(&pilih)
	}
	fmt.Print("Masukkan KODE transaksi: ")
	fmt.Scan(&Kode)
	k = codeBinarySearch(Q, Kode, pilih)
	fmt.Println("Lokasi index:", k)
	fmt.Println("-------------------------")
	if k != -1 {
		if pilih == "M" {
			fmt.Println("Nama     :", Q.masuk[k].orang)
			fmt.Println("Barang   :", Q.masuk[k].kargo.nama)
			fmt.Println("Kategori :", Q.masuk[k].kargo.kategori)
			fmt.Println("Jumlah   :", Q.masuk[k].kargo.jumlah, "unit")
		} else {
			fmt.Println("Nama     :", Q.keluar[k].orang)
			fmt.Println("Barang   :", Q.keluar[k].kargo.nama)
			fmt.Println("Kategori :", Q.keluar[k].kargo.kategori)
			fmt.Println("Jumlah   :", Q.keluar[k].kargo.jumlah, "unit")
		}
		fmt.Println("Pencarian transaksi BERHASIL.")
	} else {
		fmt.Println("Transaksi TIDAK DITEMUKAN.")
	}
}

func printTransaction(Q akun) {
	var i int
	fmt.Println("Daftar transaksi MASUK:")
	sortTransactionByCode(&Q, "M")
	for i = 0; i < Q.nMasuk; i++ {
		fmt.Print("[", i+1, "] ", Q.masuk[i].kode, " ", Q.masuk[i].orang, " ", Q.masuk[i].kargo.nama)
		fmt.Println(" ", Q.masuk[i].kargo.kategori, Q.masuk[i].kargo.jumlah)
	}
	fmt.Println("-------------------------")
	fmt.Println("Daftar transaksi KELUAR:")
	sortTransactionByCode(&Q, "K")
	for i = 0; i < Q.nKeluar; i++ {
		fmt.Print("[", i+1, "] ", Q.keluar[i].kode, " ", Q.keluar[i].orang, " ", Q.keluar[i].kargo.nama, " ")
		fmt.Println(Q.keluar[i].kargo.kategori, Q.keluar[i].kargo.jumlah)
	}
}

func displayHeader() {
	fmt.Print("\n")
	fmt.Println("=========================")
	fmt.Println("     Store Inventory    ")
	fmt.Println("    Shining Star Store   ")
	fmt.Println("      DISPLAY MENU       ")
	fmt.Println("-------------------------")
	fmt.Println("1. Full Tak Terurut")
	fmt.Println("2. Full Ascending")
	fmt.Println("3. Full Descending")
	fmt.Println("4. Category (Descending)")
	fmt.Println("0. Back")
	fmt.Println("=========================")
	fmt.Print("Pilihan anda (1/2/3/4/0): ")
}

func displayMenu(P *gudang) {
	/* I.S. Terdefinisi struct gudang P berisi n data barang
	F.S. Menampilkan header DISPLAY MENU dan input opsi secara berulang. Pengguna akan
	diarahkan ke menu sesuai dengan opsi. Perulangan berakhir ketika opsi bernilai 0 */
	var arr1, arr2, arr3 Inventory
	var opsi, N1, N2, N3 int
	var kategori string
	for {
		displayHeader()
		fmt.Scan(&opsi)
		for opsi != 1 && opsi != 2 && opsi != 3 && opsi != 4 && opsi != 0 {
			fmt.Print("Pilihan anda (1/2/3/4/0): ")
			fmt.Scan(&opsi)
		}
		switch opsi {
		case 1:
			printInventory(P.list, P.n)
		case 2:
			fullSorted(P, "Ascending")
			printInventory(P.list, P.n)
		case 3:
			fullSorted(P, "Descending")
			printInventory(P.list, P.n)
		case 4:
			fmt.Print("Masukkan kategori: ")
			fmt.Scan(&kategori)
			switch kategori {
			case "F&B":
				categorySorted(*P, kategori, &arr1, &N1)
				printInventory(arr1, N1)
			case "Med":
				categorySorted(*P, kategori, &arr2, &N2)
				printInventory(arr2, N2)
			case "ATK":
				categorySorted(*P, kategori, &arr3, &N3)
				printInventory(arr3, N3)
			default:
				fmt.Println("Kategori TIDAK VALID.")
			}
		}
		if opsi == 0 {
			break
		}
	}
}

func printInventory(arr Inventory, n int) {
	/* I.S. Terdefinisi array Inventory arr berisi n data barang
	F.S. Mencetak n data barang di piranti keluaran, 1 baris setiap 1 barang */
	var i int
	for i = 0; i < n; i++ {
		fmt.Println("[", i+1, "] ", arr[i].jumlah, arr[i].nama, arr[i].kategori)
	}
}

func fullSorted(P *gudang, sorting string) {
	/* I.S. Terdefinisi struct gudang P berisi P.n data dan string untuk menentukan
	   jenis sorting yang akan dipilih (Ascending/Descending).
	F.S. Data barang pada array di struct gudang P sudah diurutkan ASCENDING atau
	     DESCENDING (tergantung string sorting) dengan SELECTION SORT. */
	var pass, idx, i int
	var temp barang
	pass = 1
	for pass < P.n {
		idx = pass - 1
		i = pass
		for i < P.n {
			switch sorting {
			case "Ascending":
				if P.list[idx].jumlah > P.list[i].jumlah {
					idx = i
				}
			case "Descending":
				if P.list[idx].jumlah < P.list[i].jumlah {
					idx = i
				}
			}
			i++
		}
		temp = P.list[pass-1]
		P.list[pass-1] = P.list[idx]
		P.list[idx] = temp
		pass++
	}
}

func inputByCategory(P gudang, x string, list *Inventory, n *int) {
	/* I.S. Terdefinisi struct gudang P berisi P.n barang, string x yang merupakan kategori, array
	   Inventory list (terdefinisi bebas), dan n yang merupakan jumlah barang (terdefinisi bebas).
	   Proses: Memasukkan data barang dengan kategori x dari array utama ke dalam array list.
	F.S. Array Inventory list sudah berisi n data barang dengan kategori x. */
	var idx, i int
	for i = 0; i < P.n; i++ {
		if P.list[i].kategori == x {
			idx = checkItem(P, x)
			if idx == -1 {
				list[*n].nama = P.list[i].nama
				list[*n].jumlah = P.list[i].jumlah
			} else {
				list[idx].jumlah = P.list[i].jumlah
			}
			*n++
		}
	}
}

func categorySorted(P gudang, kategori string, list *Inventory, n *int) {
	/* I.S. Terdefinisi struct gudang P berisi P.n barang, string ketagori yang merupakan
	   kategori barang, array Inventory list, dan n yang merupakan jumlah barang.
	F.S. Array Inventory list yang berisi n data barang dengan kategori x sudah terurut
	secara DESCENDING dengan INSERTION SORT. */
	var pass, i int
	var temp barang
	inputByCategory(P, kategori, list, n)
	fmt.Println("Daftar barang", kategori, "di gudang:")
	pass = 1
	for pass < *n {
		i = pass
		temp = list[i]
		for i > 0 && temp.jumlah > list[i-1].jumlah {
			list[i] = list[i-1]
			i--
		}
		list[i] = temp
		pass++
	}
}

func main() {
	var Storage gudang
	var Account akun
	login()
	mainMenu(&Storage, &Account)
}
