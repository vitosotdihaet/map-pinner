function openTab(event) {
    var tabcontent = document.getElementsByClassName("tabcontent");
    for (var i = 0; i < tabcontent.length; i++) {
        tabcontent[i].style.display = "none";
    }

    document.getElementById(event.currentTarget.id + 's').style.display = "block";

    var tablinks = document.getElementsByClassName("tablinks");
    for (var i = 0; i < tablinks.length; i++) {
        tablinks[i].className = tablinks[i].className.replace(" active", "");
    }

    event.currentTarget.className += " active";
}

const tabs = document.getElementsByClassName("tablinks")
for (var tab of tabs) {
    tab.addEventListener('click', openTab)
}
tabs[0].click()


// Markers
map.on('click', Marker.addMarkerOnMapClick);
document.getElementById('showAllPoints').addEventListener('click', function(event) {
    event.preventDefault()
    drawAllMarkers()
})
document.getElementById('hideAllPoints').addEventListener('click', function(event) {
    event.preventDefault()
    hideMarkers(shownMarkers)
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

newPolygonButton = document.getElementById('newPolygon')
newPolygonButton.addEventListener('click', startNewPolygon)

// Directions
document.getElementById('showAllGraphs').addEventListener('click', function(event) {
    event.preventDefault()
    drawAllGraphs()
})

document.getElementById('hideAllGraphs').addEventListener('click', function(event) {
    event.preventDefault()
    hideShapes(shownDirections)
})

newGraphButton = document.getElementById('newGraph')
newGraphButton.addEventListener('click', startNewGraph)

