class RegionFetch {
    static async getAll() {
        return getJSON(`/api/regions?group_id=${Group.currentGroup.id}`)
    }

    static async create(region) {
        return postJSON(`/api/regions?group_id=${Group.currentGroup.id}`, JSON.stringify(region))
    }

    static async delete(region) {
        // return deleteJSON(`/api/groups/${region.id}`, '')
    }

    static async update(region) {
        // return putJSON(`/api/groups/${region.id}`, JSON.stringify(region))
    }
}