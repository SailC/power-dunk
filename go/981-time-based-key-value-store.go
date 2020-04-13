package main

type TimeMap struct {
	Map map[string][]TimeValue
}

type TimeValue struct {
	ts    int
	value string
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		Map: map[string][]TimeValue{},
	}
}

func (m *TimeMap) Set(key string, value string, timestamp int) {
	tm := m.Map
	tm[key] = append(tm[key], TimeValue{timestamp, value})
}

func (m *TimeMap) Get(key string, timestamp int) string {
	tm := m.Map
	if _, ok := tm[key]; !ok {
		return ""
	}
	tValues := tm[key]
	return findLatestValue(tValues, timestamp)
}

func findLatestValue(tValues []TimeValue, ts int) string {
	lo, hi := 0, len(tValues)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if ts == tValues[mid].ts {
			return tValues[mid].value
		}
		if ts < tValues[mid].ts {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	if hi < 0 || tValues[hi].ts > ts {
		return ""
	}
	return tValues[hi].value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
