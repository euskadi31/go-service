// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyService struct {
	Name string
}

func TestContainer(t *testing.T) {
	c := New()

	assert.False(t, c.Has("test.bad.service.name"))

	assert.Equal(t, []string{}, c.GetKeys())

	c.Set("my.service", func(c Container) interface{} {
		return &MyService{}
	})

	c.Extend("my.service", func(s *MyService) *MyService {
		s.Name = "My Service"

		return s
	})

	assert.True(t, c.Has("my.service"))

	assert.Equal(t, []string{"my.service"}, c.GetKeys())

	c.Set("my.service", func(c Container) interface{} {
		return &MyService{}
	})

	myService1 := c.Get("my.service").(*MyService)

	myService2 := c.Get("my.service").(*MyService)

	assert.Equal(t, myService1, myService2)

	assert.Equal(t, "My Service", myService1.Name)

	assert.Panics(t, func() {
		c.Set("my.service", func(c Container) interface{} {
			return &MyService{}
		})
	})

	assert.Panics(t, func() {
		c.Extend("my.service", func(s *MyService) *MyService {
			s.Name = "My Service 2"

			return s
		})
	})

	assert.Panics(t, func() {
		c.Extend("not.exists.service", func(s *MyService) *MyService {
			s.Name = "My Service 3"

			return s
		})
	})

	assert.Panics(t, func() {
		c.Get("test.bad.service.name")
	})

	var myService3 *MyService

	c.Fill("my.service", &myService3)

	assert.Equal(t, myService2, myService3)

	assert.Panics(t, func() {
		var bad string

		c.Fill("my.service", &bad)
	})
}
