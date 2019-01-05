// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultContainer(t *testing.T) {
	assert.False(t, Has("test.bad.service.name"))

	assert.Equal(t, []string{}, GetKeys())

	Set("my.service", func(c Container) interface{} {
		return &MyService{}
	})

	Extend("my.service", func(s *MyService, c Container) *MyService {
		s.Name = "My Service"

		return s
	})

	assert.True(t, Has("my.service"))

	assert.Equal(t, []string{"my.service"}, GetKeys())

	Set("my.service", func(c Container) interface{} {
		return &MyService{}
	})

	myService1 := Get("my.service").(*MyService)

	myService2 := Get("my.service").(*MyService)

	assert.Equal(t, myService1, myService2)

	assert.Equal(t, "My Service", myService1.Name)

	assert.Panics(t, func() {
		Set("my.service", func(c Container) interface{} {
			return &MyService{}
		})
	})

	assert.Panics(t, func() {
		Extend("my.service", func(s *MyService, c Container) *MyService {
			s.Name = "My Service 2"

			return s
		})
	})

	assert.Panics(t, func() {
		Extend("not.exists.service", func(s *MyService, c Container) *MyService {
			s.Name = "My Service 3"

			return s
		})
	})

	assert.Panics(t, func() {
		Get("test.bad.service.name")
	})

	var myService3 *MyService

	Fill("my.service", &myService3)

	assert.Equal(t, myService2, myService3)

	assert.Panics(t, func() {
		var bad string

		Fill("my.service", &bad)
	})
}
