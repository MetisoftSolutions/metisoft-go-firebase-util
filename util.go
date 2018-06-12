package metisoft_go_firebase_util

import (
	"fmt"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"

	"google.golang.org/api/option"
)

type ConfigurationOptions struct {
	PathToServiceAccountKey string
	DatabaseUrl             string

	FnGetFirebaseTokenForUser func(string) (string, error)
}

var configOptions ConfigurationOptions
var app *firebase.App

func defaultFnGetFirebaseTokenForUser(userId string) (token string, _ error) {
	return userId, nil
}

func SendPushNotification(userId string, title string, body string) error {
	token, err := configOptions.FnGetFirebaseTokenForUser(userId)
	if err != nil {
		return err
	}

	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		return err
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: token,
	}

	_, err = client.Send(ctx, message)
	if err != nil {
		return err
	}

	return nil
}

func Init(options ConfigurationOptions) error {
	var err error

	if options.PathToServiceAccountKey == "" {
		return fmt.Errorf("PathToServiceAccountKey in options is required")
	}

	if options.DatabaseUrl == "" {
		return fmt.Errorf("DatabaseUrl in options is required")
	}

	if options.FnGetFirebaseTokenForUser == nil {
		options.FnGetFirebaseTokenForUser = defaultFnGetFirebaseTokenForUser
	}

	configOptions = options

	opt := option.WithCredentialsFile(options.PathToServiceAccountKey)
	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	return nil
}
