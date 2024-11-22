package database

import (
	"context"
	"github.com/jay-SP/gql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectingString string = "mongodb+srv://akhil:akhil@cluster0.td5oga4.mongodb.net/test"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectingString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return &DB{client: client}
}

func (db *DB) GetJob(id string) *model.JobListing {
	jobCollection := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var jobListing model.JobListing
	err := jobCollection.FindOne(ctx, filter).Decode(&jobListing)
	if err != nil {
		log.Fatal(err)
	}
	return &jobListing
}

func (db *DB) GetJobs() []*model.JobListing {
	jobCollection := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var jobListings []*model.JobListing
	cursor, err := jobCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &jobListings); err != nil {
		panic(err)
	}
	return jobListings
}

func (db *DB) CreateJobListing(jobInfo model.CreateJobListingInput) *model.JobListing {
	jobCollection := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	insert, err := jobCollection.InsertOne(ctx, bson.M{"title": jobInfo.Title, "description": jobInfo.Description,
		"url": jobInfo.URL, "company": jobInfo.Company})
	if err != nil {
		log.Fatal(err)
	}
	insertedID := insert.InsertedID.(primitive.ObjectID).Hex()

	returnJobListing := model.JobListing{ID: insertedID, Title: jobInfo.Title, Description: jobInfo.Description}
	return &returnJobListing
}

func (db *DB) UpdateJobListing(jobId string, jobInfo model.UpdateJobListingInput) *model.JobListing {
	jobCollection := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateJobinfo := bson.M{}

	if jobInfo.Title != nil {
		updateJobinfo["title"] = jobInfo.Title
	}
	if jobInfo.Description != nil {
		updateJobinfo["description"] = jobInfo.Description
	}
	if jobInfo.URL != nil {
		updateJobinfo["url"] = jobInfo.URL
	}

	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateJobinfo}

	results := jobCollection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var jobListing model.JobListing

	if err := results.Decode(&jobListing); err != nil {
		log.Fatal(err)
	}
	return &jobListing
}

func (db *DB) DeleteJobListing(jobId string) *model.DeleteJobResponse {
	jobCollection := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	_, err := jobCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteJobResponse{DeletedJobID: jobId}
}
