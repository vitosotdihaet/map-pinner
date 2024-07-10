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
                <button class="popupDeleteButton" onclick="deleteShape(shownShapes, ${this.polygon.id})">Delete</button>
                <button class="popupUpdateButton" onclick="updateShape(shownShapes, ${this.polygon.id})">Update</button>
                `
            )
        ).openPopup()
    }

    draw() {
        if (shownShapes.has(this.polygon.id)) return
        this.mapShape.addTo(map)
        shownShapes.push(this)
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
        const index = shownShapes.indexOf(this)
        if (index > -1) {
            shownShapes.splice(index, 1)
        }
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
    shapes.forEach(shape => {
        shape.draw()
    })
}

function hideShapes(shapes) {
    shapes.forEach(shape => {
        shape.hide()
    })
}

function deleteShape(shapes, id) {
    shapes.forEach(shape => {
        if (shape.polygon.id == id) {
            shape.delete()
            return
        }
    })
}

function updateShape(shapes, id) {
    let name = document.getElementsByClassName('popupNameInput')[0].value

    shapes.forEach(shape => {
        if (shape.polygon.id == id) {
            shape.update({
                name: name,
            })
            return
        }
    })
}


async function drawAllPolygons() {
    drawShapes(polygonsToShapes(await PolygonFetch.getAll()));
}



let shownShapes = []
shownShapes.has = function (id) {
    for (let i = 0; i < this.length; i++) {
        if (this[i].polygon.id == id) return true
    }
    return false
}
