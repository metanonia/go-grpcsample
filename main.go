package main

import (
	"io/ioutil"
	"log"

	pb "github.com/metanonia/go-grpcsample/proto"
	"google.golang.org/protobuf/proto"
)

const (
	filename = "person.dat"
    filename2 = "addressbook.dat"
)

func main() {
    // 쓰기
    writeMessage()
    // 읽기
    readMessage()
}

func writeMessage() {
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	out, err := proto.Marshal(&p)
	if err != nil {
        log.Fatalf("Failed to encode Person: %v", err)
    }

    if err := ioutil.WriteFile(filename, out, 0644); err != nil {
        log.Fatalf("Failed to write Person: %v", err)
    }

    a := pb.AddressBook{}
    a.People = append(a.People, &p)
    out2, err := proto.Marshal(&a)
	if err != nil {
        log.Fatalf("Failed to encode AddressBook: %v", err)
    }

    if err := ioutil.WriteFile(filename2, out2, 0644); err != nil {
        log.Fatalf("Failed to write AddressBook: %v", err)
    }
}

func readMessage() {
	in, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatalln("Error reading file:", err)
    }

    person := &pb.Person{}

    // proto 의존성을 통해 디코딩
    if err := proto.Unmarshal(in, person); err != nil {
        log.Fatalln("Failed to parse person:", err)
    }

    log.Printf("name: %s", person.GetName())
    log.Printf("Id: %d", person.GetId())
	log.Printf("Email: %s", person.GetEmail())

    in2, err := ioutil.ReadFile(filename2)
    if err != nil {
        log.Fatalln("Error reading file:", err)
    }

    addressbook := &pb.AddressBook{}

    // proto 의존성을 통해 디코딩
    if err := proto.Unmarshal(in2, addressbook); err != nil {
        log.Fatalln("Failed to parse AddressBook:", err)
    }
    people := addressbook.GetPeople()
    for i := 0; i<len(people); i++ {
        log.Printf("name: %s", people[i].GetName())
        log.Printf("Id: %d", people[i].GetId())
        log.Printf("Email: %s", people[i].GetEmail())
    }
    
}