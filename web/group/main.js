class Group {
    static loaded = []
    static currentGroup = null

    constructor(data) {
        this.id = data.id
        this.name = data.name
        this.users = data.users
    }
}

groupInfo = localStorage.getItem('group')
if (groupInfo) {
    // TODO: check if group is still in the db
    Group.currentGroup = new Group(JSON.parse(groupInfo))
}