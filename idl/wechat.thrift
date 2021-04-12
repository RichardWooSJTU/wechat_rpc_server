struct Group {
    1:string groupID,
    2:string groupName
}

service Wechat {
    string Send(1: required i32 option, 2: required string content),
    list<Group> FetchGroups()
}