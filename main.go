package main

import (
"database/sql"
"encoding/json"
"fmt"
"log"
"net/http"
"os"

_ "github.com/lib/pq"
)

type Food struct {
ID          int    `json:"id"`
Name        string `json:"name"`
Price       float64 `json:"price"`
Description string `json:"description"`
}

var db *sql.DB

func init() {
// Database connection string
dbURL := os.Getenv("DATABASE_URL")
if dbURL == "" {
dbURL = "postgres://postgres:secret@localhost:5432/fooddb?sslmode=disable"
}

var err error
db, err = sql.Open("postgres", dbURL)
if err != nil {
log.Printf("Warning: Could not connect to database: %v", err)
} else {
// Test connection
if err := db.Ping(); err != nil {
log.Printf("Warning: Database not responding: %v", err)
} else {
log.Println("Database connected successfully")
}
}
}

func main() {
port := getEnv("PORT", "8080")
env := getEnv("ENVIRONMENT", "local")

log.Printf("=== Nigerian Food API ===")
log.Printf("Environment: %s", env)
log.Printf("Port: %s", port)
log.Printf("========================\n")

// Initialize database
if db != nil {
initDatabase()
}

http.HandleFunc("/", handleRoot)
http.HandleFunc("/health", handleHealth)
http.HandleFunc("/api/foods", handleFoods)

log.Printf("Starting server on port %s...", port)
if err := http.ListenAndServe(":"+port, nil); err != nil {
log.Fatal(err)
}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
response := map[string]string{
"service": "Nigerian Food API",
"version": "2.0",
"status":  "running",
}
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(response)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
health := map[string]string{
"status":   "healthy",
"service":  "nigerian-food-api",
"database": "connected",
}
if db != nil {
if err := db.Ping(); err != nil {
health["database"] = "disconnected"
}
}
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(health)
}

func handleFoods(w http.ResponseWriter, r *http.Request) {
// If no database, return hardcoded foods
if db == nil {
foods := []Food{
{ID: 1, Name: "Jollof Rice", Price: 1500, Description: "Delicious Nigerian jollof rice"},
{ID: 2, Name: "Pounded Yam", Price: 2000, Description: "Traditional pounded yam"},
{ID: 3, Name: "Suya", Price: 1000, Description: "Spicy grilled beef"},
}
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(foods)
return
}

// Query database
rows, err := db.Query("SELECT id, name, price, description FROM foods ORDER BY id")
if err != nil {
w.WriteHeader(http.StatusInternalServerError)
fmt.Fprintf(w, `{"error": "Database error"}`)
return
}
defer rows.Close()

var foods []Food
for rows.Next() {
var food Food
if err := rows.Scan(&food.ID, &food.Name, &food.Price, &food.Description); err != nil {
continue
}
foods = append(foods, food)
}

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(foods)
}

func initDatabase() {
createTableSQL := `
CREATE TABLE IF NOT EXISTS foods (
id SERIAL PRIMARY KEY,
name VARCHAR(100) NOT NULL,
price DECIMAL(10,2) NOT NULL,
description TEXT
);

INSERT INTO foods (name, price, description) VALUES
('Jollof Rice', 1500, 'Delicious Nigerian jollof rice'),
('Pounded Yam', 2000, 'Traditional pounded yam with egusi soup'),
('Suya', 1000, 'Spicy grilled beef'),
('Moi Moi', 500, 'Steamed bean pudding'),
('Akara', 300, 'Fried bean cakes')
ON CONFLICT DO NOTHING;
`

if _, err := db.Exec(createTableSQL); err != nil {
log.Printf("Warning: Could not initialize database: %v", err)
} else {
log.Println("Database initialized successfully")
}
}

func getEnv(key, defaultValue string) string {
if value, exists := os.LookupEnv(key); exists {
return value
}
return defaultValue
}
