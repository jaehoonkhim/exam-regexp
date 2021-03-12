package main

import (
	"fmt"
	"regexp"
)

type Credential struct {
	Name  string
	Value string
}

func GenCredential() *[]Credential {
	cs := []Credential{
		{Name: "firstName", Value: "jaehoon"},
		{Name: "lastName", Value: "kim"},
		{Name: "age", Value: "990"},
		{Name: "height", Value: "0.1"},
		{Name: "Gender", Value: "male"},
		{Name: "Lang1", Value: "English"},
		{Name: "Lang2", Value: "Go"},
		{Name: "mathGrade", Value: "3"},
		{Name: "scienceGrade", Value: "0"},
	}

	return &cs
}

func main() {

	cs := GenCredential()
	parameter := `{"firstName": "{{firstName}}","lastName": "{{lastName}}","age": {{age}},"heightInMeters": {{height}},"isMale": {{Gender}},"profile": null,"languages": [ "{{Lang1}}", "{{Lang2}}" ],"grades": {"math": "{{mathGrade}}","science": "{{scienceGrade}}"}}`
	fmt.Println(parameter)

	for _, c := range *cs {
		//fmt.Println(c.Name, ":", c.Value)
		a := fmt.Sprintf("{{2}%s}{2}", c.Name)
		rxp := regexp.MustCompile(a)
		parameter = fmt.Sprintf("%s", rxp.ReplaceAll([]byte(parameter), []byte(c.Value)))
	}

	fmt.Println(parameter)

}
