package geecache

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

//type server int
//
//func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	log.Println(r.URL.Path)
//	w.Write([]byte("Hello World!"))
//}

func TestHTTP(t *testing.T) {
	NewGroup("scores", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
	addr := "localhost:9999"
	peers := NewHTTPPool(addr)
	log.Println("geecache is running at: ", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
