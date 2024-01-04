/* Simple EchoServer in GoLang by Phu Phung, customized by Harish Rolla for SecAD*/
package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	
)

const BUFFERSIZE int = 1024
const authenticatedCode = "HZIRSJdkis@//"

type User struct {
	Username string
	Login    bool
	Key      net.Conn
}
type Input struct{
	Command string
	Message string
	User 	string
}
var allClient_conns = make(map[net.Conn]string)
var newClient = make(chan User)
var lostClient = make(chan User)
var authenticatedClient = make(map[net.Conn]User)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <port>\n", os.Args[0])
		os.Exit(0)
	}
	port := os.Args[1]
	if len(port) > 5 {
		fmt.Println("Invalid port value. Try again!")
		os.Exit(1)
	}
	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("Cannot listen on port '" + port + "'!\n")
		os.Exit(2)
	}
	fmt.Println("ChatServer in GoLang developed by Phu Phung, SecAD, revised by Harish Rolla")
	fmt.Printf("ChatServer is listening on port '%s' ...\n", port)
	go func() {
		for {
			client_conn, _ := server.Accept()
			go wait_login(client_conn)

		}
	}()

	for {
		select {
		case user := <-newClient:
			toClientData := fmt.Sprintf("%s\n",authenticatedCode)
			client_conn := user.Key
			authenticatedClient[client_conn] = user
			welcomeMessage := fmt.Sprintf("A new client '%s' connected!\n# of connected clients: %d\n", user.Username, len(authenticatedClient))
			fmt.Println(welcomeMessage)
			sendTo(client_conn, []byte(toClientData))
			sendToAll([]byte(welcomeMessage))
			
			go client_goroutine(user)

		case user := <-lostClient:
			client_conn := user.Key
			delete(authenticatedClient, client_conn)
			byemessage := fmt.Sprintf("Client %s is disconnected!\n# of clients connected: %d\n", user.Username, len(authenticatedClient))
			sendToAll([]byte(byemessage))
		default:

		}
	}

}

func wait_login(client_conn net.Conn) {
	var buffer [BUFFERSIZE]byte
	for {
		byte_received, read_error := client_conn.Read(buffer[0:])
		if read_error != nil {
			fmt.Println("DEBUG > reading error.")
			return
		}

		data := buffer[0:byte_received]
		fmt.Printf("From client: '%s'\nReceived data: %sData size: %d\n\n", client_conn.RemoteAddr().String(), data, len(data))
		authenticated, username, loginmessage := checklogin(data)
		//if len(data) >= 5 && string(data[0:5]) == "login" {
		if authenticated {
			fmt.Printf("DEBUG > login data,]. User: %s loggedin succesfully \n", username)
			currentLoggedUser := User{Username: username, Login: true, Key: client_conn}
			newClient <- currentLoggedUser
			return

		} else {
			fmt.Println("DEBUG > non-login data" + loginmessage)
			toClientData := fmt.Sprintf("non-login data Error: %s\n", loginmessage)
			sendTo(client_conn, []byte(toClientData))

		}
	}
}

func client_goroutine(user User) {
	var buffer [BUFFERSIZE]byte
	client_conn := user.Key
	for {
		byte_received, read_err := client_conn.Read(buffer[0:])
		if read_err != nil {
			fmt.Println("Error in receiving...")
			lostClient <- user
			return
		}

		data := buffer[0:byte_received]
		fmt.Printf("From client: '%s'\nReceived data: %sData size: %d\n\n", client_conn.RemoteAddr().String(), data, len(data))
		check_user_message(client_conn, data, user)
	}
}

func check_user_message(client_conn net.Conn, data []byte, user User) {
	

	var input Input
	err := json.Unmarshal(data, &input)

	if err != nil || input.Command == "" {
		fmt.Printf(`Error, Expected data {"command":"","message":""}`)
		return
	}
	check_command(input, user)

}

func show_clients(client_conn net.Conn){
	sendTo(client_conn, []byte("Users connected to the server\n"))
	for _, user := range authenticatedClient{
		sendTo(client_conn, []byte(user.Username+"\n"))
	}
}

func check_command(input Input, user User){
	switch input.Command {
	case "showclients":
		show_clients(user.Key)
	case "private":
		private_chat(input.User, input.Message, user)
	case "public":
		sendToAll([]byte("Public message from "+user.Username+"\n Message: "+input.Message))

	default:
		fmt.Printf("ERROR , Expected comands are showclients,private,public")
	}
}

func private_chat(username string, message string, user User){
	user_presence, connections :=check_user(username, user)
	if(user_presence){
		for _, client_conn := range connections{
			sendTo(client_conn, []byte("A private from "+user.Username+"\n Message:"+message))
		}
	return
	}
	sendTo(user.Key,[]byte("user that you wanted to send  is not present in server"))
	
}
func check_user(username string, user User)(bool,[]net.Conn){
	var connections []net.Conn
	var presence = false
	for _,Username := range authenticatedClient{
		if username == Username.Username {
			connections = append(connections, Username.Key)
			presence = true
		}
	}
	return presence, connections
}

func sendTo(client_conn net.Conn, data []byte) {
	_, write_err := client_conn.Write(data)
	if write_err != nil {
		fmt.Println("Error in sending...")
		return
	}
}

func sendToAll(data []byte) {
	for client_conn, _ := range authenticatedClient {
		_, write_err := client_conn.Write(data)
		if write_err != nil {
			fmt.Printf("Error in sending. \n")
			continue
		}
	}
	fmt.Printf("Sent below data to all clients:\n%s\n", data)
}

func checklogin(data []byte) (bool, string, string) {
	type Account struct {
		Username string
		Password string
	}
	var account Account
	err := json.Unmarshal(data, &account)
	if err != nil || account.Password == "" || account.Username == "" {

		return false, "", ` REQUEST : EXPECTED DATA {"username":"","password":""}`
	}

	fmt.Printf("account=%s\n", account)
	if checkaccount(account.Username, account.Password) {
		return true, account.Username, "logged in succesfully"
	}
	return false, "", "invalid user name "
}

func checkaccount(user string, password string) bool {
	if user == "harish" && password == "123456" {
		return true
	}
	if user == "rollah1" && password == "123456"{
		return true
	}
	if user == "harishrolla" && password == "test123"{
		return true
	}
	return false
}
