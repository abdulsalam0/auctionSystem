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
	BidAmount string `json:"bidamount"`
	BidderID  string `json:"bidderid"`
}

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

// making a collection of auctions
var auctions []Auction

// create a user
func registerUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User page\n")
	fmt.Println("hello world")
}

// sign in user
func loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User page\n")
	fmt.Println("hello world")
}

// view all auctions
func viewAuctions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(auctions)
	w.Header().Set("Content-Type", "application/json")
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
	fmt.Fprintf(w, "User page\n")
	fmt.Println("hello world")
}

// create an auction
func createAuction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ID = 1
	var auction Auction
	_ = json.NewDecoder(r.Body).Decode(&auction)
	auction.AuctionID = strconv.Itoa(ID)
	auctions = append(auctions, auction)
	json.NewEncoder(w).Encode(auction)

	fmt.Println("hello world")
}

// remove an auction from the database
func deleteAuction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User page\n")
	fmt.Println("hello world")
}

// viewing all the bids on a bid
func getBids(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User page\n")
	fmt.Println("hello world")
}

// place a bid on an auction
func placeBid(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User page\n")
	fmt.Println("hello world")
}

func main() {
	r := mux.NewRouter()

	auctions = append(auctions, Auction{AuctionID: "1", AuctionName: "Iphone", FirstBid: 100, SellerID: "10", AuctionStatus: "Avalible"})
	auctions = append(auctions, Auction{AuctionID: "2", AuctionName: "Laptop", FirstBid: 500, SellerID: "11", AuctionStatus: "Avalible"})

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
	r.HandleFunc("/api/user", getBids).Methods("GET")
	r.HandleFunc("/api/user", placeBid).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", r))
}
