package services

import (
	"encoding/json"
	"fmt"
	"log"
	"main/config"
	"main/types"

	"google.golang.org/api/iterator"
)

func AddUser(data string) (map[string]string, error) {
	client, ctx, _ := config.GetClient()
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
	client, ctx, _ := config.GetClient()
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
	client, ctx, _ := config.GetClient()

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
	client, ctx, _ := config.GetClient()
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
