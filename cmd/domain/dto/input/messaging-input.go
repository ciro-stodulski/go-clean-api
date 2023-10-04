package inputdto

type MessagingInput struct {
	To        string "json:\"to\""
	Subject   string "json:\"subject\""
	SubjectId string "json:\"subject_id\""
	Body      string "json:\"body\""
}
