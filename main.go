package main

import (
"encoding/json"
"log"
"net/http"
"os"
)

// Food represents a Nigerian food item
type Food struct {
ID          int    `json:"id"`
Name        string `json:"name"`
Price       float64 `json:"price"`
Description string `json:"description"`
}

func main() {
// Get configuration from environment
port := getEnv("PORT", "8080")
dbURL := getEnv("DATABASE_URL", "localhost:5432/fooddb")
env := getEnv("ENVIRONMENT", "local")

// Log startup info
log.Printf("=== Nigerian Food API ===")
log.Printf("Environment: %s", env)
log.Printf("Port: %s", port)
log.Printf("Database: %s", dbURL)
log.Printf("=======================\n")

// Register HTTP routes
http.HandleFunc("/", handleRoot)
http.HandleFunc("/health", handleHealth)
http.HandleFunc("/api/foods", handleFoods)

// Start server
log.Printf("Starting server on port %s...", port)
if err := http.ListenAndServe(":"+port, nil); err != nil {
log.Fatal(err)
}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
response := map[string]string{
"service": "Nigerian Food API",
"version": "1.0",
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
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(health)
}

func handleFoods(w http.ResponseWriter, r *http.Request) {
foods := []Food{
{ID: 1, Name: "Jollof Rice", Price: 1500, Description: "Delicious Nigerian jollof rice"},
{ID: 2, Name: "Pounded Yam", Price: 2000, Description: "Traditional pounded yam with egusi soup"},
{ID: 3, Name: "Suya", Price: 1000, Description: "Spicy grilled beef"},
{ID: 4, Name: "Moi Moi", Price: 500, Description: "Steamed bean pudding"},
{ID: 5, Name: "Akara", Price: 300, Description: "Fried bean cakes"},
}

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(foods)
}

// Helper function to get environment variables with defaults
func getEnv(key, defaultValue string) string {
if value, exists := os.LookupEnv(key); exists {
return value
}
return defaultValue
}
