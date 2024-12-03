var map = L.map('map').setView([55.76, 37.64], 5)
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy OpenStreetMap contributors',
    detectRetina: true
}).addTo(map)

const southWest = L.latLng(-89.98155760646617, -180)
const northEast = L.latLng(89.99346179538875, 180)
const bounds = L.latLngBounds(southWest, northEast)

map.setMaxBounds(bounds);
map.options.minZoom = 2

map.on('drag', function() {
    map.panInsideBounds(bounds, { animate: true });
});


const altIcon = L.icon({
    iconUrl: '/static/linemarker.png',
    iconSize: [8, 8]
})


class MapCallback {
    static current = undefined
    static default = undefined

    static set(callback) {
        map.off('click', MapCallback.current);
        MapCallback.current = callback
        map.on('click', MapCallback.current);
    }

    static assignDefault(callback) {
        MapCallback.default = callback
        if (MapCallback.current === undefined) {
            MapCallback.current = MapCallback.default
           map.on('click', MapCallback.current);
        }
    }

    static setDefault() {
        map.off('click', MapCallback.current);
        MapCallback.current = MapCallback.default
        map.on('click', MapCallback.current);
    }
}

async function addPointOnMapClick(event) {
    // don't add new mapPoint if not left mouse button is pressed
    if (event.originalEvent.button != 0) return

    let latlng = event.latlng

    let marker = new Marker(MarkerableTypes.Point, new Point({ name: '', id: 0, latitude: latlng.lat, longitude: latlng.lng}))
    let newId = await MarkerFetch.create(marker)

    marker.updateId(newId.id)
    marker.draw()

    return marker
}


function addPolygonPointOnAMapClick(event) {
    if (event.originalEvent.button != 0) return

    let latlng = event.latlng
    let accumulatedPoint = L.marker(latlng, { icon: altIcon });
    accumulatedPoint.addTo(map)
    polygonAccumulatedMarkers.push(accumulatedPoint)

    polygonAccumulatedPoints.push(new Point({ id: 0, name: '', latitude: latlng.lat, longitude: latlng.lng}))
}


function addLinePointOnMapClick(event) {
    if (event.originalEvent.button != 0) return

    let latlng = event.latlng
    let accumulatedPoint = L.marker(latlng, { icon: altIcon });
    accumulatedPoint.addTo(map)
    lineAccumulatedMarkers.push(accumulatedPoint)

    lineAccumulatedPoints.push(new Point({ id: 0, name: '', latitude: latlng.lat, longitude: latlng.lng}))
}


MapCallback.assignDefault(addPointOnMapClick)

