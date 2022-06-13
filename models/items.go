package models

import "gorm.io/gorm"

type Items struct {
	ID                         uint    `gorm:"primary key;autoIncrement" json:"id"`
	List_Item_Name             *string `json:"list_item_name"`
	List_Item_Type             *string `json:"list_item_type"`
	List_Item_Price            *int    `json:"list_item_price"`
	List_Item_Options_Colors   *string `json:"list_item_options_colors"`
	List_Item_Options_PCB      *string `json:"list_item_options_pcb"`
	List_Item_Options_Plate    *string `json:"list_item_options_plate"`
	List_Item_Options_Switches *string `json:"list_item_options_switches"`
	List_Item_Desc             *string `json:"list_item_desc"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Items{})
	return err
}
