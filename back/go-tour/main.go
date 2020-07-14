package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Animal interface {
	sayHello() string
}

type Cow struct {
	hello string
}

type Dog struct {
	hello string
}

type Cat struct {
	hello string
}

func (c Cow) sayHello() string {
	return "muuu"
}

func (d *Dog) sayHello() string {
	return "gav"
}

func (ct *Cat) sayHello() string {
	return "miu"

}

type human struct {
	Name string
	Age  int
}

type Human struct {
	name string
	age  int
}

var c int = -999

// k := 3 this not working

type IPAddr [4]byte

func main() {
	channelTest()
}

func channelTest() {
	ch1 := make(chan int)
	go pushChannel(ch1)
	for {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}

}

func pushChannel(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func toStringChecker() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func (a IPAddr) String() string {
	var str []string
	for _, v := range a {
		intAsString := strconv.Itoa(int(v))
		str = append(str, intAsString)
	}

	return strings.Join(str, ".")
}

func checkRuntimeInterfaceType() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f, ok = i.(float64) // panic
	fmt.Println(f, ok)
}

func callWithNilInterface() {
	var animal Animal
	fmt.Printf("(%v, %T)\n", animal, animal)
	if animal == nil {
		fmt.Println("is nil")
	}
}

func interfaceChecker() {
	var animal Animal = &Cat{}
	fmt.Println(animal.sayHello())

	animal = &Dog{}
	fmt.Println(animal.sayHello())

	animal = &Cow{}
	fmt.Println(animal.sayHello())

	dog := Dog{}
	fmt.Println(dog.sayHello())
}

func (h *Human) getAge() int {
	h.age += 100
	return h.age
}

func (h Human) testTypeMethod() string {
	return h.name
}

/*
* Пройденное по туру
 */
func runned() {
	//f := fibonacci()
	//for i := 0; i < 10; i++ {
	//		fmt.Println(f())
	//	}
	//fmt.Println(funcAsValue(funcAsValueTest, 97))
	//fmt.Println(wordCount("I am learning Go!"))
	//testMap()
	//pic.Show(Pic3)
	//pic.Show(Pic2)
	//pic.Show(Pic1)
	// iterDinamicArrayWithRange()
	//testMake4DinamicArray()
	//testArray()
	// testHumanCreate()
	// fmt.Println("My favorite number is", testType(1, 2))
	// a, b := returnManyType(-100, 2)
	// fmt.Println("return many type", a, b)
	// testGlobalValue()
	// baseType()
	// testDifferentTypeDivide()
	// forLoopEachTest()default:
	// testWhileLoop()
	// //initInifinityLoop()
	// testManyInitInIfInstructions(3)
	// testManyInitInIfInstructions(-3)
	// testSwitch("test1")
	// testDefer()
	// testPoint()
}

func fibonacci() func() int {
	prev, next := 0, 1
	return func() int {
		res := prev
		prev, next = next, prev+next
		return res
	}
}

func funcAsValue(fn func(s string) string, a int) string {
	ns := fmt.Sprintf("test funcAsValueTest %d", a)
	return fn(ns)
}

func funcAsValueTest(s string) string {
	return s
}

func wordCount(s string) map[string]int {
	var delimStr []string = strings.Fields(s)

	m := make(map[string]int)
	for _, v := range delimStr {
		_, ok := m[v]
		if !ok {
			m[v] = 0
		}
		m[v]++
	}

	return m
}

func testMap() {
	m := make(map[string]human)
	m["one"] = human{Name: "one", Age: 10}

	_, exist := m["one"]
	fmt.Printf("is exist %v", exist)
}

/**
func pic3(dx, dy int) [][]uint8 {
	outer := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		inner := make([]uint8, dx)
		for j := range inner {
			inner[j] = uint8(i ^ j)
		}

		outer[i] = inner
	}
	return outer
}

func pic2(dx, dy int) [][]uint8 {
	outer := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		inner := make([]uint8, dx)
		for j := range inner {
			inner[j] = uint8((i * j) / 2)
		}

		outer[i] = inner
	}
	return outer
}

func Pic1(dx, dy int) [][]uint8 {
	outer := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		inner := make([]uint8, dx)
		for j := range inner {
			inner[j] = uint8(i * j)
		}

		outer[i] = inner
	}
	return outer
}
*/

func iterDinamicArrayWithRange() {
	var s []string = make([]string, 0)
	s = append(s, "first")
	s = append(s, "second")

	fmt.Println(s)

	for i, v := range s {
		fmt.Printf("index = %d, value =  %s\n", i, v)
	}

	for _, v := range s {
		fmt.Printf("value %s\n", v)
	}

	for i := range s {
		fmt.Printf("Just index %d\n", i)
	}
}

func testMake4DinamicArray() {
	a := make([]int, 5, 6)
	fmt.Println(a, len(a), cap(a))
	a = append(a, -100)
	a = append(a, -101)
	fmt.Println(a, len(a), cap(a))

}

func testArray() {
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println(a)
	a[2] = -100
	fmt.Println(a)

	var b []int = a[1:3]
	fmt.Println(b)
}

func testHumanCreate() {
	h := human{"this is name", 11}
	fmt.Printf("name = \"%s\" and age %v \n", h.Name, h.Age)
	fmt.Printf("human %v \n", h)

	p := &h
	p.Age = 1 * 1e011
	fmt.Printf("name = \"%s\" and age %v \n", h.Name, h.Age)

}

func testPoint() {
	var p *int
	i := 1
	p = &i
	println(*p, i)
	*p = 5
	println(*p, i)
}

func testDefer() {
	defer println("this defer1")
	defer println("this defer2")
	println("test defer start")

}

func testSwitch(switchValue string) {
	switch {
	case "test" == switchValue:
		switchValue = "testvalue1"
		switchValue = "testvalue2"
	default:
		switchValue = "default"
	}

	fmt.Println(switchValue)
}

func startNT() {
	count := 10
	for i := 1; i <= count; {
		s := math.Sqrt(float64(i))
		fmt.Printf("calc value %f ; sqrt %f\n", nt(float64(i)), s)
		i++
	}
}

func nt(x float64) float64 {
	z := 4.0
	res := 0.0

	for i := 0; i < 10; i++ {
		res = z - ((z*z)-x)/(2*z)
	}

	return res
}

func testManyInitInIfInstructions(c int) {
	if a := 1; a < c {
		fmt.Println(a, c)
	} else {
		fmt.Println(c, a)
	}

}

func initInifinityLoop() {
	for {
		fmt.Println("INFINITY")
	}
}

func testWhileLoop() {
	n := 100
	i := 1
	for i < n {
		fmt.Println("While loop ", i)
		i++
	}
}

func forLoopEachTest() {
	for i := 1; i <= 100; i++ {
		fmt.Println("This is ", i)
	}

}

func baseType() {
	isBool := true
	var isBoolTo bool = true
	var a uint32 = 100
	var b rune = 1
	var complex complex64 = complex(1, 2)

	fmt.Println(isBool, isBoolTo, a, b, complex)
}

func testDifferentTypeDivide() {
	a := 1
	b := 2.0
	res := float64(a) / b
	fmt.Println(res)
}

func testGlobalValue() {
	fmt.Println(c)
}

func returnManyType(a, b int) (x, y int) {
	x = a
	y = b
	return
}

func testType(x int, y int) int {
	a := x + y
	return a
}
