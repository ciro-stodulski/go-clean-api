package messagingentity

import "time"

type MessagingEntity struct {
	To        string    "json:\"to\""
	Subject   string    "json:\"subject\""
	SubjectId string    "json:\"subject_id\""
	Body      string    "json:\"body\""
	Date      time.Time "json:\"date\""
}
