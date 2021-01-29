package controllers

import (
	"fmt"
	"new/config"
	"new/models"
	"strconv"

	"github.com/gin-gonic/gin"
	//"../Translate"
)

//	Zarządzanie Studetami

// AddStudent ...
func AddStudent(c *gin.Context) {	
	var student models.Student
	c.BindJSON(&student)
	err := models.AddStudent(&student)
	if err != nil {
		c.JSON(404, gin.H{"error:": "cos nie tak"})
	} else {
		// role := LookRole(c)
		// if(role=="dziekan" || role=="admin"){
			c.JSON(200, gin.H{"stworzono:": "pomyslnie"})
		// }

		
	}

}
// ListStudent ...
func ListStudent(c *gin.Context) {
	var student []models.Student
	err := models.GetAllStudent(&student)
	if err != nil {
		c.JSON(404, gin.H{"error:": "Nie mozna pobrac studentow"})
	} else {
			c.JSON(200, student )
		
	}
}
// DeleteStudent ...
func DeleteStudent(c *gin.Context) {	// Dodalem role 
	id := c.Params.ByName("id")
	idusun, _ := strconv.Atoi(id)
	role := LookRole(c)
	if(role=="dziekan" || role=="admin"){
		err := models.DeleteStudent(idusun) 
		if err != nil {
			c.JSON(404, gin.H{"error": "Cos poszlo nie tak"})
		} else {
			c.JSON(200, gin.H{"Deleted": id})
		}
	}
}
// EditStudent ...
func EditStudent(c *gin.Context) {		
	role := LookRole(c)
	myid := LookToken(c)
	// Cos jest nie tak w jquery z pobieraniem nowej wartosci input'ow
	id := c.Params.ByName("id")
	idedit, _ := strconv.Atoi(id)
	if(role=="dziekan" || role=="admin" || int(myid)==idedit){
		var student models.Student
		var newstudent models.Student
		c.BindJSON(&newstudent)
		err := models.EditStudent(&student,&newstudent,idedit)
		if err != nil {
			c.JSON(404, gin.H{"error": "Cos poszlo nie tak"})
		} else {
			c.JSON(200, gin.H{"Pomyslnie edytowano:": id})
		}
	}
}

// CompleteMe ...
func CompleteMe(c *gin.Context) {		// Edytowanie studenta : Dziala takze po wpisaniu jedengo elementu
	// Cos jest nie tak w jquery z pobieraniem nowej wartosci input'ow
	id := LookToken(c)
	var student models.Student
	var newstudent models.Student
	c.BindJSON(&newstudent)
	err := models.EditStudent(&student,&newstudent,int(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "Cos poszlo nie tak"})
	} else {
		c.JSON(200, gin.H{"error": "Pomyslnie edytowano!"})
	}
}
// MyID ...
func MyID(c *gin.Context) {
	id := LookToken(c)
	c.JSON(200, gin.H{"id": id})
	fmt.Printf("TWOJE ID TO :")
	fmt.Println(id)
}
// GetOne ...
func GetOne(c *gin.Context) {	// Pobieranie danych jednego studenta
	id := c.Params.ByName("id")
	idstudent, _ := strconv.Atoi(id)
	var student models.Student
	err :=models.GetStudent(&student,idstudent)
	if err != nil {
		c.JSON(404, gin.H{"error:": "Nie mozna pobrac studentow"})
	} else {
		c.JSON(200, student )
	}
}

// LookRole ...
func LookRole(c *gin.Context) (string) {
	iduser := LookToken(c)
	var role string
	config.DB.Raw("SELECT role FROM User WHERE id = ?", &iduser).Scan(&role) 
	return role
} 
// LookRolePage ...
func LookRolePage(c *gin.Context){
	role := LookRole(c);
	fmt.Println(role)
	c.JSON(200, role)
} 

// LookClass ...
func LookClass(c *gin.Context) {	// Pobieranie klasy dla nauczyciela
	iduser := LookToken(c)
	idclass := 0;
	config.DB.Raw("SELECT id FROM Class WHERE id_tutor = ?", &iduser).Scan(&idclass) 
	var student []models.Student
	config.DB.Raw("SELECT * FROM Student WHERE id_class = ?", &idclass).Scan(&student) 
	c.JSON(200, student )
	
} 
// Verification ...

// LookGrades ...
func LookGrades(c *gin.Context) {	// Pobieranie danych jednego studenta
	idtutor := LookToken(c)
	id := c.Params.ByName("id")
	idstudent, _ := strconv.Atoi(id)
	var grades []models.Grade
	err := models.GetGrades(&grades,idtutor,idstudent)
	if err != nil {
		c.JSON(404, gin.H{"error:": "Nie mozna pobrac ocen"})
	} else {
		c.JSON(200, grades )
	}
}

// AddGrade ...
func AddGrade(c *gin.Context) {	// Pobieranie danych jednego studenta
	idtutor := LookToken(c)
	idnew := uint64(idtutor)
	// id := c.Params.ByName("id")
	// idstudent, _ := strconv.Atoi(id)
	var grades models.Grade
	c.BindJSON(&grades)
	grades.ID_Teacher = idnew
	err := models.AddGrade(&grades)
	if err != nil {
		c.JSON(404, gin.H{"error:": "Nie mozna dodac oceny"})
	} else {
		c.JSON(200, grades )

	}
}
// DelGrade ...
func DelGrade(c *gin.Context) {	// Pobieranie danych jednego studenta
	id := c.Params.ByName("id")
	idusun, _ := strconv.Atoi(id)
	err := models.DelGrade(idusun)
	if err != nil {
		c.JSON(404, gin.H{"error:": "Nie mozna dodac oceny"})
	} else {
		c.JSON(200, gin.H{"Komunikat:": "Usunięto ocene" } )

	}
}
