/*
 * Copyright(C) 2020 github.com/hidu  All Rights Reserved.
 * Author: hidu (duv123+git@baidu.com)
 * Date: 2020/7/3
 */

package mod2

import (
	"runtime"

	"github.com/fsgo/mod1"
)

func Version() string {
	return mod1.Version() + "\n" + "mod2{ mod_version:1.0.0 go:" + runtime.Version() + " }"
}
