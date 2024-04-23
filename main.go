package main

import (
	"errors"
	"fmt"
)

//PembayaranBarang

func PembayaranBarang(hargaTotal float64, metodePembayaran string, dicicil bool) error {
	//>0
	if hargaTotal <= 0 {
		return errors.New("harga tidak bisa nol")
	}

	//method pembayaran valid
	metodeValid := map[string]bool{
		"cod":      true,
		"transfer": true,
		"debit":    true,
		"credit":   true,
		"gerai":    true,
	}
	//cek pembayaran valid
	if !metodeValid[metodePembayaran] {
		return errors.New("metode tidak dikenali")
	}
	//dicicil harus credit & >=500000
	if dicicil {
		if metodePembayaran != "credit" {
			return errors.New("cicilan harus menggunakan credit")
		}
		if hargaTotal < 500000 {
			return errors.New("cicilan tidak memenuhi syarat")
		}
	} else {
		//!dicicil tidak boleh credit
		if metodePembayaran == "credit" {
			return errors.New("credit harus dicicil")
		}
	}
	return nil
}

func main() {
	err := PembayaranBarang(100000, "cod", false)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("pembayaran berhasil")
	}

	err = PembayaranBarang(500000, "credit", false)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("pembayaran berhasil")
	}

	err = PembayaranBarang(200000, "credit", false)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("pembayaran berhasil")
	}
}
