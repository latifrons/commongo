package tools

import "strconv"

type BiPair struct {
	Vs map[int]string
	sv map[string]int
}

func (b *BiPair) InitDefault() {
	b.sv = make(map[string]int)
	for v, s := range b.Vs {
		b.sv[s] = v
	}
}

func (b *BiPair) AddPair(s string, v int) {
	b.sv[s] = v
	b.Vs[v] = s
}

func (b *BiPair) GetString(v int, defaultString string) string {
	if s, ok := b.Vs[v]; ok {
		return s
	}
	return defaultString
}

func (b *BiPair) GetStringOrOrigin(v int) string {
	if s, ok := b.Vs[v]; ok {
		return s
	}
	return strconv.Itoa(v)
}

func (b *BiPair) GetValue(s string, defaultValue int) int {
	if v, ok := b.sv[s]; ok {
		return v
	}
	return defaultValue
}
