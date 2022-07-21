package handler

import (
	"maths/calculator"
	"maths/constants"
	"maths/renderer"
)

type function func(operand float64, calcInterface calculator.CalcInterface)

const add constants.Operations = "add"
const subtract constants.Operations = "subtract"
const multiply constants.Operations = "multiply"
const divide constants.Operations = "divide"
const cancel constants.Operations = "cancel"
const exit constants.Operations = "exit"

type FunctionMap map[constants.Operations]function

var Registry = make(FunctionMap)

func RegisterHandler(operationString constants.Operations, operation function) bool {
	for op, _ := range Registry {
		if operationString == op {
			panic("Duplicate entry")
		}
	}
	Registry[operationString] = operation
	return false
}

func ExecuteHandler(operand float64, operationString constants.Operations, calcInterface calculator.CalcInterface) {
	if Registry[operationString] == nil {
		renderer.ViewError("Invalid Function")
		return
	}
	Registry[operationString](operand, calcInterface)
}

func HandleAddition(operand float64, calcInterface calculator.CalcInterface) {
	calcInterface.Add(operand)
	renderer.ViewValue(calcInterface.GetValue())
}

func HandleSubtraction(operand float64, calcInterface calculator.CalcInterface) {
	calcInterface.Subtract(operand)
	renderer.ViewValue(calcInterface.GetValue())
}

func HandleMultiplication(operand float64, calcInterface calculator.CalcInterface) {
	calcInterface.Multiply(operand)
	renderer.ViewValue(calcInterface.GetValue())
}

func HandleDivision(operand float64, calcInterface calculator.CalcInterface) {
	calcInterface.Divide(operand)
	renderer.ViewValue(calcInterface.GetValue())
}

func HandleCancel(operand float64, calcInterface calculator.CalcInterface) {
	calcInterface.Cancel()
	renderer.ViewValue(calcInterface.GetValue())
}

func HandleExit(operand float64, calcInterface calculator.CalcInterface) {
	calcInterface.Exit()
}

func init() {
	RegisterHandler("add", HandleAddition)
	RegisterHandler("subtract", HandleSubtraction)
	RegisterHandler("multiply", HandleMultiplication)
	RegisterHandler("divide", HandleDivision)
	RegisterHandler("cancel", HandleCancel)
	RegisterHandler("exit", HandleExit)

}
