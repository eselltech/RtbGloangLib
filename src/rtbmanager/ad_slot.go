package rtbmanager

/* package rtbManager rtb ads
 * Created by vscode.
 * User: wujunquan
 * Date: 2020/03/13
 * Time: 17:05
 */

/**
 * 广告位类型  IMG/VDO
 */
const (
	AdTypeImg = "IMG"
	AdTpyeVdo = "VDO"
)

/*
AdSlot 广告位 实体
*/
type AdSlot struct {

	/**
	 * 广告位id
	 */
	SlotID string
	/**
	 * 资源类型
	 */
	Type string

	/**
	 * 返回广告数量
	 */
	Quantity int
}
