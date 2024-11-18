class PointFetch {
    static async getAll() {
        return getData('/api/markers/points')
    }

    static async create(point) {
        return postData("/api/markers/points", JSON.stringify(point))
    }

    static async delete(point) {
        return deleteData(`/api/markers/points/${point.id}`, "")
    }

    static async update(point) {
        return putData(`/api/markers/points/${point.id}`, JSON.stringify(point))
    }
}

class Point {
    constructor(name, id, latitude, longitude) {
        this.name = name
        this.id = id
        this.latitude = latitude
        this.longitude = longitude
    }
}

class MapPoint {
    constructor(point) {
        this.point = point
        this.setupMapPoint()
    }

    setupMapPoint() {
        this.mapMarker = L.marker([this.point.latitude, this.point.longitude], { draggable: true })

        let point = this.point
        this.mapMarker.on('drag', function (event) {
            // for some reason this shit works, though it shouldn't (should be event.originalEvent.button)
            if (event.originalEvent.buttons == 1) return
            event.target.dragging.disable()
            event.target.setLatLng([point.latitude, point.longitude])
            setTimeout(() => event.target.dragging.enable());
        })

        let mapPoint = this
        this.mapMarker.on('dragend', function (event) {
            let latlng = event.target.getLatLng()
            mapPoint.update({
                latitude: latlng.lat,
                longitude: latlng.lng
            })
        })

        this.mapMarker.bindPopup(
            L.popup().setContent(
                `
                <div class="popup">
                    Name: <input type="text" class="popupNameInput" id="${this.point.id}" maxlength="255" size="10" value="${this.point.name}"/><br/>
                    Latitude: ${this.point.latitude.toFixed(4)}<br/>
                    Longitude: ${this.point.longitude.toFixed(4)}
                </div>
                <button class="popupDeleteButton" onclick="shownMapPoints.get(${this.point.id}).delete()">Delete</button>
                <button class="popupUpdateButton" onclick="shownMapPoints.get(${this.point.id}).checkAndUpdate()">Update</button>
                `
            )
        ).openPopup()
    }

    draw() {
        if (shownMapPoints.has(this.point.id)) return
        this.mapMarker.addTo(map)
        shownMapPoints.set(this.point.id, this)
    }

    checkAndUpdate() {
        let name = document.getElementsByClassName('popupNameInput')[0].value

        let latlng = this.mapMarker.getLatLng()
        this.update({
            name: name,
            latitude: latlng.lat,
            longitude: latlng.lng
        })
    }

    update(updateInfo) {
        if (updateInfo.name !== undefined) {
            this.point.name = updateInfo.name
        }
        if (updateInfo.latitude !== undefined) {
            this.point.latitude = updateInfo.latitude
        }
        if (updateInfo.longitude !== undefined) {
            this.point.longitude = updateInfo.longitude
        }

        updateInfo.id = this.point.id

        this.hide()
        this.setupMapPoint()
        this.draw()

        PointFetch.update(updateInfo)
    }

    updateId(newId) {
        this.point.id = newId
        this.setupMapPoint()
    }

    static async addMapPointOnMapClick(event) {
        // don't add new mapPoint if not left mouse button is pressed
        if (event.originalEvent.button != 0) return

        let latlng = event.latlng

        let point = new Point('', 0, latlng.lat, latlng.lng)
        let mapPoint = new MapPoint(point)

        let newId = await PointFetch.create(point)

        mapPoint.updateId(newId.id)
        mapPoint.draw()

        return mapPoint
    }

    async delete() {
        PointFetch.delete(this.point)
        this.hide()
    }

    hide() {
        map.removeLayer(this.mapMarker)
        shownMapPoints.delete(this.point.id)
    }
}

// other mapPoint related functions
function pointsToMapPoints(points) {
    if (points == null) { return [] }

    let mapPoints = []
    points.forEach(point => {
        mapPoints.push(new MapPoint(point))
    })

    return mapPoints
}


function drawMapPoints(mapPoints) {
    mapPoints.forEach(mapPoint => {
        mapPoint.draw()
    })
}

function hideMapPoints(mapPoints) {
    mapPoints.forEach((mapPoint, _) => {
        mapPoint.hide()
    })
}


async function drawAllPoints() {
    drawMapPoints(pointsToMapPoints(await PointFetch.getAll()))
}


let shownMapPoints = new Map()
