package memcached

import (
	//"fmt"
	//"time"
	"context"
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
func (r *Memcached) XClient(ctxPtr *context.Context) interface{} {
	rt := common.GetRuntime(*ctxPtr)
	return common.Bind(rt, &Client{client: memcache.New("0.0.0.0:11211")}, ctxPtr)
}

//Set the given key with the given value and expiration time.
func (c *Client) Set(key, value string, exp int32) {
	_ = c.client.Set(&memcache.Item{Key: key, Value: []byte(value), Expiration: exp})
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
