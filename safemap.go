package safemap

type Safemap[K comparable , V any] struct{
	mu sync.RMutex
	
}
