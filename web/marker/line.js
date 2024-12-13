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
            let isEditor = Role.hasAtLeastRole('editor')
            var marker = L.marker(place, { draggable: isEditor, icon: altIcon });
            if (isEditor) {
                marker.on('drag', (_) => { this.onPointDrag() })
                marker.on('dragend', (_) => { this.pullUpdate() })
            }
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
        // TODO: prevent updating if not enough rights
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
