package structs

import (
	"fmt"
)

//Task struct
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"list"`
	Description string `json:"description"`
	Created     string `json:"created"`
	Due         string `json:"due"`
	Section     string `json:"section"`
	Owners      []User `json:"owners"`
}

func (t Task) insertSQL() {
	fmt.Printf("    - Inserting task '" + t.ID + "' into database. ")
	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s VALUES (%s,%s,%s,%s,%s,%s,%s)", taskTable, t.ID, t.Title, t.Description, t.Created, t.Due, t.Owners, t.Section))
	didError(err)
	for i := 0; i < len(t.Owners); i++ {
		fmt.Printf("    - Inserting task/owner relationship into database. ")
		_, err = db.Exec(fmt.Sprintf("INSERT INTO %s VALUES (%s,%s)", userTaskTable, t.ID, t.Owners[i].ID))
		didError(err)
	}
}

func (t Task) updateSQL() {
	fmt.Printf("    - Updating task '" + t.ID + "' in database. ")
	_, err = db.Exec(fmt.Sprintf("UPDATE %s SET title='%s' description='%s' created='%s' due='%s' WHERE id='%s'", taskTable, t.Title, t.Description, t.Created, t.Due, t.ID))
	didError(err)
}
