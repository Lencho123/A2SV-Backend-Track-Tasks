package data

import (
	"context"
	"time"
)

// Helper function to creat context
func createContext() (context.Context, context.CancelFunc) {
	return  context.WithTimeout(context.Background(), 10*time.Second)
}