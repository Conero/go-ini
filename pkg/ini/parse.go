/* @ini-go V1.x
* @Joshua Conero
* @2017年10月28日 星期六
* @ini 文件解释器重写
 */

package ini

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ini 结构体
type Ini struct {
	Liner							 // 组合继承- 行处理
	FileName  string                 // ini 文件
	DataQueue map[string]interface{} // ini 解析后的数据
	IsSuccess bool                   // ini 文件是否解析成功
	FailMsg   string                 // 错误信息
	File      *File                  // 文件解析信息
}

// 入口文件
func Open(name string) *Ini {
	// 初始化对象
	inier := &Ini{
		FileName:  name,
		DataQueue: map[string]interface{}{},
		IsSuccess: false,
		FailMsg:   "",
		File: &File{
			line: 0,
		},
	}
	inier.reader()
	return inier
}

// 私有方法，文件读取
func (I *Ini) reader() {
	if len(I.FileName) > 0 {
		fs, err := os.Open(I.FileName)
		if err == nil {
			I.IsSuccess = true
			I.parseFile(fs)
		} else {
			I.FailMsg = err.Error()
		}
	}
}

// 解析文件
func (I *Ini) parseFile(fs *os.File) {
	buf := bufio.NewReader(fs)
	bba := BBAnalyze()
	for {
		line, err := buf.ReadString('\n')
		// 程序跳转前检测是否出错，出错直接中断循环，避免还没有检查错误时便继续进入循环(死循环)
		// -(2017年4月24日)新增的问题，最后一行还未完成时并提前结束
		isPanicError := false
		if err != nil {
			isPanicError = true
		}
		line = strings.TrimSpace(line)
		I.File.countLine()
		// 非错误
		if !isPanicError{
			// 空行
			// 单行注释
			if len(line) == 0 || I.isComment(line){
				continue
			}
			// 获取基键
			isBK,BK, nLine := I.getBaseKey(line)
			if isBK{	// 是基键
				bba.UpdateBaseKey(BK)
				continue
			}else if BK != ""{}
			fmt.Println(isBK, BK, nLine)
		}
		//fmt.Println(I.strTransform(line))
		if isPanicError {
			break
		}
	}
	I.DataQueue = bba.DataQueue
	// 输出值测试
	fmt.Println(I.DataQueue)
}

// 读取值
func (I *Ini) Get(key string) (bool, interface{}) {
	value, has := I.DataQueue[key]
	return has, value
}

// 读取函数为字符串
func (I *Ini) GetString(key string) string {
	anyType, exist := I.DataQueue[key]
	value := ""
	if exist {
		switch anyType.(type) {
		case string:
			value = anyType.(string)
		case int:
			value = strconv.Itoa(anyType.(int))
		}
	}
	return value
}

// 是否存在值
func (I *Ini) HasKey(key string) bool {
	_, exist := I.DataQueue[key]
	return exist
}