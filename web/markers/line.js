class Line {
    constructor(data) {
        this.name = data.name
        this.id = data.id
        this.points = data.points
    }

    JSONify() {
        return JSON.stringify(Object.fromEntries(new Map([
            ['id', this.id],
            ['name', this.name],
            ['points', this.points],
        ])))
    }
    
    onPointDrag() {
        let newLatLngs = [];
        let newPoints = []
    
        for (let marker of this.pointMarkers) {
            let latlng = marker.getLatLng()
            newLatLngs.push(latlng);
            newPoints.push(new Point({ id: 0, name: '', latitude: latlng.lat, longitude: latlng.lng}))
        }
    
        this.lineMarker.setLatLngs(newLatLngs);
        this.points = newPoints
    }
    
    setupMarker() {
        this.color = randomColor({ "luminosity": "bright", "hue": "red" })

        let coordinates = []
        let latlngs = []
        this.pointMarkers = []
    
        this.points.forEach(point => {
            coordinates.push([point.latitude, point.longitude])
            latlngs.push(new L.LatLng(point.latitude, point.longitude))
        })
    
        this.lineMarker = L.geodesic(coordinates, {
            color: this.color,
            weight: 5,
            opacity: 0.75,
            fillOpacity: 0.5
        })
    
        latlngs.forEach(place => {
            var marker = L.marker(place, { draggable: true, icon: altIcon });
            marker.on('drag', (_) => { this.onPointDrag() })
            marker.on('dragend', (_) => { this.pullUpdate() })
            this.pointMarkers.push(marker);
        })
    
        this.lineMarker.bindPopup(
            L.popup().setContent(
                `
                <div class="popup">
                    Name: <input type="text" class="popupNameInput" id="${this.id}" maxlength="255" size="10" value="${this.name}"/><br/>
                </div>
                <button class="popupDeleteButton" onclick="Marker.shown.get(MarkerableTypes.Line).get(${this.id}).delete()">Delete</button>
                <button class="popupUpdateButton" onclick="Marker.shown.get(MarkerableTypes.Line).get(${this.id}).pullUpdate()">Update</button>
                `
            )
        ).openPopup()
    }
    
    pullUpdate() {
        let name = undefined
        let nameInput = document.getElementsByClassName('popupNameInput')
    
        for (var i = 0; i < nameInput.length; i++) {
            let element = nameInput[i]
            if (element.id == this.id) {
                name = element.value
            }
        }
    
        this.update({
            name: name,
            points: this.points
        })
    }
    
    update(updateInfo) {
        if (updateInfo.name !== undefined) {
            this.name = updateInfo.name
        }
        if (updateInfo.latitude !== undefined) {
            this.latitude = updateInfo.latitude
        }
        if (updateInfo.longitude !== undefined) {
            this.longitude = updateInfo.longitude
        }
    
        updateInfo.id = this.id
    
        this.hide()
        this.setupMarker()
        this.draw()
    
        MarkerFetch.updateType(MarkerableTypes.Line, updateInfo)
    }
    
    updateId(newId) {
        this.id = newId
        this.setupMarker()
    }

        
    draw() {
        this.setupMarker()
        this.pointMarkers.forEach(marker => { marker.addTo(map) })
        this.lineMarker.addTo(map)
    }

    hide() {
        map.removeLayer(this.lineMarker)
        this.pointMarkers.forEach(marker => { map.removeLayer(marker) })
    }
}



lineAccumulatedPoints = []
lineAccumulatedMarkers = []

// TODO move main logic to main
function startNewLine(event) {
    event.preventDefault()

    for (var tab of tabs) {
        tab.removeEventListener('click', openTab)
    }

    newLineButton.removeEventListener('click', startNewLine)
    newLineButton.addEventListener('click', stopLine)
    newLineButton.innerText = "Stop"

    map.off('click', MapPoint.addMapPointOnMapClick)
    map.on('click', newLinePointOnAMap)
}

function newLinePointOnAMap(event) {
    // don't add new point if not left mouse button is pressed
    if (event.originalEvent.button != 0) return
    let latlng = event.latlng

    let marker = L.marker(latlng, { icon: altIcon });
    marker.addTo(map)
    lineAccumulatedMarkers.push(marker)

    let point = new Point('', 0, latlng.lat, latlng.lng)
    lineAccumulatedPoints.push(point)
}

async function stopLine(event) {
    event.preventDefault()

    for (var tab of tabs) {
        tab.addEventListener('click', openTab)
    }

    newLineButton.removeEventListener('click', stopLine)
    newLineButton.addEventListener('click', startNewLine)
    newLineButton.innerText = "Start a new line"

    map.off('click', newLinePointOnAMap)
    map.on('click', MapPoint.addMarkerOnMapClick)

    if (lineAccumulatedPoints.length < 2) {
        lineAccumulatedMarkers.forEach(marker => {
            map.removeLayer(marker)
        })
        lineAccumulatedPoints = []
        lineAccumulatedMarkers = []
        return
    }

    let line = new Line("", 0, lineAccumulatedPoints)
    let mapLine = new MapLine(line)

    newId = await LineFetch.create(line)
    mapLine.updateId(newId.id)
    mapLine.draw()
    
    lineAccumulatedMarkers.forEach(marker => {
        map.removeLayer(marker)
    })
    lineAccumulatedPoints = []
    lineAccumulatedMarkers = []
}