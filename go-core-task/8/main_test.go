package main

import (
	"testing"
	"time"
)

func TestCustomWaitGroup(t *testing.T) {
	tests := []struct {
		name        string
		numG        int
		expectedRes string
	}{
		{
			name:        "тест 5 горутин",
			numG:        5,
			expectedRes: "все горутины завершились",
		},
		{
			name:        "тест 0 горутин",
			numG:        0,
			expectedRes: "все горутины завершились",
		},
		{
			name:        "тест 1 горутины",
			numG:        1,
			expectedRes: "все горутины завершились",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := NewCustomWaitGroup()

			done := make(chan struct{})

			if tt.numG == 0 {
				close(done)
			} else {
				for i := 0; i < tt.numG; i++ {
					wg.Add(1)
					go func(index int) {
						defer wg.Done()
						time.Sleep(time.Second * 1)
					}(i)
				}

				go func() {
					wg.Wait()
					close(done)
				}()
			}

			select {
			case <-done:
				t.Log(tt.expectedRes)
			case <-time.After(time.Second * 2):
				t.Fatalf("тест не завершился вовремя: %s", tt.name)
			}
		})
	}
}
