package models

type Produk struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	Nama      string  `gorm:"not null" json:"nama"`
	Deskripsi string  `json:"deskripsi"`
	Harga     float64 `gorm:"not null" json:"harga"`
	Kategori  string  `gorm:"not null" json:"kategori"`
}

type Inventaris struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	IDProduk uint   `gorm:"not null" json:"id_produk"`
	Jumlah   int    `gorm:"not null" json:"jumlah"`
	Lokasi   string `gorm:"not null" json:"lokasi"`
}

type Pesanan struct {
	IDPesanan      uint   `gorm:"primaryKey" json:"id_pesanan"`
	IDProduk       uint   `gorm:"not null" json:"id_produk"`
	Jumlah         int    `gorm:"not null" json:"jumlah"`
	TanggalPesanan string `gorm:"not null" json:"tanggal_pesanan"`
}
