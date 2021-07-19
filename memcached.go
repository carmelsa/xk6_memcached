package memcached

import (
	"context"
	"fmt"
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
func (r *Memcached) XClient(ctxPtr *context.Context, server string, max int) interface{} {
	//fmt.Println(fmt.Sprintf("start connecting to server %v", server))
	rt := common.GetRuntime(*ctxPtr)
	//c := memcache.New(server)
	//c.MaxIdleConns = max

	ss := new(memcache.ServerList)
	if err := ss.SetServers(server); err != nil {
		panic(err)
	}
	c := memcache.NewFromSelector(ss)
	c.MaxIdleConns = 10
	return common.Bind(rt, &Client{c}, ctxPtr)
}

//Set the given key with the given value and expiration time.
func (c *Client) Set(key string, value string, exp int32) {
	err := c.client.Set(&memcache.Item{Key: key, Value: []byte(value), Expiration: exp})
	if err != nil {
		fmt.Println(fmt.Sprintf("error seting key %v", err))
	}
}

//FlushAll the given key with the given value and expiration time.
func (c *Client) FlushAll() {
	err := c.client.FlushAll()
	if err != nil {
		fmt.Println(fmt.Sprintf("error flush all data %v", err))
	}
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
