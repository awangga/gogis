/**
 * Author: Rolly Maulana Awangga
 * File: gogis.go
 */

package gogis

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IpApi struct {
	Status      string `json:"status,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	Region      string `json:"region,omitempty"`
	RegionName  string `json:"regionName,omitempty"`
	City        string `json:"city,omitempty"`
	Lat         string `json:"lat,omitempty"`
	Lon         string `json:"lon,omitempty"`
	Timezone    string `json:"timezone,omitempty"`
	Isp         string `json:"isp,omitempty"`
	Org         string `json:"org,omitempty"`
	As          string `json:"as,omitempty"`
	Query       string `json:"query,omitempty"`
}

type Geometry struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}

type Desa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Province     string             `bson:"province,omitempty"`
	District     string             `bson:"district,omitempty"`
	Sub_district string             `bson:"sub_district,omitempty"`
	Village      string             `bson:"village,omitempty"`
	Border       Geometry           `bson:"border,omitempty"`
}

type MongoGeometry struct {
	MongoString    string
	DBName         string
	CollectionName string
	LocationField  string
}

// Returns the sum of two numbers
func Add(a int, b int) int {
	return a + b
}

// Returns the difference between two numbers
func Subtract(a int, b int) int {
	return a - b
}

func GetPublicIP() IpApi {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		println(err.Error())
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		println(err.Error())
	}

	var ip IpApi
	json.Unmarshal(body, &ip)
	return ip
}

func GetLocation(mongog MongoGeometry, long float64, lat float64) (loc interface{}) {
	loccollection := MongoConnect(mongog.MongoString, mongog.DBName).Collection(mongog.CollectionName)
	filter := bson.M{
		mongog.LocationField: bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{long, lat},
				},
			},
		},
	}
	err := loccollection.FindOne(context.TODO(), filter).Decode(&loc)
	if err != nil {
		fmt.Printf("GetLocation: %v\n", err)
	}
	return loc

}

func MongoConnect(mongostring string, dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongostring))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}
