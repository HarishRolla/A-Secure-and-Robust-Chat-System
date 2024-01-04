# A-Secure-and-Robust-Chat-System

<a href="https://github.com/HarishRolla/A-Secure-and-Robust-Chat-System">Link to repository</a>


## 1. Introduction

I have created a comprehensive chat server application that operates on both the server and client sides. The server is capable of managing multiple client connections and enabling communication between them. It operates by listening on a designated port, authenticating clients through their usernames and passwords, and providing access to group and private chat functionalities for authenticated users. Furthermore, the application is designed with security and reliability in mind, ensuring a robust and secure user experience.


## 2. Design

The interaction begins with the user establishing a connection with the client. The clien asks the user for their username and password in a loop until the user provides valid credentials. The client then sends the credentials to the waitlogin function in JSON format

The waitlogin function sends the credentials to the checklogin function to retrieve the username and password. If the credentials are valid, the checklogin function sends them to the checkaccount function to validate the username and password. If the validation passes, the checkaccount function returns true to the checklogin function, which then sends the username, password, and authentication key to the waitlogin function.

If the validation fails at any point, the corresponding function returns false, and the waitlogin function displays an error message to the client. If the credentials are not in the expected JSON format, the checklogin function also displays an error message to the waitlogin function.

If the authentication is successful, the waitlogin function sends the authenticated key to the client and adds the connection to the allLoggedin_conns list and the user to the authenticatedUserList. The function then uses a goroutine to send a message to all connections that a new user has joined the chat.

If the authentication fails, the waitlogin function displays an error message to the client. The client then displays the menu to the user and waits for the user to enter a command. If the command is valid, the client sends it to the goroutine, which executes the corresponding function and sends the output to the client. If the command is invalid, the client displays the menu again to the user

If the authentication key check fails, the client displays an error message to the user and self-calls to ask the user for their username and password again

<img src="https://github.com/HarishRolla/A-Secure-and-Robust-Chat-System/blob/main/demoScreenshots/fig0.png " alt="Alt Text" style="width: 800px;">

The above Figure shows Design for the user login to the Chatserver through the client

The client_goroutine function is responsible for receiving data from the client, and it does so by reading data from the client_conn connection. This data is then passed to the check_user_message function, which attempts to parse the data as a JSON-formatted Input object. If the parsing is successful and the Input object has a non-empty Command field, the Input object is then passed to the check_command function to determine what action should be taken.

Check_command function takes an Input struct and a User struct as inputs, which contain information about the command and the user making the request, respectively. It then checks the Command field of the Input struct and takes appropriate actions based on the command.

If the command is "showclients", it calls the show_clients function, which takes a net.Conn as input, sends a message to the client to indicate that it is sending the list of connected clients, and then iterates through the authenticatedClient list, sending each client's username to the client who requested the list.

If the command is "private", it calls the private_chat function, passing the User struct, the recipient's username, and the message as parameters. The private_chat function checks if the recipient is present in the authenticatedClient list and sends the message to all connections associated with the recipient's username.

If the command is "public", it calls the sendToAll function, passing the message as a parameter. The sendToAll function iterates through all connections in the authenticatedClient list and sends the message to each connection.

If the command does not match any of the expected commands, the function prints an error message indicating that the command is not recognized.


## Handled datarace issues:

In the Go language chat server implementation, I encountered data race issues that needed to be addressed. To overcome these issues, I utilized channels to synchronize the access to shared resources. By using channels, I was able to ensure that only one goroutine accessed a shared resource at a time, preventing any data race issues

Overall, the use of channels was an effective solution for handling data race issues in the Go language chat server implementation, and it helped to ensure the stability and reliability of the system.


## Handled Security Issues:

In the chat systemr implementation, security issues were addressed by implementing several measures. Firstly, the username and password were checked in the client, and the data was sent to a buffer for security purposes. The buffersize was then used to send the data to the server securely

Additionally, in the server, the data was also checked for security purposes by sending it to a buffer.(multi level security) String vulnerabilities were also taken care of in every print statement to ensure the data was not susceptible to any malicious attacks.

Overall, these measures helped to address security issues and ensure that the chat server implementation was secure and reliable for its users


## 3. Test cases and Demo

When a client tries to connect with the chat server, the client needs to provide a username and password. If the user is authenticated, this is shown below figure 1.2. After logging in successfully, the menu will be displayed for the user to chat with other users connected to the chat server

### Menu:
start the message with keywords only, and do not use :- this symbol between the message
to chat with all users start with public as keyword and:- symbol and then message
to chat with a particular user, start with private and then:- symbol and then username and:- symbol and them message
EXAMPLE-private:-harish:-hi harish this is rolla how are you?
to get users connected to the network- type showclients
type help to get a menu and type exit to logout

<img src="https://github.com/HarishRolla/A-Secure-and-Robust-Chat-System/blob/main/demoScreenshots/fig1.png " alt="Alt Text" style="width: 800px;">

Fig 1.1:- a user is successfully authenticated and connected to the server with a user name and
password

If a user tries to log in with an invalid username and password, the application will not allow the user to log in to the server. This ensures that the application is secure and that only authenticated users are granted access to the chat server. This is shown in below figure 1.2

<img src="https://github.com/HarishRolla/A-Secure-and-Robust-Chat-System/blob/main/demoScreenshots/fig2.png " alt="Alt Text" style="width: 800px;">

Fig 1.2: unauthenticated user trying to log into the chat server but failed to login

The chat server is designed to handle multiple clients simultaneously, which means it can accept connections from multiple users at the same time. 

The chat server keeps track of all the clients that are currently connected and also updates the users when a new client connects or an existing client disconnects. In the given scenario, the figure 1.3 illustrates that the server is currently connected to two clients, "rollah1" and "harish." When "rollah1" connects to the server, the server sends a message to "harish" informing them of the new connection. 

This way, the chat server ensures that all the clients are aware of who is currently connected to the server, and they can start communicating with each other in real time.

<img src="https://github.com/HarishRolla/A-Secure-and-Robust-Chat-System/blob/main/demoScreenshots/fig3.png " alt="Alt Text" style="width: 800px;">

Fig 1.3:- server accepts multiple clients, rollah1 and harish users are connected to the server

The menu provides various options for the user to interact with the chat server. To send a message to all the users that are connected to the server, the user needs to use the keyword "public" followed by the symbol ":-" and then the message they want to send. The message format would be: "public:-[message]"

<img src="https://github.com/HarishRolla/A-Secure-and-Robust-Chat-System/blob/main/demoScreenshots/fig4.png " alt="Alt Text" style="width: 800px;">

Fig 1.4:- user rollah1 sends a public message to all users connected to the server

For example, in figure 1.4, "rollah1" sends a public message using the above format. All the connected users, including "rollah1" and "harish," receive the message as "Public message from rollah1" followed by the actual message.

If a user wants to send a private message to a particular user, they can use the keyword "private" followed by the symbol ":-", the username of the recipient, another ":-" symbol, and then the actual message. The message format would be:

"private:-[username]:-[message]"

As shown in Figure 1.4, "rollah1" can send a private message to "harish" using the following format:

"private:-harish:-[message]"

"harish" is logged into two systems, both systems will receive the message from "rollah1". Each system will display the message as "Private message from rollah1" followed by the actual message. This is because the chat server sends the message to the specific user, regardless of the number of systems they are logged into.

Therefore, when "rollah1" sends the private message to "harish," both of "harish's" systems will receive the message, and it will be displayed as "Private message from rollah1" followed by the actual message.


<img src="https://github.com/HarishRolla/A-Secure-and-Robust-Chat-System/blob/main/demoScreenshots/fig5.png " alt="Alt Text" style="width: 800px;">

Fig 1.5 User rollah1 sending a private message to harish

To obtain a list of all the users currently connected to the chat server, the user can type "showclients" keyword in the chat window. This will prompt the server to send a list of all the connected users to the requesting user.

As shown in Figure 1.6, the output of the "showclients" command displays a list of all the connected users, along with their respective IP addresses and ports. This command is useful for users who want to know who is currently available on the chat server, and to whom they can send messages.

<img src="https://github.com/HarishRolla/A-Secure-and-Robust-Chat-System/blob/main/demoScreenshots/fig6.png " alt="Alt Text" style="width: 800px;">

Fig 1.6: getting the users list by command showclients

To log out from the chat server from a specific system, the user can type the keyword "exit" in the chat window. This will prompt the server to log out the user from that specific system.

However, if the user is logged into multiple systems, and uses the "exit" keyword, they will only be logged out from the system where they typed the "exit" keyword. They will remain logged in to the chat server from other systems where they did not type the "exit" keyword. 

As shown in Figure 1.7, when the user "harish" types the "exit" keyword on one system, they are logged out of the server from that system, but they remain logged in to the server from the other system. The chat server sends a message to the user indicating that they have been logged out of that system.

<img src="https://github.com/HarishRolla/A-Secure-and-Robust-Chat-System/blob/main/demoScreenshots/fig7.png " alt="Alt Text" style="width: 800px;">

Fig 1.7: User Harish logged out from one of the systems but still logged in another system

As shown in Figure 1.8 below, even if the user "harish" is disconnected from the chat server on one system, they are still able to receive and send messages from the chat server on the other system.

This is because the chat server allows multiple connections from the same user, if a user is logged in from multiple systems, and one of those systems is disconnected from the chat server, the user can still communicate with the other users on the chat server from the other systems where they are still connected.

In the example shown in Figure 1.8 below, "harish" was disconnected from the chat server on one system, but they were still able to receive and send messages from the server on the other system where they were still connected.

<img src="https://github.com/HarishRolla/A-Secure-and-Robust-Chat-System/blob/main/demoScreenshots/fig8.png " alt="Alt Text" style="width: 800px;">

Fig 1.8:- harish in system 2 will get the messages even the harish logged out from system 1

The "help" command can be used at any time to display the chat server menu. As shown in Figure 1.9, when the user types "help" in the chat window, the chat server displays the menu of available commands and their usage.

This menu provides information on how to use the chat server, including how to send messages to all users, how to send private messages to specific users, how to obtain a list of currently connected users, and how to log out from the chat server.

The "help" command is a useful reference tool for users who may not be familiar with all the available commands or need a quick reminder on how to use them.

<img src="https://github.com/HarishRolla/A-Secure-and-Robust-Chat-System/blob/main/demoScreenshots/fig9.png " alt="Alt Text" style="width: 600px;">

Fig 1.9: help command gives menu at any time

As shown in Figure 1.10 below, the chat server allows users to log in from multiple systems simultaneously. In this example, the user "rollah1" is logged in to the chat server from two different systems or windows.

<img src="https://github.com/HarishRolla/A-Secure-and-Robust-Chat-System/blob/main/demoScreenshots/fig10.png " alt="Alt Text" style="width: 800px;">

Fig 1.10: user rollah1 logged into many systems.


