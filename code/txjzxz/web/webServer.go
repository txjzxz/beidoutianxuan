package web

import (
	"fabric-go-sdk/web/controller"
	"github.com/gin-gonic/gin"
)

func WebStart(app controller.Application)  {
	r := gin.Default()
	r.Static("/css","web/static/css")
	r.Static("/font-awesome","web/static/font-awesome")
	r.Static("/js","web/static/js")
	r.Static("/img","web/static/img")
	r.Static("/fonts","web/static/img")
	r.Static("email_templates","web/static/email_templates")
	r.LoadHTMLGlob("web/tpl/*")

	r.GET("/",app.Login)
	r.GET("/login.html",app.Login)
	r.POST("/login",app.CheckLogin)

	r.GET("/form_basic1.html",app.FormBasic1)
	r.GET("/form_basic2.html",app.FormBasic2)
	r.GET("/form_basic3.html",app.FormBasic3)
	r.GET("/form_basic4.html",app.FormBasic4)
	r.GET("/form_basic5.html",app.FormBasic5)
	r.GET("/index1.html",app.Index1)
	r.GET("/index2.html",app.Index2)
	r.GET("/index3.html",app.Index3)
	r.GET("/index4.html",app.Index4)
	r.GET("/queryInvoice1.html",app.QueryInvoice1)
	r.GET("/queryInvoice2.html",app.QueryInvoice2)
	r.GET("/queryInvoice3.html",app.QueryInvoice3)
	r.GET("/queryInvoice4.html",app.QueryInvoice4)
	r.GET("/queryFile1.html",app.File1)
	r.GET("/queryFile2.html",app.File2)
	r.GET("/queryFile3.html",app.File3)
	r.GET("/queryFile4.html",app.File4)

	r.GET("/web",app.QueryWeb)
	r.GET("/invoice.html",app.Invoice)
	r.GET("/invoice_print.html",app.InvoicePrint)

	r.POST("/upload1",app.Upload1)
	r.POST("/upload2",app.Upload2)
	r.POST("/upload3",app.Upload3)
	r.POST("/upload4",app.Upload4)
	r.POST("/upload5",app.Upload5)

	r.POST("/getqrcode",app.GetQRcode)
	r.POST("/queryStep1",app.Query)
	r.POST("queryFile",app.QueryFile)

	r.POST("/select1",app.Select1)
	r.POST("/select2",app.Select2)
	r.POST("/select3",app.Select3)

	r.POST("/queryByHash1",app.QueryByHash1)

	//r.GET("/index", app.Index)
// 	r.POST("/index",app.Set)
// 	r.GET("/query",app.Query)
// 	r.POST("/query",app.Get)
// 	r.POST("/upload", app.Upload)

	r.Run(":9000")
}
