package Init

import "github.com/pritamdas99/Movie_Server/Data"

type userDB map[string]Data.User

var UserList userDB

func init() {
	UserList = make(userDB)
	UserList["prishan076@gmail.com"] = Data.User{
		Id:       "prishan076@gmail.com",
		Password: "1111",
	}
	UserList["pd17021999@gmail.com"] = Data.User{
		Id:       "pd17021999@gmail.com",
		Password: "1111",
	}
	UserList["pritam@appscode.com"] = Data.User{
		Id:       "pritam@appscode.com",
		Password: "1111",
	}

}
