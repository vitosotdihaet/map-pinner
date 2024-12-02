Group.createNewGroup = async (name) => {
    let group = new Group({ name:name })

    let newId = await GroupFetch.create(group)
    group.id = newId

    Group.loaded.push(group)
    Group.populateOptions()

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
    Group.populateOptions()
}

Group.populateOptions = () => {
    groupSelect.options.length = 0
    Group.loaded.forEach(group => {
        const option = document.createElement('option');
        option.value = group.id;
        option.text = group.name;
        groupSelect.appendChild(option);
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
    const inputLength = newGroupNameInput.value.length;
    if (inputLength > 2 && inputLength < 256) {
        Group.createNewGroup(newGroupNameInput.value)
    } else {
        throw "you little bastard..."
    }
})


const groupSelect = document.getElementById("groupSelect");
groupSelect.addEventListener("change", () => {
    const selectedValue = groupSelect.value;
    console.log(selectedValue);
});

Group.reloadAll()
