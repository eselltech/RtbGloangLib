package rtbmanager

/* package rtbManager rtb ads
 * Created by vscode.
 * User: wujunquan
 * Date: 2020/03/13
 * Time: 17:05
 */

/*
AdRequest 请求实体Body
*/
type AdRequest struct {

	/**
	 * 设备编号
	 */
	DeviceID string `json:"device-uuid"`

	/**
	 * 广告位编号
	 */
	SlotID string `json:"slot-id"`

	/**
	 * 返回广告数量
	 */
	Quantity int `json:"quantity"`

	/**
	 * 资源类型
	 */
	Type string `json:"type"`
}

/*
AdReply 请求返回实体Body
*/
type AdReply struct {

	/**
	 * 返回信息
	 */

	Message string `json:"message"`

	/**
	 * 返回码
	 */
	Code int `json:"code"`

	/**
	* 返回广告列表
	 */
	RtbAds []*RtbAd `json:"payload"`
}
