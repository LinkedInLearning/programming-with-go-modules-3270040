package volt

func Vir(i, r float64) float64 {
	return i * r
}

func Vip(i, p float64) float64 {
	return p / i
}

func Vser(volts ...float64) (Vtotal float64) {
	for _, v := range volts {
		Vtotal = Vtotal + v
	}
	return
}
