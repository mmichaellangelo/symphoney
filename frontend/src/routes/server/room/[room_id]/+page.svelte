<script lang="ts">
    import { page } from "$app/state";
    import { writable } from "svelte/store";

    let connected = $state(false);

    let displayData = $state("data")

    type State = {
        requests: Array<Request>
    }
    export const msg = writable<State>({
        requests: [],
    })

    let ws: WebSocket | null = null;

    function initConn() {
        ws = new WebSocket(`ws://localhost:8080/ws/room/${page.params.room_id}/server/`)
        ws.addEventListener("open", () => {
            connected = true
        })
        ws.addEventListener("close", () => {
            connected = false
        })
        ws.addEventListener("message", (message: any) => {
            const data = JSON.parse(message.data)
            const dataParsed = JSON.parse(data)
            console.log(dataParsed.data)
            displayData = dataParsed.data
        })
    }
</script>

<h2>{page.params.room_id} server</h2>
<button onclick={initConn}>Init conn</button>

{#if connected}
<p>Connected</p>
{:else}
<p>Not connected</p>
{/if}

<p>{displayData}</p>