package models

import (
	"fmt"
	"go/rest-api/db"
)

type Event struct {
	ID          int64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Location   string `json:"location"`
	UserID      int64
}

var events = []Event{}


func (e *Event) Save() error{
// add to database later
query := `INSERT INTO events(title,description,location,date,user_id) 
 VALUES (?,?,?,?,?)`

 statement , err := db.DB.Prepare(query)
 if err!=nil {
	fmt.Print("problem with db")
	return err
 }

 defer statement.Close()
result , err := statement.Exec(e.Title , e.Description,e.Location,e.Date, e.UserID)
if err!=nil {
	return err
 }

 id , err := result.LastInsertId()
 e.ID = id
return err
}

func GetAllEvents() ([]Event , error){
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)

	if err!=nil {
		return nil, err
	 }

defer rows.Close()
var events []Event

for rows.Next(){
	var event Event
    err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Location, &event.UserID)
    if err!=nil {
        return nil,err
    }
    events = append(events, event)  //append each row to events slice
}

return events , nil
}

func GetEventByID (id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id=?"
    row := db.DB.QueryRow(query, id)
    var event Event
    err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Date, &event.Location, &event.UserID)
    if err!=nil {
        fmt.Println("No such event")
        return nil, err
    }
    // fmt.Println(event)
	return &event, nil
}

func (event Event) Update() error{
	query := `
	UPDATE events
	SET title = ? , description = ?, date = ? , location = ?
	WHERE id =?
	`

	statement , err := db.DB.Prepare(query)
	if err!= nil {
		fmt.Print("problem with db")
        return err
	}

	defer statement.Close()
	_, err = statement.Exec(event.Title, event.Description, event.Location, event.Date, event.ID)

	return err
}



func (event  Event) Delete() error{

	query := `DELETE FROM events WHERE id =?`

	statement, err := db.DB.Prepare(query)
	if err!= nil {
		fmt.Print("")
        return err
	}

	defer statement.Close()
	_, err = statement.Exec(event.ID)
	return err

}

func (e Event ) Register(userId int64) error{
	query := "INSERT INTO registrations(event_id , user_id) VALUES (?,?)"
	statement , err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()
	_,err = statement.Exec(e.ID , userId)

	return err
}

func (e Event) CancelRegistration(userId int64) error{
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?  "
	statement , err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()
	_,err = statement.Exec(e.ID , userId)

	return err


}



//use Exec what you want to change data in DB
//Query when you want to fetch data from DB