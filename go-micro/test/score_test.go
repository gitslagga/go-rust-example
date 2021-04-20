package test

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-micro/proto/score"
	"testing"
)

func TestScore(t *testing.T) {
	// 初始化消息
	scoreInfo := &score.BaseScoreInfo{}
	scoreInfo.WinCount = 10
	scoreInfo.LoseCount = 1
	scoreInfo.ExceptionCount = 2
	scoreInfo.KillCount = 2
	scoreInfo.DeathCount = 1
	scoreInfo.AssistCount = 3
	scoreInfo.Rating = 120

	// 以字符串的形式打印消息
	fmt.Println(scoreInfo.String())

	// encode, 转换成二进制数据
	data, err := proto.Marshal(scoreInfo)
	if err != nil {
		panic(err)
	}

	// decode, 将二进制数据转换成struct对象
	newScoreInfo := score.BaseScoreInfo{}
	err = proto.Unmarshal(data, &newScoreInfo)
	if err != nil {
		panic(err)
	}

	// 以字符串的形式打印消息
	fmt.Println(newScoreInfo.String())
}
