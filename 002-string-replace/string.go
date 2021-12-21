package goquestion

/**
 * @description: 删除字符串中的特殊字符
 * @param {string} ori 原字符串
 * @param {[]rune} filter 要删除的字符
 * @return {*}
 */
func Trim(ori string, filter []rune) string {
	if ori == "" {
		return ori
	}
	outs := []rune{}
	for _, item := range ori {
		if RuneInList(item, filter) {
			continue
		}
		outs = append(outs, item)
	}

	return string(outs)
}

func RuneInList(in rune, list []rune) bool {
	for _, item := range list {
		if in == item {
			return true
		}
	}
	return false
}
