package Slice_test

import (
	"ExEngine/Slice"
	"fmt"
	"testing"
	"time"
)

func TestSlice(t *testing.T) {
	s := Slice.New()
	_ = s.Append(nil)
	_ = s.Append(nil)
	err := s.Append(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s.All())
	go func() {
		for i := 0; i < 1000000; i++ {
			//time.Sleep(time.Millisecond * 1)
			err := s.Append(i)
			if err != nil {
				fmt.Println("append", err)
				return
			}
			fmt.Println("append", i)
		}
	}()
	go func() {
		for i := 0; i < 100000; i++ {
			time.Sleep(time.Millisecond * 2)
			v, err := s.Index(i)
			if err != nil {
				fmt.Println("index", i, err)
				return
			}
			fmt.Println("index", i, "->", v)
		}
	}()
	go func() {
		for i := 0; i < 1000000; i++ {
			time.Sleep(time.Millisecond * 5)
			err := s.Del(i)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("del")
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 2)
			fmt.Println("all->", s.Len())
		}
	}()
	for {
		time.Sleep(time.Millisecond)
	}
}

func TestInArr(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(Slice.InArr(s1, 5))
	s2 := []string{"1", "2", "3", "4", "5", "6", "7"}
	fmt.Println(Slice.InArr(s2, "5"))
	s3 := []float64{1.1, 2.2, 3.3, 4.4, 5.5555, 6.7, 7.99}
	fmt.Println(Slice.InArr(s3, 5.5555))
	s4 := []float32{1.1, 2.2, 3.3, 4.4, 5.5555, 6.7, 7.99}
	fmt.Println(Slice.InArr(s4, float32(5.5555)))
	type a struct{ A int }
	s5 := []a{{A: 1}, {A: 2}, {A: 3}, {A: 4}, {A: 5}, {A: 6}, {A: 7}}
	fmt.Println(Slice.InArr(s5, a{A: 5}))
	s6 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(Slice.InArr(s6, nil))
}
