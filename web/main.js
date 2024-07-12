var map = L.map('map').setView([55.76, 37.64], 5)
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy OpenStreetMap contributors',
    // detectRetina: true
}).addTo(map)

var southWest = L.latLng(-89.98155760646617, -180),
northEast = L.latLng(89.99346179538875, 180);
var bounds = L.latLngBounds(southWest, northEast);

map.setMaxBounds(bounds);
map.options.minZoom = 2

map.on('drag', function() {
    map.panInsideBounds(bounds, { animate: false });
});


var altIcon = L.icon({
    iconUrl: '/static/notmarker.png',
    iconSize: [8, 8]
})