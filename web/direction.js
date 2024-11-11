class LineFetch {
    static async getAll() {
        return getData('/api/markers/lines')
    }

    static async create(line) {
        return postData("/api/markers/lines", JSON.stringify(line))
    }

    static async delete(line) {
        return deleteData(`/api/markers/lines/${line.id}`, "")
    }

    static async update(line) {
        return putData(`/api/markers/lines/${line.id}`, JSON.stringify(line))
    }
}

class Line {
    constructor(name, id, points) {
        this.name = name
        this.id = id
        this.points = points
    }
}

class Direction {
    constructor(line) {
        this.line = line
        this.color = randomColor({ "luminosity": "bright", "hue": "red" })
        this.setupMapDirection()
    }

    updateGeodesic() {
        let newLatLngs = [];
        let newLinePoints = []

        for (let marker of this.mapMarkers) {
            let latlng = marker.getLatLng()
            newLatLngs.push(latlng);
            newLinePoints.push(new Point("", 0, latlng.lat, latlng.lng))
        }

        this.mapDirection.setLatLngs(newLatLngs);
        this.line.points = newLinePoints
    }

    setupMapDirection() {
        let coordinates = []
        let latlngs = []
        this.mapMarkers = []

        this.line.points.forEach(point => {
            coordinates.push([point.latitude, point.longitude])
            let latlng = new L.LatLng(point.latitude, point.longitude)
            latlngs.push(latlng)
        })

        this.mapDirection = L.geodesic(coordinates, {
            color: this.color,
            weight: 5,
            opacity: 0.75,
            fillOpacity: 0.5
        })

        latlngs.forEach(place => {
            var marker = L.marker(place, { draggable: true, icon: altIcon });

            marker.on('drag', (_) => {
                this.updateGeodesic();
            });
            marker.on('dragend', (_) => {
                this.checkAndUpdate()
            })
            this.mapMarkers.push(marker);
        })

        this.mapDirection.bindPopup(
            L.popup().setContent(
                `
                <div class="popup">
                    ID: ${this.line.id}<br/>
                    Name: <input type="text" class="popupNameInput" id="${this.line.id}" maxlength="255" size="10" value="${this.line.name}"/><br/>
                </div>
                <button class="popupDeleteButton" onclick="shownDirections.get(${this.line.id}).delete()">Delete</button>
                <button class="popupUpdateButton" onclick="shownDirections.get(${this.line.id}).checkAndUpdate()">Update</button>
                `
            )
        ).openPopup()
    }

    draw() {
        if (shownDirections.has(this.line.id)) return
        this.mapMarkers.forEach(marker => { marker.addTo(map) })
        this.mapDirection.addTo(map)
        shownDirections.set(this.line.id, this)
    }

    checkAndUpdate() {
        let name = undefined
        let nameInput = document.getElementsByClassName('popupNameInput')

        for (var i = 0; i < nameInput.length; i++) {
            let element = nameInput[i]
            if (element.id == this.line.id) {
                name = element.value
            }
        }

        this.update({
            name: name,
            points: this.line.points
        })
    }

    update(updateInfo) {
        if (updateInfo.name !== undefined) {
            this.line.name = updateInfo.name
        }
        if (updateInfo.latitude !== undefined) {
            this.line.latitude = updateInfo.latitude
        }
        if (updateInfo.longitude !== undefined) {
            this.line.longitude = updateInfo.longitude
        }

        updateInfo.id = this.line.id

        this.hide()
        this.setupMapDirection()
        this.draw()

        LineFetch.update(updateInfo)
    }

    updateId(newId) {
        this.line.id = newId
        this.setupMapDirection()
    }

    async delete() {
        LineFetch.delete(this.line)
        this.hide()
    }

    hide() {
        map.removeLayer(this.mapDirection)
        this.mapMarkers.forEach(marker => { map.removeLayer(marker) })
        shownDirections.delete(this.line.id)
    }
}


function linesToDirections(lines) {
    if (lines == null) { return [] }

    let directions = []
    lines.forEach(line => {
        directions.push(new Direction(line))
    })

    return directions
}

function drawDirections(directions) {
    directions.forEach(direction => {
        direction.draw()
    })
}

function hideDirections(directions) {
    directions.forEach(direction => {
        direction.hide()
    })
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

    map.off('click', Marker.addMarkerOnMapClick)
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
    map.on('click', Marker.addMarkerOnMapClick)

    if (lineAccumulatedPoints.length == 0) return

    let line = new Line("", 0, lineAccumulatedPoints)
    let direction = new Direction(line)

    newId = await LineFetch.create(line)
    direction.updateId(newId.id)
    direction.draw()
    
    lineAccumulatedMarkers.forEach(marker => {
        map.removeLayer(marker)
    })

    lineAccumulatedPoints = []
    lineAccumulatedMarkers = []
}


async function drawAllLines() {
    drawDirections(linesToDirections(await LineFetch.getAll()));
}



let shownDirections = new Map()
