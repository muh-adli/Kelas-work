// part 1
// Pengenalan function

package main

import (
	"errors"
	"fmt"
)

func main() {
	var variablename1 string = "Hello, world!"
	variablename2 := "Hello world!"

	fmt.Println(variablename1)
	fmt.Println(variablename2)

	// Primitive: Boolean, Integer, Float, String

	// Boolean
	var boolVar = true

	fmt.Printf("Type: %T Value: %v\n", boolVar, boolVar)

	// Integer
	intVar := int(5)
	intVar1 := int32(6)
	intVar2 := int64(7)

	fmt.Printf("Type: %T Value: %v\n", intVar, intVar)
	fmt.Printf("Type: %T Value: %v\n", intVar1, intVar1)
	fmt.Printf("Type: %T Value: %v\n", intVar2, intVar2)

	// Float
	floatVar := float32(5.5)
	floatVar1 := float64(6.6)

	fmt.Printf("Type: %T Value: %v\n", floatVar, floatVar)
	fmt.Printf("Type: %T Value: %v\n", floatVar1, floatVar1)

	// Bytes
	BytesVar := byte(255)

	fmt.Printf("Type: %T Value: %v\n", BytesVar, BytesVar)
	BytesVarKarakter := []byte("Hello world")

	fmt.Printf("Type: %T Value: %v\n", BytesVarKarakter, BytesVarKarakter)

	// rune
	runeVar := '♄'

	fmt.Printf("Type: %T Value: %v\n", runeVar, runeVar)

	// Complex
	ComplexVar := -7 + 3i

	fmt.Printf("Type: %T Value: %v\n", ComplexVar, ComplexVar)

	// Error
	errVar := errors.New("Error detected")

	fmt.Printf("Type: %T Value: %v\n", errVar, errVar)

	// Interface
	var myInferfaceVar interface{}

	myInferfaceVar = 5
	fmt.Printf("Type: %T Value: %v\n", myInferfaceVar, myInferfaceVar)
	myInferfaceVar = "hello world"
	fmt.Printf("Type: %T Value: %v\n", myInferfaceVar, myInferfaceVar)
}

type MethodList interface {
	myfunctions()
	myfunctions2(int) int
}
