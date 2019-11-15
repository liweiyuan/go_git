package user_entity

type Entity interface{
   Hello() string
}

type UserEntity struct{
     UserName string
}

func (user UserEntity) Hello() string{
    return user.UserName
}