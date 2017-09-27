package main

import (
    "github.com/alok87/go-design-patterns/creational/factory-method/transporter"
)

const (
    LAND = "land"
    AIR = "air"
)

func main() {

    goods := make(map[string]map[string]string)

    goods["truck"] =  make(map[string]string)
    goods["truck"]["travelBy"] = LAND
    goods["truck"]["name"] = "truck"

    goods["airplane"] =  make(map[string]string)
    goods["airplane"]["travelBy"] = AIR
    goods["airplane"]["name"] = "air"

    transporter.ShipGoods(goods)
}
