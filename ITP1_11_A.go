package main

import "fmt"

type Dice struct {
	way1, way2, way3, way4, way5, way6 int
}

func (d *Dice) rollToW() {
	d.way1, d.way3, d.way6, d.way4 = d.way3, d.way6, d.way4, d.way1
}

func (d *Dice) rollToE() {
	d.way1, d.way4, d.way6, d.way3 = d.way4, d.way6, d.way3, d.way1
}

func (d *Dice) rollToS() {
	d.way1, d.way5, d.way6, d.way2 = d.way5, d.way6, d.way2, d.way1
}

func (d *Dice) rollToN() {
	d.way1, d.way2, d.way6, d.way5 = d.way2, d.way6, d.way5, d.way1
}

func (d *Dice) top() int {
	return d.way1
}

func main() {
	v := make([]int, 6)
	for i := 0; i < 6; i++ {
		fmt.Scan(&v[i])
	}
	d := Dice{v[0], v[1], v[2], v[3], v[4], v[5]}
	var directions string
	fmt.Scan(&directions)
	for _, direction := range directions {
		switch direction {
		case 'N':
			d.rollToN()
		case 'E':
			d.rollToE()
		case 'S':
			d.rollToS()
		case 'W':
			d.rollToW()
		}
	}
	fmt.Println(d.top())
}
