package plugin

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Datasource is a mongo datasource which can respond to data queries, reports
// its health and has streaming skills.
type Datasource struct {
	host     string
	port     int
	database string
	client   *mongo.Client
}

type queryModel struct {
	QueryText     string `json:"queryText"`
	Collection    string `json:"collection"`
	QueryType     string `json:"queryType"`
	QueryLanguage string `json:"queryLanguage"`
}

type TimeSeriesRow[T any] struct {
	Timestamp time.Time `bson:"ts"`
	Name      string    `bson:"name"`
	Value     T         `bson:"value"`
}

type Optional[T any] struct {
	Value   T
	Nothing bool
}

func (opt Optional[T]) ToPointer() *T {
	if opt.Nothing {
		return nil
	}
	return &opt.Value
}
