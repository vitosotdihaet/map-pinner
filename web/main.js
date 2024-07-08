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
async function fetchData(url) {
    const response = await fetch(url)
    return await response.json()
}

async function postData(url, body) {
    const response = await fetch(url, {
        method: "POST",
        body: body,
        // headers: {
        //   "Content-type": "application/json; charset=UTF-8"
        // }
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
        return await fetchData('/api/points')
    }

    static async post(point) {
        return postData("/api/points", JSON.stringify(point))
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
    setPopup() {
        this.map_marker.bindPopup(
            L.popup().setContent(
                `<p class="popup">
                 ID: ${this.point.id}<br/>
                 Name: <input type="text" maxlength="255" size="10" value="${this.point.name}"/><br/>
                 Latitude: ${this.point.latitude.toFixed(4)}<br/>
                 Longitude: ${this.point.longitude.toFixed(4)}</p>`
            )
        ).openPopup()
    }

    constructor(point) {
        this.point = point

        this.map_marker = L.marker([point.latitude, point.longitude], { draggable: true })
        this.map_marker.on('drag', function (event) {
            // for some reason this shit work, though it shouldn't (should be event.originalEvent.button)
            if (event.originalEvent.buttons == 1) return
            event.target.dragging.disable()
            event.target.setLatLng([point.latitude, point.longitude])
            setTimeout(() => event.target.dragging.enable(), 0);
        })
        this.setPopup()
    }

    draw() {
        if (shownMarkers.has(this.point.id)) return
        this.map_marker.addTo(map)
        shownMarkers.push(this)
    }

    update(updateInfo) {
        if (updateInfo.name !== undefined) {
            this.point.name = updateInfo.name
        }
        if (updateInfo.id !== undefined) {
            this.point.id = updateInfo.id
        }
        if (updateInfo.latitude !== undefined) {
            this.point.latitude = updateInfo.latitude
        }
        if (updateInfo.longitude !== undefined) {
            this.point.longitude = updateInfo.longitude
        }

        this.setPopup()
    }
}

// other marker functions
function pointsToMarkers(points) {
    if (points == null || points.length == 0) { return }

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
    let shownMarkersCopy = [...shownMarkers]

    for (let i = 0; i < markers.length; i++) {
        const marker = markers[i]
        map.removeLayer(marker.map_marker)

        const index = shownMarkersCopy.indexOf(marker)
        if (index > -1) {
            shownMarkersCopy.splice(index, 1)
        }
    }

    shownMarkers = shownMarkersCopy
    shownMarkers.has = _has
}

async function addMarkerOnClick(event) {
    // don't add new marker if not left mouse button is pressed
    if (event.originalEvent.button != 0) return

    let latlng = event.latlng

    let point = new Point('', 0, latlng.lat, latlng.lng)
    let marker = new Marker(point)

    marker.draw()

    newId = await PointFetch.post(point)
    marker.update(newId)
}

async function drawAllPoints() {
    drawMarkers(pointsToMarkers(await PointFetch.getAll()));
}


let shownMarkers = []
function _has(id) {
    for (let i = 0; i < this.length; i++) {
        if (this[i].point.id == id) return true
    }

    return false
}

hideMarkers(shownMarkers)


map.on('click', addMarkerOnClick);
document.getElementById('showAllPoints').addEventListener('click', function(event) {
    event.preventDefault()
    drawAllPoints()
})
document.getElementById('hideAllPoints').addEventListener('click', function(event) {
    event.preventDefault()
    hideMarkers(shownMarkers)
})







// type Shape = []L.Geodesic
// shape functions

class PolygonFetch {
    static async getAll() {
        return await fetchData('/api/polygons')
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
    drawShapes(polygonsToShapes(await getAllPolygons()));
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




