export const actions = {
    create_room: async ({fetch}) => {
        console.log("Creating room")
        const res = await fetch("http://localhost:8080/room/", {
            method: "POST",
        })
        if (!res.ok) {
            console.log(res)
            return { roomID: null }
        }
        const data = await res.json()
        console.log(data)
        return {
            roomID: data.roomID || 'Unknown'
        }
    }
}