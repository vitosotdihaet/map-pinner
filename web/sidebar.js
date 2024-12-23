document.getElementById('showAllMarkers').addEventListener('click', function(event) {
    event.preventDefault()
    Marker.drawAll()
})

document.getElementById('hideAllMarkers').addEventListener('click', function(event) {
    event.preventDefault()
    Marker.hideAll()
})


newPolygonButton = document.getElementById('newPolygon')
newPolygonButton.addEventListener('click', startNewPolygon)

polygonAccumulatedPoints = []
polygonAccumulatedMarkers = []

function startNewPolygon(event) {
    event.preventDefault()

    newPolygonButton.removeEventListener('click', startNewPolygon)
    newPolygonButton.addEventListener('click', stopPolygon)
    newPolygonButton.innerText = 'Stop'

    MapCallback.set(addPolygonPointOnAMapClick)
}


async function stopPolygon(event) {
    event.preventDefault()

    newPolygonButton.removeEventListener('click', stopPolygon)
    newPolygonButton.addEventListener('click', startNewPolygon)
    newPolygonButton.innerText = 'Start a new polygon'

    MapCallback.setDefault()

    if (polygonAccumulatedPoints.length < 3) {
        polygonAccumulatedMarkers.forEach(marker => { map.removeLayer(marker) })
        polygonAccumulatedPoints = []
        polygonAccumulatedMarkers = []
        return
    }

    polygonAccumulatedPoints.push(polygonAccumulatedPoints[0])
    let marker = new Marker(MarkerableTypes.Polygon, new Polygon({ id: 0, name: '', points: polygonAccumulatedPoints }))

    newId = await MarkerFetch.create(marker)
    marker.updateId(newId.id)
    marker.draw()

    polygonAccumulatedMarkers.forEach(marker => { map.removeLayer(marker) })
    polygonAccumulatedPoints = []
    polygonAccumulatedMarkers = []
}


newLineButton = document.getElementById('newLine')
newLineButton.addEventListener('click', startNewLine)

lineAccumulatedPoints = []
lineAccumulatedMarkers = []

function startNewLine(event) {
    event.preventDefault()

    newLineButton.removeEventListener('click', startNewLine)
    newLineButton.addEventListener('click', stopLine)
    newLineButton.innerText = 'Stop'

    MapCallback.set(addLinePointOnMapClick)
}

async function stopLine(event) {
    event.preventDefault()

    newLineButton.removeEventListener('click', stopLine)
    newLineButton.addEventListener('click', startNewLine)
    newLineButton.innerText = 'Start a new line'

    MapCallback.setDefault()

    if (lineAccumulatedPoints.length < 2) {
        lineAccumulatedMarkers.forEach(marker => { map.removeLayer(marker) })
        lineAccumulatedPoints = []
        lineAccumulatedMarkers = []
        return
    }

    let marker = new Marker(MarkerableTypes.Line, new Line({ id: 0, name: '', points: lineAccumulatedPoints }))

    newId = await MarkerFetch.create(marker)
    marker.updateId(newId.id)
    marker.draw()
    
    lineAccumulatedMarkers.forEach(marker => { map.removeLayer(marker) })
    lineAccumulatedPoints = []
    lineAccumulatedMarkers = []
}
