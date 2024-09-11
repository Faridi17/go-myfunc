package go_my_func

func getHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) {
	return "Azka", "Faridi"
}

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Muhammad"
	middleName = "Azka"
	lastName = "Faridi"

	return firstName, middleName, lastName
}

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}