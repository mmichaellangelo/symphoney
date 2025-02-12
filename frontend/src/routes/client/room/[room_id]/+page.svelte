<script lang="ts">
    import { writable } from "svelte/store";
    import { page } from "$app/state";
    let connected = $state(false);

    type State = {
        requests: Array<Request>
    }
    export const msg = writable<State>({
        requests: [],
    })

    let ws: WebSocket | null = null;

    function initConn() {
        ws = new WebSocket(`ws://localhost:8080/ws/room/${page.params.room_id}/client/`)
        ws.addEventListener("open", () => {
            connected = true
        })
        ws.addEventListener("close", () => {
            connected = false
        })
        ws.addEventListener("message", (message: any) => {
            const data: Request = JSON.parse(message.data)
            console.log(data)
            msg.update((msg) => ({
                ...msg,
                requests: [data].concat(msg.requests),
            }))
        })
    }

    function handleInput(e: Event) {
        const slider = e.target as HTMLInputElement;
        console.log(slider.value)
        if (ws) {
            ws.send(slider.value)
        }
    }
</script>

<h2>{page.params.room_id} client</h2>

<button onclick={initConn}>Init conn</button>

<input type="range" min="1" max="100" value="50" class="slider" oninput={handleInput}>

{#if connected}
<p>Connected</p>
{:else}
<p>Not connected</p>
{/if}