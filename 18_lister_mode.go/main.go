package main

import "fmt"

const (
	PGaiBang  string = "丐帮"
	PMingJiao string = "明教"
)

// 抽象层
type Listener interface {
	OnFlightDo(event *Event)
	GetName() string
	GetPart() string
	GetTitle() string
}

type Notifer interface {
	AddLister(lis Listener)
	RemoveLister(lis Listener)
	Notife(event *Event)
}

// 实现层
type Event struct {
	Other   *Hero   // 打人方
	Another *Hero   // 被打方
	Noti    Notifer // 通知人
	Msg     string  // 时间消息
}

type Hero struct {
	Name string
	Part string
}

func (h *Hero) OnFlightDo(event *Event) {
	if event.Other.GetName() == h.Name || event.Another.GetName() == h.Name {
		// 是当事人双方，不做动作
		return
	}

	if event.Other.Part == h.Part {
		// 打人方是本帮派的，拍手叫好
		fmt.Println(h.GetTitle() + "：拍手叫好...")
		return
	} else {
		fmt.Println(h.GetTitle() + " 得知消息，发起反击!!!")
		// fmt.Println(h.GetTitle() + " 要揍 " + event.Other.GetTitle() + " ...")
		h.Flight(event.Other, event.Noti)
		return
	}

}

func (h *Hero) GetName() string {
	return h.Name
}

func (h *Hero) GetPart() string {
	return h.Part
}

func (h *Hero) GetTitle() string {
	return fmt.Sprintf("[%s]:%s", h.Part, h.Name)
}

func (h *Hero) Flight(another *Hero, baixiao Notifer) {
	msg := fmt.Sprint(h.GetTitle(), "把", another.GetTitle(), " 揍了...")
	event := &Event{
		Other:   h,
		Another: another,
		Noti:    baixiao,
		Msg:     msg,
	}

	baixiao.Notife(event) // 告诉百晓生
}

type BaiXiao struct {
	listerner []Listener
}

func (bx *BaiXiao) AddLister(lis Listener) {
	bx.listerner = append(bx.listerner, lis)
}

func (bx *BaiXiao) RemoveLister(lis Listener) {
	for index, listern := range bx.listerner {
		if listern.GetName() == lis.GetName() {
			bx.listerner = append(bx.listerner[0:index], bx.listerner[index+1:]...) // 这个方法相当于删除切片中的某一项
		}
	}
}

func (bx *BaiXiao) Notife(event *Event) {
	fmt.Println("[世界消息]: " + event.Msg)
	for _, listern := range bx.listerner {
		listern.OnFlightDo(event)
	}
}

// 业务逻辑层
func main() {
	h1 := &Hero{
		Name: "黄蓉",
		Part: PGaiBang,
	}
	h2 := &Hero{
		Name: "洪七公",
		Part: PGaiBang,
	}
	h3 := &Hero{
		Name: "乔峰",
		Part: PGaiBang,
	}

	h4 := &Hero{
		Name: "张无忌",
		Part: PMingJiao,
	}
	h5 := &Hero{
		Name: "灭绝师太",
		Part: PMingJiao,
	}
	h6 := &Hero{
		Name: "金毛狮王",
		Part: PMingJiao,
	}

	baixiaosheng := &BaiXiao{}
	baixiaosheng.AddLister(h1)
	baixiaosheng.AddLister(h2)
	baixiaosheng.AddLister(h3)
	baixiaosheng.AddLister(h4)
	baixiaosheng.AddLister(h5)
	baixiaosheng.AddLister(h6)

	// 发起战争
	h1.Flight(h4, baixiaosheng)
}
