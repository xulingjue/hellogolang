package template

import (
	"regexp"
)

func StringEqual(a string, b string) int {
	if a == b {
		return 1
	}
	return 0
}

func Int64Equal(a int64, b int64) int {
	if a == b {
		return 1
	}
	return 0
}

func IntEqual(a int, b int) int {
	if a == b {
		return 1
	}
	return 0
}

func RemoveHtmlTag(content string) string {
	//  正则表达式demo
	re, _ := regexp.Compile("\\<[^>]+?\\>")
	content = re.ReplaceAllString(content, "")

	if showStrlen(content) > 300 {
		content = showSubstr(content, 300) + "......"
	}
	return content
}

func showSubstr(s string, l int) string {
	if len(s) <= l {
		return s
	}
	ss, sl, rl, rs := "", 0, 0, []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			rl = 1
		} else {
			rl = 2
		}

		if sl+rl > l {
			break
		}
		sl += rl
		ss += string(r)
	}
	return ss
}

func showStrlen(s string) int {
	sl := 0
	rs := []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			sl++
		} else {
			sl += 2
		}
	}
	return sl
}
