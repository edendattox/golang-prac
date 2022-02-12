package main

// ex - 3

// type User struct {
// 	FirstName   string
// 	LastName    string
// 	PhoneNumber string
// 	Age         int
// 	BirthDate   time.Time
// }

// func main() {
// 	user := User{
// 		FirstName: "jackie",
// 		LastName:  "chan",
// 	}

// 	log.Println(user.FirstName, user.LastName)
// }

// ex-2

// func main() {
// 	var myString string
// 	myString = "Green"

// 	log.Println("myString is set to", myString)
// 	changeUsingPointer(&myString)
// 	log.Println("after fun call myString is set to", myString)

// }

// func changeUsingPointer(s *string) {
// 	log.Println("s is set to", s)
// 	newValue := "red"
// 	*s = newValue
// }

// ex - 1

// func main() {
// 	var whatToSay string
// 	var saySomethingElse string
// 	var i int

// 	whatToSay, _ = saySomething("Hello")
// 	saySomethingElse, _ = saySomething("Goodbye")
// 	i = 7

// 	log.Println(whatToSay)
// 	log.Println(saySomethingElse)
// 	log.Println(i)
// }

// func saySomething(s string) (string, string) {
// 	return s, "world"
// }
