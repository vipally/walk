package candle

type Candle struct {
	Time  uint32
	Start float32
	End   float32
	Max   float32
	Min   float32
	Mount float32
}

func (c *Candle) TopLine() float32 {
	return 0
}
func (c *Candle) BottomLine() float32 {
	return 0
}
