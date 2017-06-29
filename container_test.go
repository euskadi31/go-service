package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyService struct {
}

func TestContainer(t *testing.T) {
	c := New()

	assert.False(t, c.Has("test.bad.service.name"))

	assert.Equal(t, []string{}, c.GetKeys())

	err := c.Set("my.service", func(c *Container) interface{} {
		return &MyService{}
	})
	assert.NoError(t, err)

	assert.True(t, c.Has("my.service"))

	assert.Equal(t, []string{"my.service"}, c.GetKeys())

	err = c.Set("my.service", func(c *Container) interface{} {
		return &MyService{}
	})
	assert.NoError(t, err)

	myService1 := c.Get("my.service").(*MyService)

	myService2 := c.Get("my.service").(*MyService)

	assert.Equal(t, myService1, myService2)

	err = c.Set("my.service", func(c *Container) interface{} {
		return &MyService{}
	})
	assert.Error(t, err)

	assert.Panics(t, func() {
		c.Get("test.bad.service.name")
	})
}
