package age

type Employee struct {
	age int
}

func (e Employee) Age() int {
	return e.age
}

type Customer struct {
	age int
}

func (c Customer) Age() int {
	return c.age
}

type Person interface {
	Age() int
}

func MaxAge(people ...Person) int {
	var maxAge int
	for _, p := range people {
		if p.Age() > maxAge {
			maxAge = p.Age()
		}
	}
	return maxAge
}
