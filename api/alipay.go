package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"tobeg/global"

	"tobeg/cert"

	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func TradePrecreate(ctx *gin.Context) {

	amount := ctx.PostForm("amount")
	price, err := strconv.ParseFloat(amount, 32)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"msg": "输入数额非法",
		})
		return
	}
	client := generateClient()
	if client == nil {
		return
	}
	client.SetReturnUrl(global.Config.BaseConfig.Url + "/v1/alipay/trade/paysuccess").
		SetNotifyUrl(global.Config.BaseConfig.Url + "/v1/alipay/trade/paysuccess")

	// 通过uuid生成订单号
	orderId := uuid.New().String()
	orderId = strings.ReplaceAll(orderId, "-", "")

	bm := make(gopay.BodyMap)

	bm.Set("subject", "双汇王中王")
	bm.Set("out_trade_no", orderId)
	bm.Set("total_amount", fmt.Sprintf("%.2f", price))

	// 穿件预支付二维码
	aliRsp, err := client.TradePrecreate(context.Background(), bm)
	if err != nil {
		zap.S().Errorf("支付宝创建预支付失败")
		return
	}

	zap.S().Debugf("aliRsp: %s", *aliRsp)
	zap.S().Debugf("aliRsp.QrCode: %s", aliRsp.Response.QrCode)
	zap.S().Debugf("aliRsp.OutTradeNo: %s", aliRsp.Response.OutTradeNo)

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "成功",
		"data": map[string]string{
			"qrCode":  aliRsp.Response.QrCode,
			"tradeNo": aliRsp.Response.OutTradeNo,
		},
	})
}

// 接收支付宝支付回调
// https://opendocs.alipay.com/open/194/103296
func TradePaySuccess(ctx *gin.Context) {
	outTradeNo := ctx.PostForm("out_trade_no")
	tradeStatus := ctx.PostForm("trade_status")
	if tradeStatus == "TRADE_SUCCESS" {
		global.PayStatus[outTradeNo] = true
	}
	ctx.JSON(http.StatusOK, "success")
	timeAfterTrigger := time.After(10 * time.Second)
	<-timeAfterTrigger
	delete(global.PayStatus, outTradeNo)
	zap.S().Infof("订单号: %s 已删除", outTradeNo)
}

func generateClient() *alipay.Client {
	// 初始化支付宝客户端
	// appid：应用ID
	// privateKey 应用私钥
	// isProd 是否是正式环境
	zap.S().Info(global.Config.AlipayConfig.AppId)
	zap.S().Info(cert.Appid)
	zap.S().Info(cert.PrivateKey)
	client, err := alipay.NewClient(cert.Appid, cert.PrivateKey, true)
	if err != nil {
		zap.S().Errorf("支付宝初始化错误: %s", err.Error())
		return nil
	}

	// 关闭debug
	client.DebugSwitch = gopay.DebugOff

	// 设置支付宝公用请求参数
	client.SetLocation(alipay.LocationShanghai).
		SetCharset(alipay.UTF8).
		SetSignType(alipay.RSA2)
		// 设置证书内容
	err = client.SetCertSnByContent(cert.AppPublicContent, cert.AlipayRootContent, cert.AlipayPublicContentRSA2)
	return client
}

// 交易关闭接口
// https://opendocs.alipay.com/open/02o6e7
func TradeClose(ctx *gin.Context) {

	client := generateClient()
	if client == nil {
		return
	}
	orderId := ctx.PostForm("orderId")
	zap.S().Info(orderId)
	if orderId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "缺少OrderId",
		})
		return
	}

	bm := make(gopay.BodyMap)

	bm.Set("out_trade_no", orderId)

	// 穿件预支付二维码
	aliRsp, err := client.TradeCancel(context.Background(), bm)
	if err != nil {
		zap.S().Errorf("取消订单失败: %s", err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "成功",
		"data": map[string]string{
			"trade_no":     aliRsp.Response.TradeNo,
			"out_trade_no": aliRsp.Response.OutTradeNo,
		},
	})
}

func TradeQuery(ctx *gin.Context) {
	orderId := ctx.PostForm("orderId")
	if orderId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "缺少OrderId",
		})
		return
	}
	status := "TRADE_WAIT"
	if val, ok := global.PayStatus[orderId]; ok {
		if val == true {
			status = "TRADE_SUCCESS"
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "成功",
		"data": map[string]string{
			"status": status,
		},
	})

}

// 交易查询接口
// https://opendocs.alipay.com/open/02ekfh?scene=23
func TradeQuery2(ctx *gin.Context) {
	client := generateClient()
	if client == nil {
		return
	}
	orderId := ctx.PostForm("orderId")
	if orderId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "缺少OrderId",
		})
		return
	}
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", orderId)

	aliRsp, err := client.TradeQuery(context.Background(), bm)
	if err != nil {
		zap.S().Errorf("查询订单状态失败: %s", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "查询失败",
		})
		return
	}
	status := aliRsp.Response.TradeStatus
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "成功",
		"data": map[string]string{
			"status": status,
		},
	})
}
