var net = require('net');
 
var readlineSync = require('readline-sync');
var username;
var password;
var authenticated;

var buf = Buffer.alloc(1024);

if(process.argv.length != 4){
	console.log("Usage: node %s <host> <port>", process.argv[1]);
	process.exit(1);	
}

var host=process.argv[2];
var port=process.argv[3];

if(host.length >253 || port.length >5 ){
	console.log("Invalid host or port. Try again!\nUsage: node %s <port>", process.argv[1]);
	process.exit(1);	
}

var client = new net.Socket();
console.log("Simple telnet.js developed by harish rolla");
console.log("Connecting to: %s:%s", host, port);

client.connect(port,host, connected);

function connected(){       
	loginsync()
}
client.on("data", data => {
	//console.log("Recived data:" + data);
	if(!authenticated){
		if(data.includes("HZIRSJdkis@//")){
			console.log("Hi welcome to chat servers\n logged in sucessfully");
			authenticated= true;
			menu();
			chat();

		}else{
			console.log("authenticated failed");
			loginsync();
		}
	}
	else{
		console.log("Recived data:" + data);
	}

});
client.on("error", function(err){
	console.log("Error");
	process.exit(2);
});
client.on("close", function(data){
	console.log("connected has been disconnected");
	process.exit(3);
});

function chat(){
	var keyboard = require('readline').createInterface({
			input: process.stdin,
			output: process.stdout
	});	

	keyboard.on('line', (input) => {
		if (input.startsWith("public")){
			split = input.split(":-");
			client.write(`{"command":"public","message":"${split[1]}"}`);
		}
		else if (input.startsWith("showclients")){
			client.write(`{"command":"showclients","message":"","user":""}`);
		}
		else if (input.startsWith("private")){
			split = input.split(":-");
			client.write(`{"command":"private","message":"${split[2]}","user":"${split[1]}"}`);
		}
		else if (input === "help") {
			menu();

		}
 		else if(input === "exit"){
			client.destroy();
			console.log("disconnected!");
			process.exit();
		}else{
			console.log("please check the menu and give the messages as mentioned");
			menu();
		}
			
	});

}

function menu(){
	console.log("*******************menu********************")
	console.log("start message with key words only, and donot use :- this symbol between the message");
	console.log("to chat with all users start with public as key word and :- symbol and then message ");
	console.log("to chate with particulart user , start with private and then :- symbol and then username and :- symbol and them message  ");
	console.log("EXAMPLE-private:-harish:-hi harish this is rolla how are you?");
	console.log("to get users connected to the network- type showclients");
	console.log("type help to get menu and type exit to logout");
}

function privatechat(){
	userName= readlineSync.question('Username:\n');
	message = readlineSync.question('Enter your message that you wanted to send\n');
	client.write(message);
}
function inputValidated(logindata) {
	buf=logindata;
	if (buf.length>=5){
		return true;
	};
	return false;
};

function loginsync(){
	username= readlineSync.question('Username:\n');
	if(!inputValidated(username)){
		console.log("Username must have at least 5 characters. Please try again!");
		loginsync();
		return;
	}
	password = readlineSync.question('Password:',{
		hideEchoBack: true
	});
	if(!inputValidated(password)){
		console.log("passwor must have at least 5 characters. Please try again");
		loginsync();
		return;
	}
	var login = `{"username":"${username}","password":"${password}"}`;
	client.write(login);
};
