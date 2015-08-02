/*
Example: Semaphore

Based on Beego sessions store which was implemented using mutexes, this is a channel
based apprach.

Given an In Memory Session Store, it should Store/Update/Delete sessions concurrently
No locks. No condition variables. No callbacks
*/
package main

import (
	"sync"
)

type signal struct{}

// Based on from github.com/astaxie/beego/session
type SessionStore interface {
	Set(key string, value interface{}) //set session value
	Get(key string) interface{}        //get session value
	Delete(key string)                 //delete session value
	SessionID() string                 //back current sessionID
	Flush()                            //delete all data
}

// memory session store.
// it saved sessions in an in memory map.
// based on from github.com/astaxie/beego/session/sess_mem.go
type MemorySessionStore struct {
	SessionStore                        //implements SessionStore interface
	sessionId    string                 //session id
	session      map[string]interface{} //session store
	semaphore    chan signal            //semaphore
}

// NewMemorySessionStore: Initialize a new In Memory Session Store
func NewMemorySessionStore(sessionId string) *MemorySessionStore {
	mStore := &MemorySessionStore{
		sessionId: sessionId,
		session:   make(map[string]interface{}),
		semaphore: make(chan signal, 1),
	}
	return mStore
}

// Set: set user to memory session
func (st *MemorySessionStore) Set(key string, value interface{}) {
	st.semaphore <- signal{}
	st.session[key] = value
	<-st.semaphore
}

// Get: get user from memory session by key
func (st *MemorySessionStore) Get(key string) interface{} {
	st.semaphore <- signal{}
	user, ok := st.session[key]
	<-st.semaphore
	if ok {
		return user
	} else {
		return nil
	}
}

// Delete: delete in memory session by key
func (st *MemorySessionStore) Delete(key string) {
	st.semaphore <- signal{}
	delete(st.session, key)
	<-st.semaphore
}

// Flush: clear all users in memory session
func (st *MemorySessionStore) Flush() {
	st.semaphore <- signal{}
	st.session = make(map[string]interface{})
	<-st.semaphore
}

type User struct {
	Username     string
	Email        string
	PasswordHash string
}

func makeUsers() []User {
	users := []User{}
	for i := 0; i < 50; i++ {
		u := User{
			Username:     "User#" + string(i),
			Email:        "User" + string(i) + "@example.com",
			PasswordHash: "UsErP4S5Word!" + string(i),
		}
		users = append(users, u)
	}
	return users
}

// The following program is for educational purposes, each goroutine runs in a
// no deterministic way, so there is no way (unless we syncronize them manually) to
// ensure a (Delete/Get) call goes after a Set, but running this program with --race
// show us that there is not any race condition.
func main() {
	users := makeUsers()
	mStore := NewMemorySessionStore("memSession")
	var wg sync.WaitGroup
	for _, u := range users {
		wg.Add(1)
		go func(u User) {
			mStore.Set(u.Username, u)
			wg.Done()
		}(u)

		wg.Add(1)
		go func(u User) {
			mStore.Get(u.Username)
			wg.Done()
		}(u)

		wg.Add(1)
		go func(u User) {
			mStore.Delete(u.Username)
			wg.Done()
		}(u)
	}
	wg.Wait()
}
