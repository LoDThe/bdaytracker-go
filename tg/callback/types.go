package callback

func Init() {
	addCallback(None{})
	addCallback(OpenMenu{})
	addCallback(AddFriend{})
	addCallback(AddFromVK{})
	addCallback(FriendsList{})
	addCallback(Settings{})
	addCallback(RemoveFriend{})
	addCallback(RemoveFriendApprove{})
	addCallback(RemoveFriendCancel{})
}

type None struct {
}

type OpenMenu struct {
	Edit bool // If it's true, callback message has to be edited. Otherwise, a new message is sent.
}

type AddFriend struct {
}

type AddFromVK struct {
}

type FriendsList struct {
	Offset int // How many friends should be skipped
}

type Settings struct {
}

type RemoveFriend struct {
}

type RemoveFriendApprove struct {
	UUID string
	Name string
}

type RemoveFriendCancel struct {
	Name string
}
