You can test the program by downloading the server from one of these links:
- [link](https://assets.01-edu.org/guess-it/guess-it.zip)
- [docker-app link](https://assets.01-edu.org/guess-it/guess-it-dockerized.zip)

Copy the sudent folder conatining the student guesser and move it to the server folder.

Make the binaries inside the ai folder executables:

```bash
chmod +x *
```

Now you can run the program using one of these ways:

1. Using `node`:

```bash
npm install
node server.js
```

2. Using `docker compose`:

```bash
docker-compose up
```

To reload the server use:

```bash
docker-compose down -v
docker-compse up --build
```

3. Using `dockerfile`:

```bash
docker build -t guesser .
docker run -p 3000:3000 guesser
``` 
