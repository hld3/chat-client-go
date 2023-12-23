package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

var clients = make(map[string]*Client)

type Message struct {
	Type    string          `json:"type"`
	Content json.RawMessage `json:"content"`
}

type UserData struct {
	UserId string `json:"userId"`
}

type ChatMessage struct {
	RecipientId string `json:"recipientId"`
	Text        string `json:"text"`
}

type ErrorMessage struct {
	Message string `json:"errorMessage"`
}

type Client struct {
	Id   string // I don't think this is needed.
	Conn net.Conn
}

func StartServer(host string, port string) error {
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatal("Error listening:", err.Error())
		return err
	}
	defer listener.Close()
	log.Println("Listening on", addr, ":", port)

	// run a loop to accept all incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting:", err.Error())
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// continue the loop for a connection to handle new messages from the connection.
	for {
		buffer := make([]byte, 2048)
		length, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error reading to the buffer:", err)
		}

		var message Message
		err = json.Unmarshal(buffer[:length], &message)
		if err != nil {
			SendErrorMessage(ErrorMessage{Message: fmt.Sprint("Error unmarshaling full message:", err)}, conn)
			log.Println("Error unmarshaling message.")
		}

		switch message.Type {
		case "initialConnection":
			log.Println("New initial connection recieved.")
			var userData UserData
			err = json.Unmarshal([]byte(message.Content), &userData)
			if err != nil {
				SendErrorMessage(ErrorMessage{Message: fmt.Sprint("Error unmarshaling user data:", err)}, conn)
				log.Println("Error unmarshaling user data.")
			}

			AddClient(Client{Id: userData.UserId, Conn: conn})
		case "chatMessage":
			log.Println("New chat message recieved.")
			var chatMessage ChatMessage
			err = json.Unmarshal([]byte(message.Content), &chatMessage)
			if err != nil {
				SendErrorMessage(ErrorMessage{Message: fmt.Sprint("Error unmarshaling chat message:", err)}, conn)
				log.Println("Error unmarshaling chat message.")
			}

			// TODO do something with the message!
		}
	}
}

func AddClient(client Client) {
	clients[client.Id] = &client
}

func RemoveClient(clientId string) {
	delete(clients, clientId)
}

// TODO update with a reciever.
func BroadcastMessage(message string, sender *Client) {
	// Broadcast a message to clients
}

func SendErrorMessage(message ErrorMessage, conn net.Conn) {
	messBytes, err := json.Marshal(message)
	if err != nil {
		log.Println("Error marshaling error message:", err)
	}

	_, err = conn.Write(messBytes)
	if err != nil {
		log.Println("Error sending error message to client:", err)
	}
}
