package main

import "fmt"
import "errors"
import "sync"
import "time"


 var empty error = 	errors.New("List is empty")
 var m sync.Mutex
type Uses interface {
	createStudent(name string, age uint, grades []int) *Student
	Average() float64
	Show()
	free()
	ChangeGrade(list *Student, grades []int) *Student

}
type Student struct {
	Name   string
	Age    uint
	Grades []int
	Next   *Student
}

func (s *Student) checkGrade() bool {
	current := s
	if s == nil {
		fmt.Println(empty)
		return false
	}
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
	if s == nil {
		fmt.Println(empty)
		return 0
	}	
	var sum float64
	for _, ele := range s.Grades {
		sum += float64(ele)
	}
	return float64(sum / float64(len(s.Grades)))
}
func (s *Student) Show() {
	current := s
	if s == nil {
		fmt.Println(empty)
		return
	}
	for current != nil {
		fmt.Printf("Name:%s, Age:%d, The Grades:%d\n", current.Name, current.Age, current.Grades)
		current = current.Next
	}
}
func ProcessList(s *Student, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	defer m.Unlock()
	current := s
	if s == nil {
		fmt.Println(empty)
		return
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
	if s == nil {
		fmt.Println(empty)
		return	nlist 
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
		list := current.Next
		current.Next = nil
		current = list
	}
	return nil
}
func  ChangeGrade(s *Student, grades []int) *Student {
	current := s 
	
		current.Grades = grades
	
	return s 

}
func main() {
var wg sync.WaitGroup
var m sync.Mutex
	var list *Student
	list = list.createStudent("Ali", 18, []int{0, 60, 75})

	list = addList(list, "Amine", 19, []int{50, 75, 90})
	list.Show()
	wg.Add(1)
	start := time.Now()
	go ProcessList(list, &wg, &m)
	end := time.Since(start)
	fmt.Println("the time for the list to be processed was:",end )
	wg.Wait()
	list =list.free()
	fmt.Println("the list is free now:")
	list.Show()
	fmt.Println("Change the list:")
	newgrades:= []int{0,50,70}
	 list = ChangeGrade(list, newgrades)
}
