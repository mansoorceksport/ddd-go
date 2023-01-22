package mongo

import (
	"context"
	"github.com/google/uuid"
	"github.com/mansoorceksport/ddd-go/aggregate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//MongoRepository to manage customer with mongodb
type MongoRepository struct {
	db       *mongo.Database
	customer *mongo.Collection
}

// mongoCustomer is internal type that is used to store a CustomerAggregate
// inside this repository.
type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

func NewFromCustomer(c aggregate.Customer) mongoCustomer {
	return mongoCustomer{
		c.GetID(),
		c.GetName(),
	}
}

func (m mongoCustomer) ToAggregate() aggregate.Customer {
	c := aggregate.Customer{}
	c.SetName(m.Name)
	c.SetID(m.ID)
	return c
}

func New(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}
	db := client.Database("ddd")
	customer := db.Collection("customers")
	return &MongoRepository{
		db:       db,
		customer: customer,
	}, nil
}

func (m *MongoRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := m.customer.FindOne(ctx, bson.M{"id": id})
	var c mongoCustomer
	err := result.Decode(&c)
	if err != nil {
		return aggregate.Customer{}, err
	}

	return c.ToAggregate(), nil
}

func (m *MongoRepository) Add(ac aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	internal := NewFromCustomer(ac)

	_, err := m.customer.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoRepository) Update(aggregate.Customer) error {
	//TODO implement me
	panic("implement me")
}
