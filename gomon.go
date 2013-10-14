package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Person struct {
	Name string
	Phone string
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Art", "+10 215 264 7491"})
	err = c.Insert(&Person{"Clare", "+55 53 8402 8510"})
	if err != nil {
		panic(err)
	}


	result := Person{}
	err = c.Find(bson.M{"name": "Art"}).One(&result)
	//err = c.Find(bson.M{"name": "Clare"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Name:", result.Name)
	fmt.Println("Phone:", result.Phone)
}
