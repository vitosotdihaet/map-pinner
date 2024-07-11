var map = L.map('map').setView([55.76, 37.64], 5)

L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy OpenStreetMap contributors',
    // detectRetina: true
}).addTo(map)

var altIcon = L.icon({
    iconUrl: '/static/notmarker.png',
    iconSize: [8, 8]
})