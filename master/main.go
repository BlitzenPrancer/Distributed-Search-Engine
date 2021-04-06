package main

import (
	"log"
	"os"
)

type User struct {
	Email string
	Name  string
}

var DataBase = []User{
	{Email: "abdul.kalam@email.com", Name: "Abdul Kalam"},
	{Email: "narendra.modi@email.com", Name: "Narendra Modi"},
	{Email: "jeff.bezos@email.com", Name: "Jeff Bezos"},
	{Email: "tom.cruise@email.com", Name: "Tom Cruise"},
	{Email: "mark.zuckerberg@email.com", Name: "Mark Zuckerberg"},
	{Email: "donald.trump@email.com", Name: "Donald Trump"},
	{Email: "robin.lee@email.com", Name: "Robin Lee"},
	{Email: "will.smith@email.com", Name: "Will Smith"},
	{Email: "jaden.smith@email.com", Name: "Jaden Smith"},
	{Email: "jackie.chan@email.com", Name: "Jackie chan"},
	{Email: "arnold.schwaranzenegger@email.com", Name: "Arnold Schwaranzenegger"},
	{Email: "gal.gadot@email.com", Name: "Gal Gadot"},
	{Email: "brie.larson@email.com", Name: "Brie Larson"},
	{Email: "sachin.tendulkar@email.com", Name: "Sachin Tendulkar"},
	{Email: "rahul.gandhi@email.com", Name: "Rahul Gandhi"},
	{Email: "zoe.loe@email.com", Name: "Zoe loe"},
	{Email: "stephen.hawking@email.com", Name: "Stephen Hawking"},
	{Email: "kamla.harris@email.com", Name: "Kamla Harris"},
	{Email: "mickey.mouse@email.com", Name: "Mickey Mouse"},
	{Email: "ratan.tata@email.com", Name: "Ratan Tata"},
	{Email: "mukesh.ambani@email.com", Name: "Mukesh Ambani"},
}

// This structure has the copy of the databse
type Worker struct {
	users []User
}

func NewWorker(users []User, ch chan *User) *Worker {
	return &Worker{users: users, ch: ch}
}

func (w *Worker) Find(email string) {
	// iterating over the database
	for i := range w.users {
		user := &w.users[i]
		if user.Email == email {
			ch <- user
		}
	}
}

func main() {
	// getting input from user
	email := os.Args[1]
	// creating channels to make the workers access the database asynchornously
	ch := make(chan *User)
	// creating the worker
	w := NewWorker(DataBase, ch)
	log.Printf("Looking for %s", email)
	w.Find(email)
	if user != nil {
		log.Printf("The email %s is owned by %s", email, user.Name)
	} else {
		log.Printf("The email %s was not found", email)
	}
}
