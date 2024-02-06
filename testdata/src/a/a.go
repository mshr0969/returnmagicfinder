package a

const (
    DiscountRate      = 0.9
	DiscountThreshold = 1200
)

func f() float64 {
	price := 1000.0
	return price * DiscountRate
}

func f1() float64 {
	price := 1000.0
	return price * 0.9 // want "magic number used in return statement"
}

func f2() float64 {
	price := 1000.0
	if price > DiscountThreshold {
		return price * DiscountRate
	}
	return price
}

func f3() float64 {
	return 100 * 1.2 + 10 // want "magic number used in return statement"
}
