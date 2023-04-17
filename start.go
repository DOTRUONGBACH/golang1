package main

import (
	"context"
	"entdemo/ent"
	"entdemo/ent/car"
	"entdemo/ent/user"
	"time"

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
	// QueryUser(context.Background(), client)
	// DeleteById(context.Background(), client)

	// CreateCars(context.Background(), client)
	var user ent.User
	QueryCars(context.Background(), &user)
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
func DeleteById(ctx context.Context, client *ent.Client) error {
	err := client.User.DeleteOneID(2).Exec(ctx)

	log.Println("deleted")
	return err
}

func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {

	// Create a new car with model "Tesla".
	tesla, err := client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", tesla)

	// Create a new car with model "Ford".
	ford, err := client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", ford)

	// Create a new user, and add it the 2 cars.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").AddCars(tesla, ford).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", a8m)
	return a8m, nil
}

func QueryCars(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %w", err)
	}
	log.Println("returned cars: ", cars)

	ford, err := a8m.QueryCars().Where(car.Model("Ford")).Only(ctx)
	if err != nil {
		return fmt.Errorf("erroooor %w ", err)
	}
	log.Println(ford)
	return nil
}
