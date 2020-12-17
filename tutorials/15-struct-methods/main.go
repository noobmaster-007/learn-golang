package main

import "fmt"

// Student - Student type
type Student struct {
	name   string
	grades []int
	age    int
}

// Method that be accessed by a student object
// and returning int
func (s Student) getAge() int {
	return s.age
}

// When need change the value for the accessable object
// A pointer need to be defined in accessible section
func (s *Student) setAge(age int) {
	s.age = age
}

// Usage field to do calculation
func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

// Get max grade
func (s Student) getMaxGrade() int {
	max := 0
	for _, v := range s.grades {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	s1 := Student{"YS", []int{80, 90, 88, 87}, 30}
	fmt.Println(s1)
	s1.setAge(31)
	fmt.Println(s1)

	fmt.Println(s1.getAverageGrade())
	fmt.Println(s1.getMaxGrade())
}
