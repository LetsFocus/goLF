package websockets

import (
	"context"
	"errors"
	"io"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/LetsFocus/goLF/logger"
	"golang.org/x/time/rate"
	"nhooyr.io/websocket"
)

type WebSocket struct {
	// subscriberMessageBuffer controls the max number of messages that can be
	// queued for a subscriber before it is kicked. Defaults to 16.
	subscriberMessageBuffer int

	// publishLimiter controls the rate limit applied to the publish endpoint.
	// Defaults to one publish every 100ms with a burst of 8.
	publishLimiter *rate.Limiter
	log            logger.CustomLogger
	serveMux       http.ServeMux
	subscribersMu  sync.Mutex
	subscribers    map[*subscriber]struct{}
}

func NewWebSocket() *WebSocket {
	return &WebSocket{
		subscriberMessageBuffer: 16,
		publishLimiter:          rate.NewLimiter(rate.Every(time.Millisecond*100), 8),
		log:                     logger.CustomLogger{},
		serveMux:                http.ServeMux{},
		subscribersMu:           sync.Mutex{},
		subscribers:             make(map[*subscriber]struct{}),
	}
}

func (ws *WebSocket) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws.serveMux.ServeHTTP(w, r)
}

// subscriber represents a subscriber.
// Messages are sent on the msgChannel and if the client cannot keep up with the messages, closeSlow is called.
type subscriber struct {
	msgChannel chan []byte
	closeSlow  func()
}

func (ws *WebSocket) Publish(path string) {
	// services will validate on what method this publish is allowed

	// add handler wrapped inside context as a later feature, pass handler as input arg to the signature
	ws.serveMux.HandleFunc(path, ws.publishHandler)
}

func (ws *WebSocket) publishHandler(w http.ResponseWriter, r *http.Request) {
	body := http.MaxBytesReader(w, r.Body, 8192)
	msg, err := io.ReadAll(body)
	if err != nil {
		ws.log.Errorf("Request body too large than allowed %v bytes", 8192)

		http.Error(w, http.StatusText(http.StatusRequestEntityTooLarge), http.StatusRequestEntityTooLarge)
		return
	}

	ws.publish(msg)

	w.WriteHeader(http.StatusAccepted)
}

// publish publishes the msg to all subscribers.
// It never blocks and so messages to slow subscribers are dropped.
func (ws *WebSocket) publish(msg []byte) {
	ws.subscribersMu.Lock()
	defer ws.subscribersMu.Unlock()

	ws.publishLimiter.Wait(context.Background())

	for s := range ws.subscribers {
		select {
		case s.msgChannel <- msg:
		default:
			go s.closeSlow()
		}
	}
}

func (ws *WebSocket) Subscribe(path string) {
	// services will validate on what method this publish is allowed

	// add handler wrapped inside context as a later feature
	ws.serveMux.HandleFunc(path, ws.subscribeHandler)
}

func (ws *WebSocket) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	err := ws.subscribe(r.Context(), w, r)
	if errors.Is(err, context.Canceled) {
		return
	}

	if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
		websocket.CloseStatus(err) == websocket.StatusGoingAway {
		return
	}
	if err != nil {
		ws.log.Errorf("Error occured while subscribing: %v", err)

		return
	}
}

func (ws *WebSocket) subscribe(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var (
		mu     sync.Mutex
		wsConn *websocket.Conn
		closed bool
	)

	s := &subscriber{
		msgChannel: make(chan []byte, ws.subscriberMessageBuffer),
		closeSlow: func() {
			mu.Lock()
			defer mu.Unlock()
			closed = true

			if wsConn != nil {
				// add logger if needed
				ws.log.Errorf("connection too slow to keep up with messages: %v", websocket.StatusPolicyViolation)

				wsConn.Close(websocket.StatusPolicyViolation, "connection too slow to keep up with messages")
			}
		},
	}

	ws.addSubscriber(s)
	defer ws.deleteSubscriber(s)

	c2, err := websocket.Accept(w, r, nil)
	if err != nil {
		return err
	}

	mu.Lock()
	if closed {
		mu.Unlock()
		return net.ErrClosed
	}

	wsConn = c2
	mu.Unlock()
	defer wsConn.CloseNow()

	ctx = wsConn.CloseRead(ctx)

	for {
		select {
		case msg := <-s.msgChannel:
			er := writeWithTimeout(ctx, time.Second*5, wsConn, msg)
			if er != nil {
				// add logger
				return er
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// addSubscriber registers a subscriber.
func (ws *WebSocket) addSubscriber(s *subscriber) {
	ws.subscribersMu.Lock()
	ws.subscribers[s] = struct{}{}
	ws.subscribersMu.Unlock()
}

// deleteSubscriber deletes the given subscriber.
func (ws *WebSocket) deleteSubscriber(s *subscriber) {
	ws.subscribersMu.Lock()
	delete(ws.subscribers, s)
	ws.subscribersMu.Unlock()
}

func writeWithTimeout(ctx context.Context, timeout time.Duration, wsConn *websocket.Conn, msg []byte) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return wsConn.Write(ctx, websocket.MessageText, msg)
}
