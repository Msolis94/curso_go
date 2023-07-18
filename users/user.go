package users

import (
	"curso_go/Modelos"
	"fmt"
	"time"
)

func AltaUsuario() {
	u := new(Modelos.User)
	u.AddUser(10, "marlon", time.Now(), true)
	fmt.Println(u)
}
