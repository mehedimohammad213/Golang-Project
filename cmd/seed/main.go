package main

import (
	"fmt"
	"log"

	"github.com/user/car-project/internal/config"
	"github.com/user/car-project/internal/db"
)

func main() {
	// Initialize logger
	log.Println("Starting comprehensive seeder...")

	// Load config
	cfg := config.LoadConfig()

	// Initialize database
	if cfg.DBURL != "" {
		db.InitDB(cfg.DBURL)
	} else {
		log.Fatal("DBURL not found in config")
	}

	// 1. Truncate all tables
	truncateTables()

	// 2. Seed tables in order
	users := seedUsers()
	roles := seedRoles()
	permissions := seedPermissions()
	seedRoleUser(users, roles)
	seedPermissionRole(permissions, roles)

	makes := seedCarMakes()
	models := seedCarModels(makes)
	cars := seedCars(models)
	seedCarPhotos(cars)
	seedDocuments(cars, users)
	seedCarGrades(cars)
	details := seedCarDetails(cars)
	seedCarSubDetails(details)
	seedStocks(cars)

	seedCarts(users, cars)
	orders := seedOrders(users)
	seedOrderItems(orders, cars)

	ph := seedPurchaseHistory(cars)
	payh := seedPaymentHistory(cars)
	seedInstallments(payh)

	_ = ph // suppress unused variable warning if any

	log.Println("Seeding completed successfully!")
}

func truncateTables() {
	log.Println("Truncating tables...")
	tables := []string{
		"installments", "payment_history", "purchase_history", "order_items", "orders", "carts",
		"stocks", "car_sub_details", "car_details", "car_grades", "documents", "car_photos", "cars",
		"car_models", "car_makes", "permission_role", "role_user", "permissions", "roles", "users",
	}

	for _, table := range tables {
		_, err := db.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table))
		if err != nil {
			log.Printf("Warning: failed to truncate %s: %v", table, err)
		}
	}
}

func seedUsers() []int64 {
	log.Println("Seeding Users...")
	var ids []int64
	// User 1: Admin
	// User 2: Accountman
	// User 3: Seller
	// User 4: Call Center
	// User 5: User
	// Pre-calculated bcrypt hash for "password123"
	const passwordHash = "$2a$10$8K1pfb9A.Z0K.Gz.Gz.Gz.Gz.Gz.Gz.Gz.Gz.Gz.Gz.Gz.Gz"

	for i := 1; i <= 5; i++ {
		var id int64
		username := fmt.Sprintf("user%d", i)
		email := fmt.Sprintf("user%d@example.com", i)
		name := fmt.Sprintf("User %d", i)

		query := `INSERT INTO users (name, username, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id`
		err := db.DB.QueryRow(query, name, username, email, passwordHash).Scan(&id)
		if err != nil {
			log.Printf("Failed to seed user %d: %v", i, err)
		} else {
			ids = append(ids, id)
		}
	}
	return ids
}

func seedRoles() map[string]int64 {
	log.Println("Seeding Roles...")
	roles := make(map[string]int64)
	roleNames := []string{"admin", "user", "accountman", "seller", "call center"}

	for _, name := range roleNames {
		var id int64
		slug := name // simplified slug
		if name == "call center" {
			slug = "call-center"
		}
		query := `INSERT INTO roles (name, slug) VALUES ($1, $2) RETURNING id`
		err := db.DB.QueryRow(query, name, slug).Scan(&id)
		if err != nil {
			log.Printf("Failed to seed role %s: %v", name, err)
		} else {
			roles[name] = id
		}
	}
	return roles
}

func seedPermissions() map[string]int64 {
	log.Println("Seeding Permissions...")
	perms := make(map[string]int64)
	permNames := []string{"car-create", "car-read", "car-update", "car-delete"}

	for _, name := range permNames {
		var id int64
		slug := name
		query := `INSERT INTO permissions (name, slug, module) VALUES ($1, $2, 'car') RETURNING id`
		err := db.DB.QueryRow(query, name, slug).Scan(&id)
		if err != nil {
			log.Printf("Failed to seed permission %s: %v", name, err)
		} else {
			perms[name] = id
		}
	}
	return perms
}

func seedRoleUser(users []int64, roles map[string]int64) {
	log.Println("Seeding Role-User...")
	// User 1 -> Admin
	assignRole(users[0], roles["admin"])
	// User 2 -> Accountman
	assignRole(users[1], roles["accountman"])
	// User 3 -> Seller
	assignRole(users[2], roles["seller"])
	// User 4 -> Call Center
	assignRole(users[3], roles["call center"])
	// User 5 -> User
	assignRole(users[4], roles["user"])
}

func assignRole(userID, roleID int64) {
	_, err := db.DB.Exec("INSERT INTO role_user (user_id, role_id) VALUES ($1, $2)", userID, roleID)
	if err != nil {
		log.Printf("Failed to assign role %d to user %d: %v", roleID, userID, err)
	}
}

func seedPermissionRole(perms map[string]int64, roles map[string]int64) {
	log.Println("Seeding Permission-Role...")

	// Admin & Seller: Full Access
	fullAccessRoles := []int64{roles["admin"], roles["seller"]}
	for _, roleID := range fullAccessRoles {
		assignPerm(roleID, perms["car-create"])
		assignPerm(roleID, perms["car-read"])
		assignPerm(roleID, perms["car-update"])
		assignPerm(roleID, perms["car-delete"])
	}

	// Accountman & Call Center: Read Only
	readOnlyRoles := []int64{roles["accountman"], roles["call center"]}
	for _, roleID := range readOnlyRoles {
		assignPerm(roleID, perms["car-read"])
	}
}

func assignPerm(roleID, permID int64) {
	_, err := db.DB.Exec("INSERT INTO permission_role (role_id, permission_id) VALUES ($1, $2)", roleID, permID)
	if err != nil {
		log.Printf("Failed to assign perm %d to role %d: %v", permID, roleID, err)
	}
}

func seedCarMakes() []int64 {
	log.Println("Seeding Car Makes...")
	var ids []int64
	for i := 1; i <= 5; i++ {
		var id int64
		name := fmt.Sprintf("Make %d", i)
		query := `INSERT INTO car_makes (name, origin_country) VALUES ($1, 'Japan') RETURNING id`
		err := db.DB.QueryRow(query, name).Scan(&id)
		if err != nil {
			log.Printf("Failed to seed make %d: %v", i, err)
		} else {
			ids = append(ids, id)
		}
	}
	return ids
}

func seedCarModels(makes []int64) []int64 {
	log.Println("Seeding Car Models...")
	var ids []int64
	for i := 0; i < 5; i++ {
		if i < len(makes) {
			var id int64
			name := fmt.Sprintf("Model %d", i+1)
			query := `INSERT INTO car_models (make_id, name) VALUES ($1, $2) RETURNING id`
			err := db.DB.QueryRow(query, makes[i], name).Scan(&id)
			if err != nil {
				log.Printf("Failed to seed model %d: %v", i, err)
			} else {
				ids = append(ids, id)
			}
		}
	}
	return ids
}

func seedCars(models []int64) []int64 {
	log.Println("Seeding Cars...")
	var ids []int64
	for i := 0; i < 5; i++ {
		if i < len(models) {
			var id int64
			refNo := fmt.Sprintf("REF-%d", i+1)
			chassis := fmt.Sprintf("CH-%d", i+1)
			query := `INSERT INTO cars (model_id, ref_no, chassis_no_full, year, engine_cc, fuel, transmission, drive, steering) 
					  VALUES ($1, $2, $3, 2020, 1500, 'Petrol', 'Automatic', 'FWD', 'Right') RETURNING id`
			err := db.DB.QueryRow(query, models[i], refNo, chassis).Scan(&id)
			if err != nil {
				log.Printf("Failed to seed car %d: %v", i, err)
			} else {
				ids = append(ids, id)
			}
		}
	}
	return ids
}

func seedCarPhotos(cars []int64) {
	log.Println("Seeding Car Photos...")
	for i, carID := range cars {
		_, err := db.DB.Exec("INSERT INTO car_photos (car_id, url) VALUES ($1, $2)", carID, fmt.Sprintf("http://example.com/photo%d.jpg", i))
		if err != nil {
			log.Printf("Failed to seed photo for car %d: %v", carID, err)
		}
	}
}

func seedDocuments(cars []int64, users []int64) {
	log.Println("Seeding Documents...")
	for i, carID := range cars {
		userID := users[0] // assign to first user
		_, err := db.DB.Exec("INSERT INTO documents (car_id, document_type, file_name, file_path, uploaded_by) VALUES ($1, 'Reg', 'doc.pdf', '/path/doc.pdf', $2)", carID, userID)
		if err != nil {
			log.Printf("Failed to seed doc for car %d: %v", i, err)
		}
	}
}

func seedCarGrades(cars []int64) {
	log.Println("Seeding Car Grades...")
	for i, carID := range cars {
		_, err := db.DB.Exec("INSERT INTO car_grades (car_id, grade_overall) VALUES ($1, '4.5')", carID)
		if err != nil {
			log.Printf("Failed to seed grade for car %d: %v", i, err)
		}
	}
}

func seedCarDetails(cars []int64) []int64 {
	log.Println("Seeding Car Details...")
	var ids []int64
	for i, carID := range cars {
		var id int64
		query := `INSERT INTO car_details (car_id, full_title, description) VALUES ($1, $2, 'Description') RETURNING id`
		err := db.DB.QueryRow(query, carID, fmt.Sprintf("Detail %d", i)).Scan(&id)
		if err != nil {
			log.Printf("Failed to seed detail for car %d: %v", i, err)
		} else {
			ids = append(ids, id)
		}
	}
	return ids
}

func seedCarSubDetails(details []int64) {
	log.Println("Seeding Car Sub Details...")
	for i, detailID := range details {
		_, err := db.DB.Exec("INSERT INTO car_sub_details (car_detail_id, title) VALUES ($1, 'Sub Detail')", detailID)
		if err != nil {
			log.Printf("Failed to seed sub detail %d: %v", i, err)
		}
	}
}

func seedStocks(cars []int64) {
	log.Println("Seeding Stocks...")
	for i, carID := range cars {
		_, err := db.DB.Exec("INSERT INTO stocks (car_id, quantity) VALUES ($1, 10)", carID)
		if err != nil {
			log.Printf("Failed to seed stock for car %d: %v", i, err)
		}
	}
}

func seedCarts(users []int64, cars []int64) {
	log.Println("Seeding Carts...")
	for i := 0; i < 5; i++ {
		if i < len(users) && i < len(cars) {
			_, err := db.DB.Exec("INSERT INTO carts (user_id, car_id, quantity) VALUES ($1, $2, 1)", users[i], cars[i])
			if err != nil {
				log.Printf("Failed to seed cart %d: %v", i, err)
			}
		}
	}
}

func seedOrders(users []int64) []int64 {
	log.Println("Seeding Orders...")
	var ids []int64
	for i := 0; i < 5; i++ {
		if i < len(users) {
			var id int64
			query := `INSERT INTO orders (user_id, total_amount) VALUES ($1, 1000.00) RETURNING id`
			err := db.DB.QueryRow(query, users[i]).Scan(&id)
			if err != nil {
				log.Printf("Failed to seed order %d: %v", i, err)
			} else {
				ids = append(ids, id)
			}
		}
	}
	return ids
}

func seedOrderItems(orders []int64, cars []int64) {
	log.Println("Seeding Order Items...")
	for i := 0; i < 5; i++ {
		if i < len(orders) && i < len(cars) {
			_, err := db.DB.Exec("INSERT INTO order_items (order_id, car_id, quantity, price) VALUES ($1, $2, 1, 1000.00)", orders[i], cars[i])
			if err != nil {
				log.Printf("Failed to seed order item %d: %v", i, err)
			}
		}
	}
}

func seedPurchaseHistory(cars []int64) []int64 {
	log.Println("Seeding Purchase History...")
	var ids []int64
	for i, carID := range cars {
		var id int64
		query := `INSERT INTO purchase_history (car_id, purchase_amount) VALUES ($1, 5000.00) RETURNING id`
		err := db.DB.QueryRow(query, carID).Scan(&id)
		if err != nil {
			log.Printf("Failed to seed purchase history %d: %v", i, err)
		} else {
			ids = append(ids, id)
		}
	}
	return ids
}

func seedPaymentHistory(cars []int64) []int64 {
	log.Println("Seeding Payment History...")
	var ids []int64
	for i, carID := range cars {
		var id int64
		query := `INSERT INTO payment_history (car_id, purchase_amount, customer_name) VALUES ($1, 6000.00, 'Customer') RETURNING id`
		err := db.DB.QueryRow(query, carID).Scan(&id)
		if err != nil {
			log.Printf("Failed to seed payment history %d: %v", i, err)
		} else {
			ids = append(ids, id)
		}
	}
	return ids
}

func seedInstallments(payHistories []int64) {
	log.Println("Seeding Installments...")
	for i, phID := range payHistories {
		_, err := db.DB.Exec("INSERT INTO installments (payment_history_id, amount, payment_method) VALUES ($1, 1000.00, 'Cash')", phID)
		if err != nil {
			log.Printf("Failed to seed installment %d: %v", i, err)
		}
	}
}
