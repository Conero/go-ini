/* @ini-go V1.x
* @Joshua Conero
* @2017年10月28日 星期六
* @ini 文件相关信息
*/
package rong

// 文件解析先关信息
type File struct {
	line int			// 总行数
}
// 行计数器
func (file *File) countLine() *File {
	file.line = file.line + 1
	return file
}