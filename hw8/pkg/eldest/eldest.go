package eldest

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

func Eldest(people ...interface{}) interface{} {
	var eldest interface{}
	var maxAge int
	for _, person := range people {
		if p, ok := person.(Employee); ok {
			if p.age > maxAge {
				eldest = p
				maxAge = p.age
			}
		}

		if p, ok := person.(Customer); ok {
			if p.age > maxAge {
				eldest = p
				maxAge = p.age
			}
		}
	}
	return eldest
}

func EldestWithSwitch(people ...interface{}) interface{} {
	var eldest interface{}
	var maxAge int
	for _, person := range people {
		switch p := person.(type) {
		case Customer:
			if p.GetAge() > maxAge {
				maxAge = p.GetAge()
				eldest = p
			}
		case Employee:
			if p.GetAge() > maxAge {
				maxAge = p.GetAge()
				eldest = p
			}
		}
	}
	return eldest
}

func EldestWithGenerics[P Person](people ...P) P {
	var eldest P
	var maxAge int
	for _, p := range people {
		if p.GetAge() > maxAge {
			maxAge = p.GetAge()
			eldest = p
		}
	}
	return eldest
}
