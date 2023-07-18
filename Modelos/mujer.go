package Modelos

type Mujer struct {
	Hombre
}

func (this *Mujer) Respirar() { this.respirando = true }

func (this *Mujer) Pensar() { this.pensando = true }

func (this *Mujer) Comer() { this.comiendo = true }

func (this *Mujer) Sexo() string { return "Femenino" }
