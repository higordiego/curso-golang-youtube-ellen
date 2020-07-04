package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type porIdade []user

func (a porIdade) Len() int           { return len(a) }
func (a porIdade) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a porIdade) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type porSobrenNome []user

func (a porSobrenNome) Len() int           { return len(a) }
func (a porSobrenNome) Less(i, j int) bool { return a[i].Last < a[j].Last }
func (a porSobrenNome) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	sort.Sort(porIdade(users))
	fmt.Println("Ordenado por idade")
	harmonioza(users)
	fmt.Println("Ordenado por sobrenome")
	sort.Sort(porSobrenNome(users))
	harmonioza(users)



}

func harmonioza(x []user) {
	for i, user := range x {
		fmt.Printf("%v. \n\tnome: %v \n\tSobrenome: %v \n\tidade: %v \n", i+1, user.First, user.Last, user.Age)
		fmt.Println("\tSayings:")
		for r, v := range user.Sayings {
			fmt.Printf("\t\t %v.  %v \n", r+1, v)
		}
		fmt.Println()
	}
}
