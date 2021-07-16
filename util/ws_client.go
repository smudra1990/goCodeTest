package util

import (
	"context"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Send pings to peer with this period
const pingPeriod = 30 * time.Second

//MessageChannel is used to channel data from websocket listner.
type MessageChannel struct {
	MessageType int
	Data        []byte
}

// WebSocketClient return websocket client connection
type WebSocketClient struct {
	configStr string
	sendBuf   chan []byte
	ctx       context.Context
	ctxCancel context.CancelFunc

	mu     sync.RWMutex
	wsconn *websocket.Conn
}

// NewWebSocketClient create new websocket connection
func NewWebSocketClient(scheme string, host, path string, message chan MessageChannel) (*WebSocketClient, error) {
	conn := WebSocketClient{
		sendBuf: make(chan []byte, 1),
	}
	conn.ctx, conn.ctxCancel = context.WithCancel(context.Background())

	u := url.URL{Scheme: scheme, Host: host, Path: path}
	conn.configStr = u.String()

	go conn.listen(message)
	return &conn, nil
}

// Connect to web socket.
func (conn *WebSocketClient) Connect() *websocket.Conn {
	conn.mu.Lock()
	defer conn.mu.Unlock()
	if conn.wsconn != nil {
		return conn.wsconn
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		select {
		case <-conn.ctx.Done():
			return nil
		default:
			ws, _, err := websocket.DefaultDialer.Dial(conn.configStr, nil)
			if err != nil {
				conn.log("connect", err, fmt.Sprintf("Cannot connect to websocket: %s", conn.configStr))
				continue
			}
			conn.log("connect", nil, fmt.Sprintf("connected to websocket to %s", conn.configStr))
			conn.wsconn = ws
			return conn.wsconn
		}
	}
}

func (conn *WebSocketClient) listen(message chan MessageChannel) {
	conn.log("listen", nil, fmt.Sprintf("listen for the messages: %s", conn.configStr))
	ws := conn.Connect()
	if ws == nil {
		return
	}
	tickerMessage := struct {
		Method string `json:"method"`
		ID     int    `json:"id"`
		Params struct {
			Symbol   string `json:"symbol"`
			Currency string `json:"currency"`
		} `json:"params"`
	}{}
	tickerMessage.ID = 101
	// Get symbol details
	tickerMessage.Method = "getCurrency"
	tickerMessage.Params.Currency = "ETH"
	conn.wsconn.WriteJSON(tickerMessage)
	tickerMessage.Params.Currency = "BTC"
	conn.wsconn.WriteJSON(tickerMessage)

	//Subscribe Ticker
	tickerMessage.ID = 102
	tickerMessage.Method = "subscribeTicker"
	tickerMessage.Params.Symbol = "ETHBTC"
	conn.wsconn.WriteJSON(tickerMessage)
	tickerMessage.Params.Symbol = "BTCUSD"
	conn.wsconn.WriteJSON(tickerMessage)

	conn.wsconn.WriteJSON(tickerMessage)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-conn.ctx.Done():
			return
		case <-ticker.C:
			for {

				mType, bytMsg, err := ws.ReadMessage()
				if err != nil {
					conn.log("listen", err, "Cannot read websocket message")
					conn.closeWs()
					break
				}
				m := MessageChannel{
					MessageType: mType,
					Data:        bytMsg,
				}
				message <- m
				//conn.log("listen", nil, fmt.Sprintf("websocket msg: %x\n", bytMsg))
			}
		}
	}
}

// Stop will send close message and shutdown websocket connection
func (conn *WebSocketClient) Stop() {
	conn.ctxCancel()
	conn.closeWs()
}

// Close will send close message and shutdown websocket connection
func (conn *WebSocketClient) closeWs() {
	conn.mu.Lock()
	if conn.wsconn != nil {
		conn.wsconn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		conn.wsconn.Close()
		conn.wsconn = nil
	}
	conn.mu.Unlock()
}

// Log print log statement
// In real word I would recommend to use zerolog or any other solution
func (conn *WebSocketClient) log(f string, err error, msg string) {
	if err != nil {
		fmt.Printf("Error in func: %s, err: %v, msg: %s\n", f, err, msg)
	} else {
		fmt.Printf("Log in func: %s, %s\n", f, msg)
	}
}
