<script lang="ts">
    import { page } from "$app/state";
    import Canvas from "$lib/elements/Canvas/Canvas.svelte";
    import { onDestroy } from "svelte";

    let connected = $state(false);

    let roomData = $state<Record<string, {x: number, y: number}>>({})

    let ws: WebSocket | null = null;

    function initConn() {
        if (ws) {
            return
        }
        ws = new WebSocket(`ws://symphoney.xyz:8080/ws/room/${page.params.room_id}/server/`)
        ws.addEventListener("open", () => {
            connected = true
        })
        ws.addEventListener("close", () => {
            connected = false
        })
        ws.addEventListener("message", (message: any) => {
            const data = JSON.parse(JSON.parse(message.data))
            const pos = JSON.parse(data.data)
            roomData[data.memberID] = {x: pos.x as number, y: pos.y as number}
            roomData = roomData
        })
    }

    onDestroy(() => {
        if (ws) {
            ws.close()
        }
    })
</script>

<h2>{page.params.room_id} server</h2>
<button onclick={initConn}>Init conn</button>

{#if connected}
<p>Connected</p>
{:else}
<p>Not connected</p>
{/if}
<Canvas roomData={roomData} mode="server" mouse={null}/>