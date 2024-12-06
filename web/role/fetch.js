class RoleFetch {
    static async getAll() {
        return getJSON(`/api/roles`)
    }
}