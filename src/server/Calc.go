package main

import (
  "regexp"
  "strconv"
  "strings"
  "math"
  "math/rand"
)

type Calc struct {
  maxHintLength int
  atomRegex *regexp.Regexp
  funcRegex *regexp.Regexp
  functions []string
  inputStr string
  hintOnError string
  stackIndex int
  errors []string
}

func (self *Calc)  Init(str string) {
  self.maxHintLength = 7;

  self.atomRegex = regexp.MustCompile("^([0-9]+\\.[0-9]+|[0-9]+\\.|[0-9]+|\\.[0-9]+)")
  self.funcRegex = regexp.MustCompile("^[A-z]+")

  self.functions = append(self.functions, "rand")
  self.functions = append(self.functions, "sqrt")
  self.functions = append(self.functions, "sin")
  self.functions = append(self.functions, "cos")
  self.functions = append(self.functions, "tg")
  self.functions = append(self.functions, "ctg")
  self.functions = append(self.functions, "arcsin")
  self.functions = append(self.functions, "arccos")
  self.functions = append(self.functions, "arctg")
  self.functions = append(self.functions, "arcctg")

  self.inputStr = str
  self.hintOnError = ""
  self.stackIndex = 0

  self.PrepareInputString()
  self.CheckSomeGreatErrors()
}

func (self *Calc) ParseExpr() float32 {
  WriteToLogStr("ParseExpr with " + self.inputStr)
  if len(self.errors) > 0 {
    return 0
  }
  if self.inputStr == "" {
    return 0
  }

  return self.ParseSum()
}

func (self *Calc) PrepareInputString() {
  spacesRegex := regexp.MustCompile("[ \\t\\r\\n]+")
  self.inputStr = spacesRegex.ReplaceAllLiteralString(self.inputStr, "")
}

func (self *Calc) CheckSomeGreatErrors() {
  greatErrorsRegex := regexp.MustCompile("[\\+\\-\\/\\*]{2,}")
  if greatErrorsRegex.MatchString(self.inputStr) {
    errorBeginArr := greatErrorsRegex.FindStringIndex(self.inputStr)
    errorBegin := errorBeginArr[0]
    remains := self.inputStr[errorBegin:len(self.inputStr)]
    var hintLength int
    if len(remains) >= self.maxHintLength {
      hintLength = self.maxHintLength
    } else {
      hintLength = len(remains)
    }
    self.hintOnError = remains[0:hintLength]
    self.errors = append(self.errors, "Syntax error near " + self.hintOnError)
  }
}

func (self *Calc) RemoveHead() {
  if len(self.inputStr) > 0 {
    self.inputStr = self.inputStr[1:]
  }
}

func (self *Calc) IsHeadSymbolExpected() bool {
  return (self.inputStr[0] == '+') ||
          (self.inputStr[0] == '-') ||
          (self.inputStr[0] == '*') ||
          (self.inputStr[0] == '/') ||
          (self.inputStr[0] == ')') ||
          (self.inputStr[0] == ',')
}

func (self *Calc) RefreshHintOnError() {
  var hintLength int
  if len(self.inputStr) >= self.maxHintLength {
    hintLength = self.maxHintLength
  } else {
    hintLength = len(self.inputStr)
  }
  self.hintOnError = self.inputStr[0:hintLength]
}

func (self *Calc) ParseSum() float32 {
  WriteToLogStr("ParseSum with " + self.inputStr)
  if len(self.errors) > 0 {
    return 0
  }

  var left float32
  left = self.ParseMul()
  for len(self.inputStr) > 0 {
    if self.inputStr[0] == '+' {
      self.RemoveHead()
      mulres := self.ParseMul()
      left += mulres
    } else if self.inputStr[0] == '-' {
      self.RemoveHead()
      mulres := self.ParseMul()
      left -= mulres
    } else {
      return left
    }
  }

  return left
}

func (self *Calc) ParseMul() float32 {
  WriteToLogStr("ParseMul with " + self.inputStr)
  if len(self.errors) > 0 {
    return 0
  }

  var left float32
  left = self.ParseUnary()
  for {
    if (len(self.inputStr) > 0) && (self.inputStr[0] == '*') {
      self.RemoveHead()
      left *= self.ParseUnary()
    } else if (len(self.inputStr) > 0) && (self.inputStr[0] == '/') {
      self.RemoveHead()
      left /= self.ParseUnary()
    } else {
      return left
    }
  }
}

func (self *Calc) ParseUnary() float32 {
  WriteToLogStr("ParseUnary with " + self.inputStr)
  if len(self.errors) > 0 {
    return 0
  }
  if len(self.inputStr) == 0 {
    self.errors = append(self.errors, "Unexpected end of string")
    return 0
  }

  if self.inputStr[0] == '+' {
    self.RemoveHead()
    return self.ParseBrackets()
  } else if self.inputStr[0] == '-' {
    self.RemoveHead()
    return -self.ParseBrackets()
  } else {
    return self.ParseBrackets()
  }
}

func (self *Calc) ParseBrackets() float32 {
  WriteToLogStr("ParseBrackets with " + self.inputStr)
  WriteArrayToLogStr(self.errors)
  if len(self.errors) > 0 {
    return 0
  }

  if self.inputStr[0] == '(' {
    self.stackIndex++
    self.RefreshHintOnError()
    self.RemoveHead()
    if (len(self.inputStr) > 0) && (self.inputStr[0] == ')') {
      self.errors = append(self.errors, "Empty brackets near " + self.hintOnError)
    }
    var result float32
    result = self.ParseSum()
    if (len(self.inputStr) > 0) && (self.inputStr[0] == ')') {
      self.RemoveHead()
      self.stackIndex++
      if (self.stackIndex == 0) && (len(self.inputStr) > 0) && (self.inputStr[0] == ')') {
        self.errors = append(self.errors, "Closing brackets mismatch begins from " + self.hintOnError)
      }
    } else {
      self.errors = append(self.errors, "Opening brackets mismatch begins from " + self.hintOnError)
    }

    return result
  }

  return self.ParseFunc()
}

func (self *Calc) ParseFunc() float32 {
  WriteToLogStr("ParseFunc with " + self.inputStr)
  WriteArrayToLogStr(self.errors)
  if len(self.errors) > 0 {
    return 0
  }

  funcName := ""
  isFunctonFound := self.funcRegex.MatchString(self.inputStr)
  self.RefreshHintOnError()
  if isFunctonFound {
    funcName = self.funcRegex.FindString(self.inputStr)
    isFunctonFound = ContainsStr(self.functions, funcName)
    if isFunctonFound {
      self.inputStr = self.inputStr[len(funcName):]
      if (len(self.inputStr) == 0) {
        self.errors = append(self.errors, "Unexpected end of string")
        return 0
      }
      if self.inputStr[0] == '(' {
        if funcName == "rand" {
          return self.CalculateRand()
        } else {
          var bracketsContent float32
          bracketsContent = self.ParseBrackets()

          return self.CalculateFunction(funcName, bracketsContent)
        }
      } else {
        self.errors = append(self.errors, "Unknown identifier " + funcName)
      }
    } else {
      self.errors = append(self.errors, "Unknown identifier " + funcName)
    }
  }

  return self.ParseAtom()
}

func (self *Calc) CalculateRand() float32 {
  WriteToLogStr("CalculateRand with " + self.inputStr)
  if len(self.errors) > 0 {
    return 0
  }

  var randArgStr string
  bracketsCounter := 0
  randBracketsLength := 0
  self.stackIndex++

  for len(self.errors) == 0 {
    if self.inputStr[randBracketsLength] == '(' {
      bracketsCounter++
    } else if (self.inputStr[randBracketsLength] == ')') {
      bracketsCounter--
    }
    randBracketsLength++
    if bracketsCounter == 0 {
      break
    } else if randBracketsLength == len(self.inputStr) {
      self.errors = append(self.errors, "Opening brackets mismatch begins from " + self.hintOnError)
    }
  }

  if (randBracketsLength < 2) {
    if len(self.errors) == 0 {
      self.errors = append(self.errors, "Opening brackets mismatch begins from " + self.hintOnError)
    }
    
    return 0
  }

  randArgStr = self.inputStr[1:randBracketsLength - 1]
  self.RemoveHead()
  var argv []string
  argv = strings.Split(randArgStr, ",")

  if len(argv) > 2 {
    self.errors = append(self.errors, "Too many arguments for rand() near " + self.hintOnError)
    return 0
  } else if len(argv) < 2 {
    self.errors = append(self.errors, "Too few arguments for rand() near " + self.hintOnError)
    return 0
  }

  var argvFloat32 []float32
  i := 0
  for i < len(argv) {
    argvFloat32 = append(argvFloat32, self.ParseSum())
    self.RemoveHead()
    i++
  }

  if argvFloat32[0] >= argvFloat32[1] {
    self.errors = append(self.errors, "Invalid arguments to rand() near " + self.hintOnError)
  }

  self.stackIndex--

  return argvFloat32[0] + rand.Float32() * (argvFloat32[1] - argvFloat32[0])
}

func (self *Calc) CalculateFunction(funcName string, arg32 float32) float32 {
  if len(self.errors) > 0 {
    return 0
  }

  var result float64
  arg := float64(arg32)
  switch funcName {
  case "sin":
    result = math.Sin(self.DegToRad(arg))
  case "cos":
    result = math.Cos(self.DegToRad(arg))
  case "tg":
    result = math.Tan(self.DegToRad(arg))
  case "ctg":
    result = 1.0 / math.Tan(self.DegToRad(arg))
  case "arcsin":
    result = self.RadToDeg(math.Asin(arg))
  case "arccos":
    result = self.RadToDeg(math.Acos(arg))
  case "arctg":
    result = self.RadToDeg(math.Atan(arg))
  case "arcctg":
    result = self.RadToDeg(math.Atan(1 / arg))
  case "sqrt":
    result = math.Sqrt(arg)
  default:
    self.errors = append(self.errors, "Unknown identifier " + funcName)
    return 0
  }

  return float32(result)
}

func (self *Calc) ParseAtom() float32 {
  WriteToLogStr("ParseAtom with " + self.inputStr)
  WriteArrayToLogStr(self.errors)
  if len(self.errors) > 0 {
    return 0
  }

  self.RefreshHintOnError()
  isAtomFound := self.atomRegex.MatchString(self.inputStr)
  if isAtomFound {
    var atom string
    atom = self.atomRegex.FindString(self.inputStr)
    self.inputStr = self.inputStr[len(atom):]
    if len(self.inputStr) > 0 {
      if !self.IsHeadSymbolExpected() {
        self.errors = append(self.errors, "Error near " + self.hintOnError)
      } else if (self.stackIndex == 0) && (self.inputStr[0] == ')') {
        self.errors = append(self.errors, "Closing brackets mismatch begins from " + self.hintOnError)
      }
    }

    parsedFloat64, err := strconv.ParseFloat(atom, 32)
    parsedFloat32 := float32(parsedFloat64)
    if err != nil {
      self.errors = append(self.errors, "Parsing problem near " + self.hintOnError)

      return 0
    }
    WriteToLogStr("ParseAtom: got")
    WriteToLogFloat32(parsedFloat32)

    return parsedFloat32
  } else {
    self.errors = append(self.errors, "Error near " + self.hintOnError)

    return 0
  }

  return 1 // Something went totally wrong
}

func (self *Calc) DegToRad(deg float64) float64 {
  return deg * math.Pi / 180
}

func (self *Calc) RadToDeg(rad float64) float64 {
  return rad * 180 / math.Pi
}
