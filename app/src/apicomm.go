package main

import (
   // import standard libraries
    "context"
    "fmt"
    "log"
    "os"

    // Import the GenerativeAI package for Go
    "github.com/joho/godotenv"
    "github.com/google/generative-ai-go/genai"
    "google.golang.org/api/option"
)

func main() {
    // Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
    
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()


    // Use client.UploadFile to upload a file to the service.
    // Pass it an io.Reader.
    f, err := os.Open("jetpack.jpg")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    // Optionally set a display name.
    opts := genai.UploadFileOptions{DisplayName: "Jetpack drawing"}
    // Let the API generate a unique `name` for the file by passing an empty string.
    // If you specify a `name`, then it has to be globally unique.
    img1, err := client.UploadFile(ctx, "", f, &opts)
    if err != nil {
        log.Fatal(err)
    }

    // View the response.
    fmt.Printf("Uploaded file %s as: %q\n", img1.DisplayName, img1.URI)
} 
