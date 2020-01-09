package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

// Auction class
type Auction struct {
	RequestNum    string `json:"requestnum"`
	AuctionID     string `json:"auctionid"`
	AuctionName   string `json:"auctionName"`
	FirstBid      int    `json:"firstbid"`
	SellerID      string `json:"sellerid"`
	AuctionStatus string `json:"auctionstatus"`
}

// Bid class
type Bid struct {
	RequestNum string `json:"requestnum"`
	BidID      string `json:"bidid"`
	AuctionID  string `json:"auctionid"`
	BidAmount  int    `json:"bidamount"`
	BidderID   string `json:"bidderid"`
}

// User class
type User struct {
	RequestNum string `json:"requestnum"`
	UserID     string `json:"userid"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

// IDRequest counter
var IDRequest = 1

func createUserRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.RequestNum = strconv.Itoa(IDRequest)
	requestBody, err := json.Marshal(map[string]string{
		"requestnum": user.RequestNum,
		"userid":     user.UserID,
		"firstname":  user.FirstName,
		"lastname":   user.LastName,
		"username":   user.Username,
		"password":   user.Password,
	})
	IDRequest++

	if err != nil {
		fmt.Println(err)
	}

	result := postServerRequest(requestBody, r.URL.Path)

	json.NewEncoder(w).Encode(result)
}
func loginUserReqest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var auction Auction
	_ = json.NewDecoder(r.Body).Decode(&auction)
	auction.RequestNum = strconv.Itoa(IDRequest)
	requestBody, err := json.Marshal(map[string]string{
		"requestnum":   auction.RequestNum,
		"userid":       auction.AuctionID,
		"firstname":    auction.AuctionName,
		"auctionsatus": auction.AuctionStatus,
		"firstbid":     strconv.Itoa(auction.FirstBid),
		"sellerid":     auction.SellerID,
	})
	IDRequest++

	if err != nil {
		fmt.Println(err)
	}

	result := postServerRequest(requestBody, r.URL.Path)

	json.NewEncoder(w).Encode(result)
}
func createAuctionRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.RequestNum = strconv.Itoa(IDRequest)
	requestBody, err := json.Marshal(map[string]string{
		"requestnum": user.RequestNum,
		"userid":     user.UserID,
		"firstname":  user.FirstName,
		"lastname":   user.LastName,
		"username":   user.Username,
		"password":   user.Password,
	})
	IDRequest++

	if err != nil {
		fmt.Println(err)
	}

	result := postServerRequest(requestBody, r.URL.Path)

	json.NewEncoder(w).Encode(result)
}
func updateAuctionRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.RequestNum = strconv.Itoa(IDRequest)
	requestBody, err := json.Marshal(map[string]string{
		"requestnum": user.RequestNum,
		"userid":     user.UserID,
		"firstname":  user.FirstName,
		"lastname":   user.LastName,
		"username":   user.Username,
		"password":   user.Password,
	})
	IDRequest++

	if err != nil {
		fmt.Println(err)
	}

	result := serverRequest(requestBody, r.URL.Path)

	json.NewEncoder(w).Encode(result)
}
func deleteAuctionRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.RequestNum = strconv.Itoa(IDRequest)
	requestBody, err := json.Marshal(map[string]string{
		"requestnum": user.RequestNum,
		"userid":     user.UserID,
		"firstname":  user.FirstName,
		"lastname":   user.LastName,
		"username":   user.Username,
		"password":   user.Password,
	})
	IDRequest++

	if err != nil {
		fmt.Println(err)
	}

	result := serverRequest(requestBody, r.URL.Path)

	json.NewEncoder(w).Encode(result)
}
func creaetBidRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bid Bid
	_ = json.NewDecoder(r.Body).Decode(&bid)
	bid.RequestNum = strconv.Itoa(IDRequest)
	requestBody, err := json.Marshal(map[string]string{
		"requestnum": bid.RequestNum,
		"bidid":      bid.BidID,
		"bidderid":   bid.BidderID,
		"auctionid":  bid.AuctionID,
		"bidamount":  strconv.Itoa(bid.BidAmount),
	})
	IDRequest++

	if err != nil {
		fmt.Println(err)
	}

	result := postServerRequest(requestBody, r.URL.Path)

	json.NewEncoder(w).Encode(result)
}

func postServerRequest(requestBody []byte, endpoint string, params ...string) map[string]interface{} {
	resp, err := http.Post("http://server1:8080/"+endpoint, "application/json", bytes.NewBuffer(requestBody))
	resp.Header.Set("Content-type", "application/json")

	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	// check the content of the body
	fmt.Println(result)
	return result

}

// create a client
func createClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}

func get(client *redis.Client) error {
	value, err := client.Get("server1").Result()
	if err != nil {
		return err
	}
	fmt.Println("server 1 port", value)

	portvalue, err := client.Get("server2").Result()
	if err != nil {
		return err
	}
	fmt.Println("server 2 port", portvalue)
	return nil
}

func viewList(client *redis.Client) error {
	value, err := client.LRange("server_list", 0, -1).Result()
	if err != nil {
		return err
	}
	fmt.Println(value)
	return nil
}

func main() {
	fmt.Println("we are up and running")

	client := createClient()

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)

	// for {
	// 	err := get(client)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	err1 := viewList(client)
	// 	if err != nil {
	// 		fmt.Println(err1)
	// 	}
	// 	time.Sleep(time.Millisecond * 5000)
	// }

	r := mux.NewRouter()

	r.HandleFunc("/api/user", createUserRequest).Methods("POST")
	r.HandleFunc("/api/login", loginUserReqest).Methods("POST")

	r.HandleFunc("/api/auction", createAuctionRequest).Methods("POST")
	r.HandleFunc("/api/auction/{id}", updateAuctionRequest).Methods("PUT")
	r.HandleFunc("/api/auction/{id}", deleteAuctionRequest).Methods("DELETE")

	//r.HandleFunc("/api/auction/{id}/bid", forwardRequest).Methods("POST")

	log.Fatal(http.ListenAndServe(":9090", r))

}
