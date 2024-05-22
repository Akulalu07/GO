package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "piww7305"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	createTable(db)
	insertEmployee(db, "Alice", "Engineer", 70000)
	insertEmployee(db, "Bob", "Manager", 80000)
	getEmployees(db)
	updateEmployeeSalary(db, 1, 75000)
	deleteEmployee(db, 2)
	getEmployees(db)
}

func createTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS employees (
        id SERIAL PRIMARY KEY,
        name TEXT,
        position TEXT,
        salary NUMERIC
    );`

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Table created successfully")
}

func insertEmployee(db *sql.DB, name, position string, salary float64) {
	query := `
    INSERT INTO employees (name, position, salary)
    VALUES ($1, $2, $3)
    RETURNING id`

	var id int
	err := db.QueryRow(query, name, position, salary).Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Inserted employee with ID: %d\n", id)
}

func getEmployees(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, position, salary FROM employees")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, position string
		var salary float64

		err = rows.Scan(&id, &name, &position, &salary)
		if err != nil {
			panic(err)
		}

		fmt.Printf("ID: %d, Name: %s, Position: %s, Salary: %.2f\n", id, name, position, salary)
	}
}

func updateEmployeeSalary(db *sql.DB, id int, salary float64) {
	query := `
    UPDATE employees
    SET salary = $1
    WHERE id = $2`

	_, err := db.Exec(query, salary, id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Salary updated successfully")
}

func deleteEmployee(db *sql.DB, id int) {
	query := `
    DELETE FROM employees
    WHERE id = $1`

	_, err := db.Exec(query, id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Employee deleted successfully")
}
