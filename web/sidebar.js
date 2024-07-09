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

document.getElementsByClassName("tablinks")[1].click();

// Markers
map.on('click', Marker.addMarkerOnMapClick);
document.getElementById('showAllPoints').addEventListener('click', function(event) {
    event.preventDefault()
    drawAllMarkers()
})
document.getElementById('hideAllPoints').addEventListener('click', function(event) {
    event.preventDefault()
    hideMarkers([...shownMarkers])
})

// Shapes
document.getElementById('showAllPolygons').addEventListener('click', function(event) {
    event.preventDefault()
    drawAllPolygons()
})

document.getElementById('hideAllPolygons').addEventListener('click', function(event) {
    event.preventDefault()
    hideShapes(shownShapes)
})
