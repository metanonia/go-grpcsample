package com.metanonia.grpcsample;

import java.io.FileInputStream;
import java.io.FileOutputStream;

import com.metanonia.grpcsample.AddressbookProtos;
import com.metanonia.grpcsample.AddressbookProtos.Person.PhoneType;

public class ReadWriteTest {
    public static void main(String[] args) {
        AddressbookProtos.Person.PhoneNumber phoneNumber = 
            AddressbookProtos.Person.PhoneNumber.newBuilder()
                .setNumber("555-4321")
                .setType(PhoneType.HOME)
                .build();

        AddressbookProtos.Person person = 
            AddressbookProtos.Person.newBuilder()
            .setId(1234)
            .setName("John Doe")
            .setEmail("jdoe@example.com")
            .addPhones(0, phoneNumber)
            .build();

        AddressbookProtos.AddressBook addressBook = 
            AddressbookProtos.AddressBook.newBuilder()
                .addPeople(person)
                .build();

        try {
            FileOutputStream fos = new FileOutputStream("person.dat", false);
            person.writeTo(fos);
            fos.close();

            AddressbookProtos.Person dPerson =
                AddressbookProtos.Person.newBuilder()
                    .mergeFrom(new FileInputStream("person.dat"))
                    .build();
            
            System.out.println("name: " +  dPerson.getName());
            System.out.println("Id: " +  dPerson.getId());
            System.out.println("Email: " +  dPerson.getEmail());  

            FileOutputStream fos2 = new FileOutputStream("addressbook.dat", false);
            addressBook.writeTo(fos2);
            fos2.close();

            AddressbookProtos.AddressBook dAddress =
            AddressbookProtos.AddressBook.newBuilder()
                .mergeFrom(new FileInputStream("addressbook.dat"))
                .build();
            
            for(int i=0; i < dAddress.getPeopleCount(); i++) {
                System.out.println("name: " +  dAddress.getPeople(i).getName());
                System.out.println("Id: " +  dAddress.getPeople(i).getId());
                System.out.println("Email: " +  dAddress.getPeople(i).getEmail()); 
            }
            
        }
        catch(Exception e) {
            System.out.println(e.toString());
        }
        

        
    }
}