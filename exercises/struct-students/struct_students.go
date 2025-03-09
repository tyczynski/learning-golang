package main

import (
	"fmt"
)

type Student struct {
	ID               int
	Name             string
	Grades           []int
	EnrollmentStatus bool
}

var students []Student

func addStudent(student Student) {
	students = append(students, student)
}

func displayStudents(onlyEnrollment bool) {
	fmt.Println("--- Students ---")
	studentsToDisplay := students[:0]

	if onlyEnrollment {
		for _, student := range students {
			if student.EnrollmentStatus {
				studentsToDisplay = append(studentsToDisplay, student)
			}
		}
	} else {
		studentsToDisplay = students
	}

	for _, student := range studentsToDisplay {
		fmt.Printf("Student %s. ID: %d, Grades: %v, Enrollment: %t\n",
			student.Name, student.ID, student.Grades, student.EnrollmentStatus)
	}
}

func findStudent(ID int) *Student {
	// It does not work because student is a copy
	// for _, student := range students {
	// 	if student.ID == ID {
	// 		return &student
	// 	}
	// }

	for i := range students {
		if students[i].ID == ID {
			return &students[i]
		}
	}

	return nil
}

func updateStudent(ID int, newData map[string]any) {
	student := findStudent(ID)
	if student == nil {
		fmt.Printf("Student with ID %d not found\n", ID)
	}

	for key, value := range newData {
		switch key {
		case "Name":
			if name, ok := value.(string); ok {
				student.Name = name
			}
		case "Grades":
			if grades, ok := value.([]int); ok {
				student.Grades = grades
			}
		case "EnrollmentStatus":
			if enrollmentStatus, ok := value.(bool); ok {
				student.EnrollmentStatus = enrollmentStatus
			}
		}
	}
}

func main() {
	addStudent(Student{ID: 1, Name: "Jan Kowalski", Grades: []int{1, 2, 3}, EnrollmentStatus: true})
	addStudent(Student{ID: 2, Name: "Paweł Jumper", Grades: []int{5, 5, 5}, EnrollmentStatus: false})
	displayStudents(false)

	updateStudent(1, map[string]any{"Name": "Jan Nowak"})
	displayStudents(true)
	displayStudents(false)
}
