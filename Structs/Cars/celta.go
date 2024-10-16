package cars

import "fmt"

type Celta struct{ Car Car }

func (c *Celta) Accelerate(speed int) {
	celtinha := c.Car
	CurrentSpeed := celtinha.CurrentSpeed
	celtinha.CurrentSpeed = CurrentSpeed + speed
	fmt.Printf("Ao acelerar, a velocidade do %v era %v e agora mudou para %v.\n", celtinha.Name, CurrentSpeed, celtinha.CurrentSpeed)
}

func (c *Celta) Break(speed int) {
	celtinha := c.Car
	CurrentSpeed := celtinha.CurrentSpeed
	celtinha.CurrentSpeed = CurrentSpeed - speed
	fmt.Printf("Ao frear, a velocidade do %v era %v e agora mudou para %v.\n", celtinha.Name, CurrentSpeed, celtinha.CurrentSpeed)
}
