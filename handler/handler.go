package handler

import (
	"fmt"
	"maths/calculator"
	"maths/constants"
	"reflect"
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
		panic("Invalid Function")
	}
	Registry[operationString](operand, calcInterface)
}

func HandleAddition(operand float64, calcInterface calculator.CalcInterface) {
	calcInterface.Add(operand)
	//fmt.Println(calcInterface.GetValue())
	fmt.Println(reflect.ValueOf(calcInterface), reflect.TypeOf(calcInterface))
}

func HandleSubtraction(operand float64, calcInterface calculator.CalcInterface) {
	calcInterface.Subtract(operand)
	fmt.Println(calcInterface.GetValue())
}

func HandleMultiplication(operand float64, calcInterface calculator.CalcInterface) {
	calcInterface.Multiply(operand)
	fmt.Println(calcInterface.GetValue())
}

func HandleDivision(operand float64, calcInterface calculator.CalcInterface) {
	calcInterface.Divide(operand)
	fmt.Println(calcInterface.GetValue())
}

func HandleCancel(operand float64, calcInterface calculator.CalcInterface) {
	calcInterface.Cancel()
	fmt.Println(calcInterface.GetValue())
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
