package main

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"github.com/shiowwj/go-cointracker-crud/pkg/utils/log"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

func main() {
	fmt.Println()
	opt := option.WithCredentialsFile("../../test-cointracker-firebase-adminsdk.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal("Failed to initialise app", zap.Error(err))
	}

	a, _ := app.Auth(context.Background())
	log.Debug("HELLO firebase service", zap.Any("a", a))
}

// log.Fatal(http.ListenAndServe("localhost:9010", r))

// import (
//   "fmt"
//   "context"

//   firebase "firebase.google.com/go"
//   "firebase.google.com/go/auth"

//   "google.golang.org/api/option"
// )

// opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
// app, err := firebase.NewApp(context.Background(), nil, opt)
// if err != nil {
//   return nil, fmt.Errorf("error initializing app: %v", err)
// }
