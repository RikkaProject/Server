package util

func GeneratePlayerKey(rid int, record int) uint64 {
	var base = 4000000000000000000
	return uint64(base + (rid * 100) + record)
}
