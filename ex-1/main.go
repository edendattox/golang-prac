package main

// type Person struct {
// 	FirstName string `json:"first_name"`
// 	LastName string `json:"last_name"`
// 	HairColor string `json:"hair_color"`
// 	HasDog bool `json:"has_dog"
// }

// func main() {
// 	myJson := `
// 	[
// 		{
// 			"first_name": "John",
// 			"last_name": "Kent",
// 			"hair_color": "black",
// 			"has_dog": true
// 		},
// 		{
// 			"first_name": "Bruce",
// 			"last_name": "nai",
// 			"hair_color": "black",
// 			"has_dog": true
// 		},
// 	]`

// 	var unmarshalled []Person

// 	err := json.Unmarshall()

// }

// const numPool = 10

// func CalculateValue(intChan chan int) {
// 	RandomNumber := helpers.RandomNumber(numPool)
// 	intChan <- RandomNumber
// }

// func main() {
// 	intChan := make(chan int)
// 	defer close(intChan)

// 	go CalculateValue(intChan)

// 	num := <-intChan
// 	log.Println(num)
// }

// func main() {
// 	log.Println("Hello")

// 	var myVar helpers.SomeType
// 	myVar.TypeName = "Some name"
// 	log.Println(myVar.TypeName)

// }

//  ex - 10 interfaces

// type Animal interface {
// 	Says() string
// 	NumberOfLegs() int
// }

// type Dog struct {
// 	Name  string
// 	Breed string
// }

// type Gorilla struct {
// 	Name          string
// 	Color         string
// 	NumberOfTeeth int
// }

// func main() {

// 	dog := Dog{
// 		Name:  "Samson",
// 		Breed: "German Shepard",
// 	}

// 	// gorilla := Gorilla{
// 	// 	Name:          "King Kong",
// 	// 	Color:         "black",
// 	// 	NumberOfTeeth: "32",
// 	// }

// 	PrintInfo(dog)

// }

// func (d Dog) Says() string {
// 	return "woof"
// }

// func (d Dog) NumberOfLegs() int {
// 	return 4
// }

// func PrintInfo(a Animal) {
// 	log.Println("this animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
// }

// ex - 9 looping range

// func main() {

// myMap := make(map[string]string)
// myMap["dog"] = "dog"
// myMap["fish"] = "fish"
// myMap["hat"] = "hat"

// for i, x := range myMap {
// 	log.Println(i, x)
// }

// mySlice := []string{"dog", "cat", "fish", "toad", "sheep"}

// for _, x := range mySlice {
// 	log.Println(x)
// }

// for i := 0; i < 10; i++ {
// 	log.Println(i)
// }
// }

// ex - 8 switch

// func main() {

// 	myVar := "cat"

// 	switch myVar {
// 	case "cat":
// 		log.Println("cat is set to cat")

// 	case "dog":
// 		log.Println("cat is set to dog")

// 	case "fist":
// 		log.Println("cat is set to fish")

// 	default:
// 		log.Println("Cat is something else")

// 	}

// }

// ex - 7 if else

// func main() {

// myNum := 100
// isTrue := false

// if myNum > 99 && !isTrue {
// 	log.Println("MyNum is Greater than 99 and isTrue is set to true")
// }

// cat := "cat"

// if cat == "cat" {
// 	log.Println("Cat is cat")
// } else {
// 	log.Println("Cat is not cat")
// }

// var isTrue bool
// isTrue = true

// if isTrue {
// 	log.Println("isTrue is", isTrue)
// } else {
// 	log.Println("isFalse is", isTrue)
// }

// }

// ex - 6 slices

// func main() {

// names := []string{"one", "two", "three"}

// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

// log.Println(len(names))
// log.Println(numbers)
// log.Println(numbers[0:2])

// var mySlice []string

// mySlice = append(mySlice, "Trevor", "Rahul", "Mary")

// var a int

// a = len(mySlice) - 1

// log.Println(mySlice[1])
// log.Println(mySlice[a])

// }

// ex - 5 type of variables and map

// type User struct {
// 	FirstName string
// 	LastName  string
// }

// func main() {
// 	myMap := make(map[string]User)

// 	me := User{
// 		FirstName: "Trevor",
// 		LastName:  "Sawler",
// 	}

// 	myMap["me"] = me

// 	log.Println(myMap["me"].FirstName)

// 	//	var myOtherMap map[string]string

// 	// myMap := make(map[string]string)

// 	// myMap["dog"] = "Samson"

// 	// log.Println(myMap["dog"])
// 	// log.Println(myMap)

// }

// ex - 4

// type myStruct struct {
// 	FirstName string
// }

// func (m *myStruct) printFirstName() string {
// 	return m.FirstName
// }

// func main() {
// 	var myVar myStruct
// 	myVar.FirstName = "John"
// 	myVar2 := myStruct{
// 		FirstName: "Mary",
// 	}

// 	log.Println("myVar is set to", myVar.printFirstName())
// 	log.Println("myVar is set to", myVar2.printFirstName())
// }

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
