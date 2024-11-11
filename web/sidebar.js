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
document.getElementById('showAllLines').addEventListener('click', function(event) {
    event.preventDefault()
    drawAllLines()
})

document.getElementById('hideAllLines').addEventListener('click', function(event) {
    event.preventDefault()
    hideShapes(shownDirections)
})

newLineButton = document.getElementById('newLine')
newLineButton.addEventListener('click', startNewLine)

