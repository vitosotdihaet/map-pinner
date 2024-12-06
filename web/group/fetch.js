class GroupFetch {
    static async getAll() {
        return getJSON(`/api/groups`)
    }

    static async getById(groupId) {
        return getJSON(`/api/groups/${groupId}`)
    }

    static async create(group) {
        return postJSON(`/api/groups`, JSON.stringify(group))
    }

    static async delete(group) {
        return deleteJSON(`/api/groups/${group.id}`, '')
    }

    // static async update(group) {
    //     return putJSON(`/api/groups/${group.id}`, JSON.stringify(group))
    // }

    static async addUserToGroup(groupId, userName, roleId) {
        return postFetch(`/api/groups/${groupId}/${userName}/${roleId}`)
    }
}