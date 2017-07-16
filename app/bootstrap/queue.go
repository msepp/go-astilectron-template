package bootstrap

import (
	"log"

	"github.com/msepp/go-astilectron-template/app/message"
)

// Send sends an message into queue for delivery to renderer.
func (a *App) Send(m *message.Message) {
	a.msgQueue <- *m
}

// messageQueueFlusher sends out message from the queue to given window.
func (a *App) messageQueueFlusher() {
	for {
		select {
		case m, ok := <-a.msgQueue:
			if !ok {
				// Closed. Time to exit.
				return
			}

			if m.Key == message.RequestWindowClose {
				a.Window.Close()
				a.Renderer.Stop()

			} else if m.Key == message.RequestWindowMinimize {
				if err := a.Window.Minimize(); err != nil {
					log.Printf("Error minimizing window: %s", err)
				}

			} else {
				if err := a.Window.Send(m); err != nil {
					log.Printf("While sending to client: %s", err)
				}
			}
		}
	}
}
