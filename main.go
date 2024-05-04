package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type Message struct {
	SenderId   string
	ReceiverId string
	Content    string
	Time       time.Time
}

type User struct {
	ID       string
	Messages []Message
}

var users = map[string]*User{
	"user1": {ID: "user_1"},
	"user2": {ID: "user_2"},
	"user3": {ID: "user_3"},
	"user4": {ID: "user_4"},
	"user5": {ID: "user_5"},
}

func main() {
	fmt.Println("Enter your userId")
	var userId string
	fmt.Scanln(&userId)

	if _, exist := users[userId]; !exist {
		users[userId] = &User{ID: userId}
	}

	for {
		fmt.Println("\n1. Send message to a particular user")
		fmt.Println("2. Send message to all users")
		fmt.Println("3. Read messages")
		fmt.Println("4. Add new user")
		fmt.Println("5. Exit")

		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			sendMessageToParticularUser(userId)
		case 2:
			sendMessageToAll(userId)
		case 3:
			readMessages()
		case 4:
			addNewUser()
		case 5:
			exit()
		default:
			fmt.Println("Please enter a valid choice")
		}
	}
}

func sendMessageToParticularUser(senderId string) {
	fmt.Println("Available users:")
	for userID := range users {
		if userID == senderId {
			continue
		}
		fmt.Println(userID)
	}
	fmt.Println("Enter the receiver's Id")
	var receiverId string
	fmt.Scanln(&receiverId)
	if receiverId == senderId {
		fmt.Println("Cannot send message to yourself")
		return
	}

	fmt.Println("Enter the message")
	reader := bufio.NewReader(os.Stdin)
	msg, _ := reader.ReadString('\n')
	msg = strings.TrimSpace(msg)

	if msg == "" {
		msg = getRandomFact()
	}

	content := Message{
		SenderId:   senderId,
		ReceiverId: receiverId,
		Content:    msg,
		Time:       time.Now(),
	}

	if _, exist := users[receiverId]; exist {
		users[receiverId].Messages = append(users[receiverId].Messages, content)
		fmt.Printf("\nMessage sent from %v to %v", senderId, receiverId)
	} else {
		fmt.Println("User Does not found")
	}
}

func sendMessageToAll(senderId string) {
	fmt.Println("Enter the message")
	reader := bufio.NewReader(os.Stdin)
	msg, _ := reader.ReadString('\n')
	msg = strings.TrimSpace(msg)

	if msg == "" {
		msg = getRandomFact()
	}

	for receiverId := range users {
		if receiverId != senderId {
			content := Message{
				SenderId:   senderId,
				ReceiverId: receiverId,
				Content:    msg,
				Time:       time.Now(),
			}

			if _, exist := users[receiverId]; exist {
				users[receiverId].Messages = append(users[receiverId].Messages, content)
				fmt.Printf("\nMessage sent from %v to %v", senderId, receiverId)
			} else {
				fmt.Println("User Does not found")
			}
		}
	}
	fmt.Println("\nPress enter to continue")
	fmt.Scanln()

}

func readMessages() {
	var userId string
	fmt.Println("Enter the user id")
	fmt.Scanln(&userId)
	if user, exist := users[userId]; exist {
		fmt.Println("Your messages:")
		for _, msg := range user.Messages {
			fmt.Printf("[%s] %s: %s\n", msg.Time.Format("2006-01-02 15:04:05"), msg.SenderId, msg.Content)
		}
	} else {
		fmt.Println("User not found.")
	}
}

func getRandomFact() string {
	res, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		fmt.Println("Error fetching api:", err)
	}
	defer res.Body.Close()

	var RandomFact struct {
		Fact string `json:"fact"`
	}
	if err := json.NewDecoder(res.Body).Decode(&RandomFact); err != nil {
		fmt.Println("Error decoding cat fact response:", err)
		return "Failed to decode cat fact response"
	}

	return RandomFact.Fact
}

func addNewUser() {
	fmt.Println("Enter the userId")
	var userId string
	fmt.Scanln(&userId)
	if _, ok := users[userId]; ok {
		fmt.Println("User already exists.")
		return
	}

	users[userId] = &User{ID: userId}
	fmt.Println("User", userId, "added successfully.")
}

func exit() {
	fmt.Println("Exiting...")

	for _, user := range users {
		fmt.Printf("User %s's messages:\n", user.ID)
		for _, msg := range user.Messages {
			fmt.Printf("[%s] %s: %s\n", msg.Time.Format("2006-01-02 15:04:05"), msg.SenderId, msg.Content)
		}
	}

	os.Exit(0)
}