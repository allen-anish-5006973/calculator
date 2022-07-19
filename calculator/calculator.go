package calculator

type Calculator struct {
	output float64
}

func NewCalculator() Calculator {
	return Calculator{}
}

func (calc *Calculator) Add(operand float64) {
	calc.output += operand
}
