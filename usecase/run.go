package usecase

import (
	"log"
)

func Run() {
	CheckInit()

	for i := 0; i < Scinario.WorkerNum; i++ {
		worker := NewWorker(i)
		waitGroup.Add(1)
		go worker.Run()
	}

	go func() {
		waitGroup.Wait()
		close(resultChan)
	}()

	NewCollector().Run()
}

func CheckInit() {
	if resultChan == nil {
		log.Fatal("result channel is not initialized")
	}
	if TpPrinter == nil {
		log.Fatal("tp printer is not initialized")
	}
	if AveragePrinter == nil {
		log.Fatal("average printer is not initialized")
	}
	if ResultPrinter == nil {
		log.Fatal("result printer is not initialized")
	}
}
