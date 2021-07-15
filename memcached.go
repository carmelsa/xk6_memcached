package memcached

import (
	//"time"
	"context"
	"fmt"

	//"google.golang.org/appengine/memcache"
	"github.com/bradfitz/gomemcache/memcache"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

// Register the extension on module initialization, available to
// import from JS as "k6/x/memcached".
func init() {
	modules.Register("k6/x/memcached", new(Memcached))
}

type Memcached struct{}

type Client struct {
	client *memcache.Client
}

//XClient represents the Client constructor (i.e. `new Memcached.Client()`) and
//returns a new Memcached client object.
func (r *Memcached) XClient(ctxPtr *context.Context, server string) interface{} {
	fmt.Println(fmt.Sprintf("start connecting to server %v", server))
	rt := common.GetRuntime(*ctxPtr)
	return common.Bind(rt, &Client{client: memcache.New(server)}, ctxPtr)
}

//Set the given key with the given value and expiration time.
func (c *Client) Set(key string, value string, exp int32) {
	aa := legalKey(key)
	fmt.Println(fmt.Sprintf("legalKey resulte %v", aa))
	fmt.Println(fmt.Sprintf("key is %v", key))

	err := c.client.Set(&memcache.Item{Key: key, Value: []byte(value), Expiration: exp})
	if err != nil {
		fmt.Println(fmt.Sprintf("error seting key %v", err))
	}
}

func legalKey(key string) bool {
	fmt.Println(fmt.Sprintf("len %v", len(key)))
	if len(key) > 250 {
		return false
	}
	for i := 0; i < len(key); i++ {
		fmt.Println(i)
		if key[i] <= ' ' || key[i] == 0x7f {
			return false
		}
	}
	return true
}

func (c *Client) Ping() error {
	return c.client.Ping()
}

// Get returns the value for the given key.
func (c *Client) Get(key string) (string, error) {
	item, err := c.client.Get(key)
	if err != nil {
		return "", err
	}
	return string(item.Value), nil
}

//
//func main() {
//	//client_1 :=memcache.Client("localhost")
//	//fmt.Printf("ping %v", client_1.ping())
//	//item := &memcache.Item{
//	//	Key:   "[aaaa]",
//	//	Value: []byte("[bbbb]"),
//	//}
//	//_ = item
//	mc := memcache.New("0.0.0.0:11211")
//	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value"), Expiration: 3})
//
//	it, err := mc.Get("foo")
//	//if err := memcache.Set(c, item); err != nil {
//	//	fmt.Println(fmt.Sprint("error occurred: %v", err))
//	//}
//	fmt.Printf("test %v %v", it, err)
//
//	fmt.Println("Hello, world.")
//}
