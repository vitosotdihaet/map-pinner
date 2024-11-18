function openTab(event) {
    var tabcontent = document.getElementsByClassName("tabcontent");
    for (var i = 0; i < tabcontent.length; i++) {
        tabcontent[i].style.display = "none";
    }

    document.getElementById(event.currentTarget.id + 's').style.display = "block";

    var tablinks = document.getElementsByClassName("tablinks");
    for (var i = 0; i < tablinks.length; i++) {
        tablinks[i].className = tablinks[i].className.replace(" active", "");
    }

    event.currentTarget.className += " active";
}

const tabs = document.getElementsByClassName("tablinks")
for (var tab of tabs) {
    tab.addEventListener('click', openTab)
}
tabs[0].click()


// Markers
map.on('click', MapPoint.addMapPointOnMapClick);
document.getElementById('showAllPoints').addEventListener('click', function(event) {
    event.preventDefault()
    drawAllPoints()
})
document.getElementById('hideAllPoints').addEventListener('click', function(event) {
    event.preventDefault()
    hideMapPoints(shownMapPoints)
})

// Shapes
document.getElementById('showAllPolygons').addEventListener('click', function(event) {
    event.preventDefault()
    drawAllPolygons()
})

document.getElementById('hideAllPolygons').addEventListener('click', function(event) {
    event.preventDefault()
    hideMapPolygons(shownMapPolygons)
})

newPolygonButton = document.getElementById('newPolygon')
newPolygonButton.addEventListener('click', startNewPolygon)

// Directions
document.getElementById('showAllLines').addEventListener('click', function(event) {
    event.preventDefault()
    drawAllLines()
})

document.getElementById('hideAllLines').addEventListener('click', function(event) {
    event.preventDefault()
    hideMapPolygons(shownMapLines)
})

newLineButton = document.getElementById('newLine')
newLineButton.addEventListener('click', startNewLine)


//// Fetch and populate regions based on selected group
// async function loadRegions(groupId) {
//     const response = await fetch(`/api/groups/${groupId}/regions`); // Assuming an API endpoint to get regions by group ID
//     const regions = await response.json();
//     const regionSelect = document.getElementById('regionSelect');
//     regionSelect.innerHTML = '<option value="">Choose a region</option>';

//     regions.forEach(region => {
//         const option = document.createElement('option');
//         option.value = region.id;
//         option.text = region.name;
//         regionSelect.appendChild(option);
//     });

//     regionSelect.disabled = false;
// }

// // Event listeners for group and region selection
// document.getElementById('groupSelect').addEventListener('change', async (event) => {
//     const groupId = event.target.value;
//     if (groupId) {
//         await loadRegions(groupId);
//     } else {
//         document.getElementById('regionSelect').innerHTML = '<option value="">Choose a region</option>';
//         document.getElementById('regionSelect').disabled = true;
//     }
// });

// document.getElementById('regionSelect').addEventListener('change', (event) => {
//     const regionId = event.target.value;
//     if (regionId) {
//         openMapForRegion(regionId); // Function to display the map for the selected region
//     }
// });

loadGroups();
