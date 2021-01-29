package models

import (
	"gorm.io/gorm"
)

// Student ...
type Student struct {	
	gorm.Model
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
	Login   string `json:"login"`
	Password   string `json:"password"`
	Role string `gorm:"default:student" json:"role"` 
	Verify bool `gorm:"default:false" json:"verify"` 
	
}
// TableName ...
 func (b *User) TableName() string {
 	return "User"
 }

// Class ...
type Class struct {	
	gorm.Model
	ID_Tutor   uint64 `json:"id_tutor"`
	Name   string `json:"name"`

}
// TableName ...
 func (b *Class) TableName() string {
 	return "Class"
 }
// Subject ...
type Subject struct {	
	gorm.Model

	ID_Teacher   uint64 `json:"id_teacher"`
	Name   string `json:"name"`
}
// TableName ...
 func (b *Subject) TableName() string {
 	return "Subject"
 }

 // Grade ...
type Grade struct {	
	gorm.Model
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
	KEY string `json:"key"`
}
// TableName ...
 func (b *Verify) TableName() string {
 	return "Verify"
 }


