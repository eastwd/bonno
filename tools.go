//go:build tools

// conf/app.conf で読み込む Revel モジュールを go.mod に保持するためのファイル
package tools

import (
	_ "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app/controllers"
)
