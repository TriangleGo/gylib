package etcd

import (
	"testing"
	orgEtcd "github.com/coreos/go-etcd/etcd"
	"logger"
	"util"
	"bufio"
	"os"
	"service/action"
)

func TestEtcd(t *testing.T) {
	InitEtcd("./etcd_config.conf", action.Profile_service_range)
	//clearAll("/gy")
	listAll("/gy")
	running := true
	reader := bufio.NewReader(os.Stdin)
	for running {
		data, _, _ := reader.ReadLine()
		command := string(data)
		if command == "stop" {
			running = false
		}
		t.Log("command", command)
	}
}

func listAll(prefix string) (err error) {
	client := _default_pool.client_pool.Get().(*orgEtcd.Client)
	defer func() {
		_default_pool.client_pool.Put(client)
	}()
	resp, err := client.Get(prefix, true, true)
	logger.Infof(test.ToJsonString(resp))
	return
}

func clearAll(prefix string) (err error) {
	client := _default_pool.client_pool.Get().(*orgEtcd.Client)
	defer func() {
		_default_pool.client_pool.Put(client)
	}()
	resp, err := client.Delete(prefix, true)
	logger.Infof(test.ToJsonString(resp))
	return
}