package ModelUser

type (

	// for response list user
	ShowUser struct {
		UserID			int
		UserName		string
		UserFullname	string
		UserEmail		string
		UserPassword	string
		DateCreated		string
	}

	// for param request by user
	ReqUser struct {
		UserID			int
		UserName		string
		UserFullname	string
		UserEmail		string
		UserPassword	string
		DateCreated		string
	}
)