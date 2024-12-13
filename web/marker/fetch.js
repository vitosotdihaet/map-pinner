class MarkerFetch {
    static async getAllType(type) {
        return getJSON(`/api/markers/${type}?region_id=${Region.currentRegion.id}`)
    }

    static async getAll() {
        return getJSON(`/api/markers?region_id=${Region.currentRegion.id}`)
    }


    static async createType(type, value) {
        return postJSON(`/api/markers/${type}?region_id=${Region.currentRegion.id}`, value.JSONify())
    }

    static async create(marker) {
        return postJSON(`/api/markers/${marker.type}?region_id=${Region.currentRegion.id}`, marker.value.JSONify())
    }


    static async deleteType(type, value) {
        return deleteJSON(`/api/markers/${type}/${value.id}`, "")
    }

    static async delete(marker) {
        return deleteJSON(`/api/markers/${marker.type}/${marker.id}`, "")
    }


    static async updateType(type, updateInfo) {
        return putFetch(`/api/markers/${type}/${updateInfo.id}`, JSON.stringify(updateInfo))
    }
}
