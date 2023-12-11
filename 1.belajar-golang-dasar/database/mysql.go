package database

var koneksi string

func init() {
	koneksi = "MySql"
}

func GetDatabase() string {
	return koneksi
}
