function quit() {
    localStorage.setItem('jwt', undefined)
    localStorage.setItem('user', undefined)
    localStorage.setItem('region', undefined)
    window.location.href = '/static/auth.html'
}