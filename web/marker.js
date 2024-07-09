class PointFetch {
    static async getAll() {
        return getData('/api/points')
    }

    static async create(point) {
        return postData("/api/points", JSON.stringify(point))
    }

    static async delete(point) {
        return deleteData(`/api/points/${point.id}`, "")
    }

    static async update(point) {
        return putData(`/api/points/${point.id}`, JSON.stringify(point))
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

class Marker {
    constructor(point) {
        this.point = point
        this.setupMapMarker()
    }

    setupMapMarker() {
        this.mapMarker = L.marker([this.point.latitude, this.point.longitude])

        let point = this.point
        this.setDraggable(true)
        this.mapMarker.on('drag', function (event) {
            // for some reason this shit works, though it shouldn't (should be event.originalEvent.button)
            if (event.originalEvent.buttons == 1) return
            event.target.dragging.disable()
            event.target.setLatLng([point.latitude, point.longitude])
            setTimeout(() => event.target.dragging.enable());
        })

        let marker = this
        this.mapMarker.on('dragend', function (event) {
            let latlng = event.target.getLatLng()
            marker.update({
                latitude: latlng.lat,
                longitude: latlng.lng
            })
        })

        this.mapMarker.bindPopup(
            L.popup().setContent(
                `
                <div class="popup">
                    ID: ${this.point.id}<br/>
                    Name: <input type="text" class="popupNameInput" id="${this.point.id}" maxlength="255" size="10" value="${this.point.name}"/><br/>
                    Latitude: ${this.point.latitude.toFixed(4)}<br/>
                    Longitude: ${this.point.longitude.toFixed(4)}
                </div>
                <button class="popupDeleteButton" onclick="deleteMarker(shownMarkers, ${this.point.id})">Delete</button>
                <button class="popupUpdateButton" onclick="updateMarker(shownMarkers, ${this.point.id})">Update</button>
                `
            )
        ).openPopup()
    }

    setDraggable(value) {
        this.mapMarker.options.draggable = value
    }

    draw() {
        if (shownMarkers.has(this.point.id)) return
        this.mapMarker.addTo(map)
        shownMarkers.push(this)
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
        this.setupMapMarker()
        this.draw()

        PointFetch.update(updateInfo)
    }

    updateId(newId) {
        this.point.id = newId
        this.setupMapMarker()
    }

    static async addMarkerOnMapClick(event) {
        // don't add new marker if not left mouse button is pressed
        if (event.originalEvent.button != 0) return

        let latlng = event.latlng

        let point = new Point('', 0, latlng.lat, latlng.lng)
        let marker = new Marker(point)

        let newId = await PointFetch.create(point)

        marker.updateId(newId.id)
        marker.draw()

        return marker
    }

    async delete() {
        PointFetch.delete(this.point)
        this.hide()
    }

    hide() {
        map.removeLayer(this.mapMarker)
        const index = shownMarkers.indexOf(this)
        if (index > -1) {
            shownMarkers.splice(index, 1)
        }
    }
}

// other marker related functions
function pointsToMarkers(points) {
    if (points == null || points.length == 0) { return [] }

    let markers = []
    points.forEach(point => {
        markers.push(new Marker(point))
    })

    return markers
}


function drawMarkers(markers) {
    markers.forEach(marker => {
        marker.draw()
    })
}

function hideMarkers(markers) {
    markers.forEach(marker => {
        marker.hide()
    })
}


function deleteMarker(markers, id) {
    markers.forEach(marker => {
        if (marker.point.id == id) {
            marker.delete()
            return
        }
    })
}

function updateMarker(markers, id) {
    let name = document.getElementsByClassName('popupNameInput')[0].value

    markers.forEach(marker => {
        if (marker.point.id == id) {
            let latlng = marker.mapMarker.getLatLng()
            marker.update({
                name: name,
                latitude: latlng.lat,
                longitude: latlng.lng
            })
            return
        }
    })
}


async function drawAllMarkers() {
    drawMarkers(pointsToMarkers(await PointFetch.getAll()))
}



let shownMarkers = []
shownMarkers.has = function (id) {
    for (let i = 0; i < this.length; i++) {
        if (this[i].point.id == id) return true
    }
    return false
}
