package main

type TerroristDress struct {
	color string
}

func (c *TerroristDress) getColor() string {
	return c.color
}

func newTerroristDress() *TerroristDress {
	return &TerroristDress{color: "red"}
}
