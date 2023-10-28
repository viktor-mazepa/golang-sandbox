package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	myString := "This is clear EXAMPLE of why we search in one case only."

	searchString := strings.ToLower(myString)

	if strings.Contains(searchString, "this") {
		fmt.Println("Found it")
	}

	//other case functions
	fmt.Println(strings.ToLower(myString))
	fmt.Println(strings.ToUpper(myString))

	/*myString = strings.Title(myString)
	tempString := "EXAMPLE"
	if strings.Contains(myString, tempString) {
		myString = strings.ReplaceAll(myString, tempString, "Example")
	}*/

	fmt.Println(cases.Title(language.Und).String(strings.ToLower(myString)))

	myString = "Goodbye, cruel world!"

	fmt.Println(len(myString))
	fmt.Println(myString[3:5])
}

func stringManipulation() {
	newString := "Go is a great programming language. Go for it!"
	if strings.Contains(newString, "Go") {
		newString = strings.Replace(newString, "Go", "Golang", 1)
	}

	fmt.Println(newString)

	//string comparison
	if "Alpha" > "Absolute" {
		fmt.Println("A is greater than B")
	} else {
		fmt.Println("A is not greater than B")
	}

	badEmail := " me@here.com "
	goodEmail := strings.TrimSpace(badEmail)

	fmt.Printf("=%s=\n", goodEmail)

	str := "alpha alpha alpha alpha alpha alpha"

	str = replaceNth(str, "alpha", "beta", 3)

	fmt.Println(str)
}

func replaceNth(s, old, new string, n int) string {
	i := 0

	for j := 1; j <= n; j++ {
		x := strings.Index(s[i:], old)
		if x < 0 {
			break
		}
		i = i + x
		if j == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}

func stringsPackage() {
	cources := []string{
		"Learn GO for Beginners Crash Cource",
		"Learn Java for Beginners Crash Cource",
		"Learn Python for Beginners Crash Cource",
		"Learn C for Beginners Crash Cource",
	}

	for _, x := range cources {
		if strings.Contains(x, "GO") {
			fmt.Println("Go is found in:'", x, "'and index is", strings.Index(x, "GO"))
		}
	}

	newString := "Go is a great programming language. Go for it!"

	fmt.Println(strings.HasPrefix(newString, "Go"))
	fmt.Println(strings.HasPrefix(newString, "Python"))
	fmt.Println(strings.HasSuffix(newString, "!"))
	fmt.Println(strings.Count(newString, "Go"))

	fmt.Println(strings.Index(newString, "Go"))
	fmt.Println(strings.Index(newString, "Python"))
	fmt.Println(strings.LastIndex(newString, "Go"))
}
func stringLens() {
	courseName := "Learn GO for Beginners Crash Cource"

	for i := 0; i < 21; i++ {
		fmt.Print(string(courseName[i]))
	}
	fmt.Println()
	for i := 13; i < 21; i++ {
		fmt.Print(string(courseName[i]))
	}
	fmt.Println()

	fmt.Println("Length of course name is:", len(courseName))

	var mySlice []string

	mySlice = append(mySlice, "one")
	mySlice = append(mySlice, "two")
	mySlice = append(mySlice, "three")

	fmt.Println("Length of mySlace is:", len(mySlice), "elements")
	fmt.Println("The last element in mySlice is:", mySlice[len(mySlice)-1])
}
func simpleStringsOperations() {
	fmt.Println()
	name := "Hello world"
	fmt.Println("String", name)
	fmt.Println()

	fmt.Println("Bytes")

	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}
	fmt.Println()

	fmt.Println("Index\tRune\tString")

	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}
	fmt.Println()

	nameUA := "Привіт Світ"
	fmt.Println("Index\tRune\tString")

	for x, y := range nameUA {
		fmt.Println(x, "\t", y, "\t", string(y))
	}

	fmt.Println()
	fmt.Println("Three ways to concatenate strings:")

	h := "Hello "
	w := "world."

	// using + not very efficient
	myString := h + w
	fmt.Println(myString)

	//using fmt - more efficient
	myString = fmt.Sprintf("%s%s", h, w)
	fmt.Println(myString)

	// using stringbuilder - very efficient
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println(sb.String())

	fmt.Println()
	name = "ABCDEFGHJIKLMNOPQRSTUVXWYZ"
	fmt.Println("Getting a substring")
	fmt.Println(name[10:13])
}
