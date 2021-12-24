package main

import "sync"

func ThreadSafe_Demo1_main(){

}

type SafeMap struct {
	data map[string]interface{}
	sync.Mutex
}

func (_self *SafeMap) Put(key string, obj interface{})  {
	_self.Lock()
	defer _self.Unlock()

	_self.data[key] = obj
}

func (_self *SafeMap) Get(key string) interface{}  {
	_self.Lock()
	defer _self.Unlock()

	return _self.data[key]
}



type SafeRWMap struct {
	data map[string]interface{}
	sync.RWMutex
}

func (_self *SafeRWMap) Put(key string, obj interface{})  {
	_self.Lock()
	defer _self.Unlock()

	_self.data[key] = obj
}

func (_self *SafeRWMap) Get(key string) interface{}  {
	//读写锁，适用于 读多写少的场景，可以提升效率
	_self.RLock()
	defer _self.RUnlock()

	return _self.data[key]
}