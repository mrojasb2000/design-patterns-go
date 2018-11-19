package main

import (
	"fmt"
)

type UserFinder interface {
	FinderUser(id int32) (User, error)
}

type User struct {
	ID int32
}

type UserList []User

type UserListProxy struct {
	// Mock database
	SomeDatabase *UserList
	// Cache FIFO stack
	StackCache             UserList
	StackCapacity          int
	DidLastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.DidLastSearchUsedCache = true
		return user, nil
	}

	user, err = u.SomeDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}

	u.addUserToStack(user)

	fmt.Println("Returning user from database")
	u.DidLastSearchUsedCache = false
	return user, nil

}

func (t *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*t); i++ {
		current := (*t)[i]
		if current.ID == id {
			return current, nil
		}
	}
	return User{}, fmt.Errorf("User %d could not be found\n", id)
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackCapacity {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}

func (t *UserList) addUser(newUser User) {
	*t = append(*t, newUser)
}
