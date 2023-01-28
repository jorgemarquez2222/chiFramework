package products

type Produt struct {
	Products []DataPerson `json:"products"`
}
type DataPerson struct {
	IDProduct       int     `json:"id_producto"`
	Nombre          string  `json:"nombre"`
	PrecioPacaDolar float32 `json:"precio_paca_dolar"`
	CantidadProPaca float32 `json:"cantidad_por_paca"`
	PorcentDoalr    float32 `json:"porcent_dolar"`
	PorcentEfect    float32 `json:"porcent_efect"`
	PorcentPunto    float32 `json:"porcent_punto"`
	Cantidad        int     `json:"cantidad"`
}

func (d *DataPerson) SetCantidad(cantidad int) {
	d.Cantidad = cantidad
}

func (c Produt) FindProductIndex(idProduct int) int {
	for i, v := range c.Products {
		if v.IDProduct == idProduct {
			return i
		}
	}
	return -1
}

func (d *Produt) RemoveProduct(idProduct int) {
	index := d.FindProductIndex(idProduct)
	if index < 0 {
		return
	}
	d.Products = append(d.Products[:index], d.Products[index+1:]...)
}

func (p *Produt) AddElement(idProduct int) {
	var pr = DataPerson{
		IDProduct:       idProduct,
		Nombre:          "Jorge",
		PrecioPacaDolar: 1,
		CantidadProPaca: 1,
		PorcentDoalr:    1,
		PorcentEfect:    1,
		PorcentPunto:    1,
		Cantidad:        1,
	}
	p.Products = append(p.Products, pr)
}

func (c *Produt) AddAllQauntity(value int) {
	for i, _ := range c.Products {
		c.Products[i].SetCantidad(c.Products[i].IDProduct * value)
	}
}
