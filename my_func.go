package go_my_func

func GetHello(name string) string {
	return "Hello " + name
}

func GetFullName() (string, string) {
	return "Azka", "Faridi"
}

func GetCompleteName() (firstName, middleName, lastName string) {
	firstName = "Muhammad"
	middleName = "Azka"
	lastName = "Faridi"

	return firstName, middleName, lastName
}

func SumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func GetGoodBye(name string) string {
	return "Good Bye " + name
}