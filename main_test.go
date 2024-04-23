package main

import "testing"

func TestPembayaranBarang(t *testing.T) {
	err := PembayaranBarang(0, "cod", false)
	if err == nil {
		t.Error("Diharapkan error harga nol")
	}

	err = PembayaranBarang(100000, "barter", false)
	if err == nil {
		t.Error("Diharapkan pembayaran tidak valid")
	}

	err = PembayaranBarang(200000, "credit", false)
	if err == nil {
		t.Error("Diharapkan cicilan error")
	}

	err = PembayaranBarang(100000, "cod", false)
	if err == nil {
		t.Error("Diharapkan pembayaran tidak valid")
	}
}
