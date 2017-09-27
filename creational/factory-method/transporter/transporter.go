package transporter

import "fmt"

type Transporter interface {
    ShipGoods(n int) 
}

type TransporterOptions struct {
    name string
    travelBy string
}

type truck struct {
    options TransporterOptions
}

type airplane struct {
    options TransporterOptions
}

func (t * truck) ShipGoods(n int) {
    fmt.Println("Shipped goods:", n, t.options)
    return
}

func (a * airplane) ShipGoods(n int) {
    fmt.Println("Shipped goods:", n, a.options)
    return
}

//Single Factory-method for implementing truck
func NewTruckTransporter(options TransporterOptions) Transporter {
    return &truck{
        options: options,
    }
}

//Single Factory-method for implementing airplane
func NewAirplaneTransporter(options TransporterOptions) Transporter {
    return &airplane{
        options: options,
    }
}

type Goods struct {
    Transportfactories []Transporter
}

func ShipGoods(goods map[string]map[string]string) {
    for key, value := range goods {
        switch key {
            case "truck":
                truck := NewTruckTransporter(
                    TransporterOptions{
                        name: value["name"], 
                        travelBy: value["travelBy"],
                })
                truck.ShipGoods(5)
            case "airplane":
                airplane := NewAirplaneTransporter(
                    TransporterOptions{
                        name: value["name"],
                        travelBy: value["travelBy"],
                })
                airplane.ShipGoods(6)
            default:
                fmt.Println("No transporter available")
        }
    }
}
