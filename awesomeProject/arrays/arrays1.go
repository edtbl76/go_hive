package main

import "fmt"

func main() {
	// Initializing an array one index at a time
	var playerScores [4]int
	playerScores[0] = 43
	playerScores[1] = 7
	playerScores[2] = 32
	playerScores[3] = 65
	fmt.Println(playerScores)

	// creating an empty array
	var playerNames [5]string
	fmt.Println(playerNames)

	// creating an array w/ some initial data
	myArray := [5]int{5, 48, 32, 1, 6}
	fmt.Println(myArray)

	// accessing array values w/ indices
	triangleAngles := [3]int{30, 60, 90}
	fmt.Println(triangleAngles[2])
	sum := triangleAngles[0] + triangleAngles[1] + triangleAngles[2]
	fmt.Println(sum)

	// mutating data by accessing it via index
	myDogs := [3]string{"Frida", "Fedo", "Jegf"}
	fmt.Println(myDogs)

	myDogs[1] = "Fido"
	myDogs[2] = "Jeff"
	fmt.Println(myDogs)

	// Arrays and Slices
	myTutors := [4]string{"Kirsty", "Mishell", "Jose", "Neil"}
	fmt.Println(myTutors)

	// slice of entire array
	tutorSlice := myTutors[:]

	// new slice
	subjects := []string{"Go", "Java", "Python"}
	fmt.Println(tutorSlice)
	fmt.Println(subjects)

	// len() and cap()
	myTutors2 := []string{"Kirsty", "Mishell", "Jose", "Neil"}
	fmt.Println(myTutors2)
	fmt.Println(len(myTutors2))
	fmt.Println(cap(myTutors2))

	// append
	myTutors3 := []string{"Kirsty", "Mishell", "Jose", "Neil"}
	fmt.Println(myTutors3)
	myTutors3 = append(myTutors3, "Josh")
	fmt.Println(myTutors3)

	// arrays/slice in function
	seahawks := []string{"Tyler Lockett", "D.K. Metcalf", "Pete Carroll", "Russell Wilson"}
	changeLastElement(seahawks, "Geno Smith")

}

func changeLastElement(slice []string, str string) {
	slice[len(slice)-1] = str
	fmt.Println(slice)
}
