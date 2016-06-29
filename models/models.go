package models

/*Domain estructura con anotaciones para gorp y json*/
type Contacto struct {
	Id   int64  `db:"id" json:"id"`
	Nombre string `db:"nombre" json:"nombre"`
  Apellido string `db:"apellido" json:"apellido"`
  Celular string `db:"celular" json:"celular"`
  Mail string `db:"mail" json:"mail"`
}
