package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

type UserTask struct {
	Title string `json:"title"`
}

func SuggestTasks(c *fiber.Ctx) error {
	// Extract Authorization header
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing or invalid Authorization header")
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse the email from request body
	var body struct {
		Email string `json:"email"`
	}
	if err := c.BodyParser(&body); err != nil {
		log.Println("Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	// Fetch tasks via internal POST request
	tasks, err := getUserTasks(body.Email, token)
	if err != nil {
		log.Println("Failed to fetch user tasks:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Could not fetch user tasks")
	}
	if len(tasks) == 0 {
		return c.JSON(fiber.Map{
			"suggestions": []string{"AI could not generate any suggestions. Please try again later."},
		})
	}

	// Debug print the task
	fmt.Println("Fetched task for AI prompt:", tasks[0].Title)

	// Select any task to build prompt
	selectedTask := tasks[0].Title
	prompt := fmt.Sprintf(`Heyy gemini, what work comes after %s, can you list them in just words with just comma and nothing more`, selectedTask)

	// Prepare Gemini API call
	geminiURL := os.Getenv("GEMINI_CONN_STRING")
	geminiApiKey := os.Getenv("GEMINI_API_KEY")

	requestBody := fmt.Sprintf(`{
		"contents": [{
			"parts": [{"text": "%s"}]
		}]
	}`, prompt)

	resp, err := http.Post(geminiURL+geminiApiKey, "application/json", bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		log.Println("Error calling Gemini:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get suggestions")
	}
	defer resp.Body.Close()

	var geminiResponse GeminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&geminiResponse); err != nil {
		log.Println("Invalid Gemini response:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to parse Gemini response")
	}

	if len(geminiResponse.Candidates) == 0 || len(geminiResponse.Candidates[0].Content.Parts) == 0 {
		return c.JSON(fiber.Map{
			"suggestions": []string{"AI could not generate any suggestions. Please try again later."},
		})
	}

	text := geminiResponse.Candidates[0].Content.Parts[0].Text
	suggestions := strings.Split(text, ",")
	for i := range suggestions {
		suggestions[i] = strings.TrimSpace(suggestions[i])
	}

	return c.JSON(fiber.Map{
		"suggestions": suggestions,
	})
}

func getUserTasks(email, token string) ([]UserTask, error) {
	body := map[string]string{"email": email}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "http://localhost:8000/api/tasks", bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tasks []UserTask
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
