class UserFetch {
    static async getAll() {
        return getData('/api/users')
    }

    static async createWithNamePassword(username, password) {
        return postData(`/api/users?username=${username}&password=${password}`)
    }

    static async delete(user) {
        return deleteData(`/api/users/${user.id}`, '')
    }

    static async update(user) {
        return putData(`/api/users/${user.id}`, JSON.stringify(user))
    }

    static async getByUsernamePassword(username, password) {
        return getData(`/api/users/bynamepassword?username=${username}&password=${password}`);
    }
}

class User {
    constructor(id, name, hashedPassword) {
        this.id = id
        this.name = name
        this.hashedPassword = hashedPassword
    }
}