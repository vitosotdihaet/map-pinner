var map = L.map('map').setView([55.76, 37.64], 0)

L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy OpenStreetMap contributors'
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
        L.marker([point.latitude, point.longitude]).addTo(map)
    })
}

function drawPolygons(polygons) {
    if (polygons == null || polygons.length == 0) { return }

    polygons.forEach(polygon => {
        polygonPoints = polygon.Points

        let coordinates = []
        polygonPoints.forEach(point => {
            if (point != null) {
                coordinates.push(new L.LatLng(point.latitude, point.longitude))
            }
        })

        new L.Geodesic(coordinates, {
            color: randomColor({ "luminosity": "bright", "hue": "blue" }),
            weight: 3,
            opacity: 0.75,
            fillOpacity: 0.5
        }).addTo(map)
    })
}


let loadedPoints
let loadedPolygons

async function initializeMap() {
    loadedPoints = await getAllPoints();
    drawPoints(loadedPoints);

    loadedPolygons = await getAllPolygons();
    drawPolygons(loadedPolygons);
}

initializeMap();