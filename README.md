
# Go-CLI-Messaging

## Overview
This project is a command-line interface (CLI) messaging application developed in GoLang. It allows users to send messages to specific users or broadcast messages to all users.

## Features
Some features are \
1.Send Message to Specific User: Users can send messages to a specific user by specifying the recipient's ID.\
2.Broadcast Message: Users can broadcast messages to all users.\
3.Read Messages: Users can read their messages.\
4.Add New User: Administrators can add new users to the system.

 ## Usage
 Upon running the application, users will be prompted to enter their user ID.\
Users can then choose from the following options:\
1.Send message to a particular user\
2.Send message to all users\
3.Read messages\
4.Add new user\
5.Exit

## Dependencies
1.bufio: For reading user input.\
2.encoding/json: For encoding and decoding JSON data.\
3.fmt: For printing formatted output.\
4.net/http: For making HTTP requests.\
5.os: For interacting with the operating system.\
6.strings: For string manipulation.\
7.time: For working with time-related functions.

## Screenshots

![send message](/images/img1.png)
When we wnat to send message to any particular user

![send message](/images/img2.png)

When we wnat to send message to all the users

![send message](/images/img3.png)

When we wnat to read the messages of any particular user

![send message](/images/img4.png)

When we wnat to add new user

![send message](/images/img5.png)

When we wnat to exit the command line interface

![send message](/images/img6.png)

When we send a blank message it generate a random message from api
