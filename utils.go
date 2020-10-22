/**
 * @author      Liu Yongshuai<liuyongshuai@hotmail.com>
 * @date        2018-03-27 15:20
 */
package pkg

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var (
	LocalHostIP         = ""
	LocalHostIpArr      []string
	LocalHostIpTraceId  = ""
	preTraceID          = ""
	ScreenWidth         int
	ScreenHeight        int
)


// 在指定浮点数范围内生成随机数
func RandFloat64InRange(min, max float64) float64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()*(max-min) + min
}

// 处理请求的时间proc_time
// 时间要求time.Now().UnixNano()
func ProcTime(st, et int64) float64 {
	var t1 int64 = 0
	var t2 int64 = 0
	if st > et {
		t1 = et
		t2 = st
	} else {
		t1 = st
		t2 = et
	}
	ret := float64(t2-t1) / float64(1000*1000*1000)
	return ret
}

//生成一个假的traceId
func FakeTraceId() (traceId string) {
	for {
		ReRandSeed()
		traceId = fmt.Sprintf("%x%s%x", time.Now().UnixNano(), LocalHostIpTraceId, rand.Int63())
		if preTraceID != traceId {
			preTraceID = traceId
			break
		}
	}
	return traceId
}

//重新设置随机数种子
func ReRandSeed() {
	rand.Seed(time.Now().UnixNano())
}

//根据业务特点，过滤非法的ID并去重，一般用于批量根据ID提取信息时
func FilterIds(ids []interface{}) (ret []int64) {
	tmap := map[int64]struct{}{}
	for _, id := range ids {
		v, err := TryBestToInt64(id)
		if err != nil || v <= 0 {
			continue
		}
		tmap[v] = struct{}{}
	}
	for i := range tmap {
		ret = append(ret, i)
	}
	return
}

//返回最大的一个int型
func MaxInt64(args ...interface{}) (int64, error) {
	if len(args) <= 0 {
		return 0, ErrorInvalidInputType
	}
	var m int64 = math.MinInt64
	var tmps []int64
	for _, arg := range args {
		a, e := TryBestToInt64(arg)
		if e != nil {
			continue
		}
		tmps = append(tmps, a)
	}
	if len(tmps) <= 0 {
		return 0, ErrorInvalidInputType
	}
	for _, t := range tmps {
		if t > m {
			m = t
		}
	}
	return m, nil
}

//返回最小的一个int型
func MinInt64(args ...interface{}) (int64, error) {
	if len(args) <= 0 {
		return 0, ErrorInvalidInputType
	}
	var m int64 = math.MaxInt64
	var tmps []int64
	for _, arg := range args {
		a, e := TryBestToInt64(arg)
		if e != nil {
			continue
		}
		tmps = append(tmps, a)
	}
	if len(tmps) <= 0 {
		return 0, ErrorInvalidInputType
	}
	for _, t := range tmps {
		if t < m {
			m = t
		}
	}
	return m, nil
}