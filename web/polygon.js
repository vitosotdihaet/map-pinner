class PolygonFetch {
    static async getAll() {
        return getData('/api/markers/polygons')
    }

    static async create(polygon) {
        return postData("/api/markers/polygons", JSON.stringify(polygon))
    }

    static async delete(polygon) {
        return deleteData(`/api/markers/polygons/${polygon.id}`, "")
    }

    static async update(polygon) {
        return putData(`/api/markers/polygons/${polygon.id}`, JSON.stringify(polygon))
    }
}

class Polygon {
    constructor(name, id, points) {
        this.name = name
        this.id = id
        this.points = points
    }
}


class MapPolygon {
    constructor(polygon) {
        this.polygon = polygon
        this.color = randomColor({ "luminosity": "bright", "hue": "blue" })
        this.setupMapPolygon()
    }

    onPointDrag() {
        let newLatLngs = [];
        let newPolygonPoints = []

        for (let marker of this.mapMarkers) {
            let latlng = marker.getLatLng()
            newLatLngs.push(latlng);
            newPolygonPoints.push(new Point("", 0, latlng.lat, latlng.lng))
        }

        newLatLngs.push(newLatLngs[0])
        newPolygonPoints.push(newPolygonPoints[0])

        this.mapPolygon.setLatLngs(newLatLngs);
        this.polygon.points = newPolygonPoints
    }

    setupMapPolygon() {
        let coordinates = []
        let latlngs = []
        this.mapMarkers = []

        this.polygon.points.forEach(point => {
            coordinates.push([point.latitude, point.longitude])
            let latlng = new L.LatLng(point.latitude, point.longitude)
            latlngs.push(latlng)
        })
        latlngs.pop()

        this.mapPolygon = L.geodesic(coordinates, {
            color: this.color,
            weight: 5,
            opacity: 0.75,
            fillOpacity: 0.5
        })

        latlngs.forEach(place => {
            var marker = L.marker(place, { draggable: true, icon: altIcon });

            marker.on('drag', (_) => {
                this.onPointDrag();
            });
            marker.on('dragend', (_) => {
                this.checkAndUpdate()
            })
            this.mapMarkers.push(marker);
        })

        this.mapPolygon.bindPopup(
            L.popup().setContent(
                `
                <div class="popup">
                    ID: ${this.polygon.id}<br/>
                    Name: <input type="text" class="popupNameInput" id="${this.polygon.id}" maxlength="255" size="10" value="${this.polygon.name}"/><br/>
                </div>
                <button class="popupDeleteButton" onclick="shownMapPolygons.get(${this.polygon.id}).delete()">Delete</button>
                <button class="popupUpdateButton" onclick="shownMapPolygons.get(${this.polygon.id}).checkAndUpdate()">Update</button>
                `
            )
        ).openPopup()
    }

    draw() {
        if (shownMapPolygons.has(this.polygon.id)) return
        this.mapMarkers.forEach(marker => { marker.addTo(map) })
        this.mapPolygon.addTo(map)
        shownMapPolygons.set(this.polygon.id, this)
    }

    checkAndUpdate() {
        let name = undefined
        let nameInput = document.getElementsByClassName('popupNameInput')

        for (var i = 0; i < nameInput.length; i++) {
            let element = nameInput[i]
            if (element.id == this.polygon.id) {
                name = element.value
            }
        }

        let temp = [...this.polygon.points]
        temp.pop()

        this.update({
            name: name,
            points: temp
        })
    }

    update(updateInfo) {
        if (updateInfo.name !== undefined) {
            this.polygon.name = updateInfo.name
        }
        if (updateInfo.latitude !== undefined) {
            this.polygon.latitude = updateInfo.latitude
        }
        if (updateInfo.longitude !== undefined) {
            this.polygon.longitude = updateInfo.longitude
        }

        updateInfo.id = this.polygon.id

        this.hide()
        this.setupMapPolygon()
        this.draw()

        PolygonFetch.update(updateInfo)
    }

    updateId(newId) {
        this.polygon.id = newId
        this.setupMapPolygon()
    }

    async delete() {
        PolygonFetch.delete(this.polygon)
        this.hide()
    }

    hide() {
        map.removeLayer(this.mapPolygon)
        this.mapMarkers.forEach(marker => { map.removeLayer(marker) })
        shownMapPolygons.delete(this.polygon.id)
    }
}


function polygonsToMapPolygons(polygons) {
    if (polygons == null) { return [] }

    let mapPolygons = []
    polygons.forEach(polygon => {
        mapPolygons.push(new MapPolygon(polygon))
    })

    return mapPolygons
}

function drawMapPolygons(mapPolygons) {
    mapPolygons.forEach((mapPolygon, _) => {
        mapPolygon.draw()
    })
}

function hideMapPolygons(mapPolygons) {
    mapPolygons.forEach((mapPolygon, _) => {
        mapPolygon.hide()
    })
}


polygonAccumulatedPoints = []
polygonAccumulatedMarkers = []

// TODO move main logic to main
function startNewPolygon(event) {
    event.preventDefault()

    for (var tab of tabs) {
        tab.removeEventListener('click', openTab)
    }

    newPolygonButton.removeEventListener('click', startNewPolygon)
    newPolygonButton.addEventListener('click', stopPolygon)
    newPolygonButton.innerText = "Stop"

    map.off('click', MapPoint.addMapPointOnMapClick)
    map.on('click', newPolygonPointOnAMap)
}

function newPolygonPointOnAMap(event) {
    // don't add new point if not left mouse button is pressed
    if (event.originalEvent.button != 0) return
    let latlng = event.latlng

    let marker = L.marker(latlng, { icon: altIcon });
    marker.addTo(map)
    polygonAccumulatedMarkers.push(marker)
    let point = new Point('', 0, latlng.lat, latlng.lng)

    polygonAccumulatedPoints.push(point)
}

async function stopPolygon(event) {
    event.preventDefault()

    for (var tab of tabs) {
        tab.addEventListener('click', openTab)
    }

    newPolygonButton.removeEventListener('click', stopPolygon)
    newPolygonButton.addEventListener('click', startNewPolygon)
    newPolygonButton.innerText = "Start a new polygon"

    map.off('click', newPolygonPointOnAMap)
    map.on('click', MapPoint.addMarkerOnMapClick)

    if (polygonAccumulatedPoints.length < 3) {
        polygonAccumulatedMarkers.forEach(marker => {
            map.removeLayer(marker)
        })
        polygonAccumulatedPoints = []
        polygonAccumulatedMarkers = []
        return
    }

    polygonAccumulatedPoints.push(polygonAccumulatedPoints[0])
    let polygon = new Polygon("", 0, polygonAccumulatedPoints)
    let mapPolygon = new MapPolygon(polygon)

    newId = await PolygonFetch.create(polygon)
    mapPolygon.updateId(newId.id)
    mapPolygon.draw()
    
    polygonAccumulatedMarkers.forEach(marker => {
        map.removeLayer(marker)
    })
    polygonAccumulatedPoints = []
    polygonAccumulatedMarkers = []
}


async function drawAllPolygons() {
    drawMapPolygons(polygonsToMapPolygons(await PolygonFetch.getAll()));
}



let shownMapPolygons = new Map()
