package regexp_te

import (
	"fmt"
	"regexp"
)

//简单的在一串字符串提取一个特定位置的数字
/*
注意：匹配的正则需要用括号分组，才能提取出来，如 (.*?)
其中 (.*?) 表示已知字符串前后内容，找中间匹配到的字符串。如：
“好喝的蜜桃乌龙茶，哈哈哈”，根据前半部分“好喝的”和后半部分“，哈哈哈”，提取出中间部分“蜜桃乌龙茶”

特例：
如果想匹配英文括号时，需要特别处理，取消转义~如示例：findDestStr()

若与正则表达式匹配的字符串不存在，FindStringSubmatch() 这里会直接抛出panic，这里请注意 !!!
*/
func SimpleMatching() {
	str := "andarb:black_player_info:213 34123"
	// 正则表达式的分组，以括号()表示，每一对括号就是我们匹配到的一个文本，可以把他们提取出来。
	compliRegex := regexp.MustCompile("andarb:black_player_info:([0-9]*)")
	// FindStringSubmatch 方法是提取出匹配的字符串，然后通过[]string返回。
	// 我们可以看到，第1个匹配到的是这个字符串本身，从第2个开始，才是我们想要的字符串
	matchArr := compliRegex.FindStringSubmatch(str)
	fmt.Println(matchArr)
}
