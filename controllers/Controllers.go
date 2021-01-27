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
func AddStudent(c *gin.Context) {		// Proces dodawania studenta
	var student models.Student
	c.BindJSON(&student)
	err := models.AddStudent(&student)
	if err != nil {
		c.JSON(404, gin.H{"error:": "cos nie tak"})
	} else {
		// role := LookRole(c)
		
		// if(role=="dziekan" || role=="admin"){
			c.JSON(200, gin.H{"stworzono:": "pomyslnbie"})
		// }

		
	}

}
// ListStudent ...
func ListStudent(c *gin.Context) {	// Pobieranie wszystkich studentow
	var student []models.Student
	err := models.GetAllStudent(&student)
	if err != nil {
		c.JSON(404, gin.H{"error:": "Nie mozna pobrac studentow"})
	} else {
		

			c.JSON(200, student )
		
	}
}
// DeleteStudent ...
func DeleteStudent(c *gin.Context) {	// Usuwanie Studnetow
	// //var student Models.Student
	id := c.Params.ByName("id")
	idusun, _ := strconv.Atoi(id)
	
	err := models.DeleteStudent(idusun) // Usuwanie studenta o id=idusun
		
	if err != nil {
		c.JSON(404, gin.H{"error": "Cos poszlo nie tak"})
	} else {

		role := LookRole(c)
		
		if(role=="dziekan" || role=="admin"){
			c.JSON(200, gin.H{"Deleted": id})
		}
		
	}
}
// EditStudent ...
func EditStudent(c *gin.Context) {		// Edytowanie studenta : Dziala takze po wpisaniu jedengo elementu
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
	config.DB.Raw("SELECT role FROM User WHERE id = ?", &iduser).Scan(&role) // Pobranie passwordhash z database
	//fmt.Println(role)
	return role
} 

// LookClass ...
func LookClass(c *gin.Context) {
	iduser := LookToken(c)
	idclass := 0;
	config.DB.Raw("SELECT id FROM Class WHERE id_tutor = ?", &iduser).Scan(&idclass) // Pobranie passwordhash z database
	var student []models.Student
	config.DB.Raw("SELECT * FROM Student WHERE id_class = ?", &idclass).Scan(&student) // Pobranie passwordhash z database
	c.JSON(200, student )
	
} 
// Verification ...

// LookGrades ...
func LookGrades(c *gin.Context) {	// Pobieranie danych jednego studenta
	idtutor := LookToken(c)
	id := c.Params.ByName("id")
	idstudent, _ := strconv.Atoi(id)
	var grades []models.Grade
	fmt.Println("LISTA W LOOKGRADES")
	fmt.Println(idtutor)
	fmt.Println(idstudent)
	fmt.Println("---------------")
	err := models.GetGrades(&grades,idtutor,idstudent)
	if err != nil {
		c.JSON(404, gin.H{"error:": "Nie mozna pobrac ocen"})
	} else {
		c.JSON(200, grades )

	}
}

// AddGrade ...
func AddGrade(c *gin.Context) {	// Pobieranie danych jednego studenta
	// idtutor := LookToken(c)
	// id := c.Params.ByName("id")
	// idstudent, _ := strconv.Atoi(id)
	var grades models.Grade
	c.BindJSON(&grades)
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
