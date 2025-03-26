package storage

import (
	"kwdb/app/storage/displacement"
	"kwdb/internal/helper/informer"
	"time"
)

type workerSem struct {
	ch        chan struct{}
	lastCheck time.Time
}

var sem workerSem

func RunWatcher(policy *displacement.Policy) {

	sem = workerSem{make(chan struct{}, 1), time.Now()}

	go worker(sem.ch, policy)

	informer.InfChan <- "Запущен watcher вытеснения"

	for {
		if time.Now().After(sem.lastCheck) {
			sem.ch <- struct{}{}
			sem.lastCheck = time.Now()
		}

		time.Sleep(10 * time.Second)
	}
}

func worker(sem chan struct{}, p *displacement.Policy) {
	for range sem {
		informer.InfChan <- "Запущена очистка памяти"
		//p.Cut()
	}
}
