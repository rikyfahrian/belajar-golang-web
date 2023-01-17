package main

type Student struct {
	Id    string
	Name  string
	Grade int
}

func init() {
	students = append(students, &Student{"s001", "riky", 2})
}

var students = []*Student{}

func selectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}

	return nil
}

func GetStudent() []*Student {
	return students

}
