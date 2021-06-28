package main

import (
	"log"
	"os"
	"example.com/hw2/decoder"
)

var jsonStr = []byte(`{
    "things": [
        {
            "name": "Alice",
            "age": 37
        },
        {
            "city": "Ipoh",
            "country": "Malaysia"
        },
        {
            "name": "Bob",
            "age": 36
        },
        {
            "city": "Northampton",
            "country": "England"
        },
 		{
            "name": "Albert",
            "age": 3
        },
		{
            "city": "Dnipro",
            "country": "Ukraine"
        },
		{
            "name": "Roman",
            "age": 32
        },
		{
            "city": "New York City",
            "country": "US"
        }
    ]
}`)

func main() {
	service := decoder.NewService(log.New(os.Stdout, "INFO: ", 0))
	persons, places := service.Decode(jsonStr)
    
    service.Printlen(persons, places)
    service.Print(persons)
    service.Print(places)
}

