class User {
    static currentUser = null
    
    constructor(data) {
        this.id = data.id
        this.name = data.name
    }
}


userInfo = localStorage.getItem('user')
if (userInfo != null) {
    // TODO: check if user is still in the db
    User.currentUser = new User(JSON.parse(userInfo))
}
