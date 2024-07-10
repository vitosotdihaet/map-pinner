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
        this.setupMapDirection()
    }

    setupMapDirection() {
        let coordinates = []

        this.graph.points.forEach(point => {
            coordinates.push([point.latitude, point.longitude])
        })

        this.mapDirection = L.geodesic(coordinates, {
            color: randomColor({ "luminosity": "bright", "hue": "red" }),
            weight: 5,
            opacity: 0.75,
            fillOpacity: 0.5
        })

        this.mapDirection.bindPopup(
            L.popup().setContent(
                `
                <div class="popup">
                    ID: ${this.graph.id}<br/>
                    Name: <input type="text" class="popupNameInput" id="${this.graph.id}" maxlength="255" size="10" value="${this.graph.name}"/><br/>
                </div>
                <button class="popupDeleteButton" onclick="deleteDirection(shownDirections, ${this.graph.id})">Delete</button>
                <button class="popupUpdateButton" onclick="updateDirection(shownDirections, ${this.graph.id})">Update</button>
                `
            )
        ).openPopup()
    }

    draw() {
        if (shownDirections.has(this.graph.id)) return
        this.mapDirection.addTo(map)
        shownDirections.push(this)
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

    // static async addDirectionOnMapClick(event) {
    //     // don't add new direction if not left mouse button is pressed
    //     if (event.originalEvent.button != 0) return

    //     let latlng = event.latlng

    //     let graph = new Graph('', 0, latlng.lat, latlng.lng)
    //     let direction = new Direction(graph)

    //     let newId = await GraphFetch.create(graph)

    //     direction.updateId(newId.id)
    //     direction.draw()

    //     return direction
    // }

    async delete() {
        GraphFetch.delete(this.graph)
        this.hide()
    }

    hide() {
        map.removeLayer(this.mapDirection)
        const index = shownDirections.indexOf(this)
        if (index > -1) {
            shownDirections.splice(index, 1)
        }
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

function deleteDirection(directions, id) {
    directions.forEach(direction => {
        if (direction.graph.id == id) {
            direction.delete()
            return
        }
    })
}

function updateDirection(directions, id) {
    let name = document.getElementsByClassName('popupNameInput')[0].value

    directions.forEach(direction => {
        if (direction.graph.id == id) {
            direction.update({
                name: name,
            })
            return
        }
    })
}


async function drawAllGraphs() {
    drawDirections(graphsToDirections(await GraphFetch.getAll()));
}



let shownDirections = []
shownDirections.has = function (id) {
    for (let i = 0; i < this.length; i++) {
        if (this[i].graph.id == id) return true
    }
    return false
}
