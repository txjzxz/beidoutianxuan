package service

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	shell "github.com/ipfs/go-ipfs-api"
	"gopkg.in/gomail.v2"
	"mime/multipart"
)

func (t *ServiceSetup) Set(args []string) (string, error) {

	eventID := "eventSetNew"
	reg ,_ := regitserEvent(t.Client, args[0],eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	var tempArgs [][]byte
	for i := 2; i < len(args); i++ {
		tempArgs = append(tempArgs, []byte(args[i]))
	}


	request := channel.Request{ChaincodeID: args[0], Fcn: args[1], Args:tempArgs}
	response, err := t.Client.Execute(request)
	if err != nil {
		// 资产转移失败
		return "", err
	}

	//fmt.Println("============== response:",response)

	return string(response.TransactionID), nil
}

func (t *ServiceSetup) Get(args []string) (string, error) {
	response, err := t.Client.Query(channel.Request{ChaincodeID: args[0], Fcn: args[1], Args:[][]byte{[]byte(args[2])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}

func (t *ServiceSetup) UploadFilesOnIpfs(f *multipart.FileHeader) (string, error) {
	sh := shell.NewShell("localhost:5001")
	of,err := f.Open()
	if err != nil {
		fmt.Println(err)
	}
	cid, err := sh.Add(of)
	fmt.Println(cid)
	return cid,err
}

func (t *ServiceSetup) QueryInfo() string {
	ctx := t.SDK.ChannelContext("allchannel", fabsdk.WithOrg("Org1"), fabsdk.WithUser("User1"))
	c, err := ledger.New(ctx)
	if err != nil {
		fmt.Println("failed to create client")
	}
	block, err := c.QueryInfo()
	if err != nil {
		fmt.Println(err)
	}
	str,err := json.Marshal(block)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	return string(str)
}

func (t *ServiceSetup) QueryBlock(n uint64) (string) {
	ctx := t.SDK.ChannelContext("allchannel", fabsdk.WithOrg("Org1"), fabsdk.WithUser("User1"))
	c, err := ledger.New(ctx)
	if err != nil {
		fmt.Println("failed to create client")
	}
	block, err := c.QueryBlock(n)
	if err != nil {
		fmt.Println(err)
	}
	str,err := json.Marshal(block)
	if err != nil {
		fmt.Println(err)
	}
	return string(str)
}

func (t *ServiceSetup) QueryTransaction(n fab.TransactionID) (string) {
	ctx := t.SDK.ChannelContext("allchannel", fabsdk.WithOrg("Org1"), fabsdk.WithUser("User1"))
	c, err := ledger.New(ctx)
	if err != nil {
		fmt.Println("failed to create client")
	}
	block, err := c.QueryTransaction(n)
	if err != nil {
		fmt.Println(err)
	}
	str,err := json.Marshal(block)
	if err != nil {
		fmt.Println(err)
	}
	return string(str)
}

func (t ServiceSetup) SendEmail(step,ID string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "2603206395@qq.com")
	m.SetHeader("To", "421486469@qq.com","1352633214@qq.com","2292815551@qq.com","2506142592@qq.com")
	m.SetHeader("Subject", "信息更新")
	var ostr string
	switch step {
	case "step1":
		ostr = "施工方给混凝土企业的订单信息"
	case "step2":
		ostr = "混凝土企业的生产信息"
	case "step3":
		ostr = "物流方上传的运输信息"
	case "step4":
		ostr = "施工方需要的收货信息"
	case "step5":
		ostr = "施工方上传混凝土浇筑信息"
	}
	str := "<h1>订单编号" + ID + "<h1>\n<h2>" + ostr + "已更新，详细请查看http://42.192.209.60:9000/web?" + "id=" + ID + "&cc=" + step + "<h2>"
	m.SetBody("text/html", str)
	d := gomail.NewDialer("smtp.qq.com", 587, "2603206395@qq.com", "igpuaoelolkhdjdb")
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("DialAndSend err %v:", err)
		panic(err)
	}
	fmt.Printf("send mail success\n")
}
