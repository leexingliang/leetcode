package regular_expression

var pattern string
var matched bool
var plen int

// Pattern pattern
func Pattern(p string) {
	pattern = p
	plen = len(pattern)
}

func isMatch(text string) bool {
	matched = false
	rmatch(0, 0, text, len(text))
	return matched
}

func rmatch(ti int, pj int, text string, tlen int) {
	if matched {
		return
	}
	if pj >= plen {
		if ti >= tlen {
			matched = true
		}
		return
	}
	if pattern[pj] == '*' { // 匹配任意字符
		for k := 0; k <= tlen-ti; k++ {
			rmatch(ti+k, pj+1, text, tlen)
		}
	} else if pattern[pj] == '?' { //匹配 0 个或 1 个字符
		rmatch(ti, pj+1, text, tlen)
		rmatch(ti+1, pj+1, text, tlen)
	} else if ti < tlen && pattern[pj] == text[ti] {
		rmatch(ti+1, pj+1, text, tlen)
	}
}
