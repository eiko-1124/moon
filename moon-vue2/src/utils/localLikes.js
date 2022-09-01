export function updateLike(id) {
    let likes = getLikes()
    const index = likes.indexOf('' + id)
    if (index == -1) {
        let local = localStorage.getItem('likes')
        local = local == null ? '' : local
        local += ('#' + id)
        localStorage.setItem('likes', local)
    }
    else {
        let local = ''
        likes.splice(index, 1)
        likes.forEach(site => {
            local += ("#" + site)
        })
        localStorage.setItem('likes', local)
    }
}

function getLikes() {
    const local = localStorage.getItem('likes')
    if (local == null) return []
    const likes = local.split("#").filter((i) => i && i.trim())
    return likes
}

export function formatLikes(list) {
    const likes = getLikes()
    list.forEach(site => {
        if (likes.indexOf('' + site.id) != -1) site.likeFlag = true
        else site.likeFlag = false
    })
}