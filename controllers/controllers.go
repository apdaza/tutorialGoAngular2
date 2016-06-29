package controllers

import (
	"github.com/gin-gonic/gin"

  . "github.com/apdaza/tutorial/models"
  . "github.com/apdaza/tutorial/data"
  . "github.com/apdaza/tutorial/utils"
)

var dbmap = InitDb()

/*GetContactos funcion para obtener todos los contactos*/
func GetContactos(c *gin.Context) {
  var contactos []Contacto
  _, err := dbmap.Select(&contactos, "SELECT * FROM contacto")
  if err == nil {
    c.JSON(200, contactos)
  } else {
    c.JSON(404, gin.H{"error": "no hay contactos en la tabla"})
  }
}

/*PostContacto funcion para insertar un contacto*/
func PostContactos(c *gin.Context) {
  var contacto Contacto
  c.Bind(&contacto)
  if contacto.Nombre != "" && contacto.Apellido != "" && contacto.Celular != "" && contacto.Mail != ""{
    if insert, _ := dbmap.Exec(`INSERT INTO contacto (nombre, apellido, celular, mail) VALUES (?, ?, ?, ?)`,
                    contacto.Nombre, contacto.Apellido, contacto.Celular, contacto.Mail); insert != nil {
      contacto_id, err := insert.LastInsertId()
      if err == nil {
        content := &Contacto{
          Id:   contacto_id,
          Nombre: contacto.Nombre,
          Apellido: contacto.Apellido,
          Celular: contacto.Celular,
          Mail: contacto.Mail,
        }
        c.JSON(201, content)
      } else {
        CheckErr(err, "error en el Insert")
      }
    }
  } else {
    c.JSON(422, gin.H{"error": "campos vacios"})
  }
}
