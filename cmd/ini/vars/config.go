package vars

import (
	"fmt"
	"github.com/conero/inigo"
	"github.com/conero/uymas"
)

/**
 * @DATE        2019/6/6
 * @NAME        Joshua Conero
 * @DESCRIPIT   配置程序
 **/

var gCfg inigo.Parser

const (
	vCfgName = "./inigo.ini"
)

// 获取接口
func Cfg() inigo.Parser {
	return gCfg
}

// 欢迎语
func Welcome() {
	Br := "\r\n"
	welcomeStr :=
		"" +
			" 欢迎使用 [" + Name + "]" + Br +
			" version " + Version + "/" + Release + Br +
			" Since " + Since + Br +
			Br +
			" Base On: " + Br +
			"   . github/conero/inigo      v" + inigo.Version + "/" + inigo.Release + Br +
			"   . github/conero/uymas      v" + uymas.Version + "/" + uymas.Release + Br +
			""

	fmt.Println(welcomeStr)
}

// 初始化
func init() {
	gCfg = inigo.NewParser()
	gCfg.OpenFile(vCfgName)
}
