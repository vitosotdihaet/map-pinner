class Region {
    static loaded = []
    static currentRegion = null

    constructor(data) {
        this.id = data.id
        this.name = data.name
    }
}
