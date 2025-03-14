package main

import "fmt"

import "errors"
import "sync"

var w sync.WaitGroup

 var empty error = 	errors.New("List is empty\n")
type Uses interface {
	createStudent(name string, age uint, grades []int) *Student
	Average() float64
	Show()
	free()
}
type Student struct {
	Name   string
	Age    uint
	Grades []int
	Next   *Student
}

func (s *Student) checkGrade() bool {
	current := s
	for current != nil {
		if current.Average() < float64(50) {
			return false
		}
		current = current.Next
	}

	return true
}
func (s *Student) createStudent(name string, age uint, grades []int) *Student {

	s = &Student{
		Name:   name,
		Age:    age,
		Grades: grades,
		Next:   nil,
	}

	return s
}
func (s *Student) Average() float64 {
	
	var sum float64
	for _, ele := range s.Grades {
		sum += float64(ele)
	}
	return float64(sum / float64(len(s.Grades)))
}
func (s *Student) Show() {

	current := s
	if current == nil {
		fmt.Println(empty)
	}
	for current != nil {
		fmt.Printf("Name:%s, Age:%d, The Grades:%d\n", current.Name, current.Age, current.Grades)
		current = current.Next
	}

}
func ProcessList(s *Student, wg *sync.WaitGroup) {

	defer w.Done()

	current := s
	if current == nil {
		fmt.Println(empty)
	}
	for current != nil {
		fmt.Printf("Average: %.2f, Pass:%v\n", current.Average(), current.checkGrade())
		current = current.Next
	}
}
func addList(s *Student, name string, age uint, grades []int) *Student {
	nlist := &Student{
		Name:   name,
		Age:    age,
		Grades: grades,
		Next:   nil,
	}
	current := s
	if current == nil {
		fmt.Println(empty)
	}	
	for current.Next != nil {
		current = current.Next
	}
	current.Next = nlist
	return s

}

func (s *Student) free() *Student {

	current := s
	for current != nil {
		list := current
		current.Next = nil
		current = list
	}
	return nil
}
func main() {
	var list *Student
	list = list.createStudent("Ali", 18, []int{0, 60, 75})

	list = addList(list, "Amine", 19, []int{50, 75, 90})
	list.Show()
	w.Add(1)
	ProcessList(list, &w)
	w.Wait()
	list.free()
	fmt.Println("the list is free now:")
	list.Show()
}
