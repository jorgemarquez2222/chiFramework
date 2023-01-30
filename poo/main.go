package poo

type Person struct {
	Name  string
	Email string
}
type Persons struct {
	persons []Person
}

func New(Name string, Email string) *Person {
	return &Person{
		Name:  Name,
		Email: Email,
	}
}

func (p *Persons) AddPerson() {
	var pr = &Person{
		Name:  "otra person",
		Email: "otraPerson@gmial.com",
	}
	p.persons = append(p.persons, *pr)
}

func (p *Persons) AddPersonParam(pr *Person) {
	p.persons = append(p.persons, *pr)
}

func (p *Persons) GetPersons() []Person {
	return p.persons
}

func (c *Persons) FindProductIndex(Name string) int {
	for i, v := range c.persons {
		if v.Name == Name {
			return i
		}
	}
	return -1
}

func (ps *Persons) deleteElem(index int) {
	ps.persons = append(ps.persons[:index], ps.persons[index+1:]...)
}

func (ps *Persons) RemovePerson(Name string) {
	index := ps.FindProductIndex(Name)
	if index < 0 {
		return
	}
	ps.deleteElem(index)
}

func (ps *Persons) RemoveAllPersonByName(Name string) {
	for _, v := range ps.persons {
		index := ps.FindProductIndex(v.Name)
		if index > 0 {
			ps.deleteElem(index)
		}
	}
}
