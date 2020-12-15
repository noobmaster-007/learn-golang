# Variables & Data Types
This section will cover go lang's variable and data types

## Variables
### There are many ways to declare a variable
#### Declare variable without initialize
```
var name string
var value int
var flag bool
```
#### Declare variable with initialize
```
var name string = "hello"
var value int = 12
var flag bool = true
```
#### Short syntax to declare variable
Compiler will select the correct data type automatically.
```
name := "hello"
```

## Data Types
### Integer
#### Unsigned integer (non negative value)
* uint8 / byte (0-255) (0 - 2^8 - 1)
* uint16 (0-65535)
* uint32
* uint64
#### Signed integer
* int8 (-128 to 127)
* int16
* int32
* int64
#### Machine Dependent types
* uint (32 or 64 depends on the machine)
* int
* uintptr

### Floating Number
#### Floats
* float32 (32 bits floating number)
* float64
#### Complex
* complex64
* complex128

## Strings
* "Hello World"

## Booleans
* true
* false