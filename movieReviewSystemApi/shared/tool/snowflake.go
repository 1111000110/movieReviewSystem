package tool

import (
	"errors"
	"sync"
	"time"
)

// Snowflake 配置参数
const (
	epoch          = 1609459200000 // 起始时间戳（2021-01-01 00:00:00 UTC）
	timestampBits  = 41            // 时间戳占用位数
	dataCenterBits = 5             // 数据中心ID位数
	machineIDBits  = 5             // 机器ID位数
	sequenceBits   = 12            // 序列号位数

	maxDataCenterID = 1<<dataCenterBits - 1
	maxMachineID    = 1<<machineIDBits - 1
	maxSequence     = 1<<sequenceBits - 1

	timestampShift  = dataCenterBits + machineIDBits + sequenceBits
	dataCenterShift = machineIDBits + sequenceBits
	machineShift    = sequenceBits
)

// Snowflake 生成器结构体
type Snowflake struct {
	mu           sync.Mutex
	dataCenterID int64
	machineID    int64
	sequence     int64
	lastStamp    int64
}

// NewSnowflake 创建实例
func NewSnowflake(dataCenterID, machineID int64) (*Snowflake, error) {
	if dataCenterID < 0 || dataCenterID > maxDataCenterID {
		return nil, errors.New("invalid data center ID")
	}
	if machineID < 0 || machineID > maxMachineID {
		return nil, errors.New("invalid machine ID")
	}
	return &Snowflake{
		dataCenterID: dataCenterID,
		machineID:    machineID,
	}, nil
}

// Generate 生成唯一ID
func (s *Snowflake) Generate() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	currentStamp := time.Now().UnixNano() / 1e6 // 毫秒时间戳
	if currentStamp < s.lastStamp {
		return 0, errors.New("clock moved backwards")
	}

	if currentStamp == s.lastStamp {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 { // 当前毫秒序列号用完，等待下一毫秒
			currentStamp = s.waitNextMillis()
		}
	} else {
		s.sequence = 0
	}

	s.lastStamp = currentStamp

	// 组合各部分的位运算
	id := (currentStamp-epoch)<<timestampShift |
		s.dataCenterID<<dataCenterShift |
		s.machineID<<machineShift |
		s.sequence

	return id, nil
}

// 等待下一毫秒
func (s *Snowflake) waitNextMillis() int64 {
	current := time.Now().UnixNano() / 1e6
	for current <= s.lastStamp {
		time.Sleep(100 * time.Microsecond)
		current = time.Now().UnixNano() / 1e6
	}
	return current
}
