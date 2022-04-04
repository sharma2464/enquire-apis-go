package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/types"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func GetClient() (*firestore.Client, context.Context, error) {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("./config/petheroes-aabc5-709546fec3c0.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	firebase_client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil, ctx, err
	}

	return firebase_client, ctx, nil
}

func AddUser(data string) (map[string]string, error) {
	client, ctx, _ := GetClient()
	var user types.EnquireForm
	json.Unmarshal([]byte(data), &user)
	doc, _, err := client.Collection("users").Add(ctx, user)

	if err != nil {
		// log.Fatalf("Failed adding alovelace: %v", err)
		log.Printf("Error adding doc to `users` collection\n%s:\nError: %s", user, err)
		client.Close()
		return nil, err
	}

	if err == nil {
		res := map[string]string{
			"message": "Created",
			"id":      doc.ID,
		}
		fmt.Println(res)
		client.Close()
		return res, nil
	}
	client.Close()
	return nil, nil
}

func GetAllUsers() ([]map[string]interface{}, error) {
	client, ctx, _ := GetClient()
	iter := client.Collection("users").Documents(ctx)
	var dt []map[string]interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Error Getting docs for collection `users`\nError: %s", err)
			return nil, err
		}
		dt = append(dt, doc.Data())

	}
	client.Close()
	return dt, nil
}

func GetOneUser(id string) (interface{}, error) {
	client, ctx, _ := GetClient()

	user, err := client.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		log.Printf("Error Getting doc for ID: %s:\nError: %s", id, err)
		return nil, err
	}

	if user != nil {
		return user.Data(), nil
	}
	client.Close()
	return nil, nil
}

func DeleteOneUser(id string) (map[string]string, error) {
	client, ctx, _ := GetClient()
	fmt.Println("All users:")

	deleted, err := client.Collection("users").Doc(id).Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("Error Deleting doc for ID: %s:\nError: %s", id, err)
		return nil, err
	}

	if deleted != nil {
		return map[string]string{"message": "deleted"}, nil
	}
	client.Close()
	return nil, nil
}
