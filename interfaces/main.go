package interfaces

type Storager interface {
	Get() string
	Set(Name string)
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func New(Name string, typeI string) Storager {
	if typeI == "animal" {
		return &Animal{
			Name: Name,
		}
	}
	if typeI == "persona" {
		return &Person{
			Name: Name,
		}
	}

	return &Person{
		Name: Name,
	}
}

func (p *Person) Get() string {
	return p.Name
}

func (p *Person) Set(Name string) {
	p.Name = Name
}
func (p *Animal) Get() string {
	return p.Name
}

func (p *Animal) Set(Name string) {
	p.Name = Name
}

func Exec(s Storager, Name string) string {
	s.Set(Name)
	return s.Get()
}
