class RoleFetch {
    static async getAll() {
        return getJSON(`/api/roles`)
    }

    static async isOwner() {
        if ((await getFetch('/api/roles/is-admin')).ok) {
            return true
        }
        return false
    }
}