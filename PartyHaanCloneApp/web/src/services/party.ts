const getList = () => {
    //TODO call get party list 
    return [{
        img: "https://www.scdn.co/i/_global/open-graph-default.png",
        name: "Spotify",
        totalMember: 5,
        currentMember: 0,
        id: 1,
    },
    {
        img: "https://www.scdn.co/i/_global/open-graph-default.png",
        name: "Spotify",
        totalMember: 5,
        currentMember: 0,
        id: 2,
    },
    {
        img: "https://www.scdn.co/i/_global/open-graph-default.png",
        name: "Spotify",
        totalMember: 5,
        currentMember: 0,
        id: 3
    }]
}

const create = (partyName: string, totalMember: number) => {
    //TODO call create party backend

}


export {
    getList,
    create,
}