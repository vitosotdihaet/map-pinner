const currentRegionId = 1

class MarkerFetch {
    static async getAllType(type) {
        return getData(`/api/markers/${type}?regionId=${currentRegionId}`)
    }

    static async getAll() {
        return getData(`/api/markers?regionId=${currentRegionId}`)
    }


    static async createType(type, value) {
        return postData(`/api/markers/${type}?regionId=${currentRegionId}`, value.JSONify())
    }

    static async create(marker) {
        return postData(`/api/markers/${marker.type}?regionId=${currentRegionId}`, marker.value.JSONify())
    }


    static async deleteType(type, value) {
        return deleteData(`/api/markers/${type}/${value.id}`, "")
    }

    static async delete(marker) {
        return deleteData(`/api/markers/${marker.type}/${marker.id}`, "")
    }


    static async updateType(type, updateInfo) {
        return putData(`/api/markers/${type}/${updateInfo.id}`, JSON.stringify(updateInfo))
    }
}


// all the markerable types
const MarkerableTypes = {
    Point: 'point',
    Polygon: 'polygon',
    Line: 'line',
}

/**
 * Each markerable entity class should have these functions:
 * - constructor(data) - construct the class, data is a JS Object
 * - JSONify() - return a JSON string of the entity
 * - pullUpdate() - read update data (from somewhere) and apply it
 * - updateId(newId) - set entity's id with a new value
 * - draw() - add the entity to the global map
 * - hide() - hide the entity from the global map
 */
const MarkerableClasses = new Map([
    [MarkerableTypes.Point, Point],
    [MarkerableTypes.Polygon, Polygon],
    [MarkerableTypes.Line, Line],
])


// each Marker.value is an instance of a class that represents a markerable entity
class Marker {
    // list of shown markers 
    static shown = new Map(Object.values(MarkerableTypes).map(type => [type, new Map()]))

    // static async drawType(type) {
    //     const markerablesData = await MarkerFetch.getAllType(type)
    //     if (markerablesData === null) { return }
    //     markerablesData.forEach(markerableData => {
    //         const markerable = new (MarkerableClasses.get(type))(markerableData)
    //         let marker = new Marker(type, markerable)
    //         marker.draw()
    //     });
    // }

    static async drawAll() {
        const markerablesData = await MarkerFetch.getAll()
        if (markerablesData === null) { return }

        // every first value is a type, every second is the JSON of a markerable
        for (let i = 0; i < markerablesData.length; i += 2) {
            const type = markerablesData[i]
            const data = markerablesData[i + 1]
            const markerable = new (MarkerableClasses.get(type))(data)
            let marker = new Marker(type, markerable)
            marker.draw()
        }
    }

    static hideType(type) {
        for (let [_, marker] of Marker.shown.get(type)) {
            marker.value.hide()
        }
        Marker.shown.set(type, new Map())
    }

    static hideAll() {
        Object.values(MarkerableTypes).forEach(type => {
            Marker.hideType(type)
        })
    }



    constructor(type, value) {
        this.type = type
        this.value = value
        this.id = this.value.id
    }

    pullUpdate() {
        this.value.pullUpdate()
    }

    updateId(newId) {
        this.id = newId
        this.value.updateId(this.id)
    }

    delete() {
        this.hide()
        MarkerFetch.delete(this)
    }

    draw() {
        if (Marker.shown.get(this.type).has(this.id)) return
        this.value.draw()
        Marker.shown.get(this.type).set(this.id, this)
    }

    hide() {
        Marker.shown.delete(this.id)
        this.value.hide()
    }
}