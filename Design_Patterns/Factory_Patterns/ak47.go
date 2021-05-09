package main

type ak47 struct {
	gun
}

func (a *ak47) new() iGun {
	a.gun.setName("ak47")
	a.gun.setPower(100)

	return a
}
