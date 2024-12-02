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
        return getFetch(`/users/bynamepassword?username=${username}&password=${password}`);
    }

    static async getCurrent() {
        return getFetch('/users/current-user')
    }
}
