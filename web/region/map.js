regionInfo = localStorage.getItem('region')
if (regionInfo != null) {
    Region.currentRegion = new Region(JSON.parse(regionInfo))
}

