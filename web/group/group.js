Group.createNewGroup = async (name) => {
    let group = new Group({ name:name })

    let newId = await GroupFetch.create(group)
    group.id = newId.id

    Group.loaded.push(group)
    Group.addOption(group)

    groupSelect.value = group.id
    groupSelect.dispatchEvent(new Event('change'))

    return group
}

Group.reloadAll = async () => {
    Group.loaded = []

    groupsData = null
    try {
        groupsData = await GroupFetch.getAll();
    } catch (error) {
        console.error(`${error}`)
    }

    if (groupsData !== null) {
        groupsData.forEach(groupData => {
            Group.loaded.push(new Group(groupData))
        })
    }

    Group.populateSelect()
}

Group.addOption = (group) => {
    const option = document.createElement('option');
    option.value = group.id;
    option.text = group.name;
    groupSelect.appendChild(option);
}

Group.populateSelect = () => {
    groupSelect.options.length = 0
    Group.loaded.forEach(group => {
        Group.addOption(group) 
    });
    groupSelect.dispatchEvent(new Event('change'))
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
        const newGroup = Group.createNewGroup(newGroupNameInput.value)
        groupSelect.value = newGroup.id;
        groupSelect.dispatchEvent(new Event('change'))
        newGroupNameInput.value = ''
    }
})


const groupSelect = document.getElementById("groupSelect");
groupSelect.addEventListener('change', async () => {
    const groupId = groupSelect.value

    if (groupId != '') {
        Group.currentGroup = new Group(await GroupFetch.getById(groupId))
        activateRegions()
        Region.reloadAll()
    } else {
        Group.currentGroup = null
        deactivateRegions()
        Region.reloadAll()
    }
});

Group.reloadAll()
