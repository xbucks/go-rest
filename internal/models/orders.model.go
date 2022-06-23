package models

import (
	"time"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceInfo struct {
	Name        string
	UpTime      time.Time
	Environment string
	Version     string
}

func (s ServiceInfo) MarshalZerologObject(e *zerolog.Event) {
	e.Str("name", s.Name).
		Str("environment", s.Environment).
		Time("started", s.UpTime).
		Str("version", s.Version)
}

type Order struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"order_id"`
	LastUpdatedAt string             `bson:"last_updated_at,omitempty"`
	Products      []Product          `bson:"products,omitempty"`
}

type Product struct {
	Name      string `bson:"name,omitempty"`
	UpdatedAt string `bson:"updated_at,omitempty"`
	Price     uint   `bson:"price,omitempty"`
	Status    string `bson:"status,omitempty"`
	Remarks   string `bson:"remarks,omitempty"`
}
