package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Auction struct {
	AuctionID     string `json:"auctionid"`
	AuctionName   string `json:"auctionName"`
	FirstBid      int    `json:"firstbid"`
	SellerID      string `json:"sellerid"`
	AuctionStatus string `json:"auctionstatus"`
}

type Bid struct {
	BidID     string `json:"bidid"`
	AuctionID string `json:"auctionid"`
	BidAmount int    `json:"bidamount"`
	BidderID  string `json:"bidderid"`
}

type User struct {
	UserID    string `json:"userid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

// making all the difderent collection
var auctions []Auction
var users []User
var bids []Bid

var IDUser = 1
var IDAuction = 1
var IDBid = 1

// create a user
func registerUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.UserID = strconv.Itoa(IDUser)
	users = append(users, user)
	fmt.Println(user)
	json.NewEncoder(w).Encode(user)
	IDUser++

	// TODO respond with the jwt token

	fmt.Println("created user")
}

// sign in user
func loginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jwt = "this is the jwt token"
	var errorMsg = "User error"
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	for _, item := range users {
		if item.Username == user.Username && item.Password == user.Password {
			// TODO give back the JWT
			json.NewEncoder(w).Encode(jwt)
			return
		}
	}
	json.NewEncoder(w).Encode(errorMsg)

}

// view all auctions
func viewAuctions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auctions)
	fmt.Println(auctions)
	fmt.Println("Sending List of Auctions")
}

// view auction by ID
func viewAuctionByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range auctions {
		if item.AuctionID == params["id"] {
			json.NewEncoder(w).Encode(item)
			fmt.Println("sent one auction")
			return
		}
	}
	json.NewEncoder(w).Encode(&Auction{})
}

// update an auction
func updateAuctions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range auctions {
		if item.AuctionID == params["id"] {
			//slice to remove the element
			auctions = append(auctions[:index], auctions[index+1:]...)
			//adding a new element with new info
			var auction Auction
			_ = json.NewDecoder(r.Body).Decode(&auction)
			auction.AuctionID = params["id"]
			auctions = append(auctions, auction)
			json.NewEncoder(w).Encode(auction)
			fmt.Println("Update Auction")
			return
		}
	}
}

// create an auction
func createAuction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var auction Auction
	_ = json.NewDecoder(r.Body).Decode(&auction)
	auction.AuctionID = strconv.Itoa(IDAuction)
	auctions = append(auctions, auction)
	json.NewEncoder(w).Encode(auction)
	IDAuction++
	fmt.Println("added New auction")
}

// remove an auction from the database
func deleteAuction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range auctions {
		if item.AuctionID == params["id"] {
			auctions = append(auctions[:index], auctions[index+1:]...)
			break
		}
	}
	fmt.Println(auctions)
	json.NewEncoder(w).Encode(auctions)
	fmt.Println("deleted Auction")
}

// viewing the bids on an auction
func getBids(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var bidList []Bid
	for _, item := range bids {
		if item.AuctionID == params["id"] {
			bidList = append(bidList, item)
		}
	}
	json.NewEncoder(w).Encode(bidList)
	fmt.Println("hello world")
}

// place a bid on an auction
func placeBid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var bid Bid
	_ = json.NewDecoder(r.Body).Decode(&bid)
	bid.BidID = strconv.Itoa(IDBid)
	bid.AuctionID = params["id"]

	// add the bid to the list
	bids = append(bids, bid)

	json.NewEncoder(w).Encode(bid)
	IDBid++
}

func main() {
	r := mux.NewRouter()

	// creating fake data
	auctions = append(auctions, Auction{AuctionID: "9", AuctionName: "Iphone", FirstBid: 100, SellerID: "10", AuctionStatus: "Avalible"})
	auctions = append(auctions, Auction{AuctionID: "10", AuctionName: "Laptop", FirstBid: 500, SellerID: "11", AuctionStatus: "Avalible"})

	// fake users
	users = append(users, User{UserID: "10", FirstName: "Laptop", LastName: "500", Username: "11", Password: "Avalible"})
	users = append(users, User{UserID: "11", FirstName: "abdul", LastName: "aboubakar", Username: "abdul123", Password: "wassup"})

	// user endpoints
	r.HandleFunc("/api/user", registerUser).Methods("POST")
	r.HandleFunc("/api/login", loginUser).Methods("POST")

	// auctions endpoints
	r.HandleFunc("/api/auctions", viewAuctions).Methods("GET")
	r.HandleFunc("/api/auction/{id}", viewAuctionByID).Methods("GET")
	r.HandleFunc("/api/auction/{id}", updateAuctions).Methods("PUT")
	r.HandleFunc("/api/auction", createAuction).Methods("POST")
	r.HandleFunc("/api/auction/{id}", deleteAuction).Methods("DELETE")

	// bid endpoints
	r.HandleFunc("/api/auction/{id}/bids", getBids).Methods("GET")
	r.HandleFunc("/api/auction/{id}/bid", placeBid).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", r))
}