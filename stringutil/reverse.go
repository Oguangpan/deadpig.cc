/*
stringutil 包含有用于处理字符串的工具函数. 这句是抄的.

使用方法举例

stringutil.Reverse("字符串文本")

输出的结果是

"本文串符字"

*/
package stringutil

// Reverse 将其实参数字符串以符文为单位左右反转.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
