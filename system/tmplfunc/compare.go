package tmplfunc

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

	if len(content) > 300 {
		content = content[:300] + "……"
	}
	return content
}
