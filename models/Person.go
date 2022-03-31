package models

import (
	"github.com/jinzhu/gorm"
	sql "goweb/database"
)

type Person struct {
	gorm.Model
	Id        int    `gorm:"primary_key" json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

func (p *Person) AddPerson() (id int64, err error) {
	rs := sql.GetDb().Create(p)
	//rs, err := db.SqlDB.Exec("INSERT INTO person(first_name, last_name) VALUES (?, ?)", p.FirstName, p.LastName)
	println(rs.Value)
	return
}

func (p *Person) ModPerson() (ra int64, err error) {
	rs := sql.GetDb().Exec("UPDATE person SET first_name = ?, last_name = ? WHERE id = ?", p.FirstName, p.LastName, p.Id)
	if err != nil {
		return
	}
	ra = rs.RowsAffected
	return
}

func (p *Person) DelPerson() (ra int64, err error) {
	rs := sql.GetDb().Exec("DELETE FROM person WHERE id = ?", p.Id)
	if err != nil {
		return
	}
	ra = rs.RowsAffected
	return
}

func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)
	rows, err := sql.GetDb().Raw("SELECT id, first_name, last_name FROM person").Rows()
	println(err)
	defer rows.Close()

	if err != nil {
		return
	} else {
		println(err)
	}

	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return
	} else {
		println(err)
	}
	return
}

// get person
func (p *Person) GetPerson() (err error) {
	res := sql.GetDb().Raw("SELECT id, first_name, last_name FROM person WHERE id=?", p.Id)
	scan := res.Scan(&p)
	println(scan)
	//res.Scan(
	//	&p.Id,
	//	&p.FirstName,
	//	&p.LastName,
	//)
	return
}
