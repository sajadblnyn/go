package main

import "fmt"

type Score struct {
	Score  float32
	Weight float32
}
type ScoresCalculator interface {
	CalculateAvgScores() float32
}
type Student struct {
	Id   int
	Name string
}
type SchoolStudent struct {
	Student
	Scores []float32
}
type UniversityStudent struct {
	Student
	Scores []Score
}

func (student SchoolStudent) CalculateAvgScores() float32 {
	var sum float32
	for _, v := range student.Scores {
		sum = sum + v
	}
	return sum / float32(len(student.Scores))
}

func (student *UniversityStudent) CalculateAvgScores() float32 {
	var sum float32
	for _, v := range student.Scores {
		sum = sum + (v.Score * v.Weight)
	}
	return sum / float32(len(student.Scores))
}

func main() {

	var student1 ScoresCalculator = SchoolStudent{Student: Student{Id: 1, Name: "mamad"}, Scores: []float32{12, 20, 17, 15.5}}

	var student2 ScoresCalculator = &UniversityStudent{Student: Student{Id: 1, Name: "mamad"}, Scores: []Score{Score{Score: 12, Weight: 2}, Score{Score: 15, Weight: 1}, Score{Score: 20, Weight: 3}}}

	// student3 := new(SchoolStudent)
	// student3.Id = 2

	fmt.Println(student1.CalculateAvgScores())

	fmt.Println(student2.CalculateAvgScores())

}
