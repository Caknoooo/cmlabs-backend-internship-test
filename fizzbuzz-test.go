package main

func test(res int) {
	for i := 1; i <= res; i++ {
		if i % 15 == 0 {
			println("FizzBuzz")
		} else if i % 3 == 0 {
			println("Fizz")
		} else if i % 5 == 0 {
			println("Buzz")
		}
	}
}

func main() {
	test(100)
}