package resist

func Rvp(v, p float64) float64 {
	return (v * v) / p
}

func Riv(i, v float64) float64 {
	return v / i
}

func Rip(i, p float64) float64 {
	return p / (i * i)
}

func Rser(resists ...float64) (Rtotal float64) {
	for _, r := range resists {
		Rtotal = Rtotal + r
	}
	return
}

func Rpara(resists ...float64) (Rtotal float64) {
	for _, r := range resists {
		Rtotal = Rtotal + (1/r)
	}
	return 1/Rtotal
}