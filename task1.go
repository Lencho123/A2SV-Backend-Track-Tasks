package main

import (
 "bufio"
 "fmt"
 "os"
 "strconv"
 "strings"
)

// Function to calculate average from a slice of float64
func calculateAverage(grades []float64) float64 {
 var total float64
 for _, grade := range grades {
  total += grade
 }
 return total / float64(len(grades))
}

func main() {
 reader := bufio.NewReader(os.Stdin)

 // Ask for student name
 fmt.Print("Enter your name: ")
 studentName, _ := reader.ReadString('\n')
 studentName = strings.TrimSpace(studentName)

 // Ask for number of subjects
 fmt.Print("Enter the number of subjects you have taken: ")
 subjectCountStr, _ := reader.ReadString('\n')
 subjectCountStr = strings.TrimSpace(subjectCountStr)

 subjectCount, err := strconv.Atoi(subjectCountStr)
 if err != nil || subjectCount <= 0 {
  fmt.Println("Invalid number of subjects.")
  return
 }

 // Map to store subject name → grade
 subjectGrades := make(map[string]float64)
 var grades []float64

 // Loop to input subjects and grades
 for i := 1; i <= subjectCount; i++ {
  fmt.Printf("Enter the name of subject #%d: ", i)
  subject, _ := reader.ReadString('\n')
  subject = strings.TrimSpace(subject)

  fmt.Printf("Enter the grade for %s (0-100): ", subject)
  gradeStr, _ := reader.ReadString('\n')
  gradeStr = strings.TrimSpace(gradeStr)

  grade, err := strconv.ParseFloat(gradeStr, 64)

  if err != nil || grade < 0 || grade > 100 {

   fmt.Println("Invalid grade. Must be a number between 0 and 100.")

   i-- // repeat this subject
   continue
  }

  subjectGrades[subject] = grade
  grades = append(grades, grade)
 }

 // Calculate average
 average := calculateAverage(grades)

 // Display results
 fmt.Println("\n📄 Report Card")
 fmt.Printf("Student Name: %s\n", studentName)
 fmt.Println("Subjects and Grades:")
 for subject, grade := range subjectGrades {
  fmt.Printf("  - %s: %.2f\n", subject, grade)
 }
 fmt.Printf("Average Grade: %.2f\n", average)
}