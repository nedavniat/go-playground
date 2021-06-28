package decoder

import (
	"encoding/json"
	"os"
	"fmt"
)


type Person struct {
	Name string
	Age int
}

type Place struct {
	City, Country string
}

type Joint struct {
	Name, City, Country string
	Age int
}

type jsonStruct struct {
	Things []Joint
}

type Decoder interface {
	Decode(data []byte) ([]Person, []Place)
	Print(interface{})
	Printlen(persons []Person, places []Place)
}

type Logger interface {
	Println(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type Service struct {
	log Logger
}

func NewService (l Logger) *Service {
	return &Service {
		log: l,
	}
}

func (s *Service) Decode(data []byte) ([]Person, []Place){
	var result jsonStruct
	var places []Place = make([]Place, 0, len(result.Things))
	var persons []Person = make([]Person, 0, len(result.Things))
	
	if err := json.Unmarshal(data, &result); err != nil {
		s.log.Fatalf("Critical error %v", err)
		os.Exit(1)
	}
	fmt.Println(result)

	for _, obj := range result.Things {
		if (obj.Name != "") {
			persons = append(persons, Person{
				Name: obj.Name,
				Age: obj.Age,
			})
			continue
		}
		places = append(places, Place{
			City: obj.City,
			Country: obj.Country,
		})
	}

	return persons, places
}

func (s *Service) Print(v interface{}) {
	s.log.Println(v)
}

func (s *Service) Printlen(persons []Person, places []Place) {
	s.log.Println(len(persons), len(places))
}
