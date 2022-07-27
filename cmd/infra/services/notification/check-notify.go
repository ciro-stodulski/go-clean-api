package notificationservice

func (ns notificationService) CheckNotify(msg string) (string error) {

	err := ns.NotificationProto.Verify(msg)

	if err != nil {
		return err
	}

	return nil
}
