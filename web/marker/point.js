class Point {
    constructor(data) {
        this.id = data.id
        this.name = data.name
        this.latitude = data.latitude
        this.longitude = data.longitude
    }

    JSONify() {
        return JSON.stringify(Object.fromEntries(new Map([
            ['id', this.id],
            ['name', this.name],
            ['latitude', this.latitude],
            ['longitude', this.longitude],
        ])));
    }

    setupMarker() {
        this.marker = L.marker([this.latitude, this.longitude], { draggable: true })

        let lat, lng = (this.latitude, this.longitude)
        this.marker.on('drag', function (event) {
            // for some reason this shit works, though should be event.originalEvent.button
            if (event.originalEvent.buttons == 1) return
            event.target.dragging.disable()
            event.target.setLatLng([lat, lng])
            setTimeout(() => event.target.dragging.enable());
        })

        let marker = this.marker
        this.marker.on('dragend', function (event) {
            let latlng = event.target.getLatLng()
            marker.update({
                latitude: latlng.lat,
                longitude: latlng.lng
            })
        })

        this.marker.bindPopup(
            L.popup().setContent(
                `
                <div class="popup">
                    Name: <input type="text" class="popupNameInput" id="${this.id}" maxlength="255" size="10" value="${this.name}"/><br/>
                    Latitude: ${this.latitude.toFixed(4)}<br/>
                    Longitude: ${this.longitude.toFixed(4)}
                </div>
                <button class="popupDeleteButton" onclick="Marker.shown.get(MarkerableTypes.Point).get(${this.id}).delete()">Delete</button>
                <button class="popupUpdateButton" onclick="Marker.shown.get(MarkerableTypes.Point).get(${this.id}).pullUpdate()">Update</button>
                `
            )
        ).openPopup()
    }

    pullUpdate() {
        let name = document.getElementsByClassName('popupNameInput')[0].value
        let latlng = this.marker.getLatLng()

        this.update({
            name: name,
            latitude: latlng.lat,
            longitude: latlng.lng
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
        // update the popup with new info
        this.setupMarker()
        this.draw()

        MarkerFetch.updateType(MarkerableTypes.Point, updateInfo)
    }

    updateId(newId) {
        this.id = newId
        this.setupMarker()
    }

    draw() {
        this.setupMarker()
        this.marker.addTo(map)
    }

    hide() {
        if (this.marker !== undefined) {
            map.removeLayer(this.marker)
        }
    }
}
