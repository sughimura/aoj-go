package main

import "fmt"

type Dice struct {
	label1, label2, label3, label4, label5, label6 int
}

func (d *Dice) rollToW() {
	tmp := d.label1
	d.label1 = d.label3
	d.label3 = d.label6
	d.label6 = d.label4
	d.label4 = tmp
}

func (d *Dice) rollToE() {
	tmp := d.label1
	d.label1 = d.label4
	d.label4 = d.label6
	d.label6 = d.label3
	d.label3 = tmp
}

func (d *Dice) rollToS() {
	tmp := d.label1
	d.label1 = d.label5
	d.label5 = d.label6
	d.label6 = d.label2
	d.label2 = tmp
}

func (d *Dice) rollToN() {
	tmp := d.label1
	d.label1 = d.label2
	d.label2 = d.label6
	d.label6 = d.label5
	d.label5 = tmp
}

func (d *Dice) top() int {
	return d.label1
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
