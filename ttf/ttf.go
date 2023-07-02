// 字体来源：https://github.com/lxgw/LxgwWenKai/
package ttf

import (
	_ "embed"
)

//go:embed LXGWWenKai-Bold.ttf
var WenKaiBold []byte

//go:embed LXGWWenKai-Light.ttf
var WenKaiLight []byte

//go:embed LXGWWenKai-Regular.ttf
var WenKaiRegular []byte

//go:embed LXGWWenKaiMono-Bold.ttf
var WenKaiMonoBold []byte

//go:embed LXGWWenKaiMono-Light.ttf
var WenKaiMonoLight []byte

//go:embed LXGWWenKaiMono-Regular.ttf
var WenKaiMonoRegular []byte
