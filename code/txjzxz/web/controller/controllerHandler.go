package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/skip2/go-qrcode"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

var cuser User

var status Status

var data1 step1Order

var data2 step2

var data3 step3

var data4 step4

var data5 step5

var indexData block

func (app *Application) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"Flag": true,
	})
}

func (app *Application) Invoice(c *gin.Context) {
	c.HTML(http.StatusOK, "invoice.html", cuser)
}

func (app *Application) FormBasic1(c *gin.Context) {
	c.HTML(http.StatusOK, "form_basic1.html", cuser)
}
func (app *Application) FormBasic2(c *gin.Context) {
	c.HTML(http.StatusOK, "form_basic2.html", cuser)
}
func (app *Application) FormBasic3(c *gin.Context) {
	c.HTML(http.StatusOK, "form_basic3.html", cuser)
}
func (app *Application) FormBasic4(c *gin.Context) {
	c.HTML(http.StatusOK, "form_basic4.html", cuser)
}
func (app *Application) FormBasic5(c *gin.Context) {
	c.HTML(http.StatusOK, "form_basic5.html", cuser)
}

func (app *Application) Index1(c *gin.Context) {
	str := app.Setup.QueryInfo()

	fmt.Println(str)

	json.Unmarshal([]byte(str), &indexData)

	str1 := str[41:85]
	str2 := str[108:152]

	fmt.Println(indexData)

	fmt.Println(str1)
	fmt.Println(str2)

	data := &struct {
		BlockNumber  uint64
		Hash         string
		PreviousHash string
		Result       string
		LoginName    string
		IsAdmin      bool
	}{
		BlockNumber:  indexData.BCI.Height,
		Hash:         str1,
		PreviousHash: str2,
		Result:       "查询结果待显示",
		LoginName:    cuser.LoginName,
		IsAdmin:      cuser.IsAdmin,
	}

	c.HTML(http.StatusOK, "index1.html", data)
}

func (app *Application) Index2(c *gin.Context) {
	str := app.Setup.QueryInfo()

	fmt.Println(str)

	json.Unmarshal([]byte(str), &indexData)

	str1 := str[41:85]
	str2 := str[108:152]

	fmt.Println(indexData)

	fmt.Println(str1)
	fmt.Println(str2)

	data := &struct {
		BlockNumber  uint64
		Hash         string
		PreviousHash string
		Result       string
		LoginName    string
		IsAdmin      bool
	}{
		BlockNumber:  indexData.BCI.Height,
		Hash:         str1,
		PreviousHash: str2,
		Result:       "查询结果待显示",
		LoginName:    cuser.LoginName,
		IsAdmin:      cuser.IsAdmin,
	}
	fmt.Println(data)
	c.HTML(http.StatusOK, "index2.html", data)
}

func (app *Application) Index3(c *gin.Context) {
	str := app.Setup.QueryInfo()

	fmt.Println(str)

	json.Unmarshal([]byte(str), &indexData)

	str1 := str[41:85]
	str2 := str[108:152]

	fmt.Println(indexData)

	fmt.Println(str1)
	fmt.Println(str2)

	data := &struct {
		BlockNumber  uint64
		Hash         string
		PreviousHash string
		Result       string
		LoginName    string
		IsAdmin      bool
	}{
		BlockNumber:  indexData.BCI.Height,
		Hash:         str1,
		PreviousHash: str2,
		Result:       "查询结果待显示",
		LoginName:    cuser.LoginName,
		IsAdmin:      cuser.IsAdmin,
	}
	fmt.Println(data)
	c.HTML(http.StatusOK, "index3.html", data)
}

func (app *Application) Index4(c *gin.Context) {
	str := app.Setup.QueryInfo()

	fmt.Println(str)

	json.Unmarshal([]byte(str), &indexData)

	str1 := str[41:85]
	str2 := str[108:152]

	fmt.Println(indexData)

	fmt.Println(str1)
	fmt.Println(str2)

	data := &struct {
		BlockNumber  uint64
		Hash         string
		PreviousHash string
		Result       string
		LoginName    string
		IsAdmin      bool
	}{
		BlockNumber:  indexData.BCI.Height,
		Hash:         str1,
		PreviousHash: str2,
		Result:       "查询结果待显示",
		LoginName:    cuser.LoginName,
		IsAdmin:      cuser.IsAdmin,
	}
	fmt.Println(data)
	c.HTML(http.StatusOK, "index4.html", data)
}

func (app *Application) QueryInvoice1(c *gin.Context) {
	c.HTML(http.StatusOK, "queryInvoice1.html", cuser)
}
func (app *Application) QueryInvoice2(c *gin.Context) {
	c.HTML(http.StatusOK, "queryInvoice2.html", cuser)
}
func (app *Application) QueryInvoice3(c *gin.Context) {
	c.HTML(http.StatusOK, "queryInvoice3.html", cuser)
}
func (app *Application) QueryInvoice4(c *gin.Context) {
	c.HTML(http.StatusOK, "queryInvoice4.html", cuser)
}
func (app *Application) InvoicePrint(c *gin.Context) {
	c.HTML(http.StatusOK, "invoice_print.html", cuser)
}

func (app *Application) File1(c *gin.Context) {
	c.HTML(http.StatusOK, "queryFile1.html", cuser)
}
func (app *Application) File2(c *gin.Context) {
	c.HTML(http.StatusOK, "queryFile2.html", cuser)
}
func (app *Application) File3(c *gin.Context) {
	c.HTML(http.StatusOK, "queryFile3.html", cuser)
}
func (app *Application) File4(c *gin.Context) {
	c.HTML(http.StatusOK, "queryFile4.html", cuser)
}

func (app *Application) Select1(c *gin.Context) {
	org := c.PostForm("org")

	str := app.Setup.QueryInfo()

	fmt.Println(str)

	json.Unmarshal([]byte(str), &indexData)

	str1 := str[41:85]
	str2 := str[108:152]

	fmt.Println(indexData)

	fmt.Println(str1)
	fmt.Println(str2)

	data := &struct {
		BlockNumber  uint64
		Hash         string
		PreviousHash string
		Result       string
		LoginName    string
		IsAdmin      bool
	}{
		BlockNumber:  indexData.BCI.Height,
		Hash:         str1,
		PreviousHash: str2,
		Result:       "查询结果待显示",
		LoginName:    cuser.LoginName,
		IsAdmin:      cuser.IsAdmin,
	}
	fmt.Println(data)

	switch org {
	case "1":
		c.HTML(http.StatusOK, "index4.html", data)
	case "2":
		c.HTML(http.StatusOK, "index1.html", data)
	case "6":
		c.HTML(http.StatusOK, "index3.html", data)
	}
}

func (app *Application) Select2(c *gin.Context) {
	org := c.PostForm("org")

	str := app.Setup.QueryInfo()

	fmt.Println(str)

	json.Unmarshal([]byte(str), &indexData)

	str1 := str[41:85]
	str2 := str[108:152]

	fmt.Println(indexData)

	fmt.Println(str1)
	fmt.Println(str2)

	data := &struct {
		BlockNumber  uint64
		Hash         string
		PreviousHash string
		Result       string
		LoginName    string
		IsAdmin      bool
	}{
		BlockNumber:  indexData.BCI.Height,
		Hash:         str1,
		PreviousHash: str2,
		Result:       "查询结果待显示",
		LoginName:    cuser.LoginName,
		IsAdmin:      cuser.IsAdmin,
	}
	fmt.Println(data)

	switch org {
	case "1":
		c.HTML(http.StatusOK, "index4.html", data)
	case "2":
		c.HTML(http.StatusOK, "index1.html", data)
	case "6":
		c.HTML(http.StatusOK, "index3.html", data)
	}
}

func (app *Application) Select3(c *gin.Context) {
	org := c.PostForm("org")

	str := app.Setup.QueryInfo()

	fmt.Println(str)

	json.Unmarshal([]byte(str), &indexData)

	str1 := str[41:85]
	str2 := str[108:152]

	fmt.Println(indexData)

	fmt.Println(str1)
	fmt.Println(str2)

	data := &struct {
		BlockNumber  uint64
		Hash         string
		PreviousHash string
		Result       string
		LoginName    string
		IsAdmin      bool
	}{
		BlockNumber:  indexData.BCI.Height,
		Hash:         str1,
		PreviousHash: str2,
		Result:       "查询结果待显示",
		LoginName:    cuser.LoginName,
		IsAdmin:      cuser.IsAdmin,
	}
	fmt.Println(data)

	switch org {
	case "1":
		c.HTML(http.StatusOK, "index4.html", data)
	case "2":
		c.HTML(http.StatusOK, "index1.html", data)
	case "6":
		c.HTML(http.StatusOK, "index3.html", data)
	}
}

func (app *Application) CheckLogin(c *gin.Context) {
	name := c.PostForm("username")
	pw := c.PostForm("pw")
	value := c.PostForm("select")

	var flag bool
	for _, user := range users {
		if user.LoginName == name && user.Password == pw && user.Org == value {
			cuser = user
			flag = true
			break
		}
	}

	str := app.Setup.QueryInfo()

	fmt.Println(str)

	json.Unmarshal([]byte(str), &indexData)

	str1 := str[41:85]
	str2 := str[108:152]

	fmt.Println(indexData)

	fmt.Println(str1)
	fmt.Println(str2)

	data := &struct {
		CurrentUser  User
		IsAdmin      bool
		LoginName    string
		Value        string
		Flag         bool
		BlockNumber  uint64
		Hash         string
		PreviousHash string
		Result       string
	}{
		CurrentUser:  cuser,
		IsAdmin:      cuser.IsAdmin,
		LoginName:    cuser.LoginName,
		Value:        cuser.Org,
		Flag:         flag,
		BlockNumber:  indexData.BCI.Height,
		Hash:         str1,
		PreviousHash: str2,
		Result:       "查询结果待显示",
	}

	if flag {
		if value == "1" {
			c.HTML(http.StatusOK, "select.html", data)
		} else if value == "2" {
			c.HTML(http.StatusOK, "select.html", data)
		} else if value == "3" {
			c.HTML(http.StatusOK, "index2.html", data)
		} else if value == "4" {
			c.HTML(http.StatusOK, "index2.html", data)
		} else if value == "5" {
			c.HTML(http.StatusOK, "index2.html", data)
		} else if value == "6" {
			c.HTML(http.StatusOK, "select.html", data)
		}
	} else {
		data.Flag = false
		data.CurrentUser.LoginName = name
		c.HTML(http.StatusPreconditionFailed, "login.html", data)
	}

}

func (app *Application) Upload1(c *gin.Context) {
	id := c.PostForm("id")
	date := c.PostForm("date")
	projectName := c.PostForm("projectName")
	projectLocation := c.PostForm("projectLocation")
	nameOfConstruction := c.PostForm("nameOfConstruction")
	nameOfSupplier := c.PostForm("nameOfSupplier")
	nameOfSupervision := c.PostForm("nameOfSupervision")
	time1 := c.PostForm("time1")
	time2 := c.PostForm("time2")
	estimatedTime := time1 + time2
	pouringPosition := c.PostForm("pouringPosition")
	StrengthGradeOfConcrete := c.PostForm("StrengthGradeOfConcrete")
	volume := c.PostForm("volume")
	pumpingMethod := c.PostForm("pumpingMethod")
	slump := c.PostForm("slump")
	performance := c.PostForm("performance")

	var hash string
	f, err := c.FormFile("f1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		res, err := app.Setup.UploadFilesOnIpfs(f)
		if err != nil {
			fmt.Println("upload file failed,err:%v", err)
		}
		hash = res
		fmt.Println("hash: " + hash)
	}

	name := c.PostForm("name")
	//password := c.PostForm("password")

	a := []string{"step1", "set", id, date, projectName, projectLocation, nameOfConstruction, nameOfSupplier, nameOfSupervision, estimatedTime,
		pouringPosition, StrengthGradeOfConcrete, volume, pumpingMethod, slump, performance, hash, name}
	ret, err := app.Setup.Set(a)
	if err != nil {
		fmt.Println(err)
	}

	data := &struct {
		Id                      string
		Date                    string
		ProjectName             string
		ProjectLocation         string
		NameOfConstruction      string
		NameOfSupplier          string
		NameOfSupervision       string
		EstimatedTime           string
		PouringPosition         string
		StrengthGradeOfConcrete string
		Volume                  string
		PumpingMethod           string
		Slump                   string
		Performance             string
		Hash                    string
		Name                    string
		LoginName               string
		IsAdmin                 bool
	}{
		Id:                      id,
		Date:                    date,
		ProjectName:             projectName,
		ProjectLocation:         projectLocation,
		NameOfConstruction:      nameOfConstruction,
		NameOfSupplier:          nameOfSupplier,
		NameOfSupervision:       nameOfSupervision,
		EstimatedTime:           estimatedTime,
		PouringPosition:         pouringPosition,
		StrengthGradeOfConcrete: StrengthGradeOfConcrete,
		Volume:                  volume,
		PumpingMethod:           pumpingMethod,
		Slump:                   slump,
		Performance:             performance,
		Hash:                    hash,
		Name:                    name,
		LoginName:               cuser.LoginName,
		IsAdmin:                 cuser.IsAdmin,
	}

	url := `http://42.192.209.60:9000/web?id=` + id + `&cc=step1`
	//png, err := json.Marshal(url)
	//if err != nil {
	//	fmt.Println(err)
	//}
	qrcodename := id + `step1`
	qrcode.WriteFile(url, qrcode.Medium, 256, "./QRcode/"+qrcodename+".png")

	fmt.Println("<--- 添加信息　--->：", ret)
	app.Setup.SendEmail("step1", id)
	c.HTML(http.StatusOK, "invoice1.html", data)
}

func (app *Application) Upload2(c *gin.Context) {
	id := c.PostForm("id")
	date := c.PostForm("date")
	projectName := c.PostForm("projectName")
	projectLocation := c.PostForm("projectLocation")
	nameOfConstruction := c.PostForm("nameOfConstruction")
	time1 := c.PostForm("time1")
	time2 := c.PostForm("time2")
	estimatedTime := time1 + time2
	pouringPosition := c.PostForm("pouringPosition")
	StrengthGradeOfConcrete := c.PostForm("StrengthGradeOfConcrete")
	volume := c.PostForm("volume")
	pumpingMethod := c.PostForm("pumpingMethod")
	slump := c.PostForm("slump")
	name := c.PostForm("name")

	var hash1 string
	var hash2 string
	var hash3 string

	f1, err := c.FormFile("f1")
	if err != nil {
		fmt.Println(err)
	}
	f2, err := c.FormFile("f2")
	if err != nil {
		fmt.Println(err)
	}
	f3, err := c.FormFile("f3")
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		res1, err := app.Setup.UploadFilesOnIpfs(f1)
		if err != nil {
			fmt.Printf("upload file failed,err:%v\n", err)
		}
		res2, err := app.Setup.UploadFilesOnIpfs(f2)
		if err != nil {
			fmt.Printf("upload file failed,err:%v\n", err)
		}
		res3, err := app.Setup.UploadFilesOnIpfs(f3)
		if err != nil {
			fmt.Printf("upload file failed,err:%v\n", err)
		}
		hash1 = res1
		hash2 = res2
		hash3 = res3
	}

	a := []string{"step2", "set", id, date, estimatedTime, projectName, projectLocation, nameOfConstruction, pouringPosition,
		StrengthGradeOfConcrete, volume, pumpingMethod, slump, hash1, hash2, hash3, name}

	ret, err := app.Setup.Set(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("<--- 添加信息　--->：", ret)

	data := &struct {
		Id                      string
		Date                    string
		ProjectName             string
		ProjectLocation         string
		NameOfConstruction      string
		EstimatedTime           string
		PouringPosition         string
		StrengthGradeOfConcrete string
		Volume                  string
		PumpingMethod           string
		Slump                   string
		Hash1                   string
		Hash2                   string
		Hash3                   string
		Name                    string
		LoginName               string
		IsAdmin                 bool
	}{
		Id:                      id,
		Date:                    date,
		ProjectName:             projectName,
		ProjectLocation:         projectLocation,
		NameOfConstruction:      nameOfConstruction,
		EstimatedTime:           estimatedTime,
		PouringPosition:         pouringPosition,
		StrengthGradeOfConcrete: StrengthGradeOfConcrete,
		Volume:                  volume,
		PumpingMethod:           pumpingMethod,
		Slump:                   slump,
		Hash1:                   hash1,
		Hash2:                   hash2,
		Hash3:                   hash3,
		Name:                    name,
		LoginName:               cuser.LoginName,
		IsAdmin:                 cuser.IsAdmin,
	}

	status.Order = true
	url := `http://42.192.209.60:9000/web?id=` + id + `&cc=step2`
	//png, err := json.Marshal(url)
	//if err != nil {
	//	fmt.Println(err)
	//}
	qrcodename := id + `step2`
	qrcode.WriteFile(url, qrcode.Medium, 256, "./QRcode/"+qrcodename+".png")
	app.Setup.SendEmail("step1", id)

	c.HTML(http.StatusOK, "invoice2.html", data)
}

func (app *Application) Upload3(c *gin.Context) {
	id := c.PostForm("id")
	ptime1 := c.PostForm("ptime1")
	ptime2 := c.PostForm("ptime2")
	produceTime := ptime1 + ptime2
	projectName := c.PostForm("projectName")
	projectLocation := c.PostForm("projectLocation")
	nameOfConstruction := c.PostForm("nameOfConstruction")
	time1 := c.PostForm("time1")
	time2 := c.PostForm("time2")
	estimatedTime := time1 + time2
	pouringPosition := c.PostForm("pouringPosition")
	StrengthGradeOfConcrete := c.PostForm("StrengthGradeOfConcrete")
	theCarVolume := c.PostForm("theCarVolume")
	totalVolume := c.PostForm("totalVolume")
	planVolume := c.PostForm("planVolume")
	distance := c.PostForm("distance")
	pumpingMethod := c.PostForm("pumpingMethod")
	slump := c.PostForm("slump")
	name := c.PostForm("name")

	var hash string
	f1, err := c.FormFile("f1")
	if err != nil {
		fmt.Println(err)
	}
	res1, err := app.Setup.UploadFilesOnIpfs(f1)
	if err != nil {
		fmt.Printf("upload file failed,err:%v\n", err)
	}
	hash = res1

	a := []string{"step3", "set", id, produceTime, estimatedTime, projectName, projectLocation, nameOfConstruction, pouringPosition,
		StrengthGradeOfConcrete, theCarVolume, totalVolume, planVolume, pumpingMethod, slump, distance, hash, name}

	ret, err := app.Setup.Set(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("<--- 添加信息　--->：", ret)

	data := &struct {
		Id                      string
		ProductionTime          string
		ProjectName             string
		ProjectLocation         string
		NameOfConstruction      string
		EstimatedTime           string
		PouringPosition         string
		StrengthGradeOfConcrete string
		CarVolume               string
		SumVolume               string
		PlanVolume              string
		PumpingMethod           string
		Slump                   string
		Distance                string
		Hash                    string
		Name                    string
		LoginName               string
		IsAdmin                 bool
	}{
		Id:                      id,
		ProductionTime:          produceTime,
		ProjectName:             projectName,
		ProjectLocation:         projectLocation,
		NameOfConstruction:      nameOfConstruction,
		EstimatedTime:           estimatedTime,
		PouringPosition:         pouringPosition,
		StrengthGradeOfConcrete: StrengthGradeOfConcrete,
		CarVolume:               theCarVolume,
		SumVolume:               totalVolume,
		PlanVolume:              planVolume,
		PumpingMethod:           pumpingMethod,
		Slump:                   slump,
		Distance:                distance,
		Hash:                    hash,
		Name:                    name,
		LoginName:               cuser.LoginName,
		IsAdmin:                 cuser.IsAdmin,
	}

	status.Order = true
	url := `http://42.192.209.60:9000/web?id=` + id + `&cc=step3`
	//png, err := json.Marshal(url)
	//if err != nil {
	//	fmt.Println(err)
	//}
	qrcodename := id + `step3`
	qrcode.WriteFile(url, qrcode.Medium, 256, "./QRcode/"+qrcodename+".png")
	app.Setup.SendEmail("step3", id)

	c.HTML(http.StatusOK, "invoice3.html", data)
}

func (app *Application) Upload4(c *gin.Context) {
	id := c.PostForm("id")
	ptime1 := c.PostForm("ptime1")
	ptime2 := c.PostForm("ptime2")
	produceTime := ptime1 + ptime2
	time1 := c.PostForm("atime1")
	time2 := c.PostForm("atime2")
	estimatedTime := time1 + time2
	projectName := c.PostForm("projectName")
	projectLocation := c.PostForm("projectLocation")
	nameOfConstruction := c.PostForm("nameOfConstruction")
	nameOfSupplier := c.PostForm("nameOfSupplier")
	nameOfSupervision := c.PostForm("nameOfSupervision")
	pouringPosition := c.PostForm("pouringPosition")
	StrengthGradeOfConcrete := c.PostForm("StrengthGradeOfConcrete")
	volume := c.PostForm("volume")
	pumpingMethod := c.PostForm("pumpingMethod")
	slump := c.PostForm("slump")
	name := c.PostForm("name")

	var hash string
	f1, err := c.FormFile("f1")
	if err != nil {
		fmt.Println(err)
	}
	res1, err := app.Setup.UploadFilesOnIpfs(f1)
	if err != nil {
		fmt.Printf("upload file failed,err:%v\n", err)
	}
	hash = res1

	check := Check{
		IsQualified: true,
		Reason:      "",
	}

	b := []string{"step1", "get", id}
	response, err := app.Setup.Get(b)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(response), &data1)
	if err != nil {
		fmt.Println(err)
	}

	str := data1.OrderData.ConcreteSlump
	fmt.Println(str)
	reg1 := regexp.MustCompile(`[0-9]+`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}
	fmt.Println(reg1)

	result1 := reg1.FindAllStringSubmatch(str, -1)
	fmt.Println(result1)
	number1, err := strconv.Atoi(result1[0][0])
	if err != nil {
		fmt.Println(err)
	}
	number2, err := strconv.Atoi(result1[1][0])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(number1)
	fmt.Println(number2)

	actualNumber, err := strconv.Atoi(slump)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(actualNumber)

	if actualNumber < number1 {
		check.IsQualified = false
		check.Reason = "false,Reason: The slump is too low!"
	} else if actualNumber > number2 {
		check.IsQualified = false
		check.Reason = "false,Reason: The slump is too high!"
	} else {
		check.IsQualified = true
		check.Reason = "true,Reason: Qualified!"
	}

	a := []string{"step4", "set", id, produceTime, estimatedTime, projectName, projectLocation, nameOfConstruction, nameOfSupplier, nameOfSupervision, pouringPosition,
		StrengthGradeOfConcrete, volume, pumpingMethod, slump, hash, check.Reason, name}

	ret, err := app.Setup.Set(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("<--- 添加信息　--->：", ret)

	data := &struct {
		Id                      string
		ProductionTime          string
		EstimatedTime           string
		OrderID                 string
		ProjectName             string
		ProjectLocation         string
		NameOfConstruction      string
		NameOfSupplier          string
		NameOfSupervision       string
		PouringPosition         string
		StrengthGradeOfConcrete string
		Volume                  string
		PumpingMethod           string
		Slump                   string
		Hash                    string
		Name                    string
		IsQualified             bool
		Reason                  string
		LoginName               string
		IsAdmin                 bool
	}{
		Id:                      id,
		ProductionTime:          produceTime,
		EstimatedTime:           estimatedTime,
		ProjectName:             projectName,
		ProjectLocation:         projectLocation,
		NameOfConstruction:      nameOfConstruction,
		NameOfSupplier:          nameOfSupplier,
		NameOfSupervision:       nameOfSupervision,
		PouringPosition:         pouringPosition,
		StrengthGradeOfConcrete: StrengthGradeOfConcrete,
		Volume:                  volume,
		PumpingMethod:           pumpingMethod,
		Slump:                   slump,
		Hash:                    hash,
		Name:                    name,
		IsQualified:             check.IsQualified,
		Reason:                  check.Reason,
		LoginName:               cuser.LoginName,
		IsAdmin:                 cuser.IsAdmin,
	}

	url := `http://42.192.209.60:9000/web?id=` + id + `&cc=step4`
	//png, err := json.Marshal(url)
	//if err != nil {
	//	fmt.Println(err)
	//}
	qrcodename := id + `step4`
	qrcode.WriteFile(url, qrcode.Medium, 256, "./QRcode/"+qrcodename+".png")
	app.Setup.SendEmail("step4", id)

	c.HTML(http.StatusOK, "invoice4.html", data)
}

func (app *Application) Upload5(c *gin.Context) {
	id := c.PostForm("id")
	dtime1 := c.PostForm("dtime1")
	dtime2 := c.PostForm("dtime2")
	dischargeTime := dtime1 + dtime2
	weather := c.PostForm("weather")
	temperature := c.PostForm("temperature")
	projectName := c.PostForm("projectName")
	projectLocation := c.PostForm("projectLocation")
	nameOfConstruction := c.PostForm("nameOfConstruction")
	nameOfSupplier := c.PostForm("nameOfSupplier")
	nameOfSupervision := c.PostForm("nameOfSupervision")
	pouringPosition := c.PostForm("pouringPosition")
	StrengthGradeOfConcrete := c.PostForm("StrengthGradeOfConcrete")
	volume := c.PostForm("volume")
	pumpingMethod := c.PostForm("pumpingMethod")
	name := c.PostForm("name")

	var hash1 string
	f1, err := c.FormFile("f1")
	if err != nil {
		fmt.Println(err)
	}
	res1, err := app.Setup.UploadFilesOnIpfs(f1)
	if err != nil {
		fmt.Printf("upload file failed,err:%v\n", err)
	}
	hash1 = res1

	var hash2 string
	f2, err := c.FormFile("f2")
	if err != nil {
		fmt.Println(err)
	}
	res2, err := app.Setup.UploadFilesOnIpfs(f2)
	if err != nil {
		fmt.Printf("upload file failed,err:%v\n", err)
	}
	hash2 = res2

	a := []string{"step5", "set", id, dischargeTime, weather, temperature, projectName, projectLocation, nameOfConstruction, nameOfSupplier, nameOfSupervision, pouringPosition,
		StrengthGradeOfConcrete, volume, pumpingMethod, hash1, hash2, name}

	ret, err := app.Setup.Set(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("<--- 添加信息　--->：", ret)

	data := &struct {
		Id                      string
		DischargeTime           string
		Weather                 string
		Temperature             string
		ProjectName             string
		ProjectLocation         string
		NameOfConstruction      string
		NameOfSupplier          string
		NameOfSupervision       string
		PouringPosition         string
		StrengthGradeOfConcrete string
		Volume                  string
		PumpingMethod           string
		Hash1                   string
		Hash2                   string
		Name                    string
		LoginName               string
		IsAdmin                 bool
	}{
		Id:                      id,
		DischargeTime:           dischargeTime,
		Weather:                 weather,
		Temperature:             temperature,
		ProjectName:             projectName,
		ProjectLocation:         projectLocation,
		NameOfConstruction:      nameOfConstruction,
		NameOfSupplier:          nameOfSupplier,
		NameOfSupervision:       nameOfSupervision,
		PouringPosition:         pouringPosition,
		StrengthGradeOfConcrete: StrengthGradeOfConcrete,
		Volume:                  volume,
		PumpingMethod:           pumpingMethod,
		Hash1:                   hash1,
		Hash2:                   hash2,
		Name:                    name,
		LoginName:               cuser.LoginName,
		IsAdmin:                 cuser.IsAdmin,
	}

	url := `http://42.192.209.60:9000/web?id=` + id + `&cc=step5`
	//png, err := json.Marshal(url)
	//if err != nil {
	//	fmt.Println(err)
	//}
	qrcodename := id + `step5`
	qrcode.WriteFile(url, qrcode.Medium, 256, "./QRcode/"+qrcodename+".png")
	app.Setup.SendEmail("step5", id)

	c.HTML(http.StatusOK, "invoice5.html", data)
}

func (app *Application) GetQRcode(c *gin.Context) {
	filename := c.PostForm("id") + c.PostForm("cc") + ".png"
	ContentType := "image/png"
	if err != nil {
		fmt.Println("read fail", err)
	}
	c.Data(http.StatusOK, ContentType, data)

}

func (app *Application) Query(c *gin.Context) {
	id := c.PostForm("id")

	var uploadIndex int

	uploadIndex = 0

	reStep1 := []string{"step1", "get", id}
	reStep2 := []string{"step2", "get", id}
	reStep3 := []string{"step3", "get", id}
	reStep4 := []string{"step4", "get", id}
	reStep5 := []string{"step5", "get", id}

	step1Response, err := app.Setup.Get(reStep1)
	if err == nil {
		uploadIndex = 1
	}
	err = json.Unmarshal([]byte(step1Response), &data1)
	if err != nil {
		fmt.Println(err)
	}

	step2Response, err := app.Setup.Get(reStep2)
	if err == nil {
		uploadIndex = 2
	}
	err = json.Unmarshal([]byte(step2Response), &data2)
	if err != nil {
		fmt.Println(err)
	}

	step3Response, err := app.Setup.Get(reStep3)
	if err == nil {
		uploadIndex = 3
	}
	err = json.Unmarshal([]byte(step3Response), &data3)
	if err != nil {
		fmt.Println(err)
	}

	var isQualified bool
	isQualified = true
	step4Response, err := app.Setup.Get(reStep4)
	if err == nil {
		uploadIndex = 4
		if data4.OrderData.IsQualified == "true" {
			isQualified = true
		} else {
			isQualified = false
		}
	}
	err = json.Unmarshal([]byte(step4Response), &data4)
	if err != nil {
		fmt.Println(err)
	}

	step5Response, err := app.Setup.Get(reStep5)
	if err == nil {
		uploadIndex = 5
	}
	err = json.Unmarshal([]byte(step5Response), &data5)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(uploadIndex)

	anotherData := &struct {
		/*   此部分为查询信息需要返回的数据项     */
		Id                      string // step1
		Date                    string
		ProjectName             string
		ProjectLocation         string
		NameOfConstruction      string
		NameOfSupplier          string
		NameOfSupervision       string
		EstimatedTime           string
		PouringPosition         string
		StrengthGradeOfConcrete string
		Volume                  string
		PumpingMethod           string
		Slump                   string
		Performance             string
		Step1Hash1              string
		Step1Name               string
		Step2Hash1              string // step2
		Step2Hash2              string
		Step2Hash3              string
		Step2Name               string
		ProductionTime          string // step3
		CarVolume               string
		SumVolume               string
		Distance                string
		Step3Hash1              string
		Step3Name               string
		ArriveTime              string // step4
		ActualSlump             string
		Step4Hash1              string
		IsQualified             bool
		Step4Name               string
		DischargeTime           string // step5
		Weather                 string
		Temperature             string
		Step5Hash1              string
		Step5Hash2              string
		Step5Name               string
		LoginName               string // ID信息
		IsAdmin                 bool
		UploadIndex             int
	}{
		Id:                      data1.Id, // step1
		Date:                    data1.OrderData.Date,
		ProjectName:             data1.OrderData.ProjectName,
		ProjectLocation:         data1.OrderData.ProjectLocation,
		NameOfConstruction:      data1.OrderData.ConstructionOrg,
		NameOfSupplier:          data1.OrderData.SupplyOrg,
		NameOfSupervision:       data1.OrderData.SupervisorOrg,
		EstimatedTime:           data1.OrderData.ScheduledTime,
		PouringPosition:         data1.OrderData.PouringPosition,
		StrengthGradeOfConcrete: data1.OrderData.GradesOfConcrete,
		Volume:                  data1.OrderData.ConcreteVolume,
		PumpingMethod:           data1.OrderData.ConcretePumpingMethod,
		Slump:                   data1.OrderData.ConcreteSlump,
		Performance:             data1.OrderData.ConcretePerformance,
		Step1Hash1:              data1.OrderData.Hash,
		Step1Name:               data1.OrderData.Principal,
		Step2Hash1:              data2.OrderData.Hash1, // step2
		Step2Hash2:              data2.OrderData.Hash2,
		Step2Hash3:              data2.OrderData.Hash3,
		Step2Name:               data2.OrderData.Principal,
		ProductionTime:          data3.OrderData.ProductionTime, // step3
		CarVolume:               data3.OrderData.TheCarSupplyVolume,
		SumVolume:               data3.OrderData.SumVolume,
		Distance:                data3.OrderData.SupplyDistance,
		Step3Hash1:              data3.OrderData.Hash,
		Step3Name:               data3.OrderData.Yardman,
		ArriveTime:              data4.OrderData.ArriveTime, // step4
		ActualSlump:             data4.OrderData.ConcreteActualSlump,
		Step4Hash1:              data4.OrderData.Hash,
		IsQualified:             isQualified,
		Step4Name:               data4.OrderData.ConstructionPrincipal,
		DischargeTime:           data5.OrderData.ConcreteDischargeTime, // step5
		Weather:                 data5.OrderData.Weather,
		Temperature:             data5.OrderData.Temperature,
		Step5Hash1:              data5.OrderData.Hash1,
		Step5Hash2:              data5.OrderData.Hash2,
		Step5Name:               data5.OrderData.ConstructionPrincipal,
		LoginName:               cuser.LoginName,
		IsAdmin:                 cuser.IsAdmin,
		UploadIndex:             uploadIndex,
	}
	fmt.Println("<--- 查询信息　--->：", anotherData)

	c.HTML(http.StatusOK, "invoiceAll.html", anotherData)
}

func (app *Application) QueryWeb(c *gin.Context) {
	id := c.Query("id")
	cc := c.Query("cc")
	a := []string{cc, "get", id}
	response, err := app.Setup.Get(a)
	fmt.Println(response)
	if err != nil {
		fmt.Println(err)
		c.HTML(404, "404.html", nil)
		return
	}

	switch cc {
	case "step1":
		err = json.Unmarshal([]byte(response), &data1)
		if err != nil {
			fmt.Println(err)
		}
		anotherData := &struct {
			Id                      string
			Date                    string
			ProjectName             string
			ProjectLocation         string
			NameOfConstruction      string
			NameOfSupplier          string
			NameOfSupervision       string
			EstimatedTime           string
			PouringPosition         string
			StrengthGradeOfConcrete string
			Volume                  string
			PumpingMethod           string
			Slump                   string
			Performance             string
			Hash                    string
			Name                    string
			LoginName               string
			IsAdmin                 bool
		}{
			Id:                      data1.Id,
			Date:                    data1.OrderData.Date,
			ProjectName:             data1.OrderData.ProjectName,
			ProjectLocation:         data1.OrderData.ProjectLocation,
			NameOfConstruction:      data1.OrderData.ConstructionOrg,
			NameOfSupplier:          data1.OrderData.SupplyOrg,
			NameOfSupervision:       data1.OrderData.SupervisorOrg,
			EstimatedTime:           data1.OrderData.ScheduledTime,
			PouringPosition:         data1.OrderData.PouringPosition,
			StrengthGradeOfConcrete: data1.OrderData.GradesOfConcrete,
			Volume:                  data1.OrderData.ConcreteVolume,
			PumpingMethod:           data1.OrderData.ConcretePumpingMethod,
			Slump:                   data1.OrderData.ConcreteSlump,
			Performance:             data1.OrderData.ConcretePerformance,
			Hash:                    data1.OrderData.Hash,
			Name:                    data1.OrderData.Principal,
			LoginName:               cuser.LoginName,
			IsAdmin:                 cuser.IsAdmin,
		}
		fmt.Println("<--- 查询信息　--->：", anotherData)
		c.HTML(http.StatusOK, "invoice1.html", anotherData)
	case "step2":
		err = json.Unmarshal([]byte(response), &data2)
		fmt.Println(data2)
		if err != nil {
			fmt.Println(err)
		}
		anotherData := &struct {
			Id                      string
			Date                    string
			ProjectName             string
			ProjectLocation         string
			NameOfConstruction      string
			EstimatedTime           string
			PouringPosition         string
			StrengthGradeOfConcrete string
			Volume                  string
			PumpingMethod           string
			Slump                   string
			Hash1                   string
			Hash2                   string
			Hash3                   string
			Name                    string
			LoginName               string
			IsAdmin                 bool
		}{
			Id:                      data2.Id,
			Date:                    data2.OrderData.Date,
			ProjectName:             data2.OrderData.ProjectName,
			ProjectLocation:         data2.OrderData.ProjectLocation,
			NameOfConstruction:      data2.OrderData.ConstructionOrg,
			EstimatedTime:           data2.OrderData.ScheduledTime,
			PouringPosition:         data2.OrderData.PouringPosition,
			StrengthGradeOfConcrete: data2.OrderData.GradesOfConcrete,
			Volume:                  data2.OrderData.ConcreteVolume,
			PumpingMethod:           data2.OrderData.ConcretePumpingMethod,
			Slump:                   data2.OrderData.ConcreteSlump,
			Hash1:                   data2.OrderData.Hash1,
			Hash2:                   data2.OrderData.Hash2,
			Hash3:                   data2.OrderData.Hash3,
			Name:                    data2.OrderData.Principal,
			LoginName:               cuser.LoginName,
			IsAdmin:                 cuser.IsAdmin,
		}
		fmt.Println("<--- 查询信息　--->：", anotherData)
		c.HTML(http.StatusOK, "invoice2.html", anotherData)
	case "step3":
		err = json.Unmarshal([]byte(response), &data3)
		if err != nil {
			fmt.Println(err)
		}
		anotherData := &struct {
			Id                      string
			ProductionTime          string
			ProjectName             string
			ProjectLocation         string
			NameOfConstruction      string
			EstimatedTime           string
			PouringPosition         string
			StrengthGradeOfConcrete string
			CarVolume               string
			SumVolume               string
			PlanVolume              string
			PumpingMethod           string
			Slump                   string
			Distance                string
			Hash                    string
			Name                    string
			LoginName               string
			IsAdmin                 bool
		}{
			Id:                      data3.Id,
			ProductionTime:          data3.OrderData.ProductionTime,
			ProjectName:             data3.OrderData.ProjectName,
			ProjectLocation:         data3.OrderData.ProjectLocation,
			NameOfConstruction:      data3.OrderData.ConstructionOrg,
			EstimatedTime:           data3.OrderData.ScheduledTime,
			PouringPosition:         data3.OrderData.PouringPosition,
			StrengthGradeOfConcrete: data3.OrderData.GradesOfConcrete,
			CarVolume:               data3.OrderData.TheCarSupplyVolume,
			SumVolume:               data3.OrderData.SumVolume,
			PlanVolume:              data3.OrderData.PlanVolume,
			PumpingMethod:           data3.OrderData.ConcretePumpingMethod,
			Slump:                   data3.OrderData.ConcreteSlump,
			Distance:                data3.OrderData.SupplyDistance,
			Hash:                    data3.OrderData.Hash,
			Name:                    data3.OrderData.Yardman,
			LoginName:               cuser.LoginName,
			IsAdmin:                 cuser.IsAdmin,
		}
		fmt.Println("<--- 查询信息　--->：", anotherData)
		c.HTML(http.StatusOK, "invoice3.html", anotherData)
	case "step4":
		err = json.Unmarshal([]byte(response), &data4)
		if err != nil {
			fmt.Println(err)
		}
		var IsQualified bool
		str := data4.OrderData.IsQualified[0:5]
		if str == "false" {
			IsQualified = false
		} else {
			IsQualified = true
		}

		anotherData := &struct {
			Id                      string
			ProductionTime          string
			EstimatedTime           string
			OrderID                 string
			ProjectName             string
			ProjectLocation         string
			NameOfConstruction      string
			NameOfSupplier          string
			NameOfSupervision       string
			PouringPosition         string
			StrengthGradeOfConcrete string
			Volume                  string
			PumpingMethod           string
			Slump                   string
			Hash                    string
			Name                    string
			IsQualified             bool
			Reason                  string
			LoginName               string
			IsAdmin                 bool
		}{
			Id:                      data4.ArriveId,
			ProductionTime:          data4.OrderData.ProductionTime,
			EstimatedTime:           data4.OrderData.ArriveTime,
			OrderID:                 data4.ArriveId,
			ProjectName:             data4.OrderData.ProjectName,
			ProjectLocation:         data4.OrderData.ProjectLocation,
			NameOfConstruction:      data4.OrderData.ConstructionOrg,
			PouringPosition:         data4.OrderData.PouringPosition,
			StrengthGradeOfConcrete: data4.OrderData.GradesOfConcrete,
			Volume:                  data4.OrderData.ConcreteVolume,
			PumpingMethod:           data4.OrderData.ConcretePumpingMethod,
			Slump:                   data4.OrderData.ConcreteActualSlump,
			Hash:                    data4.OrderData.Hash,
			Name:                    data4.OrderData.ConstructionPrincipal,
			IsQualified:             IsQualified,
			Reason:                  data4.OrderData.IsQualified,
			LoginName:               cuser.LoginName,
			IsAdmin:                 cuser.IsAdmin,
		}
		fmt.Println("<--- 查询信息　--->：", anotherData)
		c.HTML(http.StatusOK, "invoice4.html", anotherData)
	case "step5":
		err = json.Unmarshal([]byte(response), &data5)
		if err != nil {
			fmt.Println(err)
		}
		anotherData := &struct {
			Id                      string
			DischargeTime           string
			Weather                 string
			Temperature             string
			ProjectName             string
			ProjectLocation         string
			NameOfConstruction      string
			NameOfSupplier          string
			NameOfSupervision       string
			PouringPosition         string
			StrengthGradeOfConcrete string
			Volume                  string
			PumpingMethod           string
			Hash1                   string
			Hash2                   string
			Name                    string
			LoginName               string
			IsAdmin                 bool
		}{
			Id:                      data5.Id,
			DischargeTime:           data5.OrderData.ConcreteDischargeTime,
			Weather:                 data5.OrderData.Weather,
			Temperature:             data5.OrderData.Temperature,
			ProjectName:             data5.OrderData.ProjectName,
			ProjectLocation:         data5.OrderData.ProjectLocation,
			NameOfConstruction:      data5.OrderData.ConstructionOrg,
			NameOfSupplier:          data5.OrderData.SupplyOrg,
			NameOfSupervision:       data5.OrderData.SupervisorOrg,
			PouringPosition:         data5.OrderData.PouringPosition,
			StrengthGradeOfConcrete: data5.OrderData.GradesOfConcrete,
			Volume:                  data5.OrderData.ConcreteVolume,
			PumpingMethod:           data5.OrderData.ConcretePumpingMethod,
			Hash1:                   data5.OrderData.Hash1,
			Hash2:                   data5.OrderData.Hash2,
			Name:                    data5.OrderData.ConstructionPrincipal,
			LoginName:               cuser.LoginName,
			IsAdmin:                 cuser.IsAdmin,
		}
		fmt.Println("<--- 查询信息　--->：", anotherData)
		c.HTML(http.StatusOK, "invoice5.html", anotherData)
	}
}

func (app *Application) QueryFile(c *gin.Context) {
	id := c.PostForm("id")

	sh := shell.NewShell("localhost:5001")

	read, err := sh.Cat(id)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(read)
	ContentType := http.DetectContentType(body)
	if err != nil {
		fmt.Println(err)
	}
	c.Header("Content-Disposition", "attachment-file")
	c.Data(http.StatusOK, ContentType, body)
}

func (app *Application) QueryByHash1(c *gin.Context) {
	h1 := c.PostForm("hash1")
	h2 := c.PostForm("hash2")
	org := c.PostForm("cc")

	str := app.Setup.QueryInfo()

	fmt.Println(str)

	json.Unmarshal([]byte(str), &indexData)

	str1 := str[41:85]
	str2 := str[108:152]

	var response string

	if h2 != "" {
		ih1, err := strconv.ParseUint(h1, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(ih1)
		response = app.Setup.QueryBlock(ih1)
	} else if h1 != "" {
		ih2 := fab.TransactionID(h2)
		response = app.Setup.QueryTransaction(ih2)
		fmt.Println(response)
	}
	data := &struct {
		BlockNumber  uint64
		Hash         string
		PreviousHash string
		Result       string
		LoginName    string
		IsAdmin      bool
	}{
		BlockNumber:  indexData.BCI.Height,
		Hash:         str1,
		PreviousHash: str2,
		Result:       response,
		LoginName:    cuser.LoginName,
		IsAdmin:      cuser.IsAdmin,
	}

	switch org {
	case "index1":
		c.HTML(http.StatusOK, "index1.html", data)
	case "index2":
		c.HTML(http.StatusOK, "index2.html", data)
	case "index3":
		c.HTML(http.StatusOK, "index3.html", data)
	case "index4":
		c.HTML(http.StatusOK, "index4.html", data)
	}
}

// func (app *Application) Set(c *gin.Context)  {
// 	name := c.PostForm("propertyname")
// 	amount := c.PostForm("amount")
// 	a:=[]string{"set",name,amount}
// 	ret,err := app.Setup.Set(a)
// 	if (err != nil) {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("<--- 添加信息　--->：", ret)
// 	c.HTML(http.StatusOK,"output.html",ret)
// }
//
// func (app *Application) Query(c *gin.Context)  {
// 	c.HTML(http.StatusOK,"query.html",nil)
// }
//
// func (app *Application) Get(c *gin.Context)  {
// 	name := c.PostForm("queryname")
//
// 	a := []string{"get",name}
// 	response, err := app.Setup.Get(a)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
//
// 	fmt.Println("<--- 查询信息　--->：", response)
// 	c.HTML(http.StatusOK,"queryOutput.html",gin.H{
// 		"name" : name,
// 		"response" : response,
// 	})
// }
//
// func (app *Application) Upload(c *gin.Context)  {
//	f, err := c.FormFile("file1")
//	fmt.Println("I am Upload")
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"error": err.Error(),
//		})
//	}else {
//		res ,err := app.Setup.UploadFilesOnIpfs(f)
//		if err != nil {
//			fmt.Println("upload file failed,err:%v",err)
//		}
//		c.JSON(http.StatusOK,gin.H{
//			"hash" : res,
//		})
//	}
// }

