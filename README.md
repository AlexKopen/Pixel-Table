# Pixel Table
Pixel Table is a cryptocurrency trading bot simulation engine.  It enables the ability to run trading algorithms
against historical data for 100+ stable coin trading pairs.

The engine for this project is built using Go and its corresponding web application client is 
created with Angular.

## Instructions to Use
### 1. Build plugins
The engine leverages Go's [plugin](https://golang.org/pkg/plugin/) package to dynamically load 
the trading logic into the engine.

Navigate to the `/trade-bot` directory and build the `trade-bot-example.go` file as `trade-bot.so`.  This step must be 
completed after every change to the plugin.
```
cd simulator/plugins && go build -buildmode=plugin -o trade-bot.so
```

### 2. Start the server
Upon completing the trade simulations, the engine outputs the results to `output.json`.  
[Gorilla](https://github.com/gorilla/websocket) is used to listen for changes to `output.json` 
and then broadcast the results to the client application through a websocket.

Navigate to the `/server` directory and start the websocket with `output.json` as the first argument.
```
cd server && go run main.go ../output.json
```

### 3. Start the client
An Angular web application is used to display changes broadcast by the websocket.  First, navigate to the `/client` directory and install project dependencies.
```
cd client && npm install
```

Second, start the application.  This will run the web app at [localhost:4200](http://localhost:4200)
```
npm start
```

### 4. Run the simulation
Navigate to `/simulator/pkg` and run `main.go`.
```
cd simulator && go run main.go
```

The engine will process all USDT trading pairs for each coin defined in `/simulator/pkg/constants.go` for the selected date range.
