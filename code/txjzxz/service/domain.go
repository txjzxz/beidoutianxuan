package service

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	fabAPI "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"time"
)

type ServiceSetup struct {
	ChaincodeID1 string
	ChaincodeID2 string
	ChaincodeID3 string
	ChaincodeID4 string
	ChaincodeID5 string
	Client       *channel.Client
	SDK          *fabsdk.FabricSDK
}

func regitserEvent(client *channel.Client, chaincodeID, eventID string) (fabAPI.Registration, <-chan *fabAPI.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("注册链码事件失败: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fabAPI.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到链码事件: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return nil
}

