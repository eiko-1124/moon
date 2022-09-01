export function formatTime(list) {
    for (let site of list) {
        const oldTime = new Date(site.time).getTime();
        site.time = timestampToTime(oldTime)
    }
}

function timestampToTime(timestamp) {
    var date = new Date(timestamp);//时间戳为10位需*1000，时间戳为13位的话不需乘1000
    var Y = date.getFullYear() + '-';
    var M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
    var D = date.getDate() + ' ';
    var h = date.getHours() + ':';
    var m = date.getMinutes() + ':';
    var s = date.getSeconds();
    return Y + M + D + h + m + s;
}

export function formatDate(list) {
    for (let site of list) {
        const oldTime = new Date(site.time).getTime();
        site.time = timestampToDate(oldTime)
    }
}

function timestampToDate(timestamp) {
    var date = new Date(timestamp);//时间戳为10位需*1000，时间戳为13位的话不需乘1000
    var Y = date.getFullYear() + '-';
    var M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1) + '-';
    var D = date.getDate();
    return Y + M + D;
}

export function formatTpye(list) {
    for (let site of list) {
        if (site.atype == 1) site.type = "技术"
        if (site.atype == 2) site.type = "杂文"
        if (site.atype == 3) site.type = "故事"
    }
}

export function formatTags(list) {
    for (let site of list) {
        site.tagArr = site.tags.split("#").filter((i) => i && i.trim())
    }
}

export function formatdays(list) {
    list.forEach(site => {
        const timestamp = new Date(site.time).getTime();
        const date = new Date(timestamp)
        const now = new Date()
        let time
        if (now.getFullYear() - date.getFullYear() >= 1) {
            time = now.getFullYear() - date.getFullYear()
            site.days = time + "年前"
        }
        else if (now.getMonth() - date.getMonth() >= 1) {
            time = now.getMonth() - date.getMonth()
            site.days = time + "个月前"
        } else if (now.getDate() - date.getDate() >= 1) {
            time = now.getDate() - date.getDate()
            site.days = time + "天前"
        } else if (now.getHours() - date.getHours() >= 1) {
            time = now.getHours() - date.getHours()
            site.days = time + "小时前"
        } else if (now.getMinutes() - date.getMinutes() >= 1) {
            time = now.getMinutes() - date.getMinutes()
            site.days = time + "分钟前"
        } else {
            time = now.getSeconds() - date.getSeconds()
            time = time == 0 ? 1 : time
            site.days = time + "秒前"
        }
    });
}