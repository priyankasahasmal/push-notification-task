package main

import (
	"context"
	"testing"
)

// TestSendPushNotification is a basic unit test
func TestSendPushNotification(t *testing.T) {
	// Replace with a **valid test token** from your device
	testToken := "your-test-device-token"

	ctx := context.Background()
	resp, err := SendPushNotification(ctx, testToken, "Test Title", "Test Body")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if resp == "" {
		t.Errorf("Expected a response ID, but got empty string")
	}
}
