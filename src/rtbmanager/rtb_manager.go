package rtbmanager

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/* package rtbManager rtb ads
 * Created by vscode.
 * User: wujunquan
 * Date: 2020/03/13
 * Time: 17:35
 */

/*
RtbManager RTB管理
*/
type RtbManager struct {

	/**
	 * appid
	 */
	AppID string

	/**
	 * appkey
	 */
	AppKey string

	/**
	 * 设备ID
	 */
	DeviceID string
}

/*
NewRtbManager 初始化方法
appid:由esell 提供
appkey:由esell 提供
deviceID:已经入库屏效宝的设备ID
*/
func NewRtbManager(appid, appkey, deviceID string) *RtbManager {

	rtbManager := new(RtbManager)
	rtbManager.AppID = appid
	rtbManager.AppKey = appkey
	rtbManager.DeviceID = deviceID
	return rtbManager
}

/*
payload 生成请求体方法
adSlot:请求目标广告位
*/
func (rtbManager *RtbManager) payload(adSlot *AdSlot) string {

	adRequest := &AdRequest{
		DeviceID: rtbManager.DeviceID,
		Type:     adSlot.Type,
		Quantity: adSlot.Quantity,
		SlotID:   adSlot.SlotID,
	}

	data, err := json.Marshal(adRequest)

	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}

	return string(data)

}

/*
RequestAds 请求广告方法
adSlot: 从目标广告位请求广告
AdReply:请求广告返回
*/
func (rtbManager *RtbManager) RequestAds(adSlot *AdSlot) *AdReply {

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	fmt.Println("时间戳---->" + timestamp)

	payload := rtbManager.payload(adSlot)

	fmt.Println("post payload--->" + payload)

	sign := rtbManager.sign(payload, timestamp)

	fmt.Println("sign->" + sign)

	postURL := rtbManager.createURL(timestamp, sign)

	fmt.Println(postURL)

	resp, err := http.Post(postURL,
		"application/x-www-form-urlencoded",
		strings.NewReader("payload="+payload))
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		// handle error
		adReply := &AdReply{Code: -2, Message: fmt.Sprintf("请求失败-返回错误:%s", err.Error())}
		return adReply
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		adReply := &AdReply{Code: -1, Message: fmt.Sprintf("请求返回流读取错误:%s", err.Error())}
		return adReply
	}

	fmt.Println("请成功求返回结果--->" + string(body))

	defer resp.Body.Close()

	adReply := new(AdReply)

	err = json.Unmarshal([]byte(string(body)), adReply)
	if err != nil {
		adReply.Message = fmt.Sprintf("响应成功-但是json解析错误:%s", err.Error())
		adReply.Code = -3
	}

	return adReply

}

/*
createURL 连接拼接方法
timestamp: 时间戳
sign:签名
*/
func (rtbManager *RtbManager) createURL(timestamp string, sign string) string {
	var buffer bytes.Buffer
	buffer.WriteString("http://api6.pingxiaobao.com/rtb/subscribe.shtml?")
	buffer.WriteString("appid=")
	buffer.WriteString(rtbManager.AppID)
	buffer.WriteString("&")
	buffer.WriteString("sequence=")
	buffer.WriteString(timestamp)
	buffer.WriteString("&")
	buffer.WriteString("timestamp=")
	buffer.WriteString(timestamp)
	buffer.WriteString("&")
	buffer.WriteString("uuid=")
	buffer.WriteString(rtbManager.DeviceID)
	buffer.WriteString("&")
	buffer.WriteString("version=")
	buffer.WriteString("1.3")
	buffer.WriteString("&")
	buffer.WriteString("sign=")
	buffer.WriteString(sign)
	urlStr := buffer.String()
	return urlStr
}

/*
sign 计算签名
timestamp: 时间戳
*/
func (rtbManager *RtbManager) sign(payload, timestamp string) string {
	var buffer bytes.Buffer
	buffer.WriteString("appid=")
	buffer.WriteString(rtbManager.AppID)
	buffer.WriteString("&")
	buffer.WriteString("appkey=")
	buffer.WriteString(rtbManager.AppKey)
	buffer.WriteString("&")
	buffer.WriteString("payload=")
	buffer.WriteString(payload)
	buffer.WriteString("&")
	buffer.WriteString("sequence=")
	buffer.WriteString(timestamp)
	buffer.WriteString("&")
	buffer.WriteString("timestamp=")
	buffer.WriteString(timestamp)
	buffer.WriteString("&")
	buffer.WriteString("uuid=")
	buffer.WriteString(rtbManager.DeviceID)
	buffer.WriteString("&")
	buffer.WriteString("version=")
	buffer.WriteString("1.3")
	signstr := buffer.String()
	return strMd5(signstr)
}

/*
strMd5:字符串转MD5
*/
func strMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
