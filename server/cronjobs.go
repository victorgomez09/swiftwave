package server

import (
	HAPROXY_MANAGER "keroku/m/haproxy_manager"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *Server) InitCronJobs() {
	go s.MovePendingApplicationsToImageGenerationQueueCronjob()
	go s.MoveRedeployPendingApplicationsToImageGenerationQueueCronjob()
	go s.MoveDeployingPendingApplicationsToDeployingQueueCronjob()
	go s.ProcessIngressRulesRequestCronjob()
}

// Move `pending` applications to `image generation queue` for building docker image
func (s *Server) MovePendingApplicationsToImageGenerationQueueCronjob() {
	for {
		// Get all pending applications
		var applications []Application
		tx := s.DB_CLIENT.Where("status = ?", ApplicationStatusPending).Find(&applications)
		if tx.Error != nil {
			log.Println(tx.Error)
			time.Sleep(5 * time.Second)
			continue
		}
		// Move them to image generation queue
		for _, application := range applications {
			log.Println("Moving application to image generation queue: ", application.ServiceName)
			err := s.DB_CLIENT.Transaction(func(tx *gorm.DB) error {
				// Update status
				application.Status = ApplicationStatusBuildingImageQueued
				tx2 := tx.Save(&application)
				if tx2.Error != nil {
					log.Println(tx2.Error)
					return tx2.Error
				}
				// Create log record
				logRecord := ApplicationBuildLog{
					ID:            uuid.New().String(),
					ApplicationID: application.ID,
					Logs:          "Queued for image generation",
					Time:          time.Now(),
				}
				tx3 := tx.Create(&logRecord)
				if tx3.Error != nil {
					log.Println(tx3.Error)
					return tx3.Error
				}
				// Enqueue
				err := s.AddServiceToDockerImageGenerationQueue(application.ServiceName, logRecord.ID)
				if err != nil {
					log.Println(err)
					return err
				}
				s.AddLogToApplicationBuildLog(logRecord.ID, "Successfully enqueued for image generation", "info")
				return nil
			})
			if err != nil {
				log.Println("Error while moving pending applications to image generation queue: ", err)
			}
		}
		time.Sleep(10 * time.Second)
	}
}

// Move `redeploy_pending` applications to `image generation queue` for building docker image
func (s *Server) MoveRedeployPendingApplicationsToImageGenerationQueueCronjob() {
	for {
		// Get all pending applications
		var applications []Application
		tx := s.DB_CLIENT.Where("status = ?", ApplicationStatusRedeployPending).Find(&applications)
		if tx.Error != nil {
			log.Println(tx.Error)
			time.Sleep(5 * time.Second)
			continue
		}
		// Move them to image generation queue
		for _, application := range applications {
			log.Println("Moving application to image generation queue: ", application.ServiceName)
			err := s.DB_CLIENT.Transaction(func(tx *gorm.DB) error {
				// Update status
				application.Status = ApplicationStatusBuildingImageQueued
				tx2 := tx.Save(&application)
				if tx2.Error != nil {
					log.Println(tx2.Error)
					return tx2.Error
				}
				// Create log record
				logRecord := ApplicationBuildLog{
					ID:            uuid.New().String(),
					ApplicationID: application.ID,
					Logs:          "Queued for image generation",
					Time:          time.Now(),
				}
				tx3 := tx.Create(&logRecord)
				if tx3.Error != nil {
					log.Println(tx3.Error)
					return tx3.Error
				}
				// Enqueue
				err := s.AddServiceToDockerImageGenerationQueue(application.ServiceName, logRecord.ID)
				if err != nil {
					log.Println(err)
					return err
				}
				s.AddLogToApplicationBuildLog(logRecord.ID, "Successfully enqueued for image generation", "info")
				return nil
			})
			if err != nil {
				log.Println("Error while moving pending applications to image generation queue: ", err)
			}
		}
		time.Sleep(10 * time.Second)
	}
}

// Move `deploying_pending` applications to `deploying queue` for deploying the application
func (s *Server) MoveDeployingPendingApplicationsToDeployingQueueCronjob() {
	for {
		// Get all pending applications
		var applications []Application
		tx := s.DB_CLIENT.Where("status = ?", ApplicationStatusDeployingPending).Find(&applications)
		if tx.Error != nil {
			log.Println(tx.Error)
			time.Sleep(5 * time.Second)
			continue
		}
		// Move them to image generation queue
		for _, application := range applications {
			log.Println("Moving application to deploying-queue: ", application.ServiceName)
			err := s.DB_CLIENT.Transaction(func(tx *gorm.DB) error {
				// Update status
				application.Status = ApplicationStatusDeployingQueued
				tx2 := tx.Save(&application)
				if tx2.Error != nil {
					log.Println(tx2.Error)
					return tx2.Error
				}

				// Enqueue
				err := s.AddServiceToDeployQueue(application.ServiceName)
				if err != nil {
					log.Println(err)
					return err
				}
				return nil
			})
			if err != nil {
				log.Println("Error while moving deploying_pending applications to deploying queue: ", err)
			}
		}
		time.Sleep(10 * time.Second)
	}
}

// Process ingress rules request - `pending` , `delete_pending` status records
func (s *Server) ProcessIngressRulesRequestCronjob() {
	for {
		var ingressRules []IngressRule
		tx := s.DB_CLIENT.Where("status = ? OR status = ?", IngressRuleStatusPending, IngressRuleStatusDeletePending).Find(&ingressRules)
		if tx.Error != nil {
			log.Println(tx.Error)
			time.Sleep(5 * time.Second)
			continue
		}
		for _, ingressRule := range ingressRules {
			transaction_id, err := s.HAPROXY_MANAGER.FetchNewTransactionId()
			if err != nil {
				log.Println(err)
				continue
			}
			if ingressRule.Status == IngressRuleStatusPending {
				// add backend
				backend_name := s.HAPROXY_MANAGER.GenerateBackendName(ingressRule.ServiceName, int(ingressRule.ServicePort))
				// skip if backend already exists - check db for service name and port and status != pending and id != ingressRule.ID
				backendDoesNotExist := false
				var ingressRuleCheck IngressRule
				tx := s.DB_CLIENT.Where("id != ? AND service_name = ? AND service_port = ? AND status != ?", ingressRule.ID, ingressRule.ServiceName, ingressRule.ServicePort, IngressRuleStatusPending).First(&ingressRuleCheck)
				if tx.Error != nil {
					if tx.Error == gorm.ErrRecordNotFound {
						backendDoesNotExist = true
					}
				}
				// if backend does not exist, create it
				if backendDoesNotExist {
					_, err := s.HAPROXY_MANAGER.AddBackend(transaction_id, ingressRule.ServiceName, int(ingressRule.ServicePort), 1)
					if err != nil {
						s.HAPROXY_MANAGER.DeleteTransaction(transaction_id)
						log.Println(err)
						continue
					}
				}
				// create backend switch rule
				if ingressRule.Protocol == HTTPSProtcol {
					err = s.HAPROXY_MANAGER.AddHTTPSLink(transaction_id, backend_name, ingressRule.DomainName)
					if err != nil {
						s.HAPROXY_MANAGER.DeleteTransaction(transaction_id)
						log.Println(err)
						continue
					}
				} else if ingressRule.Protocol == HTTPProtcol && ingressRule.Port == 80 {
					err = s.HAPROXY_MANAGER.AddHTTPLink(transaction_id, backend_name, ingressRule.DomainName)
					if err != nil {
						s.HAPROXY_MANAGER.DeleteTransaction(transaction_id)
						log.Println(err)
						continue
					}
				} else {
					var listenerMode HAPROXY_MANAGER.ListenerMode
					if ingressRule.Protocol == TCPProtcol {
						listenerMode = HAPROXY_MANAGER.TCPMode
					} else {
						listenerMode = HAPROXY_MANAGER.HTTPMode
					}
					err = s.HAPROXY_MANAGER.AddTCPLink(transaction_id, backend_name, int(ingressRule.Port), ingressRule.DomainName, listenerMode, s.RESTRICTED_PORTS)
					if err != nil {
						s.HAPROXY_MANAGER.DeleteTransaction(transaction_id)
						log.Println(err)
						continue
					}
				}
				// commit transaction
				err = s.HAPROXY_MANAGER.CommitTransaction(transaction_id)
				if err != nil {
					s.HAPROXY_MANAGER.DeleteTransaction(transaction_id)
					log.Println(err)
					continue
				}
				// update status
				ingressRule.Status = IngressRuleStatusApplied
				tx2 := s.DB_CLIENT.Save(&ingressRule)
				if tx2.Error != nil {
					log.Println(tx2.Error)
				}
			} else if ingressRule.Status == IngressRuleStatusDeletePending {
				backend_name := s.HAPROXY_MANAGER.GenerateBackendName(ingressRule.ServiceName, int(ingressRule.ServicePort))
				// delete backend switch rule
				if ingressRule.Protocol == HTTPSProtcol {
					err = s.HAPROXY_MANAGER.DeleteHTTPSLink(transaction_id, backend_name, ingressRule.DomainName)
					if err != nil {
						s.HAPROXY_MANAGER.DeleteTransaction(transaction_id)
						log.Println(err)
						continue
					}
				} else if ingressRule.Protocol == HTTPProtcol && ingressRule.Port == 80 {
					err = s.HAPROXY_MANAGER.DeleteHTTPLink(transaction_id, backend_name, ingressRule.DomainName)
					if err != nil {
						s.HAPROXY_MANAGER.DeleteTransaction(transaction_id)
						log.Println(err)
						continue
					}
				} else {
					err = s.HAPROXY_MANAGER.DeleteTCPLink(transaction_id, backend_name, int(ingressRule.Port), ingressRule.DomainName, s.RESTRICTED_PORTS)
					if err != nil {
						s.HAPROXY_MANAGER.DeleteTransaction(transaction_id)
						log.Println(err)
						continue
					}
				}
				// ensure this backend is not used by any other ingress rules
				// check by service name and port and status != delete pending
				backendUsedByOther := true
				var ingressRuleCheck IngressRule
				tx := s.DB_CLIENT.Where("id != ? AND service_name = ? AND service_port = ? AND status != ?", ingressRule.ID, ingressRule.ServiceName, ingressRule.ServicePort, IngressRuleStatusDeletePending).First(&ingressRuleCheck)
				if tx.Error != nil {
					if tx.Error == gorm.ErrRecordNotFound {
						backendUsedByOther = false
					}
				}
				if !backendUsedByOther {
					// delete backend
					err = s.HAPROXY_MANAGER.DeleteBackend(transaction_id, backend_name)
					if err != nil {
						s.HAPROXY_MANAGER.DeleteTransaction(transaction_id)
						log.Println(err)
						continue
					}
				}
				// commit transaction
				err = s.HAPROXY_MANAGER.CommitTransaction(transaction_id)
				if err != nil {
					s.HAPROXY_MANAGER.DeleteTransaction(transaction_id)
					log.Println(err)
					continue
				}
				// delete ingress rule
				tx2 := s.DB_CLIENT.Delete(&ingressRule)
				if tx2.Error != nil {
					log.Println(tx2.Error)					
				}
			}
		}
		time.Sleep(10 * time.Second)
	}
}