class GroupFetch {
    static async getAll() {
        return getData('/api/groups')
    }

    static async create(group) {
        return postData('/api/groups', JSON.stringify(group))
    }

    static async delete(group) {
        return deleteData(`/api/groups/${group.id}`, '')
    }

    static async update(group) {
        return putData(`/api/groups/${group.id}`, JSON.stringify(group))
    }
}


class Group {
    static loaded = new Map()

    static async createNewGroup() {
        let group = new Group('', 0)

        let newId = await GroupFetch.create(group)
        group.id = newId

        Group.loaded.push(group)

        return group
    }

    static async reloadAll() {
        const groupsData = await GroupFetch.getAll();

        Group.loaded = new Map()
        groupsData.forEach(groupData => {
            Group.loaded.push(new Group(groupData))
        })

        Group.populateOptions()
    }

    static populateOptions() {
        const groupSelect = document.getElementById('groupSelect');

        Group.loaded.forEach(group => {
            const option = document.createElement('option');
            option.value = group.id;
            option.text = group.name;
            groupSelect.appendChild(option);
        });
    }

    constructor(id, name) {
        this.id = id
        this.name = name
    }
}


document.getElementById('createNewGroup').addEventListener('click', function(event) {
    Group.createNewGroup()
})

Group.reloadAll()
