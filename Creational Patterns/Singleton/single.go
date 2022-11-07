package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &single{}
			fmt.Println("Create new Instance")
		} else {
			fmt.Println("Got existing Instance")
		}
	} else {
		fmt.Println("Got existing Instance")
	}
	return singleInstance
}
