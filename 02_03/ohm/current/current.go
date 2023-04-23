package current

func Ivr(v, r float64) float64 {
	return v / r
}

func Ivp(v, p float64) float64 {
	return p / v
}

func Ipara(currents ...float64) (Itotal float64) {
	for _, i := range currents {
		Itotal = Itotal + i
	}
	return
}