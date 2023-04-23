package volt

// Vir returns the voltage value when current i and resistance r are known
func Vir(i, r float64) float64 {
	return i * r
}

// Vip returns the voltage value when current i and power p are known
func Vip(i, p float64) float64 {
	return p / i
}

// Vser calculates total circuit voltage in serial
func Vser(volts ...float64) (Vtotal float64) {
	for _, v := range volts {
		Vtotal = Vtotal + v
	}
	return
}
