package gg

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
)

type measureStringer interface {
	MeasureString(s string) (w, h float64)
}

func splitOnSpace(x string) []string {
	var result []string
	pi := 0
	ps := false
	for i, c := range x {
		s := unicode.IsSpace(c)
		if s != ps && i > 0 {
			result = append(result, x[pi:i])
			pi = i
		}
		ps = s
	}
	result = append(result, x[pi:])
	return result
}

func wordWrap(m measureStringer, s string, width float64) []string {
	var result []string
	for _, line := range strings.Split(s, "\n") {
		fields := splitOnSpace(line)

		if len(fields)%2 == 1 {
			fields = append(fields, "")
		}

		x := ""
		for i := 0; i < len(fields); i += 2 {
			w, _ := m.MeasureString(x + fields[i])
			if w > width {
				if x == "" {
					result = append(result, fields[i])
					x = ""
					continue
				} else {
					result = append(result, x)
					x = ""
				}
			}
			x += fields[i] + fields[i+1]
		}
		if x != "" {
			result = append(result, x)
		}
	}
	for i, line := range result {
		result[i] = strings.TrimSpace(line)
	}
	return result
}

func wordNewWrap(m measureStringer, s string, count int, width float64) []string {
	var result []string
	fields := strings.Split(strings.ToUpper(s), ",")
	rand.Seed(time.Now().Unix())
	pre := ""
	for i := 0; i <= count; i++ {
		x := ""
		for {
			if pre != "" {
				x += pre
				pre = ""
			}
			if len(x) > 0 {
				x += " "
			}
			field := fields[rand.Intn(len(fields))]
			w,_ := m.MeasureString(x + field)
			if w < width {
				x += field
			} else if w == width {
				x += field
				break
			} else {
				more := w-width
				fieldLen := len(field)
				for j := fieldLen; j>0; j-- {
					mw,_ := m.MeasureString(string([]byte(field)[j:fieldLen]))
					if mw > more {
						x += string([]byte(field)[:j])
						pre = string([]byte(field)[j+1:fieldLen])
						break
					}
				}
				break
			}

		}
		if x != "" {
			result = append(result, x)
		}
	}

	return result
}