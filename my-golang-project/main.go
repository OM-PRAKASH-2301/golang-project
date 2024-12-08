package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	// MongoDB connection setup
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Replace with your MongoDB URI
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Ping the MongoDB server
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Could not ping MongoDB server:", err)
	}

	// Get a handle for the "formData" collection in the "testDB" database
	collection = client.Database("testDB").Collection("formData")
	fmt.Println("Connected to MongoDB!")
}

func main() {
	// Serve static files (CSS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Serve the HTML form
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Go HTML Form</title>
			<link rel="stylesheet" href="/static/style.css">
		</head>
		<body>
			<div class="form-container">
				<h1>Sample Form</h1>
				<form action="/submit" method="POST">
					<div class="form-group">
						<label for="name">Name:</label>
						<input type="text" id="name" name="name" required>
					</div>
					<div class="form-group">
						<label for="email">Email:</label>
						<input type="email" id="email" name="email" required>
					</div>
					<div class="form-group">
						<label for="message">Message:</label>
						<textarea id="message" name="message" rows="5" required></textarea>
					</div>
					<button type="submit">Submit</button>
				</form>
			</div>
		</body>
		</html>`
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	})

	// Handle form submission
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Parse form data
			name := r.FormValue("name")
			email := r.FormValue("email")
			message := r.FormValue("message")

			// Create a document to insert into MongoDB
			doc := bson.M{
				"name":    name,
				"email":   email,
				"message": message,
				"created": time.Now(),
			}

			// Insert the document into the collection
			_, err := collection.InsertOne(context.TODO(), doc)
			if err != nil {
				http.Error(w, "Error storing data in database", http.StatusInternalServerError)
				log.Println("Database insertion error:", err)
				return
			}

			// Response to the client
			response := "Form submitted successfully!<br>" +
				"<strong>Name:</strong> " + name + "<br>" +
				"<strong>Email:</strong> " + email + "<br>" +
				"<strong>Message:</strong> " + message
			w.Write([]byte(response))
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	// Start the server
	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
