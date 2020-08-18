package user

// Information is the user's information retrieved from the client
type Information struct {
	Name       string                 `json:"name"`
	ModuleData map[string]interface{} `json:"module_data"`
}

// // A Reminder is something the user asked for Olivia to remember
// type Reminder struct {
// 	Reason string `json:"reason"`
// 	Date   string `json:"date"`
// }

// userInformation is a map which is the cache for user information
var userInformation = map[string]Information{}

// UpdateUserInformation requires the token of the user and a function which gives the actual
// information and returns the new information.
func UpdateUserInformation(token string, predicat func(Information) Information) {
	if ui := userInformation[token]; ui.ModuleData == nil {
		ui.ModuleData = make(map[string]interface{})
		userInformation[token] = ui
	}
	userInformation[token] = predicat(userInformation[token])
}

// SetUserInformation sets the user's information by its token.
func SetUserInformation(token string, information Information) {
	userInformation[token] = information
}

// GetUserInformation returns the information of a user with his token
func GetUserInformation(token string) Information {
	return userInformation[token]
}
