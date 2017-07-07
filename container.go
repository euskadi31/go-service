package service

import (
	"errors"
	"fmt"
	"sync"
)

// ContainerFunc type
type ContainerFunc func(container *Container) interface{}

// Container struct
type Container struct {
	values   map[string]ContainerFunc // contains the original closure to generate the service
	services map[string]interface{}   // contains the instantiated services
	mtx      *sync.RWMutex
}

// New constructor
func New() *Container {
	return &Container{
		services: make(map[string]interface{}),
		values:   make(map[string]ContainerFunc),
		mtx:      &sync.RWMutex{},
	}
}

// Set service
func (c *Container) Set(name string, f ContainerFunc) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if _, ok := c.services[name]; ok {
		return errors.New("Cannot overwrite initialized service")
	}

	c.values[name] = f

	return nil
}

// Has service exists
func (c *Container) Has(name string) bool {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	if _, ok := c.values[name]; ok {
		return true
	}

	return false
}

// Get service
func (c *Container) Get(name string) interface{} {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if _, ok := c.values[name]; !ok {
		panic(fmt.Sprintf("The service does not exist: %s", name))
	}

	if _, ok := c.services[name]; !ok {
		c.services[name] = c.values[name](c)
	}

	return c.services[name]
}

// GetKeys of all services
func (c *Container) GetKeys() []string {
	c.mtx.RLock()
	defer c.mtx.RUnlock()

	keys := make([]string, 0)

	for k := range c.values {
		keys = append(keys, k)
	}

	return keys
}
