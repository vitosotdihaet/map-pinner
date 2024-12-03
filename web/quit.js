function quit() {
    localStorage.setItem('jwt', undefined)
    window.location.href = '/static/auth.html'
}