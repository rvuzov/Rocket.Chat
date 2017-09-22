package rocketchat

type apiMethod string

const (
	kapiUsersCreate      apiMethod = "/api/v1/users.create"
	kapiUsersCreateToken           = "/api/v1/users.createToken"
	kapiUsersSetAvatar             = "/api/v1/users.setAvatar"
	kapiUsersUpdate                = "/api/v1/users.update"
)
