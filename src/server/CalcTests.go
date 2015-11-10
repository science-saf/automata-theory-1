package main

import "fmt"

type CalcTests struct {
	errors []string
}

func (self *CalcTests) run() {
	self.TestSingleValue()
}

func (self *CalcTests) TestSingleValue() {
	self.checkOneExpression("1", 1)
	self.checkOneExpression("1.", 1)
	self.checkOneExpression("1123", 1123)
	self.checkOneExpression("1123.", 1123)
	self.checkOneExpression(".0", 0)
	self.checkOneExpression(".1", 0.1)
	self.checkOneExpression(".13254678", 0.13254678)
	self.checkOneExpression("1.0", 1)
	self.checkOneExpression("1.2", 1.2)
	self.checkOneExpression("1123.0", 1123)
	self.checkOneExpression("1123.88977866", 1123.88977866)
	self.checkOneExpression("+2", 2)
	self.checkOneExpression("+258", 258)
	self.checkOneExpression("+258", 258)
	self.checkOneExpression("+258.45", 258.45)
	self.checkOneExpression("-3", -3)
	self.checkOneExpression("-3.0", -3.0)
	self.checkOneExpression("-3.78", -3.78)
}

func (self *CalcTests) checkOneExpression(str string, result float32) {
	calc := new(Calc)
	calc.Init(str)
	self.AssertFloat32(result, calc.ParseExpr())
	self.AssertInt(0, len(calc.errors))
}

func (self *CalcTests) AssertFloat32(expected float32, actual float32) {
	if expected != actual {
		panic(fmt.Sprintf("expected: %f actual: %f", expected, actual))
	}
}

func (self *CalcTests) AssertInt(expected int, actual int) {
	if expected != actual {
		panic(fmt.Sprintf("expected: %d actual: %d", expected, actual))
	}
}
