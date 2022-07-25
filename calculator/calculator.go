package calculator

import "os"

type ArithmeticInterface interface {
	Add(operand float64)
	Subtract(operand float64)
	Multiply(operand float64)
	Divide(operand float64)
}

type UtilityInterface interface {
	Cancel()
	Exit()
}

type ViewInterface interface {
	GetValue() float64
}

type CalcInterface interface {
	ArithmeticInterface
	ViewInterface
	UtilityInterface
}

type Calculator struct {
	output float64
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (calc *Calculator) Add(operand float64) {
	calc.output += operand
}

func (calc *Calculator) Subtract(operand float64) {
	calc.output -= operand
}

func (calc *Calculator) Multiply(operand float64) {
	calc.output *= operand
}

func (calc *Calculator) Divide(operand float64) {
	if operand == 0 {
		panic("division by zero is not possible")
	}
	calc.output /= operand
}

func (calc Calculator) GetValue() float64 {
	return calc.output
	
}

func (calc *Calculator) Cancel() {
	calc.output = 0.0
}

func (calc *Calculator) Exit() {
	os.Exit(0)
}
