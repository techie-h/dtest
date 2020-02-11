package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/gddo/httputil/header"
	"github.com/gorilla/mux"
	pb "github.com/techie-h/dtest/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	Message string `json:"message"`
	Rand    int    `json:"rand"`
}

type Input struct {
	Message string `json:"message"`
}

func main() {
	conn, err := grpc.Dial("reverse-service:3000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	addClient := pb.NewReverseServiceClient(conn)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UFT-8")

		if r.Header.Get("Content-Type") != "" {
			value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
			if value != "application/json" {
				msg := "Content-Type header is not application/json"
				http.Error(w, msg, http.StatusBadRequest)
				return
			}
		}

		if r.Body == http.NoBody {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode("No data submitted")
			return
		}

		var i Input
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		if i.Message == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode("message attributed not provided")
			return
		}

		ctx, cancel := context.WithTimeout(context.TODO(), time.Minute)
		defer cancel()

		var reversedMessage string
		req := &pb.Request{Query: i.Message}
		if resp, err := addClient.Compute(ctx, req); err == nil {
			reversedMessage = resp.Result
		} else {
			msg := fmt.Sprintf("Internal server error: %s", err.Error())
			json.NewEncoder(w).Encode(msg)
		}

		respone := Response{Message: reversedMessage, Rand: generateRandomNumber()}
		responeJson, _ := json.Marshal(respone)

		// Convert bytes to string
		jsonString := string(responeJson)
		json.NewEncoder(w).Encode(jsonString)
	}).Methods("POST")

	fmt.Println("Application is running on : 8080 .....")
	http.ListenAndServe(":8080", router)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UFT-8")
	json.NewEncoder(w).Encode("Server is running")
}

func generateRandomNumber() int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	return random.Intn(10000)
}
