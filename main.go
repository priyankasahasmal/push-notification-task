package main

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

// Function to send push notification
func SendPushNotification(ctx context.Context, token string, title, body string) (string, error) {
	// Initialize Firebase app
	opt := option.WithCredentialsFile("ServiceAccountKey2.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return "", fmt.Errorf("error initializing Firebase app: %v", err)
	}

	// Get Messaging client
	client, err := app.Messaging(ctx)
	if err != nil {
		return "", fmt.Errorf("error getting Messaging client: %v", err)
	}

	// it will create message
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: token, // device token
	}

	// this is sending message
	response, err := client.Send(ctx, message)
	if err != nil {
		return "", fmt.Errorf("error sending message: %v", err)
	}
	return response, nil
}

// Entry point
func main() {
	ctx := context.Background()

	//  isse actual device token se replcae krna ha from firebase client app
	deviceToken := "e2robMiHje3UgH-YWAwea6:APA91bEmXKcoQAFrheDwRSRpmxCOnvJ54R_MAu32QQYtK86MMsp4JWBCAjnRr41gjttCpiobqUYAYugYu8fGq5-YGubNz5Rzieqyho51vgsxx-RxfU9W-xo"

	response, err := SendPushNotification(ctx, deviceToken, "Hello 👋 Iam priyanka", "This is my first task!")
	if err != nil {
		fmt.Println("❌ Failed to send:", err)
		return
	}

	fmt.Println("✅ Successfully sent message:", response)
}
