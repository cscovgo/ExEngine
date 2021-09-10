package Map_test

import (
	"ExEngine/Map"
	"fmt"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	s := Map.New()
	go func() {
		for i := 0; i < 100000; i++ {
			time.Sleep(time.Millisecond * 1)
			err := s.Set(fmt.Sprint(i), i)
			if err != nil {
				fmt.Println("set", err)
				return
			}
			fmt.Println("set", i, "<-", i)
		}
	}()
	go func() {
		for i := 0; i < 100000; i++ {
			time.Sleep(time.Millisecond * 3)
			v, ok, err := s.Get(fmt.Sprint(i))
			if err != nil {
				fmt.Println("get", err)
				return
			}
			fmt.Println("get", i, "->", ok, v)
		}
	}()
	go func() {
		for i := 0; i < 100000; i++ {
			time.Sleep(time.Millisecond * 10)
			err := s.Del(fmt.Sprint(i))
			if err != nil {
				fmt.Println("del", err)
				return
			}
			fmt.Println("del", i)
		}
	}()
	for {
		time.Sleep(time.Millisecond * 1)
	}
}
