package user

type Role string

type User struct {
    ID    string
    Name  string
    Roles []Role
}
