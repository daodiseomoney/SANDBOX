package types

const (
	// ModuleName defines the module name
	ModuleName = "achilles"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_achilles"
)

var (
	ParamsKey = []byte("p_achilles")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
