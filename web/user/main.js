class User {
    static currentUser = null
    
    constructor(data) {
        this.id = data.id
        this.name = data.name
    }
}

User.currentUser = new User(JSON.parse(localStorage.getItem("user")))

const userNameLabel = document.getElementById('userName')
userNameLabel.innerHTML = User.currentUser.name