package ShardMap_test

import (
	"ExEngine/Map/ShardMap"
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	m := ShardMap.New()
	m.Set("key1", "value1")
	m.Set("key2", "value2")
	m.Set("key3", "value3")
	m.Set("key4", "value4")
	fmt.Println(m.Get("key1"))
	fmt.Println(m.Get("key2"))
	fmt.Println(m.Get("key3"))
	fmt.Println(m.Get("key4"))
	m.Del("key4")
	fmt.Println(m.Get("key4"))
	fmt.Println(m.Len())
	fmt.Println(m.Keys())
}

func BenchmarkShardMap(b *testing.B) {
	num := 10000
	kv := genKv(num)
	m := ShardMap.New()
	for k, v := range kv {
		m.Set(k, v)
	}
	b.ResetTimer()
	for i := 0; i < 5; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			b.N = 1000000
			wg := sync.WaitGroup{}
			wg.Add(b.N * 2)
			for i := 0; i < b.N; i++ {
				e := kv[rand.Intn(num)]
				go func(k, v interface{}) {
					m.Set(k, v)
					wg.Done()
				}(e.k, e.v)
				go func(k string) {
					_, _ = m.Get(k)
					wg.Done()
				}(e.k)
			}
			wg.Wait()
		})
	}
}

type kv struct {
	k string
	v string
}

func genKv(num int) (kvs []kv) {
	for i := 0; i < num; i++ {
		k := fmt.Sprint(i)
		h := md5.New()
		h.Write([]byte(k))
		kvs = append(kvs, kv{
			k: k,
			v: string(h.Sum(nil)),
		})
	}
	return
}
