package cart // definición del paquete actual

import "errors" // importamos paquete de la librería estándar

const maxItems = 5 // no se exporta

type Item string // tipo personalizado

type Order struct {
	ID    string
	items []Item // slice: array de longitud variable
	total float64
}

func (o *Order) AddItem(i Item, price float64) error {
	if len(o.items) >= maxItems {
		return errors.New("the order can not hold any more items")
	}
	o.items = append(o.items, i)
	o.total += price

	return nil
}
