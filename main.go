package main

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

type SomeRepository interface {
	GetData() string
}

type SomeRepositoryImpl struct{}

func (r *SomeRepositoryImpl) GetData() string {
	// Здесь происходит запрос к базе данных
	return "data"
}

type SomeRepositoryProxy struct {
	repository SomeRepository
	cache      *redis.Pool
}

// если не задан addr(пустая строка), то будет использоваться localhost:6379
func NewSomeRepositoryProxy(addr string) *SomeRepositoryProxy {
	// инициализируем параметры подключения к Redis
	if addr == "" {
		addr = "localhost:6379"
	}

	// Инициализация пула подключений к Redis
	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
	}

	repo := &SomeRepositoryImpl{}

	return &SomeRepositoryProxy{
		repository: repo,
		cache:      redisPool,
	}
}

func (r *SomeRepositoryProxy) GetData() string {
	conn := r.cache.Get()
	defer conn.Close()

	// Здесь происходит проверка наличия данных в кэше
	cachedValue, err := redis.String(conn.Do("GET", "data"))

	// Если данные есть в кэше, то они возвращаются
	if err == nil {
		//fmt.Println("вернули данные из кеша")
		return cachedValue
	}

	// Если данных нет в кэше(или редис не доступен), то они запрашиваются у оригинального объекта и сохраняются в кэш
	realValue := r.repository.GetData()

	_, err = conn.Do("SET", "data", realValue)
	if err != nil {
		log.Println(err)
	}

	//fmt.Println("вернули реальные данные и сохранили в кеш")
	return realValue
}

/* func main() {
	p := NewSomeRepositoryProxy("")
	fmt.Println(p.GetData())
	fmt.Println(p.GetData())
} */
