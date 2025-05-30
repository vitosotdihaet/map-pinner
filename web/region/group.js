Region.createNewRegion = async (name) => {
    let region = new Region({ name:name })

    let newId = await RegionFetch.create(region)
    region.id = newId.id

    if (Region.loaded.length == 0) {
        regionButtons.innerHTML = ''
    }
    Region.loaded.push(region)
    Region.addButton(region)

    return region
}

Region.reloadAll = async () => {
    Region.loaded = []

    regionData = null
    try {
        regionData = await RegionFetch.getAll();
    } catch (error) {
        console.error(`could not fetch regions: ${error}`)
    }

    if (regionData !== null) {
        regionData.forEach(regionData => {
            Region.loaded.push(new Region(regionData))
        })
    }

    Region.populateButtons()
}


const regionButtons = document.getElementById('regionButtons')
const noMapsLabel = document.getElementById('noMapsLabel')

Region.addButton = (region) => {
    const button = document.createElement('button');
    button.type = 'button';
    button.id = region.id;
    button.name = 'regionButton';
    button.textContent = region.name;

    button.addEventListener('click', () => {
        localStorage.setItem('region', JSON.stringify(region))
        window.location.href = '/static/map.html'
    });

    regionButtons.appendChild(button);
    regionButtons.appendChild(document.createElement('br'));
}

Region.populateButtons = async () => {
    regionButtons.innerHTML = ''

    if (Region.loaded.length == 0) {
        const noMapsLabel = document.createElement('p')
        noMapsLabel.id = 'noMaps'
        noMapsLabel.textContent = 'No maps in this group'
        regionButtons.appendChild(noMapsLabel)
    } else {
        Region.loaded.forEach((region) => {
            Region.addButton(region)
        });
    }
}


const newRegionNameInput = document.getElementById('regionName');
const newRegionButton = document.getElementById('createNewRegion');
newRegionButton.disabled = true;

newRegionNameInput.addEventListener('input', () => {
    if (!Role.hasAtLeastRole('editor')) { return }
    const inputLength = newRegionNameInput.value.length;
    if (inputLength >= 1 && inputLength <= 255) {
        newRegionButton.disabled = false;
    } else {
        newRegionButton.disabled = true;
    }
});

newRegionButton.addEventListener('click', function(event) {
    if (!Role.hasAtLeastRole('editor')) { return }
    const inputLength = newRegionNameInput.value.length;
    if (inputLength >= 1 && inputLength <= 255) {
        Region.createNewRegion(newRegionNameInput.value)
        newRegionNameInput.value = ''
    }
})


regionsDiv = document.getElementById('region')
function hideRegions() {
    regionsDiv.hidden = true
}

function unhideRegions() {
    regionsDiv.hidden = false
}


groupSelect.dispatchEvent(new Event('change'))


