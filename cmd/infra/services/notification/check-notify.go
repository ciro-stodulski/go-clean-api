package notificationservice

import domainexceptions "go-clean-api/cmd/domain/exceptions"

func (ns notificationService) CheckNotify(msg string) (*domainexceptions.ApplicationException, error) {

	err := ns.NotificationProto.Verify(msg)

	return nil, err
}
