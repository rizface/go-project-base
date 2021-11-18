package app

/*

	pada folder ini berisi model - model yang akan migrate ke database

*/

type Users struct {
	Email    string `gorm:"type:not null;unique"`
	Username string `gorm:"not null;unique"`
}
