# py4go
transpile python to golang

python code:
```python
def pi():
    return 3.1415
print(2 * pi())

x = 1
def print_x():
    print(x)
    if False: x = 0
print_x()

# Comment x function of x*x or not?
x = ( x + 1) * (x + 3) * 5**8
print(x)
```

transpiler Go code:
```go
package main

func pi() {
	return 3.1415
}
func print_x() {
	Print(x)
	if False {
		x = 0
	}
}
func main() {
	Print(2 * pi())
	x = 1
	print_x()
	x = (x + 1) * (x + 3) * math.Pow(5, 8)
	Print(x)
}
```
