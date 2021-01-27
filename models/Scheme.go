package models

import (
	"gorm.io/gorm"
)

// Student ...
type Student struct {	
	gorm.Model
	//ID uint64 `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Surname  string `json:"surname"`
	Age string `json:"age"`
	Role string `gorm:"default:student" json:"role"` 
	ID_Class uint64 `json:"id_class"` 
}
// TableName ...
func (b *Student) TableName() string {
	return "Student"
}
	
// User ...
type User struct {	
	gorm.Model
	//ID uint64 `gorm:"primaryKey" json:"id"`
	Login   string `json:"login"`
	Password   string `json:"password"`
	Role string `gorm:"default:student" json:"role"` 
	Verify bool `gorm:"default:false" json:"verify"` 


	// Name  string `json:"name"`
	// Surname  string `json:"surname"`
	// Age string `json:"age"`
	// Role string `gorm:"default:student" json:"role"` 
	
}
// TableName ...
 func (b *User) TableName() string {
 	return "User"
 }

// Class ...
type Class struct {	
	gorm.Model
	//ID uint64 `gorm:"primaryKey" json:"id"`
	ID_Tutor   uint64 `json:"id_tutor"`
	Name   string `json:"name"`

	// Name  string `json:"name"`

}
// TableName ...
 func (b *Class) TableName() string {
 	return "Class"
 }
// Subject ...
type Subject struct {	
	gorm.Model
	//ID uint64 `gorm:"primaryKey" json:"id"`
	ID_Teacher   uint64 `json:"id_teacher"`
	Name   string `json:"name"`

	// Name  string `json:"name"`

}
// TableName ...
 func (b *Subject) TableName() string {
 	return "Subject"
 }

 // Grade ...
type Grade struct {	
	gorm.Model
	//ID uint64 `gorm:"primaryKey" json:"id"`
	ID_Teacher   uint64 `json:"id_teacher"`
	ID_Student   uint64 `json:"id_student"`
	Number   uint64 `json:"number"`
	Comment string `json:"comment"`

	// Name  string `json:"name"`

}
// TableName ...
 func (b *Grade) TableName() string {
 	return "Grade"
 }
// Verify ...
type Verify struct {	
	gorm.Model
	//ID uint64 `gorm:"primaryKey" json:"id"`
	KEY string `json:"key"`


	// Name  string `json:"name"`

}
// TableName ...
 func (b *Verify) TableName() string {
 	return "Verify"
 }


