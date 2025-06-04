package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/seyamibrahim/students-api/internal/config"
	"github.com/seyamibrahim/students-api/internal/types"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)

	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
	
	)`)
	if err != nil {
		return nil, err
	}
	return &Sqlite{Db: db}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {

	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?,?,?)")

	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM students WHERE id = ? LIMIT 1")

	if err != nil {
		return types.Student{}, err

	}

	defer stmt.Close()

	var student types.Student
	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("not student found with id %s", fmt.Sprint(id))
		}

		return types.Student{}, fmt.Errorf("query error : %s", err)
	}

	return student, err
}

func (s *Sqlite) GetStudents() ([]types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email,age FROM students ")

	if err != nil {
		return nil, err

	}

	defer stmt.Close()

	var students []types.Student
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var student types.Student
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, err

		}
		students = append(students, student)

	}
	return students, nil
}

func (s *Sqlite) DeleteStudent(id int64) error{

	stmt, err := s.Db.Prepare("DELETE FROM students WHERE id == ?")


	if err != nil{
		return fmt.Errorf("failed to prepare")
	}


	defer stmt.Close()

	res, err := stmt.Exec(id)

	if err != nil {
		return  fmt.Errorf("failed to execute delete %w", err)
	}


	rowsaffect, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retreive rowaffect %w", err)
	}

	if rowsaffect == 0{
		return fmt.Errorf("no student found with id %d", id)
	}

	return  nil

}