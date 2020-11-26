package database

type InputUser struct {
	TelegramID int32 `json:"telegram_id"`
	U2ID       int32 `json:"u2_id"`
}

type OutputUser struct {
	TelegramID       int32  `json:"telegram_id"`
	U2ID             int32  `json:"u2_id"`
	VerificationCode string `json:"verification_code"`
	Verified         bool   `json:"verified"`
}

func ToOutputUser(user *User) *OutputUser {
	if user == nil {
		return nil
	}

	outputUser := &OutputUser{
		TelegramID:       user.TelegramID,
		U2ID:             user.U2ID,
		VerificationCode: user.VerificationCode,
		Verified:         user.Verified,
	}
	return outputUser
}

func ToOutputUsers(users []*User) []*OutputUser {
	if len(users) == 0 {
		return nil
	}

	outputUsers := make([]*OutputUser, 0, len(users))
	for _, v := range users {
		outputUsers = append(outputUsers, ToOutputUser(v))
	}
	return outputUsers
}

func ToUser(inputUser *InputUser) *User {
	if inputUser == nil {
		return nil
	}

	user := &User{
		TelegramID:       inputUser.TelegramID,
		U2ID:             inputUser.U2ID,
		VerificationCode: "",
		Verified:         false,
	}
	return user
}

func ToUsers(inputUsers []*InputUser) []*User {
	if inputUsers == nil {
		return nil
	}

	users := make([]*User, 0, 100)
	for _, v := range inputUsers {
		users = append(users, ToUser(v))
	}
	return users
}
