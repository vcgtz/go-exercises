/*
Create a program that allows a teacher to input grades for multiple students
in different subjects.

Use structs to represent students and their respective grades. Then, use maps
to store each student's data and calculate their average grade.
*/
package main

import "fmt"

func main() {
	type student struct {
		name  string
		grade int
	}

	students := map[string]student{}

	for i := 0; i < 3; i++ {
		var s student

		fmt.Println("Enter the name of the student:", i+1)
		_, nameErr := fmt.Scanf("%s", &s.name)

		if nameErr != nil {
			fmt.Println("There was an error reading the student", i, ":", nameErr)
		}

		fmt.Println("Enter the grade of:", s.name)
		_, gradeErr := fmt.Scanf("%d", &s.grade)

		if gradeErr != nil {
			fmt.Println("There was an error reading the grade of", s.name, ":", gradeErr)
		}

		fmt.Println("Student added:", s.name)
		students[s.name] = s
	}

	average := 0

	for _, value := range students {
		average += value.grade
	}

	fmt.Println("Average:", average/len(students))
}
