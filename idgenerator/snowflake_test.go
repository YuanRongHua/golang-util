package idgenerator

import "testing"

func TestInitSnowFlakeWorker(t *testing.T) {
	var workId int64= 1
	err := InitSnowFlakeWorker(workId)
	if err != nil {
		t.Errorf("TestInitSnowFlakeWorker failed, %v",err.Error())
	}
}

func TestSnowFlakeId(t *testing.T) {
	var workId int64= 1
	InitSnowFlakeWorker(workId)
	snowFlakeId:=SnowFlakeId()
	if snowFlakeId==0{
		t.Error("TestSnowFlakeId failed")
	}else{
		t.Logf("snowFlakeId:%v",snowFlakeId)
	}
}