package storage

import (
	"fmt"
)

// Store interface for different types of storage implementation
type Store interface {
	Write(s string)
	Read()
}

// DiskOptions for the disk
type DiskOptions struct{}

// DiskStorage for the disk
type DiskStorage struct {
	options DiskOptions
}

// MemoryOptions sss
type MemoryOptions struct{}

type InMemoryStorage struct {
	options MemoryOptions
}

type FileOptions struct{}

type DefaultFileStorage struct {
	options FileOptions
}

func (t *DiskStorage) Write(s string) {
	fmt.Println(s)
}

func (t *InMemoryStorage) Write(s string) {
	fmt.Println(s)
}

func (t *DefaultFileStorage) Write(s string) {
	fmt.Println(s)
}

func (t *DiskStorage) Read() {
	fmt.Println("Reading data")
}

func (t *InMemoryStorage) Read() {
	fmt.Println("Reading data")
}

func (t *DefaultFileStorage) Read() {
	fmt.Println("Reading data")
}

const (
	Disk   = 1
	Memory = 2
)

type StorageType int

func newDisk() Store {
	return &DiskStorage{
		options: DiskOptions{},
	}
}

func newMemory() Store {
	return &InMemoryStorage{
		options: MemoryOptions{},
	}
}

func newFile() Store {
	return &DefaultFileStorage{
		options: FileOptions{},
	}
}

//NewStore is a Factory Method to create objects for each of the storage types
//Objects are being created using this intermediate factory-method
//without having to specify exact type(storage-type in this case) of the object
func NewStore(t StorageType) Store {
	switch t {
	case Disk:
		return newDisk()
	case Memory:
		return newMemory()
	default:
		return newFile()
	}
}
