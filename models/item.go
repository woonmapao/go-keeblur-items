package models

type Item struct {
	ID          string  `json:"id,omitempty" bson:"_id,omitempty"`
	ItemName    string  `json:"item_name" bson:"item_name"`
	ItemType    string  `json:"item_type" bson:"item_type"`
	ItemPrice   float64 `json:"item_price" bson:"item_price"`
	ItemOptions struct {
		Colors   []string `json:"colors,omitempty" bson:"colors,omitempty"`
		Plate    []string `json:"plate,omitempty" bson:"plate,omitempty"`
		Pcb      []string `json:"pcb,omitempty" bson:"pcb,omitempty"`
		Switches []string `json:"switches,omitempty" bson:"switches,omitempty"`
	} `json:"item_options" bson:"item_options"`
	ItemDesc string `json:"item_desc" bson:"item_desc"`
}
