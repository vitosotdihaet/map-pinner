class RoleFetch {
    static async getAll() {
        return getJSON(`/api/roles/all`)
    }

    static async isOwner() {
        if ((await getFetch('/api/roles/is-admin')).ok) {
            return true
        }
        return false
    }

    static async getRole(group_id) {
        return getJSON(`/api/roles/${group_id}`)
    }
}