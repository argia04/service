package main

import (
   "fmt"
   "os" // untuk pilihan keluar
)

const capasityArray int = 1000 // kapasitas array adalah 1000

type datasparePart struct {
   nama   string
   kode   string
   jumlah int
   harga  int
}

type tabData struct {
   sparePart [capasityArray]datasparePart
   pelanggan [capasityArray]dataPelanggan
   transaksi [capasityArray]riwayatTransaksi
   total     int
}

type dataPelanggan struct {
   nama  string
   kode  string
   waktu string
}

type riwayatTransaksi struct {
   kodeTransaksi int
   kodePelanggan int
   kodeSparepart int
}

func tambahDataSparepart(data *tabData) {
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡")
   fmt.Print("Masukkan kode barang: ")
   fmt.Scan(&data.sparePart[data.total].kode)
   fmt.Print("Masukkan nama barang: ")
   fmt.Scan(&data.sparePart[data.total].nama)
   fmt.Print("Masukkan jumlah barang: ")
   fmt.Scan(&data.sparePart[data.total].jumlah)
   fmt.Print("Masukkan harga barang per item: ")
   fmt.Scan(&data.sparePart[data.total].harga)
   fmt.Println("")
   data.total++

   updateHargaData(data)
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡")
}

func tambahDataPelanggan(data *tabData) {
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡")
   fmt.Print("Masukkan nama pelanggan: ")
   fmt.Scan(&data.pelanggan[data.total].nama)
   fmt.Print("Masukkan kode pelanggan: ")
   fmt.Scan(&data.pelanggan[data.total].kode)
   fmt.Print("Masukkan Waktu: ")
   fmt.Scan(&data.pelanggan[data.total].waktu)
   fmt.Println("")
   data.total++
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡")
}

func tambahDataTransaksi(data *tabData) {
   fmt.Print("Masukkan kode transaaksi: ")
   fmt.Scan(&data.transaksi[data.total].kodeTransaksi)
   fmt.Print("Masukkan jumlah spare-part: ")
   fmt.Scan(&data.transaksi[data.total].kodePelanggan)
   fmt.Print("Masukkan harga spare-part per item: ")
   fmt.Scan(&data.transaksi[data.total].kodeSparepart)
   fmt.Println("")
   data.total++
}

func updateHargaData(data *tabData) {
   for i := 0; i < data.total; i++ {
      data.sparePart[i].harga *= data.sparePart[i].jumlah
   }
}

func tampilkanDataSparepart(data *tabData) {
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡ ")
   for i := 0; i < data.total; i++ {
      fmt.Println("Data barang ke:", i+1)
      fmt.Println("kode barang:", data.sparePart[i].kode)
      fmt.Println("Nama barang:", data.sparePart[i].nama)
      fmt.Println("Jumlah barang:", data.sparePart[i].jumlah)
      fmt.Println("harga barang total:", data.sparePart[i].harga)
      fmt.Println("")
   }
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡ ")
}

func tampilkanDataPelanggan(data *tabData) {
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡ ")
   for i := 0; i < data.total; i++ {
      fmt.Println("Data barang ke:", i+1)
      fmt.Println("kode pelanggan:", data.pelanggan[i].kode)
      fmt.Println("Nama Pelanggan:", data.pelanggan[i].nama)
      fmt.Println("waktu:", data.pelanggan[i].waktu)
      fmt.Println("")
   }
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡ ")
}

func tampilkanDataTransaksi(data *tabData) {
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡ ")
   for i := 0; i < data.total; i++ {
      fmt.Println("Data barang ke:", i+1)
      fmt.Println("Kode Transaksi:", data.transaksi[i].kodeTransaksi)
      fmt.Println("Kode Pelanggan:", data.transaksi[i].kodePelanggan)
      fmt.Println("Kode Spare-part:", data.transaksi[i].kodeSparepart)
      fmt.Println("")
   }
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡ ")
}

func cariDataKode(data tabData, kode string) int {
   // Sequential Search berdasarkan kode
   for i := 0; i < data.total; i++ {
      if data.sparePart[i].kode == kode {
         return i
      }
   }
   return -1
}

func cariDataWaktu(data tabData, waktu string) int {
   // Sequential Search berdasarkan waktu
   for i := 0; i < data.total; i++ {
      if data.pelanggan[i].waktu == waktu {
         return i
      }
   }
   return -1
}

func tampilkanDataDicariKode(data tabData) {
   var kode string
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡")
   fmt.Println("Pencarian Kode")
   for true {
      fmt.Print("Kode barang: ")
      fmt.Scan(&kode)
      if cariDataKode(data, kode) != -1 {
         fmt.Println("Nama barang:", data.sparePart[cariDataKode(data, kode)].nama)
         fmt.Println("Jumlah barang:", data.sparePart[cariDataKode(data, kode)].jumlah)
         fmt.Println("Harga barang:", data.sparePart[cariDataKode(data, kode)].harga)
         fmt.Println("")
      } else {
         fmt.Println("Data tidak ditemukan")
         break
      }
   }
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡")
}

func tampilkanDataDicariWaktu(data tabData) {
   var waktu string

   fmt.Println("Pencarian Waktu")
   for true {
      fmt.Print("Waktu: ")
      fmt.Scan(&waktu)
      if cariDataWaktu(data, waktu) != -1 {
         fmt.Println("Nama Pelanggan:", data.pelanggan[cariDataWaktu(data, waktu)].nama)
         fmt.Println("Kode Pelanggan:", data.pelanggan[cariDataWaktu(data, waktu)].kode)
         fmt.Println("Waktu:", data.pelanggan[cariDataWaktu(data, waktu)].waktu)
         fmt.Println("")
      } else {
         fmt.Println("Data tidak ditemukan")
         break
      }
   }
}

func sortJumlahSparepart(data *tabData) {
   for i := 0; i < (*data).total-1; i++ {
      max := i
      for j := i + 1; j < (*data).total; j++ {
         if (*data).sparePart[max].jumlah < (*data).sparePart[j].jumlah {
            max = j
         }
      }
      temp := (*data).sparePart[i]
      (*data).sparePart[i] = (*data).sparePart[max]
      (*data).sparePart[max] = temp
   }
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡")
   fmt.Println("Urutan data berdasarkan jumlah spare-part paling sering diganti")
   for i := 0; i < data.total; i++ {
      fmt.Printf("urutan %d: %s -> %d buah \n", i+1, data.sparePart[i].nama, data.sparePart[i].jumlah)
   }
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡")
}

func hapusDataPelanggan(data *tabData) {
   var kode string
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡")
   fmt.Println("Hapus data berdasarkan kode")
   for true {
      fmt.Print("Kode barang: ")
      fmt.Scan(&kode)
      if cariDataKode(*data, kode) != -1 {
         (*data).pelanggan[cariDataKode(*data, kode)] = (*data).pelanggan[(*data).total-1]
         (*data).total--
         fmt.Println("Data berhasil dihapus")
         break
      } else {
         fmt.Println("Data tidak ditemukan")
         break
      }
   }
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡")
}

func hapusDataTransaksi(data *tabData) {
   var kode string
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡")
   fmt.Println("Hapus data berdasarkan kode")
   for true {
      fmt.Print("Kode barang: ")
      fmt.Scan(&kode)
      if cariDataKode(*data, kode) != -1 {
         (*data).transaksi[cariDataKode(*data, kode)] = (*data).transaksi[(*data).total-1]
         (*data).total--
         fmt.Println("Data berhasil dihapus")
         break
      } else {
         fmt.Println("Data tidak ditemukan")
         break
      }
   }
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡")
}

func hapusDataSparepart(data *tabData) {
   var kode string
   fmt.Println("Hapus data berdasarkan kode")
   for true {
      fmt.Print("Kode barang: ")
      fmt.Scan(&kode)
      if cariDataKode(*data, kode) != -1 {
         (*data).sparePart[cariDataKode(*data, kode)] = (*data).sparePart[(*data).total-1]
         (*data).total--
         fmt.Println("Data berhasil dihapus")
         break
      } else {
         fmt.Println("Data tidak ditemukan")
         break
      }
   }
}
func editDataSparepart(data *tabData) {
   var kode string
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡")
   fmt.Println("Edit data Spare-Part")
   fmt.Print("Masukkan kode barang yang akan diedit: ")
   fmt.Scan(&kode)

   index := cariDataKode(*data, kode)
   if index != -1 {
      fmt.Println("Data saat ini:")
      fmt.Println("Kode barang:", data.sparePart[index].kode)
      fmt.Println("Nama barang:", data.sparePart[index].nama)
      fmt.Println("Jumlah barang:", data.sparePart[index].jumlah)
      fmt.Println("Harga barang:", data.sparePart[index].harga)
      fmt.Println("")

      fmt.Print("Masukkan nama barang baru: ")
      fmt.Scan(&data.sparePart[index].nama)
      fmt.Print("Masukkan jumlah barang baru: ")
      fmt.Scan(&data.sparePart[index].jumlah)
      fmt.Print("Masukkan harga barang per item baru: ")
      fmt.Scan(&data.sparePart[index].harga)

      fmt.Println("Data berhasil diupdate")
   } else {
      fmt.Println("Data tidak ditemukan")
   }
   fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡")
}

func editDataPelanggan(data *tabData) {
   var kode string
   fmt.Println("Edit data Pelanggan")
   fmt.Print("Masukkan kode pelanggan yang akan diedit: ")
   fmt.Scan(&kode)

   index := cariDataKode(*data, kode)
   if index != -1 {
      fmt.Println("Data saat ini:")
      fmt.Println("Kode pelanggan:", data.pelanggan[index].kode)
      fmt.Println("Nama pelanggan:", data.pelanggan[index].nama)
      fmt.Println("Waktu:", data.pelanggan[index].waktu)
      fmt.Println("")

      fmt.Print("Masukkan nama pelanggan baru: ")
      fmt.Scan(&data.pelanggan[index].nama)
      fmt.Print("Masukkan waktu baru: ")
      fmt.Scan(&data.pelanggan[index].waktu)

      fmt.Println("Data berhasil diupdate")
   } else {
      fmt.Println("Data tidak ditemukan")
   }
}

func editDataTransaksi(data *tabData) {
   var kode string
   fmt.Println("Edit data Transaksi")
   fmt.Print("Masukkan kode transaksi yang akan diedit: ")
   fmt.Scan(&kode)

   index := cariDataKode(*data, kode)
   if index != -1 {
      fmt.Println("Data saat ini:")
      fmt.Println("Kode transaksi:", data.transaksi[index].kodeTransaksi)
      fmt.Println("Kode pelanggan:", data.transaksi[index].kodePelanggan)
      fmt.Println("Kode spare-part:", data.transaksi[index].kodeSparepart)
      fmt.Println("")

      fmt.Print("Masukkan kode pelanggan baru: ")
      fmt.Scan(&data.transaksi[index].kodePelanggan)
      fmt.Print("Masukkan kode spare-part baru: ")
      fmt.Scan(&data.transaksi[index].kodeSparepart)

      fmt.Println("Data berhasil diupdate")
   } else {
      fmt.Println("Data tidak ditemukan")
   }
}
func main() {
   var data tabData
   var pilihan int

   for true {
      fmt.Println("☆.。.:*・°☆.。.:*・°Hai, silahkan pilih! ☆.。.:*・°☆.。.:*・°☆ ")
      fmt.Println("♡ ´･ᴗ･ `♡  (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ (っ◕‿◕)っ  ♡ ´･ᴗ･ `♡ ")
      fmt.Println("･ᴗ･ 1. Tambah data Spare-Part･ᴗ･ ")
      fmt.Println("･ᴗ･ 2. Edit data Spare-Part ･ᴗ･ ")
      fmt.Println("･ᴗ･ 3. Hapus data Spare-Part ･ᴗ･ ")
      fmt.Println("･ᴗ･ 4. Tambah data Pelanggan ･ᴗ･ ")
      fmt.Println("･ᴗ･ 5. Edit data Pelanggan ･ᴗ･ ")
      fmt.Println("･ᴗ･ 6. Hapus data Pelanggan ･ᴗ･ ")
      fmt.Println("･ᴗ･ 7. Tambah data Transaksi ･ᴗ･ ")
      fmt.Println("･ᴗ･ 8. Edit data Transaksi ･ᴗ･ ")
      fmt.Println("･ᴗ･ 9. Hapus data Transaksi ･ᴗ･ ")
      fmt.Println("･ᴗ･ 10. Cari data Pelanggan berdasarkan Waktu ･ᴗ･ ")
      fmt.Println("･ᴗ･ 11. Cari data Pelanggan berdasarkan kode Spare-Part ･ᴗ･ ")
      fmt.Println("･ᴗ･ 12. Tampilkan data Spare Part yang Paling Sering Diganti ･ᴗ･ ")
      fmt.Println("･ᴗ･ 13. Tampilkan data Spare-Part ･ᴗ･ ")
      fmt.Println("･ᴗ･ 14. Tampilkan data Pelanggan ･ᴗ･ ")
      fmt.Println("･ᴗ･ 15. Tampilkan data Transaksi ･ᴗ･ ")
      fmt.Println("･ᴗ･ 16. Keluar ･ᴗ･ ")
      fmt.Print("Pilihanmu ? ･ᴗ･ ･ᴗ･ ･ᴗ･   ")

      fmt.Scan(&pilihan)
      fmt.Println("")

      switch pilihan {
      case 1:
         tambahDataSparepart(&data)
      case 2:
         editDataSparepart(&data)
      case 3:
         hapusDataSparepart(&data)
      case 4:
         tambahDataPelanggan(&data)
      case 5:
         editDataPelanggan(&data)
      case 6:
         hapusDataPelanggan(&data)
      case 7:
         tambahDataTransaksi(&data)
      case 8:
         editDataTransaksi(&data)
      case 9:
         hapusDataTransaksi(&data)
      case 10:
         tampilkanDataDicariWaktu(data)
      case 11:
         tampilkanDataDicariKode(data)
      case 12:
         sortJumlahSparepart(&data)
      case 13:
         tampilkanDataSparepart(&data)
      case 14:
         tampilkanDataPelanggan(&data)
      case 15:
         tampilkanDataTransaksi(&data)
      case 16:
         os.Exit(0)
      default:
         fmt.Println("Pilihan tidak tersedia")
      }
      fmt.Println("")
   }
}
