module bonno

go 1.27

require (
	github.com/go-gorp/gorp/v3 v3.1.0
	github.com/mattn/go-sqlite3 v2.0.1+incompatible
	github.com/revel/modules v1.1.0
	github.com/revel/revel v1.1.0
)

require (
	github.com/bradfitz/gomemcache v0.0.0-20220106215444-fb4bf637b56d // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/gomodule/redigo v1.8.8 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/inconshreveable/log15 v0.0.0-20201112154412-8562bdadbbac // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/revel/config v1.0.0 // indirect
	github.com/revel/log15 v2.11.20+incompatible // indirect
	github.com/revel/pathtree v0.0.0-20140121041023-41257a1839e9 // indirect
	github.com/xeonx/timeago v1.0.0-rc4 // indirect
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5 // indirect
	golang.org/x/sys v0.0.0-20221013171732-95e765b1cc43 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/stack.v0 v0.0.0-20141108040640-9b43fcefddd0 // indirect
)

// revel/modules が古い go-sqlite3 v2.0.1+incompatible を間接参照するため、最新の v1 系に固定する
replace github.com/mattn/go-sqlite3 => github.com/mattn/go-sqlite3 v1.14.52
