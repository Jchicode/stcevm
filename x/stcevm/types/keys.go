package types

const (
	// ModuleName defines the module name
	ModuleName = "stcevm"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_stcevm"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	DbKey = "Db/value/"
)

const (
	KvKey = "Kv/value/"
)
