import { getInfo, getSigns, getLinks, getArticleSum, getMoodSum, getLinkSum, getTagSum, getCommentSum, GetAbout, GetTags, GetNear, GetDiary } from "@/api/apis/user"

export const actions = {
    setInfo(contenxt) {
        getInfo(
            (Res) => {
                contenxt.commit('setInfo', Res)
            },
            (Err) => {
                console.log(Err);
            }
        );
    },
    setSigns(contenxt) {
        getSigns((Res) => {
            contenxt.commit('setSigns', Res)
        },
            (Err) => {
                console.log(Err);
            })
    },
    setLinks(contenxt) {
        getLinks((Res) => {
            contenxt.commit('setLinks', Res)
        }, (Err) => {
            console.log(Err)
        })
    },
    getSum(contenxt) {
        getArticleSum(Res => {
            contenxt.commit('setArticleSum', Res)
        }, Err => {
            console.log(Err)
        })
        getMoodSum(Res => {
            contenxt.commit('setMoodSum', Res)
        }, Err => {
            console.log(Err)
        }), getLinkSum(Res => {
            contenxt.commit('setLinkSum', Res)
        }, Err => {
            console.log(Err)
        }), getTagSum(Res => {
            contenxt.commit('setTagSum', Res)
        }, Err => {
            console.log(Err)
        }), getCommentSum(Res => {
            contenxt.commit('setCommentSum', Res)
        }, Err => {
            console.log(Err)
        })
    },
    getAbout(content) {
        GetAbout(Res => {
            content.commit('setRander', Res.rander)
            content.commit('setFormat', Res.format)
        }, Err => {
            console.log(Err)
        })
    },
    getTags(content) {
        GetTags(Res => {
            if (Res.res == 1) {
                content.commit('setTags', Res.tags)
            }
        }, Err => {
            console.log(Err)
        })
    },
    getNear(content) {
        GetNear(0, Res => {
            if (Res.res == 1) {
                content.commit('setNearDate', Res.near)
                content.commit('setNearAdiary', Res.near)
            }
        }, Err => {
            console.log(Err)
        })
        GetNear(1, Res => {
            if (Res.res == 1) {
                content.commit('setNearUdiary', Res.near)
            }
        }, Err => {
            console.log(Err)
        })
    },
    getDiary(content) {
        GetDiary(0, Res => {
            if (Res.res == 1) {
                content.commit('setAdiary', Res.diary)
            }
        }, Err => {
            console.log(Err)
        })
        GetDiary(1, Res => {
            if (Res.res == 1) {
                content.commit('setUdiary', Res.diary)
            }
        }, Err => {
            console.log(Err)
        })
    }
}