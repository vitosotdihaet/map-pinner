let allRoles

(async function initializeRoles() {
    allRoles = new Map(Object.entries(await RoleFetch.getAll()));
})();