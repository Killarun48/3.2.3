package main

import (
	"reflect"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestNewSomeRepositoryProxy(t *testing.T) {
	repository := NewSomeRepositoryProxy("")
	if reflect.TypeOf(repository) != reflect.TypeOf(&SomeRepositoryProxy{}) {
		t.Errorf("NewSomeRepositoryProxy() = %v, want %v", reflect.TypeOf(repository), reflect.TypeOf(&SomeRepositoryProxy{}))
	}

	if reflect.TypeOf(repository.repository) != reflect.TypeOf(&SomeRepositoryImpl{}) {
		t.Errorf("NewSomeRepositoryProxy() = %v, want %v", reflect.TypeOf(repository.repository), reflect.TypeOf(&SomeRepositoryImpl{}))
	}

	if reflect.TypeOf(repository.cache) != reflect.TypeOf(&redis.Pool{}) {
		t.Errorf("NewSomeRepositoryProxy() = %v, want %v", reflect.TypeOf(repository.cache), reflect.TypeOf(&redis.Pool{}))
	}

}

func TestSomeRepositoryProxy_GetData(t *testing.T) {
	repository := NewSomeRepositoryProxy("")
	data := repository.GetData()
	if data != "data" {
		t.Errorf("repository.GetData() = %v, want %v", data, "data")
	}
	data = repository.GetData()
	if data != "data" {
		t.Errorf("repository.GetData() = %v, want %v", data, "data")
	}
}

func TestData(t *testing.T) {
	main()
}
func TestSomeRepositoryProxy_GetData_DontWorkingRedis(t *testing.T) {
	repository := NewSomeRepositoryProxy("abc")
	data := repository.GetData()
	if data != "data" {
		t.Errorf("repository.GetData() = %v, want %v", data, "data")
	}
	data = repository.GetData()
	if data != "data" {
		t.Errorf("repository.GetData() = %v, want %v", data, "data")
	}
}

