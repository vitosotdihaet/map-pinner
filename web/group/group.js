Group.createNewGroup = async (name) => {
    let group = new Group({ name:name })

    let newId = await GroupFetch.create(group)
    group.id = newId.id

    Group.loaded.push(group)
    Group.addButton(group)

    return group
}

Group.reloadAll = async () => {
    try {
        const groupsData = await GroupFetch.getAll();
        if (groupsData === null) { return }

        Group.loaded = []
        
        groupsData.forEach(groupData => {
            Group.loaded.push(new Group(groupData))
        })
    } catch (error) {}
    Group.populateButtons()
}

Group.addButton = (group) => {
    const option = document.createElement('option');
    option.value = group.id;
    option.text = group.name;
    groupSelect.appendChild(option);
}

Group.populateButtons = () => {
    groupSelect.options.length = 1
    Group.loaded.forEach(group => {
        Group.addButton(group) 
    });
}



const newGroupNameInput = document.getElementById('groupName');
const newGroupButton = document.getElementById('createNewGroup');
newGroupButton.disabled = true;

newGroupNameInput.addEventListener('input', () => {
    const inputLength = newGroupNameInput.value.length;
    if (inputLength > 2 && inputLength < 256) {
        newGroupButton.disabled = false;
    } else {
        newGroupButton.disabled = true;
    }
});

newGroupButton.addEventListener('click', function(event) {
    event.preventDefault()
    const inputLength = newGroupNameInput.value.length;
    if (inputLength > 2 && inputLength < 256) {
        Group.createNewGroup(newGroupNameInput.value)
        newGroupNameInput.value = ''
    }
})


const groupSelect = document.getElementById("groupSelect");
groupSelect.addEventListener("change", async () => {
    const groupId = groupSelect.value;

    if (groupId == '') {
        Group.currentGroup = null
        deactivateRegions()
    } else {
        Group.currentGroup = new Group(await GroupFetch.getById(groupId))
        activateRegions()
        Region.reloadAll()
    }
});

Group.reloadAll()
