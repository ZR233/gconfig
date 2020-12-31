package cerr

import "fmt"

var (
	ErrConfigNotInit = fmt.Errorf("配置文件未初始化")
)
