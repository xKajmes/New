package models

import (
	"fmt"
	"new/config"
)

// GetGrades ...
func GetGrades(g *[]Grade, idtutor uint, idstudent int) (err error){
	
	//config.DB.Raw("SELECT * FROM Grade WHERE id_tutor = ? AND id_student = ?", idtutor, idstudent).Scan(&g).Error
	// if err = config.DB.Where("id_tutor = ? AND id_student = ?", idtutor, idstudent).Find(g).Error; err != nil {
	// 	return err
	// }
	// return nil
	fmt.Println("LISTA W GET GRADES")
	fmt.Println(idtutor)
	fmt.Println(idstudent)
	fmt.Println("---------------")
	config.DB.Where("id_teacher = ? AND id_student = ?", idtutor, idstudent).Find(g)
	return nil
}
// AddGrade ...
func AddGrade(g *Grade) (err error){	// Pobieranie danych jednego studenta
	// config.DB.Create(&g)
	if err = config.DB.Create(&g).Error; err != nil {
		return err
	}
	return nil
}
// DelGrade ...
func DelGrade(id int) (err error){	// Pobieranie danych jednego studenta
	// config.DB.Create(&g)
	if err =  config.DB.Delete(&Grade{}, id).Error; err != nil {
		return err
	}
	return nil
}