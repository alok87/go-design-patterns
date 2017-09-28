package main

import (
    "github.com/alok87/go-design-patterns/creational/factory-method/storage"
)


func main() {

    diskStorage := storage.NewStore(storage.Disk)
    diskStorage.Write("disco")
    diskStorage.Read()

}
