// Copyright 2018 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package service

var defaultContainer = New()

// Set service
func Set(name string, f ContainerFunc) {
	defaultContainer.Set(name, f)
}

// Has service exists
func Has(name string) bool {
	return defaultContainer.Has(name)
}

// Get service
func Get(name string) interface{} {
	return defaultContainer.Get(name)
}

// GetKeys of all services
func GetKeys() []string {
	return defaultContainer.GetKeys()
}

// Fill dst
func Fill(name string, dst interface{}) {
	defaultContainer.Fill(name, dst)
}

// Extend service
func Extend(name string, f ExtenderFunc) {
	defaultContainer.Extend(name, f)
}
