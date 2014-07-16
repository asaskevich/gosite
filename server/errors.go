package server

// Return string equvalent for error ID
func GetMessage(id int) string {
	switch id{
	case NOT_AN_EMAIL: return "NOT_AN_EMAIL: It's not an e-mail! I'm sure in it."
	case CANT_WRITE_TO_DB:return "CANT_WRITE_TO_DB: You want to write to DB? Sorry, but not now."
	case DB_UNAVAILABLE:return "DB_UNAVAILABLE: DB is busy. Very-very."
	case INVALID_NAME: return "INVALID_NAME: Hey, I think, that your name very short or long!"
	case INVALID_FEEDBACK: return "INVALID_FEEDBACK: Feedback should be long as whole story. Write more, please."
	case OK: return "OK: Hmm, really OK!"
	default: return "What do you do?"
	}
}
