package main

func getGun(gunType string) iGun {
	if gunType == "ak47" {
		t := new(ak47)
		return t.new()
	}

	return nil
}
