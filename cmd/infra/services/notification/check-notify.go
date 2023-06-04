package notificationservice

func (ns notificationService) CheckNotify(msg string) error {

	err := ns.NotificationProto.Verify(msg)

	return err
}
