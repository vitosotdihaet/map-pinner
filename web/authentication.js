const userNameInput = document.getElementById("userName")
const passwordInput = document.getElementById("password")

function inputsAreOk() {
    return !(
        (passwordInput.value.length < 8 || userNameInput.value.length < 8) ||
        (passwordInput.value.length > 72 || userNameInput.value.length > 32)
    )
}

function toggleButtons() {
    if (inputsAreOk()) {
        loginButton.disabled = false
        registerButton.disabled = false
    } else {
        loginButton.disabled = true
        registerButton.disabled = true
    }
}

userNameInput.addEventListener('input', toggleButtons)
passwordInput.addEventListener('input', toggleButtons)

const loginButton = document.getElementById("loginButton")
loginButton.addEventListener("click", async () => {
    const userName = userNameInput.value
    const password = passwordInput.value

    try {
        response = await UserFetch.getByUsernamePassword(userName, password);

        if (response.ok) {
            response = await response.json()
            localStorage.setItem("jwt", response.token)
            localStorage.setItem("user", JSON.stringify(response.user))
            window.location.href = "/static/group.html"
        } else {
            document.getElementById("errorMessage").style.display = "block"
        }
    } catch (error) {
        console.error("Error during login:", error)
    }
});


const registerButton = document.getElementById("registerButton")
registerButton.addEventListener("click", async () => {
    const userName = userNameInput.value
    const password = passwordInput.value

    try {
        const response = await UserFetch.createWithNamePassword(userName, password);

        if (response.id) {
            location.reload()
        } else {
            document.getElementById("errorMessage").style.display = "block"
        }
    } catch (error) {
        console.error("Error during registration:", error)
    }
});


loginButton.disabled = true
registerButton.disabled = true
