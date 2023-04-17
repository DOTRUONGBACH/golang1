package main

import (
	"context"
	"entdemo/ent"
	"entdemo/ent/user"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:my-secret-pw@tcp(127.0.0.1:3307)/bach?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// CreateUser(context.Background(), client)
	QueryUser(context.Background(), client)
	DeleteById(context.Background(), client)
}
func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(20).
		SetName("bachpro").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}
func QueryUser(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
	u, err := client.User.Query().Where(user.Age(20)).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned:  ", u)
	return u, nil
}
func DeleteById(ctx context.Context, client *ent.Client) (error) {
	err := client.User.DeleteOneID(2).Exec(ctx)
	
	log.Println("deleted")
	return err
}
