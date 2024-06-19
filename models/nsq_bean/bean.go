package nsq_bean

type NsqBean struct {
	Date       string // 日期
	Categories string // 分类
}

type OrderDelay struct {
	OrderNo        string
	Topic          string
	Times          int64
	UndertakenType int
}
