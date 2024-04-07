package websockets

import (
	"bids/database"
	"bids/models"
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Auction struct {
	id                  string
	currentHighestOffer models.Offer
	Clients             map[*Client]bool
	Server              *Server
	Offer               chan models.Offer
	Broadcast           chan []byte
	End                 int64
	Stop                chan bool
	AddUser             chan *Client
	RemoveUser          chan *Client
}

func CreateAuction(name string, server *Server) (*Auction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	auctionCollection := database.GetCollection(database.DB, "auctions")
	id, _ := primitive.ObjectIDFromHex(name)
	filter := bson.D{{"_id", id}}
	opts := options.FindOne().SetProjection(bson.D{{"end", 1}, {}}).SetSort(bson.D{{"offers", 1}})
	auctionCollection.FindOne(ctx, filter)
	return &Auction{
		id:                  name,
		currentHighestOffer: offer,
		Clients:             make(map[*Client]bool),
		Server:              server,
		Offer:               make(chan models.Offer),
		End:                 end,
		Stop:                make(chan bool),
		AddUser:             make(chan *Client),
		RemoveUser:          make(chan *Client),
	}, nil
}
func (r *Auction) AddClient(client *Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	auctionCollection := database.GetCollection(database.DB, "auctions")
	id, _ := primitive.ObjectIDFromHex(r.id)
	filter := bson.D{{"_id", id}, {"users", client.UserID}}
	err := auctionCollection.FindOne(ctx, filter).Err()
	if err != nil {
		client.WriteMess <- []byte("unauthorized")
		return
	}
	client.WriteMess <- []byte(r.id)
	r.Clients[client] = true
}
func (r *Auction) RemoveClient(client *Client) {
	delete(r.Clients, client)
}
func (r *Auction) GetClient(client string) *Client {
	for i := range r.Clients {
		if i.UserID == client {
			return i
		}
	}
	return nil
}
func (r *Auction) endAuction() {
	message := []byte("end")
	for client := range r.Clients {
		client.WriteMess <- message
	}
}
func (r *Auction) sendOffer(offer models.Offer) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	auctionCollection := database.GetCollection(database.DB, "auction")
	if offer.Price > r.currentHighestOffer.Price {
		r.currentHighestOffer = offer
		id, _ := primitive.ObjectIDFromHex(r.id)
		filter := bson.D{{"_id", id}}
		update := bson.M{"$push": bson.M{"offers": offer}}
		_, err := auctionCollection.UpdateOne(ctx, filter, update)
		if err == nil {
			data, _ := json.Marshal(offer)
			for client := range r.Clients {
				client.WriteMess <- data
			}
		} else {
			r.GetClient(offer.Sender).WriteMess <- []byte("error")
		}

	}

}
func (r *Auction) RunAuction() {

	for {
		if time.Now().Unix() == r.End {
			r.endAuction()
			return
		}
		select {
		case offer := <-r.Offer:
			r.sendOffer(offer)
		case <-r.Stop:
			return
		default:
		}
		time.Sleep(time.Millisecond)
	}
}
