package etcd

import (
	"testing"
)

func TestInitDiscover(t *testing.T) {
	InitEtcd("./etcd_config.conf")
}