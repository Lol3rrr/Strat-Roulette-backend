package database

import (
	"time"

	"github.com/sirupsen/logrus"
)

func (s *session) loadCredsAndReconnect() error {
	data, err := s.loadCreds()
	if err != nil {
		logrus.Errorf("Could not load Credentials: %s \n", err)
	}

	go func() {
		for {
			leaseTime := data.LeaseDuration

			_, err := s.VaultSession.Renew(data.LeaseID, leaseTime)
			if err != nil {
				for {
					waitTime := 1

					_, err := s.loadCreds()
					if err == nil {
						err = s.connect()
						if err != nil {
							logrus.Errorf("Could not connect to Database: %s \n", err)
						} else {
							break
						}
					} else {
						logrus.Errorf("Could not load new Credentials: %s \n", err)
					}

					logrus.Errorf("Trying again in %d Seconds \n", waitTime)
					time.Sleep(time.Duration(waitTime) * time.Second)
					waitTime = waitTime * 2
				}
			}

			sleepTime := time.Duration(leaseTime/2) * time.Second
			time.Sleep(sleepTime)
		}
	}()

	return nil
}
