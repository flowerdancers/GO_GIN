package config

type Log struct {
	Level          string `yaml: "level"`
	Prefix         int    `yaml: "prefix" `
	dDirectory     string `yaml: "directory"`
	Show_line      bool   `yaml: "show_line"`      //是否显示行号
	Log_in_console bool   `yaml: "log_in_console"` // 是否显示打印的路径
}