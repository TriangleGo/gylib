package etcd

import "testing"

func TestRegisterService(t *testing.T) {
	InitEtcd("./etcd_config.conf")
	err := _default_pool.registerService()
	t.Log(err)
}