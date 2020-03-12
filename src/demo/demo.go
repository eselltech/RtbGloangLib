package main

import (
	"fmt"
	"rtbmanager"
)

/* package rtbManager rtb ads
 * Created by vscode.
 * User: wujunquan
 * Date: 2020/03/13
 * Time: 18:35
 */
func main() {

	/*
		此测试信息，媒体调试填入自己的对接信息
			APPID：y8i2vjbexf845h1i
			APPKEY：ry4od91w45hmx6pu1mnr7leceogm94t8

			广告位ID:25076009
			ktv1 入库设备
	*/

	/*
	 根据Esell 提供的信息 初始化一个实列
	 AppID
	 AppKey
	 请求的设备ID(此设备必须提前入库)
	*/
	rtbManagerInstance := rtbmanager.NewRtbManager("gidrk4ow3ervp32o", "k61fesdzetuv3ss7hh6eh715v6kxr8zr", "EA9863E9ADA4")

	/*
		初始化广告位
		广告位ID ESell提供的
		广告位类型 IMG/VDO 两种
		Quantity:一次性请求广告的数量，相同广告 track-url 不一样，按次结算。
	*/
	adSolt := &rtbmanager.AdSlot{SlotID: "25075926", Type: rtbmanager.AdTypeImg, Quantity: 4}

	/*
	 请求返回实体
	*/
	adReply := rtbManagerInstance.RequestAds(adSolt)

	//返回实体
	fmt.Printf(adReply.Message)

	//广告集合
	//adReply.RtbAds
}
