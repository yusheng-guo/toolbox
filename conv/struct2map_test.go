package conv

import (
	"fmt"
	"testing"
	"time"
)

type TransformationEffectReq struct {
	Callback       string  `json:"callback" map:"callback,omitempty"`               // 接收的__CALLBACK__替换后的http地址中的callback参数 3 付费
	EventType      int     `json:"event_type" map:"event_type,omitempty"`           // 事件类型 13位毫秒级时间戳
	EventTime      int64   `json:"event_time" map:"event_time,omitempty"`           // 事件时间 13位毫秒级时间戳
	PurchaseAmount float64 `json:"purchase_amount" map:"purchase_amount,omitempty"` // 金额 单位元 可保留两位小数
}

func TestStruct2Map(t *testing.T) {
	now := time.Now().UnixMilli()
	req := TransformationEffectReq{
		Callback:       "f2adshkfhg15akwqerqf+asd",
		EventType:      3,
		EventTime:      now,
		PurchaseAmount: 10,
	}
	res, err := Struct2Map(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("res: ", res)
}
