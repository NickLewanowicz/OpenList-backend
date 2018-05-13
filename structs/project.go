package structs

import (
	"fmt"
)

//Project struct
type Project struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Created  string    `json:"created"`
	Owners   []User    `json:"owners"`
	Sections []Section `json:"sections"`
}

func (p Project) insertSQL() {
	fmt.Printf("    - Inserting project '" + p.ID + "' into database. ")
	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s VALUES (%s,%s,%s)", projTable, p.ID, p.Title, p.Created))
	didError(err)
	for i := 0; i < len(p.Owners); i++ {
		fmt.Printf("    - Inserting proj/owner relationship into database. ")
		_, err = db.Exec(fmt.Sprintf("INSERT INTO %s VALUES (%s,%s)", userProjTable, p.ID, p.Owners[i].ID))
		didError(err)
	}
}

func (p Project) updateSQL() {
	fmt.Printf("    - Updating project '" + p.ID + "' in database. ")
	_, err = db.Exec(fmt.Sprintf("UPDATE %s SET title='%s' created='%s'", projTable, p.Title, p.Created))
	didError(err)
}
