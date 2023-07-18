package main

// import (
// 	"fmt"
// )

// func Database(act string, ID int) []*Note {
// 	switch act {
// 	case "GET":
// 		var search string
// 		if ID != 0 {
// 			search = fmt.Sprintf("WHERE %d", ID)
// 		}
// 		rows, err := db.Query("Select * From Note %s", search)
// 		errlog(err)
// 		defer rows.Close()

// 		var notes []*Note
// 		for rows.Next() {
// 			note := &Note{}
// 			err = rows.Scan(&n.ID, &n.Title, &n.Text)
// 			errlog(err)
// 			notes = append(notes, note)
// 		}
// 		err = rows.Err()
// 		errlog(err)
// 		return notes

// 	case "POST":
// 		_, err = db.Exec("INSERT INTO Note (`list`,`log`) VALUES (`%s`,`%s`)", n.Title, n.Text)
// 		errlog(err)

// 	case "PUT":
// 		_, err = db.Exec("UPDATE Note SET list = %s,log = %s WHERE id IS %d", n.Title, n.Text, ID)
// 		errlog(err)

// 	case "PATCH":
// 		var field, patch string
// 		if n.Title != "" {
// 			field = "list"
// 			patch = n.Title
// 		} else if n.Text != "" {
// 			field = "Text"
// 			patch = n.Text
// 		}
// 		_, err = db.Exec("UPDATE Note SET %s = %s WHERE id IS %d", field, patch, ID)
// 		errlog(err)

// 	case "DELETE":
// 		_, err = db.Exec("DELETE FROM Note WHERE id = %d", ID)
// 		errlog(err)

// 	}
// 	return nil
// }
