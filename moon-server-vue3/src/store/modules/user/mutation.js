export const mutations = {
    setName(state, name) {
        state.name = name
    },
    putNav(state) {
        state.navStatus = !state.navStatus
    },
    setInfo(state, info) {
        state.info = info
    },
    setInfoMessage(state, [key, value]) {
        state.info[key] = value
    },
    setSigns(state, signs) {
        state.signs = signs
    },
    addSign(state) {
        state.signs.push({
            id: '',
            content: ''
        })
    },
    setSign(state, [sign, index]) {
        state.signs[index].content = sign
    },
    setLinks(state, links) {
        state.links = links
    },
    setLink(state, [value, index, key]) {
        state.links[index][key] = value
    },
    addLink(state) {
        state.links.push({
            name: '',
            flink: ''
        })
    },
    setArticleSum(state, sum) {
        state.sum.article = sum
    },
    setMoodSum(state, sum) {
        state.sum.mood = sum
    },
    setLinkSum(state, sum) {
        state.sum.link = sum
    },
    setTagSum(state, sum) {
        state.sum.tag = sum
    },
    setCommentSum(state, sum) {
        state.sum.comment = sum
    }
    , setActiveName(state, value) {
        state.activeName = value
    }, setActiveBtn(state, value) {
        state.activeBtn = value
    },
    setRander(state, value) {
        state.about.rander = value
    },
    setFormat(state, value) {
        state.about.format = value
    },
    setTags(state, tags) {
        let tagArr = []
        for (let tag of tags) {
            tagArr.push({
                value: tag.num,
                name: tag.tag
            })
        }
        state.tags = tagArr
    },
    setNearDate(state, near) {
        for (let site of near) {
            state.nearDate.push(site.date)
        }
    },
    setNearAdiary(state, near) {
        for (let site of near) {
            state.nearAdiary.push(site.num)
        }
        if (state.nearAdiary.length == state.nearUdiary.length) state.nearState = 1
    },
    setNearUdiary(state, near) {
        for (let site of near) {
            state.nearUdiary.push(site.num)
        }
        if (state.nearAdiary.length == state.nearUdiary.length) state.nearState = 1
    },
    setNearState(state) {
        state.nearState = 0
    },
    setAdiary(state, diary) {
        state.aDiary = diary
    },
    setUdiary(state, diary) {
        state.uDiary = diary
    },
    setViewCache(state, [alike, cache]) {
        state.viewAlive = alike
        if (cache != null && alike != cache.name) {
            for (let i = 0; i < state.viewCache.length; i++) {
                if (state.viewCache[i].name == alike) {
                    state.viewCache.splice(i, 1)
                    break
                }
            }
            if (state.viewCache.length >= 5) { state.viewCache.shift() }
            state.viewCache.push(cache)
        }
    },
    delCache(state, name) {
        for (let i = 0; i < state.viewCache.length; i++) {
            if (state.viewCache[i].name == name) {
                state.viewCache.splice(i, 1)
                break
            }
        }
    }
}