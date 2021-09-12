package main

import (
	"fabric-go-sdk/sdkInit"
	"fabric-go-sdk/service"
	"fabric-go-sdk/web"
	"fabric-go-sdk/web/controller"
	"fmt"
	"os"
)

const (
	cc_name1    = "step1"
	cc_version1 = "4.0.0"
	cc_name2    = "step2"
	cc_version2 = "2.0.0"
	cc_name3    = "step3"
	cc_version3 = "2.0.0"
	cc_name4    = "step4"
	cc_version4 = "1.0.0"
	cc_name5    = "step5"
	cc_version5 = "3.0.0"
)

var App controller.Application

func main() {
	// init orgs information

	orgs := []*sdkInit.OrgInfo{
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org1",
			OrgMspId:      "Org1MSP",
			OrgUser:       "User1",
			OrgPeerNum:    2,
			OrgAnchorFile: "/root/go/src/fabric-go-sdk/fixtures/channel-artifacts/allOrg1MSPanchors.tx",
		},
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org2",
			OrgMspId:      "Org2MSP",
			OrgUser:       "User1",
			OrgPeerNum:    2,
			OrgAnchorFile: "/root/go/src/fabric-go-sdk/fixtures/channel-artifacts/allOrg2MSPanchors.tx",
		},
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org3",
			OrgMspId:      "Org3MSP",
			OrgUser:       "User1",
			OrgPeerNum:    2,
			OrgAnchorFile: "/root/go/src/fabric-go-sdk/fixtures/channel-artifacts/allOrg3MSPanchors.tx",
		},
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org4",
			OrgMspId:      "Org4MSP",
			OrgUser:       "User1",
			OrgPeerNum:    2,
			OrgAnchorFile: "/root/go/src/fabric-go-sdk/fixtures/channel-artifacts/allOrg4MSPanchors.tx",
		},
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org5",
			OrgMspId:      "Org5MSP",
			OrgUser:       "User1",
			OrgPeerNum:    2,
			OrgAnchorFile: "/root/go/src/fabric-go-sdk/fixtures/channel-artifacts/allOrg5MSPanchors.tx",
		},
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org6",
			OrgMspId:      "Org6MSP",
			OrgUser:       "User1",
			OrgPeerNum:    2,
			OrgAnchorFile: "/root/go/src/fabric-go-sdk/fixtures/channel-artifacts/allOrg6MSPanchors.tx",
		},
	}

	// init sdk env info
	info := sdkInit.SdkEnvInfo{
		ChannelID:         "allchannel",
		ChannelConfig:     "/root/go/src/fabric-go-sdk/fixtures/channel-artifacts/allchannel.tx",
		Orgs:              orgs,
		OrdererAdminUser:  "Admin",
		OrdererOrgName:    "OrdererOrg",
		OrdererEndpoint:   "orderer0.example.com",
		ChaincodeID1:      cc_name1,
		ChaincodeID2:      cc_name2,
		ChaincodeID3:      cc_name3,
		ChaincodeID4:      cc_name4,
		ChaincodeID5:      cc_name5,
		ChaincodePath:     "/root/go/src/fabric-go-sdk/chaincode/step4",
		ChaincodeVersion1: cc_version1,
		ChaincodeVersion2: cc_version2,
		ChaincodeVersion3: cc_version3,
		ChaincodeVersion4: cc_version4,
		ChaincodeVersion5: cc_version5,
	}

	// sdk setup
	sdk, err := sdkInit.Setup("config.yaml", &info)
	if err != nil {
		fmt.Println(">> SDK setup error:", err)
		os.Exit(-1)
	}

	//// create channel and join
	//if err := sdkInit.CreateAndJoinChannel(&info); err != nil {
	//     fmt.Println(">> Create channel and join error:", err)
	//     os.Exit(-1)
	//}
	//
	// create chaincode lifecycle
	//if err := sdkInit.CreateCCLifecycle(&info, 3, false, sdk); err != nil {
	//fmt.Println(">> create chaincode lifecycle error: %v", err)
	//os.Exit(-1)
	//}

	// invoke chaincode set status
	fmt.Println(">> 通过链码外部服务设置链码状态......")

	if err := info.InitService(info.ChaincodeID1, info.ChaincodeID2, info.ChaincodeID3, info.ChaincodeID4, info.ChaincodeID5, info.ChannelID, info.Orgs[0], sdk); err != nil {

		fmt.Println("InitService successful")
		os.Exit(-1)
	}

	serviceSetup := service.ServiceSetup{
		ChaincodeID1: info.ChaincodeID1,
		ChaincodeID2: info.ChaincodeID2,
		ChaincodeID3: info.ChaincodeID3,
		ChaincodeID4: info.ChaincodeID4,
		ChaincodeID5: info.ChaincodeID5,
		Client:       info.Client,
		SDK:          sdk,
	}

	App = controller.Application{
		Setup: &serviceSetup,
	}

	fmt.Println(">> 设置链码状态完成")

	a := []string{info.ChaincodeID1, "set", "C20210815", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "24", "25", "26"}
	ret, err := App.Setup.Set(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("<--- 添加信息　--->：", ret)

	//b:=[]string{info.ChaincodeID2,"set","C20210816","2","3","4","5","6","7","8","9","10","11","12","13"}
	//ret, err = App.Setup.Set(b)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("<--- 添加信息　--->：", ret)

	a = []string{info.ChaincodeID1, "get", "C20210815"}
	response, err := App.Setup.Get(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("<--- 查询信息　--->：", response)

	//b = []string{info.ChaincodeID2,"get","C20210816"}
	//response, err = App.Setup.Get(b)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("<--- 查询信息　--->：", response)

	web.WebStart(App)
}

