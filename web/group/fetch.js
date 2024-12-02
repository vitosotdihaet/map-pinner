class GroupFetch {
    static async getAll() {
        return getJSON('/api/groups')
    }

    static async create(group) {
        return postJSON('/api/groups', JSON.stringify(group))
    }

    static async delete(group) {
        return deleteJSON(`/api/groups/${group.id}`, '')
    }

    static async update(group) {
        return putJSON(`/api/groups/${group.id}`, JSON.stringify(group))
    }
}