package examples

import (
	"fmt"
	"semicolonGo/mylambdas"
)

var k string = "test"

func average(xs []float64) float64 {
	panic("Not Implemented")
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	evenNumbers := mylambdas.Filter(numbers, func(num int) bool {
		return num%2 == 0
	})

	function := func(num int) int {
		return num + 10
	}
	plusTen := mylambdas.Map(numbers, function)

	fmt.Println(evenNumbers)
	fmt.Println(plusTen)

	/*
		fmt.Println("Hello World")
		fmt.Println("length", len("Hello World"))
		fmt.Println("Hello World"[1])
		fmt.Println(numbers.Subtract(5, 4))
		fmt.Println(numbers.Add(5, 4, "a"))
		fmt.Println(true && true)
		fmt.Println(true && false)
		fmt.Println(false && false)
		fmt.Println(true || true)
		fmt.Println(true || false)
		fmt.Println(false || false)
		fmt.Println(false || true)
		fmt.Println(!true)

		fmt.Println("a")
		dogsName := "Max"
		fmt.Println("My dog's name is", dogsName)
		f()
		const m = false
		fmt.Println(m)
		//fmt.Print("Enter a number: ")
		//var input float64
		//fmt.Scanf("%f", &input)
		//output := input * 2
		//fmt.Println(output)

		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println(i, "even")
			} else {
				fmt.Println(i, "odd")
			}
		}
		x := make(map[string]int)
		x["key"] = 10
		fmt.Println(x["key"])
		fmt.Println(x["a"])

		//elements := make(map[string]string)
		//elements["H"] = "Hydrogen"
		//elements["He"] = "Helium"
		//elements["Li"] = "Lithium"
		//elements["Be"] = "Beryllium"
		//elements["B"] = "Boron"
		//elements["C"] = "Carbon"
		//elements["N"] = "Nitrogen"
		//elements["O"] = "Oxygen"
		//elements["F"] = "Fluorine"
		//elements["Ne"] = "Neon"
		//fmt.Println(elements["Li"])
		//fmt.Println(elements["Un"])
		//name, ok := elements["Un"]
		//fmt.Println(name, ok)

		elements := map[string]map[string]string{
			"H": {
				"name":  "Hydrogen",
				"state": "gas",
			},
			"He": map[string]string{
				"name":  "Helium",
				"state": "gas",
			},
			"Li": {
				"name":  "Lithium",
				"state": "solid",
			},
			"Be": map[string]string{
				"name":  "Beryllium",
				"state": "solid",
			},
			"B": map[string]string{
				"name":  "Boron",
				"state": "solid",
			},
			"C": map[string]string{
				"name":  "Carbon",
				"state": "solid",
			},
			"N": map[string]string{
				"name":  "Nitrogen",
				"state": "gas",
			},
			"O": map[string]string{
				"name":  "Oxygen",
				"state": "gas",
			},
			"F": map[string]string{
				"name":  "Fluorine",
				"state": "gas",
			},
			"Ne": map[string]string{
				"name":  "Neon",
				"state": "gas",
			},
		}

		if elem, ok := elements["Li"]; ok {
			fmt.Println(elem["name"], elem["state"])
		}

		o, p := d()
		fmt.Println(o, p)

		l := 0
		increment := func() int {
			l++
			return l
		}
		fmt.Println(increment())
		fmt.Println(increment())

		nextEven := makeEvenGenerator()
		fmt.Println(nextEven()) // 0
		fmt.Println(nextEven()) // 2
		fmt.Println(nextEven()) // 4

		fmt.Println(factorial(5))
		first()
		second()
		defer second()
		first()
		defer func() {
			str := recover()
			fmt.Println(str)
		}()
		average([]float64{1, 2})

	*/

}

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

func f() {
	fmt.Println(k)
}

func d() (int, int) {
	return 5, 6
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	fmt.Println("first call", i)
	return func() (ret uint) {
		ret = i
		i += 2
		fmt.Println("second call", i)
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
