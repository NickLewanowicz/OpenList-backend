package structs

import (
	"fmt"
)

//Section struct
type Section struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Position string `json:"position"`
	Tasks    []Task `json:"tasks"`
}

func (s Section) insertSQL() {
	fmt.Printf("    - Inserting project '" + s.ID + "' into database. ")
	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s VALUES (%s,%s,%s)", sectTable, s.ID, s.Title, s.Position))
	didError(err)
	for i := 0; i < len(s.Tasks); i++ {
		fmt.Printf("    - Inserting task/section relationship into database. ")
		_, err = db.Exec(fmt.Sprintf("INSERT INTO %s VALUES (%s,%s)", taskSectTable, s.ID, s.Tasks[i].ID))
		didError(err)
	}
}

func (s Section) updateSQL() {
	fmt.Printf("    - Updating project '" + s.ID + "' in database. ")
	_, err = db.Exec(fmt.Sprintf("UPDATE %s SET title='%s' position='%s'", projTable, s.Title, s.Position))
	didError(err)
}
