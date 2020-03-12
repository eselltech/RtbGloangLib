package rtbmanager

/* package rtbManager rtb ads
 * Created by vscode.
 * User: wujunquan
 * Date: 2020/03/13
 * Time: 17:15
 */

/*
RtbAd RTB广告实体
*/
type RtbAd struct {

	/**
	 * 广告id
	 */
	AdID int `json:"ad-id"`

	/**
	 * 显示时长 单位秒
	 */
	ShowTime int `json:"show-time"`

	/**
	 * 动态上报接口
	 */
	TrackURL string `json:"track-url"`

	/**
	 * 广告位id
	 */
	SlotID string `json:"slot-id"`

	/**
	 * 宽
	 */
	Width int `json:"width"`

	/**
	 * 文件大小
	 */
	FileSize int `json:"file-size"`

	/**
	 * 签名
	 */
	Sign string `json:"sign"`

	/**
	 * 过期时间 格式:"yyyy-MM-dd HH:mm:ss"
	 */
	ExpireTime string `json:"expire-time"`

	/**
	 * 类型
	 */
	Type string `json:"type"`

	/**
	 * 资源路径（相同素材请缓存,高峰流量5毛钱1G）
	 */
	URL string `json:"url"`

	/**
	 * 高
	 */
	Height int `json:"height"`
}
