package game

// Represents a resource that can be locked and unlocked.
type Locker interface {
	Lock()
	Unlock()
}

// Represents a resource that can be locked and unlocked for both reading and writing.
type RWLocker interface {
	Locker

	RLock()
	RUnlock()
}

// Lock the resouce for writing before executing protectedWrite and then release the lock. Only one write operation can happen at a time.
func WriteLock(locker Locker, protectedWrite func()) {
	locker.Lock()
	defer locker.Unlock()
	protectedWrite()
}

// Lock the resouce for reading before executing protectedRead and then release the lock. Multiple read operations can happen at the same time.
func ReadLock(locker RWLocker, protectedRead func()) {
	locker.RLock()
	defer locker.RUnlock()
	protectedRead()
}
