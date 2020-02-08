package user_service

import (
	"fmt"
	"context"
	"github.com/apache/thrift/lib/go/thrift"
	"user-edge-service/app/component/user-service/user"
)

func GetUserByName(username string) (*user.UserInfo, error) {
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket("172.16.57.110:9090")
	if err != nil {
		return nil,fmt.Errorf("NewTSocket failed. err: [%v]\n", err)
	}
	transport, err = thrift.NewTBufferedTransportFactory(8192).GetTransport(transport)
	if err != nil {
		return nil, fmt.Errorf("NewTransport failed. err: [%v]\n", err)
	}
	defer transport.Close()

	protocolFactory := thrift.NewTCompactProtocolFactory()
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	if err := transport.Open(); err != nil {
		return nil,fmt.Errorf("Transport.Open failed. err: [%v]\n", err)
	}
	client := user.NewUserServiceClient(thrift.NewTStandardClient(iprot, oprot))
	userinfo,err := client.GetUserByName(context.Background(), username)
	fmt.Println("GetUserByName-5")
	if err != nil {
		return nil,fmt.Errorf("GetUserByName failed. err: [%v]\n", err)
	}
	return userinfo,nil
}
