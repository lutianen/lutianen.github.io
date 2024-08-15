package main

import "fmt"

// iGun.go
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// gun.go
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) setPower(power int) {
	g.power = power
}
func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getPower() int {
	return g.power
}

// ak47.go
type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 Gun",
			power: 4,
		},
	}
}

// musket.go
type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket Gun",
			power: 1,
		},
	}
}

// GunFunctory.go
func GetGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	} else if gunType == "musket" {
		return newMusket(), nil
	}

	return nil, fmt.Errorf("Wrong gun type passed")
}

// Client.go
func main() {
	pGun, _ := GetGun("ak47")
	pGun.getName()
}
