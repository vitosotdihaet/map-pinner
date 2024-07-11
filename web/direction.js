class GraphFetch {
    static async getAll() {
        return getData('/api/graphs')
    }

    static async create(graph) {
        return postData("/api/graphs", JSON.stringify(graph))
    }

    static async delete(graph) {
        return deleteData(`/api/graphs/${graph.id}`, "")
    }

    static async update(graph) {
        return putData(`/api/graphs/${graph.id}`, JSON.stringify(graph))
    }
}

class Graph {
    constructor(name, id, points) {
        this.name = name
        this.id = id
        this.points = points
    }
}

class Direction {
    constructor(graph) {
        this.graph = graph
        this.color = randomColor({ "luminosity": "bright", "hue": "red" })
        this.setupMapDirection()
    }

    updateGeodesic() {
        let newLatLngs = [];
        let newGraphPoints = []

        for (let marker of this.mapMarkers) {
            let latlng = marker.getLatLng()
            newLatLngs.push(latlng);
            newGraphPoints.push(new Point("", 0, latlng.lat, latlng.lng))
        }

        this.mapDirection.setLatLngs(newLatLngs);
        this.graph.points = newGraphPoints
    }

    setupMapDirection() {
        let coordinates = []
        let latlngs = []
        this.mapMarkers = []

        this.graph.points.forEach(point => {
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
                    ID: ${this.graph.id}<br/>
                    Name: <input type="text" class="popupNameInput" id="${this.graph.id}" maxlength="255" size="10" value="${this.graph.name}"/><br/>
                </div>
                <button class="popupDeleteButton" onclick="shownDirections.get(${this.graph.id}).delete()">Delete</button>
                <button class="popupUpdateButton" onclick="shownDirections.get(${this.graph.id}).checkAndUpdate()">Update</button>
                `
            )
        ).openPopup()
    }

    draw() {
        if (shownDirections.has(this.graph.id)) return
        this.mapMarkers.forEach(marker => { marker.addTo(map) })
        this.mapDirection.addTo(map)
        shownDirections.set(this.graph.id, this)
    }

    checkAndUpdate() {
        let name = undefined
        let nameInput = document.getElementsByClassName('popupNameInput')

        for (var i = 0; i < nameInput.length; i++) {
            let element = nameInput[i]
            if (element.id == this.graph.id) {
                name = element.value
            }
        }

        this.update({
            name: name,
            points: this.graph.points
        })
    }

    update(updateInfo) {
        if (updateInfo.name !== undefined) {
            this.graph.name = updateInfo.name
        }
        if (updateInfo.latitude !== undefined) {
            this.graph.latitude = updateInfo.latitude
        }
        if (updateInfo.longitude !== undefined) {
            this.graph.longitude = updateInfo.longitude
        }

        updateInfo.id = this.graph.id

        this.hide()
        this.setupMapDirection()
        this.draw()

        GraphFetch.update(updateInfo)
    }

    updateId(newId) {
        this.graph.id = newId
        this.setupMapDirection()
    }

    async delete() {
        GraphFetch.delete(this.graph)
        this.hide()
    }

    hide() {
        map.removeLayer(this.mapDirection)
        this.mapMarkers.forEach(marker => { map.removeLayer(marker) })
        shownDirections.delete(this.graph.id)
    }
}


function graphsToDirections(graphs) {
    if (graphs == null) { return [] }

    let directions = []
    graphs.forEach(graph => {
        directions.push(new Direction(graph))
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


graphAccumulatedPoints = []
graphAccumulatedMarkers = []

// TODO move main logic to main
function startNewGraph(event) {
    event.preventDefault()

    for (var tab of tabs) {
        tab.removeEventListener('click', openTab)
    }

    newGraphButton.removeEventListener('click', startNewGraph)
    newGraphButton.addEventListener('click', stopGraph)
    newGraphButton.innerText = "Stop"

    map.off('click', Marker.addMarkerOnMapClick)
    map.on('click', newGraphPointOnAMap)
}

function newGraphPointOnAMap(event) {
    // don't add new point if not left mouse button is pressed
    if (event.originalEvent.button != 0) return
    let latlng = event.latlng

    let marker = L.marker(latlng, { icon: altIcon });
    marker.addTo(map)
    graphAccumulatedMarkers.push(marker)

    let point = new Point('', 0, latlng.lat, latlng.lng)
    graphAccumulatedPoints.push(point)
}

async function stopGraph(event) {
    event.preventDefault()

    for (var tab of tabs) {
        tab.addEventListener('click', openTab)
    }

    newGraphButton.removeEventListener('click', stopGraph)
    newGraphButton.addEventListener('click', startNewGraph)
    newGraphButton.innerText = "Start a new graph"

    map.off('click', newGraphPointOnAMap)
    map.on('click', Marker.addMarkerOnMapClick)

    if (graphAccumulatedPoints.length == 0) return

    let graph = new Graph("", 0, graphAccumulatedPoints)
    let direction = new Direction(graph)

    newId = await GraphFetch.create(graph)
    direction.updateId(newId.id)
    direction.draw()
    
    graphAccumulatedMarkers.forEach(marker => {
        map.removeLayer(marker)
    })

    graphAccumulatedPoints = []
    graphAccumulatedMarkers = []
}


async function drawAllGraphs() {
    drawDirections(graphsToDirections(await GraphFetch.getAll()));
}



let shownDirections = new Map()
