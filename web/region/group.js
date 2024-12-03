Region.createNewRegion = async (name) => {
    let region = new Region({ name:name })

    let newId = await RegionFetch.create(region)
    region.id = newId.id

    Region.loaded.push(region)
    Region.addButton(region)

    return region
}

Region.reloadAll = async () => {
    try {
        const regionData = await RegionFetch.getAll();
        if (regionData === null) { return }
        
        Region.loaded = []

        regionData.forEach(regionData => {
            Region.loaded.push(new Region(regionData))
        })
    } catch (error) {
        console.error('epic fail')
    }
    Region.populateButtons()
}

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

    regionButtonsContainer.appendChild(button);
    regionButtonsContainer.appendChild(document.createElement('br'));
}

Region.populateButtons = async () => {
    regionButtonsContainer.innerHTML = ''
    Region.loaded.forEach((region) => {
        Region.addButton(region)
    });
}


const newRegionNameInput = document.getElementById('regionName');
const newRegionButton = document.getElementById('createNewRegion');
newRegionButton.disabled = true;

newRegionNameInput.addEventListener('input', () => {
    const inputLength = newRegionNameInput.value.length;
    if (inputLength > 2 && inputLength < 256) {
        newRegionButton.disabled = false;
    } else {
        newRegionButton.disabled = true;
    }
});

newRegionButton.addEventListener('click', function(event) {
    event.preventDefault()
    const inputLength = newRegionNameInput.value.length;
    if (inputLength > 2 && inputLength < 256) {
        Region.createNewRegion(newRegionNameInput.value)
        newRegionNameInput.value = ''
    } else {
        throw "you little bastard..."
    }
})


const regionButtonsContainer = document.getElementById('regionButtons');

regionsDiv = document.getElementById('region')
function deactivateRegions() {
    regionsDiv.style.pointerEvents = 'none'
    regionsDiv.style.opacity = '0.5'
    regionsDiv.style.filter = 'grayscale(100%)'
}

function activateRegions() {
    regionsDiv.style.pointerEvents = 'auto'
    regionsDiv.style.opacity = '1'
    regionsDiv.style.filter = 'none'
}


deactivateRegions()
if (Region.currentGroup != null) {
    activateRegions()
    Region.reloadAll()
}


