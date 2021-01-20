package main

import "fmt"

func main() {

  var matriks1 [10][10]int
  var matriks2 [10][10]int
  var hasil [10][10]int
  var i, j, k, m, n, p, q, jumlah int

  jumlah = 0

  fmt.Print("Masukkan jumlah baris matriks pertama: ")
  fmt.Scanln(&m)
  fmt.Print("Masukkan jumlah kolom matriks pertama: ")
  fmt.Scanln(&n)

  fmt.Print("Masukkan jumlah baris matriks kedua: ")
  fmt.Scanln(&p)
  fmt.Print("Masukkan jumlah kolom matriks kedua: ")
  fmt.Scanln(&q)

  if n != p {
    fmt.Println("Matriks tidak dapat dikalikan satu sama lain.")
  } else {

    fmt.Println("Masukkan elemen matriks pertama: ")
    for i = 0; i < m; i++ {
      for j = 0; j < n; j++ {
        fmt.Scan(&matriks1[i][j])
      }
    }

    fmt.Println("Masukkan elemen matriks kedua: ")
    for i = 0; i < p; i++ {
      for j = 0; j < q; j++ {
        fmt.Scan(&matriks2[i][j])
      }
    }

    for i = 0; i < m; i++ {
      for j = 0; j < q; j++ {
        for k = 0; k < p; k++ {
          jumlah = jumlah + matriks1[i][k]*matriks2[k][j]
        }
        hasil[i][j] = jumlah
        jumlah = 0
      }
    }

    fmt.Println("Hasil perkalian matriks: ")
    for i = 0; i < m; i++ {
      for j = 0; j < n; j++ {
        fmt.Print(hasil[i][j], "\t")
      }
      fmt.Print("\n")
    }
  }
}