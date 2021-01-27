package models

import (
	"fmt"
	"new/config"
)

// GetAllStudent ...
func GetAllStudent(s *[]Student) (err error) {
	fmt.Printf("AUUUU MAMY TO 1")
	// var c *gin.Context
	// role := config.LookRole(c)
	// teacherid := config.LookToken(c)
	// if(role=="teacher"){
	// 	fmt.Printf("AUUUU MAMY TO 2")
	// 	class := &Class{}
	// 	//config.DB.Where("id_tutor <> ?", teacherid).Find(&class)
	// 	idclass := config.DB.Select("id").Where("id_tutor <> ?", teacherid).First(&class)
	// 	if err = config.DB.Where("id_class <> ?", idclass).Find(s).Error; err != nil {
	// 		return err
	// 	}
	// 	return nil
	// }
	if err = config.DB.Find(s).Error; err != nil {
		return err
	}
	return nil
}
// AddStudent ..
func AddStudent(s *Student) (err error){
	if err = config.DB.Create(s).Error; err != nil {
		return err
	}
	return nil
}
// DeleteStudent ..
func DeleteStudent(idusun int) (err error){
	if err =  config.DB.Delete(&Student{}, idusun).Error; err != nil {
		return err
	}
	if err =  config.DB.Delete(&User{}, idusun).Error; err != nil {
		return err
	}
	return nil
}
// EditStudent ..
func EditStudent(s *Student,new *Student, id int) (err error){
	if err = config.DB.Model(s).Where("ID=?", id).Updates(new).Error; err != nil {
		return err
	}
	return nil
}
// GetStudent ..
func GetStudent(s *Student, id int) (err error){
	if err = config.DB.First(s, id).Error; err != nil {
		return err
	}
	return nil
}

// Login ...
func Login(u *User) (err string){
	nohash := u.Password
	hash, _ := config.HashPassword(nohash)
	u.Password = hash
	szukaj := config.DB.Where("login = ?", u.Login).Find(&u)	// Jeżeli użytkownik istnieje
	if(szukaj.RowsAffected==1){
		var password1 string
		config.DB.Raw("SELECT password FROM User WHERE login = ?", &u.Login).Scan(&password1) // Pobranie passwordhash z database
		match := config.CheckPasswordHash(nohash, password1)
		if(match==true){	// Jeżeli hasła się zgadzają
			
			// token, err := auth.CreateToken(u.ID)
			// if err != nil {
			//    return "Błąd z tokenem"
			// }

			
			return "Zalogowano"
			//fmt.Fprintf("Token: "+ token)
		}
			return "Hasla sie nie zgadzaja"
	}
		return "Uzytkownik nie istnieje"

	// if err = Config.DB.First(s, id).Error; err != nil {
	// 	return err
	// }
	 
}
// Register ..
func Register(u *User) (err string){
	nohash := u.Password
	hash, _ := config.HashPassword(nohash)
	u.Password = hash
	szukaj := config.DB.Where("login = ?", u.Login).Find(&u)	// Jeżeli użytkownik istnieje
	if(szukaj.RowsAffected==0){
		config.DB.Create(&u)
		//err = "Zarejestrowano"
		
		return "Zarejestrowano"
	}
		//err = "Podany login juz istnieje"
		return "Podany login juz istnieje"
}
// EditUser ...
func EditUser(u *User,new *User, id int) (err error){
	if err = config.DB.Model(u).Where("ID=?", id).Updates(new).Error; err != nil {
		return err
	}
	return nil
}
