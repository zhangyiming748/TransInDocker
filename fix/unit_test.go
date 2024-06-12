package fix

import (
	"testing"
)

func TestReplace(t *testing.T) {
	// [33mDidyoumean[1mImpregnateme.[22m[0m让我怀孕吧
	str := "\u001B[33mDidyoumean\u001B[1mImpregnateme.\u001B[22m\u001B[0m让我怀孕吧"

	// 使用正则表达式匹配所有标点符号
	Fix33m(str)

}
