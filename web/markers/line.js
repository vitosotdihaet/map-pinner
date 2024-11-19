// class LineFetch {
//     static async getAll() {
//         return getData('/api/markers/lines')
//     }

//     static async create(line) {
//         return postData("/api/markers/lines", JSON.stringify(line))
//     }

//     static async delete(line) {
//         return deleteData(`/api/markers/lines/${line.id}`, "")
//     }

//     static async update(line) {
//         return putData(`/api/markers/lines/${line.id}`, JSON.stringify(line))
//     }
// }

class Line {
//     constructor(name, id, points) {
//         this.name = name
//         this.id = id
//         this.points = points
//     }
}

// class MapLine {
//     constructor(line) {
//         this.line = line
//         this.color = randomColor({ "luminosity": "bright", "hue": "red" })
//         this.setupMapLine()
//     }

//     updateGeodesic() {
//         let newLatLngs = [];
//         let newLinePoints = []

//         for (let marker of this.mapMarkers) {
//             let latlng = marker.getLatLng()
//             newLatLngs.push(latlng);
//             newLinePoints.push(new Point("", 0, latlng.lat, latlng.lng))
//         }

//         this.mapLine.setLatLngs(newLatLngs);
//         this.line.points = newLinePoints
//     }

//     setupMapLine() {
//         let coordinates = []
//         let latlngs = []
//         this.mapMarkers = []

//         this.line.points.forEach(point => {
//             coordinates.push([point.latitude, point.longitude])
//             let latlng = new L.LatLng(point.latitude, point.longitude)
//             latlngs.push(latlng)
//         })

//         this.mapLine = L.geodesic(coordinates, {
//             color: this.color,
//             weight: 5,
//             opacity: 0.75,
//             fillOpacity: 0.5
//         })

//         latlngs.forEach(place => {
//             var marker = L.marker(place, { draggable: true, icon: altIcon });

//             marker.on('drag', (_) => {
//                 this.updateGeodesic();
//             });
//             marker.on('dragend', (_) => {
//                 this.checkAndUpdate()
//             })
//             this.mapMarkers.push(marker);
//         })

//         this.mapLine.bindPopup(
//             L.popup().setContent(
//                 `
//                 <div class="popup">
//                     Name: <input type="text" class="popupNameInput" id="${this.line.id}" maxlength="255" size="10" value="${this.line.name}"/><br/>
//                 </div>
//                 <button class="popupDeleteButton" onclick="shownMapLines.get(${this.line.id}).delete()">Delete</button>
//                 <button class="popupUpdateButton" onclick="shownMapLines.get(${this.line.id}).checkAndUpdate()">Update</button>
//                 `
//             )
//         ).openPopup()
//     }

//     draw() {

//         if (shownMapLines.has(this.line.id)) return
//         this.mapMarkers.forEach(marker => { marker.addTo(map) })
//         this.mapLine.addTo(map)
//         shownMapLines.set(this.line.id, this)
//     }

//     checkAndUpdate() {
//         let name = undefined
//         let nameInput = document.getElementsByClassName('popupNameInput')

//         for (var i = 0; i < nameInput.length; i++) {
//             let element = nameInput[i]
//             if (element.id == this.line.id) {
//                 name = element.value
//             }
//         }

//         this.update({
//             name: name,
//             points: this.line.points
//         })
//     }

//     update(updateInfo) {
//         if (updateInfo.name !== undefined) {
//             this.line.name = updateInfo.name
//         }
//         if (updateInfo.latitude !== undefined) {
//             this.line.latitude = updateInfo.latitude
//         }
//         if (updateInfo.longitude !== undefined) {
//             this.line.longitude = updateInfo.longitude
//         }

//         updateInfo.id = this.line.id

//         this.hide()
//         this.setupMapLine()
//         this.draw()

//         LineFetch.update(updateInfo)
//     }

//     updateId(newId) {
//         this.line.id = newId
//         this.setupMapLine()
//     }

//     async delete() {
//         LineFetch.delete(this.line)
//         this.hide()
//     }

//     hide() {
//         map.removeLayer(this.mapLine)
//         this.mapMarkers.forEach(marker => { map.removeLayer(marker) })
//         shownMapLines.delete(this.line.id)
//     }
// }


// function linesToMapLines(lines) {
//     if (lines == null) { return [] }

//     let mapLines = []
//     lines.forEach(line => {
//         mapLines.push(new MapLine(line))
//     })

//     return mapLines
// }

// function drawMapLines(mapLines) {
//     mapLines.forEach(mapLine => {
//         mapLine.draw()
//     })
// }

// function hideMapLines(mapLines) {
//     mapLines.forEach(mapLine => {
//         mapLine.hide()
//     })
// }


// lineAccumulatedPoints = []
// lineAccumulatedMarkers = []

// // TODO move main logic to main
// function startNewLine(event) {
//     event.preventDefault()

//     for (var tab of tabs) {
//         tab.removeEventListener('click', openTab)
//     }

//     newLineButton.removeEventListener('click', startNewLine)
//     newLineButton.addEventListener('click', stopLine)
//     newLineButton.innerText = "Stop"

//     map.off('click', MapPoint.addMapPointOnMapClick)
//     map.on('click', newLinePointOnAMap)
// }

// function newLinePointOnAMap(event) {
//     // don't add new point if not left mouse button is pressed
//     if (event.originalEvent.button != 0) return
//     let latlng = event.latlng

//     let marker = L.marker(latlng, { icon: altIcon });
//     marker.addTo(map)
//     lineAccumulatedMarkers.push(marker)

//     let point = new Point('', 0, latlng.lat, latlng.lng)
//     lineAccumulatedPoints.push(point)
// }

// async function stopLine(event) {
//     event.preventDefault()

//     for (var tab of tabs) {
//         tab.addEventListener('click', openTab)
//     }

//     newLineButton.removeEventListener('click', stopLine)
//     newLineButton.addEventListener('click', startNewLine)
//     newLineButton.innerText = "Start a new line"

//     map.off('click', newLinePointOnAMap)
//     map.on('click', MapPoint.addMarkerOnMapClick)

//     if (lineAccumulatedPoints.length < 2) {
//         lineAccumulatedMarkers.forEach(marker => {
//             map.removeLayer(marker)
//         })
//         lineAccumulatedPoints = []
//         lineAccumulatedMarkers = []
//         return
//     }

//     let line = new Line("", 0, lineAccumulatedPoints)
//     let mapLine = new MapLine(line)

//     newId = await LineFetch.create(line)
//     mapLine.updateId(newId.id)
//     mapLine.draw()
    
//     lineAccumulatedMarkers.forEach(marker => {
//         map.removeLayer(marker)
//     })
//     lineAccumulatedPoints = []
//     lineAccumulatedMarkers = []
// }


// async function drawAllLines() {
//     drawMapLines(linesToMapLines(await LineFetch.getAll()));
// }



// let shownMapLines = new Map()
