package main

import (
	"google.golang.org/grpc"
	"logger"
	"node/proto"
	"golang.org/x/net/context"
	"node/handlers"
	"time"
	"node/message"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(":8090", opts...)
	if err != nil {
		logger.Errorf("dial error %v.", err)
	}
	defer conn.Close()

	client := proto.NewNodeClient(conn)

	content := handlers.EchoRequest{
		Timestamp: time.Now().Unix(),
	}

	key, err := message.CacheMsg(content)
	if err != nil {
		logger.Errorf("cache message content error %v.", err)
	}

	reqMsg := &proto.Message{
		Action: "listAction",
		Key:    key,
	}
	//respMsg, err := client.Call(context.Background(), reqMsg)
	//logger.Debugf("resp = %s, err = %v.", json.ToString(respMsg), err)
	//var respstr string
	//message.GetMsgString(respMsg.Key, &respstr)
	//logger.Debugf("response content %s.", respstr)

	reqMsg.Action = "listAction"
	client.Call(context.Background(), reqMsg)

}
