package sentry

import (
	"fmt"

	"github.com/getsentry/raven-go"
)

func init() {
	raven.SetDSN("http://634a78c491f54d2e9666e2ff36e0d747:1bd683754e9f4c38838e50f0b2b28d49@192.168.1.100:9000/2")
}

// func TestError() {
// 	var count = 100 * 100
// 	var wg sync.WaitGroup

// 	for i := 0; i < count; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			err := fmt.Errorf("Error test: %d", time.Now().Unix())
// 			log.Info(raven.CaptureError(err, map[string]string{"editor": "vscode"}), ":", i)
// 			wg.Done()
// 		}(i)
// 	}

// 	wg.Wait()

// 	time.Sleep(time.Second * 10)
// 	log.Info("Finish")
// }

func TestCapture() {
	p := raven.NewPacket("packet", nil)
	id, ch := raven.Capture(p, nil)
	<-ch
	fmt.Println(id)
}
