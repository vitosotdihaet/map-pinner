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


function drawPoints(points) {
    if (points == null || points.length == 0) { return }

    points.forEach(point => {
        let marker = [point.latitude, point.longitude]
        loadedMarkerPoints[L.marker(marker).addTo(map)] = marker
    })
}

function drawPolygons(polygons) {
    if (polygons == null || polygons.length == 0) { return }

    polygons.forEach(polygon => {
        polygonPoints = polygon.Points

        let coordinates = []
        polygonPoints.forEach(point => {
            if (point != null) {
                marker = new L.LatLng(point.latitude, point.longitude)
                coordinates.push(marker)
            }
        })

        loadedShapePolygons[coordinates] = polygonPoints

        new L.Geodesic(coordinates, {
            color: randomColor({ "luminosity": "bright", "hue": "blue" }),
            weight: 3,
            opacity: 0.75,
            fillOpacity: 0.5
        }).addTo(map)
    })
}


let loadedMarkerPoints = {}
let loadedShapePolygons = {}

async function showAllPoints() {
    loadedPoints = await getAllPoints();
    drawPoints(loadedPoints);
}

async function showAllPolygons() {
    loadedPolygons = await getAllPolygons();
    drawPolygons(loadedPolygons);
}


document.getElementById('showAllPoints').addEventListener('click', function(event) {
    event.preventDefault()
    showAllPoints()
})

document.getElementById('showAllPolygons').addEventListener('click', function(event) {
    event.preventDefault()
    showAllPolygons()
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
