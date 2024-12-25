package driver

type Interface interface {
	SetValue(key string, value string, ttl int) bool
	GetValue(key string) (*Cell, bool)
	DeleteValue(key string) bool
	UpdateValue(key string) bool
	HasKey(key string) bool
	Truncate()
	Init()
	Lock()
	Unlock()
	GetDriver() string
	Info() string
	GetVaultMap() map[string]*Cell
}
