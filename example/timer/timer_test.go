package timer

import (
	"testing"
	"time"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/example"
	"github.com/davyxu/golog"
)

var log *golog.Logger = golog.New("test")

func TestTimer(t *testing.T) {

	signal := test.NewSignalTester(t)

	queue := cellnet.NewEventQueue()

	queue.StartLoop()

	const testTimes = 3

	var count int = testTimes

	cellnet.NewTimer(queue, time.Second, func(t *cellnet.Timer) {
		log.Debugln("timer 1 sec tick")

		signal.Done(1)

		count--

		if count == 0 {
			t.Stop()
			signal.Done(2)
		}
	})

	for i := 0; i < testTimes; i++ {
		signal.WaitAndExpect("timer not tick", 1)
	}

	signal.WaitAndExpect("timer not stop", 2)

}

func TestDelay(t *testing.T) {

	signal := test.NewSignalTester(t)

	queue := cellnet.NewEventQueue()

	queue.StartLoop()

	log.Debugln("delay 1 sec begin")

	queue.DelayPost(time.Second, func() {

		log.Debugln("delay done")
		signal.Done(1)
	})

	signal.WaitAndExpect("delay not work", 1)
}