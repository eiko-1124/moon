export const requestSuccess = (config) => {
    let token = ''
    const cookies = document.cookie.split(";");
    for (let i = 0; i < cookies.length; i++) {
        const cookie = cookies[i].trim();
        if (cookie.indexOf("token=") == 0) {
            token = cookie.substring(6, cookie.length);
            break;
        }
    }
    config.headers.Authorization = "Bearer " + token
    return config
}