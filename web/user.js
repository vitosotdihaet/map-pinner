class UserFetch {
    static async getAll() {
        return getJSON('/users')
    }

    static async createWithNamePassword(username, password) {
        return postJSON(`/users?username=${username}&password=${password}`)
    }

    static async delete(user) {
        return deleteJSON(`/users/${user.id}`)
    }

    static async update(user) {
        return putJSON(`/users/${user.id}`, JSON.stringify(user))
    }

    static async getByUsernamePassword(username, password) {
        return getJSON(`/users/bynamepassword?username=${username}&password=${password}`);
    }

    static async validateToken() {
        return postFetch('/users/validate-token')
    }
}

class User {
    constructor(id, name, hashedPassword) {
        this.id = id
        this.name = name
        this.hashedPassword = hashedPassword
    }
}