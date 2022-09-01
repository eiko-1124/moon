export const state = () => {
    return {
        name: '',
        navStatus: false,
        info: {
            name: '',
            notice: '',
            illustate: '',
            sentence: '',
            qq: '',
            github: ''
        },
        signs: [],
        links: [],
        sum: {
            article: 0,
            mood: 0,
            link: 0,
            comment: 0,
            tag: 0
        },
        activeName: '',
        activeBtn: '',
        about: {
            rander: '',
            format: ''
        },
        tags: [],
        nearDate: [],
        nearUdiary: [],
        nearAdiary: [],
        nearState: 0,
        aDiary: [],
        uDiary: [],
        viewAlive: '',
        viewCache: []
    }
}