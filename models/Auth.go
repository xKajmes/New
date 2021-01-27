package models

import (
	"fmt"
	"new/config"
)

// KeyDB ...
func KeyDB(v *Verify) (err error){
	
	if err = config.DB.Create(v).Error; err != nil {
		return err
	}
	return nil
}
// LookVerify ...
func LookVerify(id uint)(err bool){
	var look bool
	config.DB.Raw("SELECT verify FROM User WHERE id = ?", &id).Scan(&look) // Pobranie id uzytkownika do verify
	fmt.Println("auuu")
	fmt.Println(look)
	return look
	
}
