package main

import "fmt"

func main() {
	ak47 := getGun("ak47")

	print(ak47)
}

func print(gun iGun) {
	fmt.Printf("Gunname :%s, GunPower: %d", gun.getName(), gun.getPower())

}
