//go:build tools

// revel コマンドのバージョン固定用。
// revel/cmd v1.1.2 が依存する古い golang.org/x/tools は新しい Go で panic するため、x/tools を新しいバージョンに上げている。
// インストール: (cd tools/revel-cmd && go install github.com/revel/cmd/revel)
package tools

import _ "github.com/revel/cmd/revel"
