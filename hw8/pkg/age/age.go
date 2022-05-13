package age

type Employee struct {
	age int
}

func (e Employee) GetAge() int {
	return e.age
}

type Customer struct {
	age int
}

func (c Customer) GetAge() int {
	return c.age
}

type Person interface {
	GetAge() int
}

func MaxAge(people ...Person) int {
	var maxAge int
	for _, p := range people {
		if p.GetAge() > maxAge {
			maxAge = p.GetAge()
		}
	}
	return maxAge
}
