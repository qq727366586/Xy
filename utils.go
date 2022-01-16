/**
 *@Author luojunying
 *@Date 2022-01-16 22:36
 */
package Xy

//首字节
func firstChar(str string) uint8 {
	if str == "" {
		panic("The length of the string cant not be 0 ")
	}
	return str[0]
}

//末尾字节
func lastChar(str string) uint8 {
	if str == "" {
		panic("The length of the string cant not be 0 ")
	}
	return str[len(str)-1]
}
