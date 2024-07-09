// navbar
function openTab(evt, tabName) {
    var i, tabcontent, tablinks;

    tabcontent = document.getElementsByClassName("tabcontent");
    for (i = 0; i < tabcontent.length; i++) {
        tabcontent[i].style.display = "none";
    }

    tablinks = document.getElementsByClassName("tablinks");
    for (i = 0; i < tablinks.length; i++) {
        tablinks[i].className = tablinks[i].className.replace(" active", "");
    }

    document.getElementById(tabName).style.display = "block";
    evt.currentTarget.className += " active";
}

document.getElementsByClassName("tablinks")[0].click();



// requests
async function getData(url) {
    const response = await fetch(url)
    return await response.json()
}

async function postData(url, body) {
    const response = await fetch(url, {
        method: "POST",
        body: body,
    });
    return await response.json()
}

async function deleteData(url, body) {
    const response = await fetch(url, {
        method: "DELETE",
        body: body,
    });
    return await response.json()
}

async function putData(url, body) {
    const response = await fetch(url, {
        method: "PUT",
        body: body,
    });
    return await response.json()
}


var map = L.map('map').setView([55.76, 37.64], 5)

L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy OpenStreetMap contributors',
    // detectRetina: true
}).addTo(map)


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
                <button class="popupDeleteMarkerButton" onclick="deleteMarker(shownMarkers, ${this.point.id})">Delete</button>
                <button class="popupUpdateMarkerButton" onclick="updateMarker(shownMarkers, ${this.point.id})">Update</button>
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
    for (let i = 0; i < markers.length; i++) {
        markers[i].hide()
    }
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
    console.log(name)

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


async function drawAllPoints() {
    drawMarkers(pointsToMarkers(await PointFetch.getAll()))
}


let shownMarkers = []
shownMarkers.has = function (id) {
    for (let i = 0; i < this.length; i++) {
        if (this[i].point.id == id) return true
    }

    return false
}

map.on('click', Marker.addMarkerOnMapClick);
document.getElementById('showAllPoints').addEventListener('click', function(event) {
    event.preventDefault()
    drawAllPoints()
})
document.getElementById('hideAllPoints').addEventListener('click', function(event) {
    event.preventDefault()
    hideMarkers([...shownMarkers])
})






class PolygonFetch {
    static async getAll() {
        return await getData('/api/polygons')
    }
}


function polygonsToShapes(polygons) {
    if (polygons == null || polygons.length == 0) { return }

    let shapes = []
    polygons.forEach(polygon => {
        polygonPoints = polygon.Points

        let coordinates = []
        polygonPoints.forEach(point => {
            if (point != null) {
                coordinates.push(new L.LatLng(point.latitude, point.longitude))
            }
        })

        shapes.push(new L.Geodesic(coordinates, {
            color: randomColor({ "luminosity": "bright", "hue": "blue" }),
            weight: 3,
            opacity: 0.75,
            fillOpacity: 0.5
        }))
    })

    return shapes
}

function drawShapes(shapes) {
    shapes.forEach(shape => {
        shownShapes.push(shape)
        shape.addTo(map)
    })
}

function hideShapes(shapes) {
    let shownShapesCopy = [...shownShapes]

    for (let i = 0; i < shapes.length; i++) {
        const shape = shapes[i]
        map.removeLayer(shape)

        const index = shownShapesCopy.indexOf(shape)
        if (index > -1) {
            shownShapesCopy.splice(index, 1)
        }
    }

    shownShapes = shownShapesCopy
}

async function drawAllPolygons() {
    drawShapes(polygonsToShapes(await PolygonFetch.getAll()));
}

let shownShapes = []


document.getElementById('showAllPolygons').addEventListener('click', function(event) {
    event.preventDefault()
    drawAllPolygons()
})

document.getElementById('hideAllPolygons').addEventListener('click', function(event) {
    event.preventDefault()
    hideShapes(shownShapes)
})




