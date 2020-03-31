package main

import (
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type timer struct {
	User string
	Hours int
	Minutes int
	Secounds int
	Nanosecounds float64
}

type result struct {
	StartUser string
	StartHours int
	StartMinutes int
	StartSecounds int
	StartNanosecounds float64
	DiffHours int
	DiffMinutes int
	DiffSecounds int
	DiffNanosec float64
	EndHours int
	EndMinutes int
	EndSecounds int
	EndNanosecounds float64
}

func add() {

	clientoptions := options.Client().ApplyURI("mongodb://localhost:5000")

	client, err := mongo.Connect(context.TODO(), clientoptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("time").Collection("timers")
	
	username := user()

	hrTime, minTime, secTime, nanoTime := getTime()

	userDetail := timer{username, hrTime, minTime, secTime, nanoTime}

	var result_add timer

	find := bson.D{{"user", username}}

	error := collection.FindOne(context.TODO(), find).Decode(&result_add)

	if error == nil {
		log.Fatal("Timer for the User already running.\nPlease stop that session first.")
	}

	_, err_insert := collection.InsertOne(context.TODO(), userDetail)

	if err_insert != nil {
		log.Fatal(err_insert)
	}

	//fmt.Println("ID is ",result.InsertedID)
	
	close(client)
	
	fmt.Println("User timer start at ", hrTime,":",minTime,":",secTime,".",nanoTime)
}

func close(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("Connection to MongoDB closed")
	
}

func diff() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:5000")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("time").Collection("timers")

	user := user()

	find := bson.D{{"user", user}}

	var results timer
	error := collection.FindOne(context.TODO(), find).Decode(&results)

	if error != nil {
		log.Fatal("Say Hi first to start Timer")
	}

	//fmt.Println("TIMERS: ", results)

	startHours := results.Hours
	startMin := results.Minutes
	startSec := results.Secounds
	startNano := results.Nanosecounds

	endHours, endMinutes, endSecounds, endNano := getTime()

	diffHours := endHours - startHours
	diffMinutes := endMinutes - startMin
	diffSecounds := endSecounds - startSec 
	diffNano :=  endNano - startNano 

	if diffMinutes < 0{
		diffMinutes = 60 + diffMinutes
	}

	if diffSecounds < 0 {
		diffSecounds = 60 + diffSecounds
	}

	if diffNano < 0 {
		diffNano = 1.0e+09 + diffNano
	}

	_, err_delete := collection.DeleteOne(context.TODO(), find)

	if err_delete != nil {
		log.Fatal(err_delete)
	}

	
	resultDatabase := client.Database("time").Collection("result")


	sheet := result{user, startHours, startMin, startSec, startNano,
				diffHours, diffMinutes, diffSecounds, diffNano, endHours,
				endMinutes, endSecounds, endNano}

	_, err_insert := resultDatabase.InsertOne(context.TODO(), sheet)

	if err_insert != nil {
		log.Fatal(err_insert)
	}

	//fmt.Println("ID is ",added.InsertedID)

	close(client)

	fmt.Println("HR: ",diffHours, " Min: ",diffMinutes," Sec: ",diffSecounds," Nano: ",diffNano)

}
