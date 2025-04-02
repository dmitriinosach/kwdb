package app

type line struct {
	key      string
	defVal   string
	validate func() bool
}

var configMatches = map[string]line{
	"Host": {
		key:    "TCP_HOST",
		defVal: "localhost",
	},
	"Port": {
		key:    "TCP_PORT",
		defVal: "712",
	},
	"HttpHost": {
		key:    "HTTP_HOST",
		defVal: "localhost",
	},
	"HttpPort": {
		key:    "HTTP_PORT",
		defVal: "712",
	},
	"Driver": {
		key:    "DATABASE_DRIVER",
		defVal: "hashmap",
	},
	"Partitions": {
		key:    "DATABASE_DRIVER_PARTITIONS",
		defVal: "hashmap",
		validate: func() bool {
			if Config.Partitions > 100 || Config.Partitions < 1 {
				return false
			}
			return true
		},
	},
	"LogPath": {
		key:    "LOG_PATH",
		defVal: "data/logs",
	},
	"MemLimit": {
		key:    "MEM_LIMIT",
		defVal: "250",
	},
}
