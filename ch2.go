package main

import (
    "fmt"
)

type Saiyan struct {
    *Person
    Power int
    Father *Saiyan
}

type Person struct {
    Name string
}

func main() {
    goku := &Saiyan {
        Person: &Person{"Gohan"},
        Power: 9000,
        Father: &Saiyan {
            Person: &Person{ Name: "Goku" },
            Power: 9000,
            Father: nil,
        },
    }
    super(goku)
    goku.alsoSuper()
    goku.introduce()
    fmt.Println(goku)
    goku.exit()
    goku.Person.exit()
}

func (s *Saiyan) alsoSuper() {
    s.Power += 1
}

func (p *Person) exit() {
    fmt.Printf("Cya %s\n", p.Name)
}

func (p *Saiyan) exit() {
    fmt.Printf("Bye %s\n", p.Name)
}

func (p *Person) introduce() {
    fmt.Printf("Hello %s\n", p.Name)
}

func super(s *Saiyan) {
    s.Power = s.Power + 12
}
