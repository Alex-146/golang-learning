package main

import "fmt"

type User struct {
	name string
	age int
	friends []*User
}

func (u User) isAdult() bool {
	return u.age >= 18
}

func (u *User) addFriend(f *User) bool {
	if u.hasFriend(f) {
		return false
	}
	u.friends = append(u.friends, f)
	f.addFriend(u)
	return true
}

func (u User) hasFriend(f *User) bool {
	for _, friend := range u.friends {
		if friend == f {
			return true
		}
	}
	return false
}

func (u *User) friendsCount() int {
	return len(u.friends)
}

func main() {
	// one()
	two()
}

func printUser(user User) {
	fmt.Println(user)
	fmt.Println(user.name, user.age)
}

func validateUser(user *User) {
	if user.name == "" {
		user.name = "Kinda randomized"
	}

	if user.age < 0 {
		user.age = 0
	}
}

func one() {
	a := User {age: -1}
	validateUser(&a)
	printUser(a)

	b := &User {age: -2}
	validateUser(b)
	printUser(*b)
}

func two() {
	var dima = User { name: "Dima", age: 20 }
	var cucumber = User { name: "Cucumber", age: 146 }

	fmt.Println("Dima has cucumber:", dima.hasFriend(&cucumber))
	fmt.Println("Cucumber added:", dima.addFriend(&cucumber))
	fmt.Println("Cucumber added again:", dima.addFriend(&cucumber))
	fmt.Println("Dima has cucumber:", dima.hasFriend(&cucumber))
	fmt.Println(dima.friendsCount())

	fmt.Println("Cucumber has dima:", cucumber.hasFriend(&dima))
	fmt.Println("Dima added again:", cucumber.addFriend(&dima))
	fmt.Println(cucumber.friendsCount())
}