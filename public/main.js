var map = L.map('map').setView([55.76, 37.64], 0);

L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; OpenStreetMap contributors'
}).addTo(map);

async function fetchGeospatialData(url) {
    const response = await fetch(url);
    return await response.json();
}


async function drawPoints() {
    const points = await fetchGeospatialData('/api/points');
    if (points == null || points.length == 0) { return }

    points.forEach(point => {
        L.marker([point.latitude, point.longitude]).addTo(map);
    });
}

async function drawPolygons() {
    const polygons = await fetchGeospatialData('/api/polygons');
    polygons.forEach(polygon => {
        polygonPoints = polygon.Points

        console.log(polygonPoints)

        let coordinates = []
        polygonPoints.forEach(point => {
            if (point != null) {
                console.log(point)
                coordinates.push(new L.LatLng(point.latitude, point.longitude))
            }
        })

        new L.Geodesic(coordinates, {
            color: randomColor({ "luminosity": "bright", "hue": "blue" }),
            weight: 3,
            opacity: 0.75,
            fillOpacity: 0.5
        }).addTo(map);
    })
}

drawPoints();
drawPolygons();
