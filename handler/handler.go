package handler

import (
	"maths/calculator"
	"maths/constants"
	"maths/parser"
	"maths/renderer"
	"os"
)

type function func(operand float64, calcInterface calculator.CalcInterface)

type FunctionMap map[constants.Operations]function

var Registry = make(FunctionMap)

var CommandList []parser.Parser

func RegisterHandler(operationString constants.Operations, operation function) bool {
	for op, _ := range Registry {
		if operationString == op {
			return true
		}
	}
	Registry[operationString] = operation
	return false
}

func ExecuteHandler(operand float64, operationString constants.Operations, calcInterface calculator.CalcInterface) {
	if Registry[operationString] == nil {
		renderer.ViewError(os.Stdout, "Invalid Function")
		return
	}
	newCommand := parser.NewParser(operationString, operand)
	if operationString != "repeat" {
		CommandList = append(CommandList, *newCommand)
	}
	Registry[operationString](operand, calcInterface)
}

func HandleAddition(operand float64, calcInterface calculator.CalcInterface) {
	calcInterface.Add(operand)
	renderer.ViewValue(os.Stdout, calcInterface.GetValue())
}

func HandleRepeat(count float64, calcInterface calculator.CalcInterface) {
	a := len(CommandList)
	for i := a - int(count); i < a; i++ {
		ExecuteHandler(CommandList[i].Operand, CommandList[i].Operator, calcInterface)
	}
}

func HandleSubtraction(operand float64, calcInterface calculator.CalcInterface) {
	calcInterface.Subtract(operand)
	renderer.ViewValue(os.Stdout, calcInterface.GetValue())
}

func HandleMultiplication(operand float64, calcInterface calculator.CalcInterface) {
	calcInterface.Multiply(operand)
	renderer.ViewValue(os.Stdout, calcInterface.GetValue())
}

func HandleDivision(operand float64, calcInterface calculator.CalcInterface) {
	calcInterface.Divide(operand)
	renderer.ViewValue(os.Stdout, calcInterface.GetValue())
}

func HandleCancel(operand float64, calcInterface calculator.CalcInterface) {
	calcInterface.Cancel()
	renderer.ViewValue(os.Stdout, calcInterface.GetValue())
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
	RegisterHandler("repeat", HandleRepeat)
	RegisterHandler("exit", HandleExit)

}
