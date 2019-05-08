// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package service

import (
	"fmt"
	"log"
	"reflect"
	"sync"
)

// Container interface
type Container interface {
	// Set service
	Set(name string, f ContainerFunc)
	// Has service exists
	Has(name string) bool
	// Get service
	Get(name string) interface{}
	// GetKeys of all services
	GetKeys() []string
	// Fill dst
	Fill(name string, dst interface{})
	// Extend service
	Extend(name string, f ExtenderFunc)
}

// ContainerFunc type
type ContainerFunc func(container Container) interface{}

// ExtenderFunc type
type ExtenderFunc interface{}

type container struct {
	values   map[string]ContainerFunc   // contains the original closure to generate the service
	extends  map[string][]reflect.Value // contains the extends closure
	services map[string]interface{}     // contains the instantiated services
	mtx      *sync.RWMutex
}

// New constructor
func New() Container {
	return &container{
		services: make(map[string]interface{}),
		values:   make(map[string]ContainerFunc),
		extends:  make(map[string][]reflect.Value),
		mtx:      &sync.RWMutex{},
	}
}

// Set service
func (c *container) Set(name string, f ContainerFunc) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if _, ok := c.services[name]; ok {
		log.Panic("Cannot overwrite initialized service")
	}

	c.values[name] = f
}

// Has service exists
func (c *container) Has(name string) bool {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	if _, ok := c.values[name]; ok {
		return true
	}

	return false
}

// Get service
func (c *container) Get(name string) interface{} {
	c.mtx.RLock()
	_, ok := c.values[name]
	c.mtx.RUnlock()
	if !ok {
		panic(fmt.Sprintf("The service does not exist: %s", name))
	}

	c.mtx.RLock()
	_, ok = c.services[name]
	c.mtx.RUnlock()
	if !ok {
		v := c.values[name](c)

		c.mtx.Lock()
		c.services[name] = v
		c.mtx.Unlock()

		// apply extends to service
		if extends, ok := c.extends[name]; ok {
			for _, extend := range extends {
				result := extend.Call([]reflect.Value{
					reflect.ValueOf(v),
					reflect.ValueOf(c),
				})

				c.mtx.Lock()
				c.services[name] = result[0].Interface()
				c.mtx.Unlock()
			}
		}
	}

	c.mtx.RLock()
	defer c.mtx.RUnlock()

	return c.services[name]
}

// GetKeys of all services
func (c *container) GetKeys() []string {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	keys := make([]string, 0)

	for k := range c.values {
		keys = append(keys, k)
	}

	return keys
}

// Fill dst
func (c *container) Fill(name string, dst interface{}) {
	obj := c.Get(name)

	if err := fill(obj, dst); err != nil {
		log.Panic(err)
	}
}

// Extend service
func (c *container) Extend(name string, f ExtenderFunc) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if _, ok := c.services[name]; ok {
		log.Panic("Cannot extend initialized service")
	}

	if _, ok := c.values[name]; !ok {
		log.Panicf("Cannot extend %s service", name)
	}

	c.extends[name] = append(c.extends[name], reflect.ValueOf(f))
}

func fill(src, dest interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			d := reflect.TypeOf(dest)
			s := reflect.TypeOf(src)
			err = fmt.Errorf("the fill destination should be a pointer to a `%s`, but you used a `%s`", s, d)
		}
	}()

	reflect.ValueOf(dest).Elem().Set(reflect.ValueOf(src))

	return err
}
