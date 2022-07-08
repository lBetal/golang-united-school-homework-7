package coverage

import (
	"os"
	"testing"
	"time"
	"reflect"
	"errors"
	"strconv"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
	panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
	panic(err)
	}
}

// WRITE YOUR CODE BELOW

func Test_len(t *testing.T) {
	people := People{}

	if people.Len() != 0 {
		t.Errorf("test for people len empty")
	}

	people = append(people, Person{})


	if people.Len() != 1 {
		t.Errorf("test for people len 1")
	}

	people = append(people, Person{})


	if people.Len() != 2 {
		t.Errorf("test for people len 1")
	}
}

func Test_less(t *testing.T) {
	people := People{}
	date := time.Time{}

	people = append(people, Person{"B", "B", date})
	people = append(people, Person{"B", "B", date.Add(1 * time.Minute)})
	people = append(people, Person{"A", "B", date})
	people = append(people, Person{"A", "A", date})
	people = append(people, Person{"A", "A", date})

	if people.Less(0, 1) {
		t.Errorf("test for people less by Birthday")
	}

	if people.Less(0, 2) {
		t.Errorf("test for people Less by FirstName")
	}

	if people.Less(2, 3) {
		t.Errorf("test for people Less by LastName")
	}

	if people.Less(3, 4) {
		t.Errorf("test for people Less Eq")
	}
}

func Test_swap(t *testing.T) {
	var people People
	date := time.Time{}

	people1 := Person{"Adam", "Adam", date}
	people2 := Person{"Eva", "Eva", date}

	people = append(people, people1)
	people = append(people, people2)

	people.Swap(0, 1)

	if people[0] != people2 || people[1] != people1 {
		t.Errorf("test for people swap")
	}
}

func Test_new(t *testing.T) {
	actual, err := New("A")
	if actual != nil || !errors.Is(err, strconv.ErrSyntax) {
		t.Errorf("test for string error")
	}

	actual, err = New("1 1 \n 2")
	if actual != nil || err.Error() != "Rows need to be the same length" {
		t.Errorf("test for wrong matrix error")
	}

	actual, err = New("1 1 \n 2 3")
	expect := &Matrix{2, 2, []int{1, 1, 2, 3}}

	if actual.cols != expect.cols || actual.rows != expect.rows || !reflect.DeepEqual(actual.data, expect.data) {
		t.Errorf("test for empty matrix")
	}
}

func Test_rows(t *testing.T) {
	m := &Matrix{2, 2, []int{1, 2, 3, 4}}
	expect := [][]int{{1, 2}, {3, 4}}

	row := m.Rows()

	if !reflect.DeepEqual(row, expect) {
		t.Errorf("test for rows matrix")
	}
}

func Test_cols(t *testing.T) {
	m := &Matrix{2, 2, []int{1, 2, 3, 4}}
	expect := [][]int{{1, 3}, {2, 4}}

	col := m.Cols()

	if !reflect.DeepEqual(col, expect) {
		t.Errorf("test for cols matrix")
	}
}

func Test_set(t *testing.T) {
	m := &Matrix{2, 2, []int{1, 2, 3, 4}}
	expect := []int{1, 2, 3, 5}

	set := m.Set(1, 1, 5)

	if !set || !reflect.DeepEqual(m.data, expect) {
		t.Errorf("test for set matrix")
	}

	set = m.Set(-1, 1, 5)
	if set {
		t.Errorf("test for set error matrix")
	}
}