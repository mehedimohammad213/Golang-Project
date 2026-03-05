package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL          string
	Port           string
	JWTSecret      string
	JWTExpiryHours int
	// RAG / OpenAI
	OpenAIAPIKey      string
	RAGEmbeddingModel string
	RAGChatModel      string
	RAGTopK           int
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSL := os.Getenv("DB_SSLMODE")

	dbURL := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " sslmode=" + dbSSL

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "default-secret"
	}

	jwtExpiryHours := 24
	if expiry := os.Getenv("JWT_EXPIRY_HOURS"); expiry != "" {
		if val, err := strconv.Atoi(expiry); err == nil {
			jwtExpiryHours = val
		}
	}

	openAIKey := os.Getenv("OPENAI_API_KEY")
	ragEmbedModel := os.Getenv("RAG_EMBEDDING_MODEL")
	if ragEmbedModel == "" {
		ragEmbedModel = "text-embedding-3-small"
	}
	ragChatModel := os.Getenv("RAG_CHAT_MODEL")
	if ragChatModel == "" {
		ragChatModel = "gpt-4o-mini"
	}
	ragTopK := 5
	if k := os.Getenv("RAG_TOP_K"); k != "" {
		if val, err := strconv.Atoi(k); err == nil && val > 0 {
			ragTopK = val
		}
	}

	return &Config{
		DBURL:             dbURL,
		Port:              port,
		JWTSecret:         jwtSecret,
		JWTExpiryHours:    jwtExpiryHours,
		OpenAIAPIKey:      openAIKey,
		RAGEmbeddingModel: ragEmbedModel,
		RAGChatModel:      ragChatModel,
		RAGTopK:           ragTopK,
	}
}
