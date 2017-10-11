package rocketchat

type apiMethod string

const (
	kapiUsersCreate      apiMethod = "/api/v1/users.create"
	kapiUsersCreateToken           = "/api/v1/users.createToken"
	kapiUsersSetAvatar             = "/api/v1/users.setAvatar"
	kapiUsersUpdate                = "/api/v1/users.update"
	kapiUsersLogin                 = "/api/v1/login"

	kapiChannelPostMessage apiMethod = "/api/v1/chat.postMessage"
)
