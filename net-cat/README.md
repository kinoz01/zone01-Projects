## Subject
This project consists on recreating the `NetCat` in a Server-Client Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

## Usage

If no port is specified the chat server wil run on the `8989` port by default:

```bash
go run . 
```

You can choose another port using:

```bash
go run . <poort>
```

> To change your username during the chat use the keyword `/name`.

