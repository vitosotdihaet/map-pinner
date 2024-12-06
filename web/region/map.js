regionInfo = localStorage.getItem('region')
if (regionInfo) {
    // TODO: check if region is still in the db
    Region.currentRegion = new Region(JSON.parse(regionInfo))
}

region = document.getElementById('region')
region.innerHTML = Region.currentRegion.name
