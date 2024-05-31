package database

import (
	"database/sql"
	"ejaw_test_case/pkg/config"
	"ejaw_test_case/pkg/utils"
	"fmt"
)

func InitDB(db *sql.DB) error {
	createUsersTable := `
	 CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			role VARCHAR(100) NOT NULL
		);
	`
	if _, err := db.Exec(createUsersTable); err != nil {
		return fmt.Errorf("error creating users table: %v", err)
	}

	var userCount int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&userCount)
	if err != nil {
		return fmt.Errorf("error checking users table: %v", err)
	}

	username := config.Get().DefaultAdminUsername
	password := config.Get().DefaultAdminPassword
	role := "admin"

	if userCount == 0 {
		password, err = utils.HashPassword(password)
		if err != nil {
			return fmt.Errorf("error hashing password: %v", err)
		}
		_, err = db.Exec("INSERT INTO users (username, password, role) VALUES ($1, $2, $3)",
			username, password, role)
		if err != nil {
			return fmt.Errorf("error inserting default admin user: %v", err)
		}
	}

	query := `
    CREATE TABLE IF NOT EXISTS sellers (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        phone VARCHAR(20) NOT NULL UNIQUE
    );

    CREATE TABLE IF NOT EXISTS products (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        description TEXT NOT NULL,
        price DECIMAL(10, 2) NOT NULL,  
        seller_id INTEGER NOT NULL,
        FOREIGN KEY (seller_id) REFERENCES sellers(id)
    );

    CREATE TABLE IF NOT EXISTS customers (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        phone VARCHAR(20) NOT NULL UNIQUE
    );

    CREATE TABLE IF NOT EXISTS orders (
        id SERIAL PRIMARY KEY,
        customer_id INTEGER NOT NULL,
        order_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (customer_id) REFERENCES customers(id)
    );

    CREATE TABLE IF NOT EXISTS order_items (
        order_id INTEGER NOT NULL,
        product_id INTEGER NOT NULL,
        quantity INTEGER DEFAULT 1,
        PRIMARY KEY (order_id, product_id),
        FOREIGN KEY (order_id) REFERENCES orders(id),
        FOREIGN KEY (product_id) REFERENCES products(id)
    );
    `
	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("error creating sellers, products, customers, orders, order_items tables: %v", err)
	}
	return nil
}
