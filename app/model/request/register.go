package request

/*

	pada folder ini berisi model / struct yang digunakan untuk menyimpan request dari client

	- required, digunakan untuk malakukan validasi, apakah sebuah input yang dikirim kosong atau tidak

*/

type User struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
}
