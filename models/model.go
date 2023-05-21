package models

import "time"

type Barang struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Nama       string    `json:"nama"`
	Deskripsi  string    `json:"deskripsi"`
	JumlahStok int       `json:"jumlahstok"`
	Harga      int       `json:"harga"`
	Kategori   string    `json:"kategori"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
