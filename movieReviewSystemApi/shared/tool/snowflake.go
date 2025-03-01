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

	// 调整取余的选项
	modulus       int64 // 取余的模数
	remainder     int64 // 希望的余数
	modAdjustment bool  // 是否启用取余调整
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
		dataCenterID:  dataCenterID,
		machineID:     machineID,
		modAdjustment: false,
	}, nil
}

// SetModOptions 设置取余选项
func (s *Snowflake) SetModOptions(modulus, remainder int64) error {
	if modulus <= 0 {
		return errors.New("modulus must be positive")
	}
	if remainder < 0 || remainder >= modulus {
		return errors.New("remainder must be in [0, modulus)")
	}
	s.modulus = modulus
	s.remainder = remainder
	s.modAdjustment = true
	return nil
}

// Generate 生成唯一ID
func (s *Snowflake) Generate() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	currentStamp := s.getTimestamp()
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

	// 调整ID以满足取余条件
	if s.modAdjustment {
		id = s.adjustMod(id)
	}

	return id, nil
}

// 调整ID以满足取余条件
func (s *Snowflake) adjustMod(id int64) int64 {
	// 如果ID已经满足取余条件，直接返回
	if (id % s.modulus) == s.remainder {
		return id
	}

	// 计算需要调整的增量
	delta := (s.modulus - ((id - s.remainder) % s.modulus)) % s.modulus
	return id + delta
}

// 获取当前时间戳（毫秒）
func (s *Snowflake) getTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

// 等待下一毫秒
func (s *Snowflake) waitNextMillis() int64 {
	current := s.getTimestamp()
	for current <= s.lastStamp {
		time.Sleep(100 * time.Microsecond)
		current = s.getTimestamp()
	}
	return current
}
