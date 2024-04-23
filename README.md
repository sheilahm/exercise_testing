exercise testing BE01

buat func PembayaranBarang
parameter : hargaTotal (float64), metode pembayaran (string), dicicil (bool)

func PembayaranBarang(ht float64, m string, cicil bool) error

- cek hargaTotal > 0 apabila <= 0 maka return error "harga tidak bisa nol"

- cek metode pembayaran: "cod", "transfer", "debit", "credit", "gerai" kalo tidak yg diatas maka return error "metode tidak dikenali"

- cek dicicil atau tidak, kalo true, metode harus credit, dan hargaTotal harus >= 500.000 return error "cicilan tidak memenuhi syarat"

kalo false, metode bukan credit. return error "credit harus dicicil"
