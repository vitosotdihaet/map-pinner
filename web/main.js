var map = L.map('map').setView([55.76, 37.64], 5)

L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy OpenStreetMap contributors',
    // detectRetina: true
}).addTo(map)


async function fetchData(url) {
    const response = await fetch(url)
    return await response.json()
}


async function getAllPoints() {
    return await fetchData('/api/points')
}

async function getAllPolygons() {
    return await fetchData('/api/polygons')
}

// type Marker = L.Marker
// marker functions
function pointsToMarkers(points) {
    if (points == null || points.length == 0) { return }

    let markers = []
    points.forEach(point => {
        markers.push(L.marker([point.latitude, point.longitude]))
    })

    return markers
}

function drawMarkers(markers) {
    markers.forEach(marker => {
        shownMarkers.push(marker)
        marker.addTo(map)
    })
}

function hideMarkers(markers) {
    let shownMarkersCopy = [...shownMarkers]

    for (let i = 0; i < markers.length; i++) {
        const marker = markers[i]
        map.removeLayer(marker)

        const index = shownMarkersCopy.indexOf(marker)
        if (index > -1) {
            shownMarkersCopy.splice(index, 1)
        }
    }

    shownMarkers = shownMarkersCopy
}

async function drawAllPoints() {
    drawMarkers(pointsToMarkers(await getAllPoints()));
}


// type Shape = []L.Geodesic
// shape functions
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



let shownMarkers = []
let shownShapes = []



document.getElementById('showAllPoints').addEventListener('click', function(event) {
    event.preventDefault()
    drawAllPoints()
})

document.getElementById('hideAllPoints').addEventListener('click', function(event) {
    event.preventDefault()
    hideMarkers(shownMarkers)
})


document.getElementById('showAllPolygons').addEventListener('click', function(event) {
    event.preventDefault()
    drawAllPolygons()
})

document.getElementById('hideAllPolygons').addEventListener('click', function(event) {
    event.preventDefault()
    hideShapes(shownShapes)
})


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
