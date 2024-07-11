class PolygonFetch {
    static async getAll() {
        return getData('/api/polygons')
    }

    static async create(polygon) {
        return postData("/api/polygons", JSON.stringify(polygon))
    }

    static async delete(polygon) {
        return deleteData(`/api/polygons/${polygon.id}`, "")
    }

    static async update(polygon) {
        return putData(`/api/polygons/${polygon.id}`, JSON.stringify(polygon))
    }
}

class Polygon {
    constructor(name, id, points) {
        this.name = name
        this.id = id
        this.points = points
    }
}

class Shape {
    constructor(polygon) {
        this.polygon = polygon
        this.setupMapShape()
    }

    setupMapShape() {
        let coordinates = []

        this.polygon.points.forEach(point => {
            coordinates.push([point.latitude, point.longitude])
        })

        this.mapShape = L.geodesic(coordinates, {
            color: randomColor({ "luminosity": "bright", "hue": "blue" }),
            weight: 5,
            opacity: 0.75,
            fillOpacity: 0.5
        })

        this.mapShape.bindPopup(
            L.popup().setContent(
                `
                <div class="popup">
                    ID: ${this.polygon.id}<br/>
                    Name: <input type="text" class="popupNameInput" id="${this.polygon.id}" maxlength="255" size="10" value="${this.polygon.name}"/><br/>
                </div>
                <button class="popupDeleteButton" onclick="shownShapes.get(${this.polygon.id}).delete()">Delete</button>
                <button class="popupUpdateButton" onclick="shownShapes.get(${this.polygon.id}).checkAndUpdate()">Update</button>
                <button class="popupChangeButton" onclick="shownShapes.get(${this.polygon.id}).change()">Change</button>
                `
            )
        ).openPopup()
    }

    draw() {
        if (shownShapes.has(this.polygon.id)) return
        this.mapShape.addTo(map)
        shownShapes.set(this.polygon.id, this)
    }

    checkAndUpdate() {
        let name = document.getElementsByClassName('popupNameInput')[0].value
        this.update({
            name: name,
        })
    }

    update(updateInfo) {
        if (updateInfo.name !== undefined) {
            this.polygon.name = updateInfo.name
        }
        if (updateInfo.latitude !== undefined) {
            this.polygon.latitude = updateInfo.latitude
        }
        if (updateInfo.longitude !== undefined) {
            this.polygon.longitude = updateInfo.longitude
        }

        updateInfo.id = this.polygon.id

        this.hide()
        this.setupMapShape()
        this.draw()

        PolygonFetch.update(updateInfo)
    }

    updateId(newId) {
        this.polygon.id = newId
        this.setupMapShape()
    }

    // static async addShapeOnMapClick(event) {
    //     // don't add new shape if not left mouse button is pressed
    //     if (event.originalEvent.button != 0) return

    //     let latlng = event.latlng

    //     let polygon = new Polygon('', 0, latlng.lat, latlng.lng)
    //     let shape = new Shape(polygon)

    //     let newId = await PolygonFetch.create(polygon)

    //     shape.updateId(newId.id)
    //     shape.draw()

    //     return shape
    // }

    async delete() {
        PolygonFetch.delete(this.polygon)
        this.hide()
    }

    hide() {
        map.removeLayer(this.mapShape)
        shownShapes[this.polygon.id].delete()
    }
}


function polygonsToShapes(polygons) {
    if (polygons == null) { return [] }

    let shapes = []
    polygons.forEach(polygon => {
        shapes.push(new Shape(polygon))
    })

    return shapes
}

function drawShapes(shapes) {
    shapes.forEach((shape, _) => {
        shape.draw()
    })
}

function hideShapes(shapes) {
    shapes.forEach((shape, _) => {
        shape.hide()
    })
}



async function drawAllPolygons() {
    drawShapes(polygonsToShapes(await PolygonFetch.getAll()));
}



let shownShapes = new Map()
