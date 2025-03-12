<script lang="ts">
    import { writable } from "svelte/store";
    import { page } from "$app/state";
    import Canvas from "$lib/elements/Canvas/Canvas.svelte";
    import type { PageProps } from "./$types";

    let connected = $state(false);
    let roomData = $state<Record<string, {x: number, y: number}>>({})
    let mouse = $state<{x: number, y: number}>({x: 0, y: 0})

    let { data }: PageProps = $props();

    type State = {
        requests: Array<Request>
    }
    export const msg = writable<State>({
        requests: [],
    })

    let ws: WebSocket | null = null;

    function initConn() {
        if (ws) {
            return
        }
        ws = new WebSocket(`ws://symphoney.xyz:8080/ws/room/${page.params.room_id}/client/`)
        ws.addEventListener("open", () => {
            connected = true
        })
        ws.addEventListener("close", () => {
            connected = false
        })
        ws.addEventListener("message", (message: any) => {
            // const data: Request = JSON.parse(message.data)
            // console.log(data)
            // msg.update((msg) => ({
            //     ...msg,
            //     requests: [data].concat(msg.requests),
            // }))
        })
    }

    // This isn't running!

    $effect(() => {
        const data = mouse
        if (ws) {
            const mouseData = JSON.stringify(data)
            ws.send(mouseData)
        }
    })
</script>

<h2>{data.roomID} client</h2>

<button onclick={initConn}>Init conn</button>

{#if connected}
<p>Connected</p>
{:else}
<p>Not connected</p>
{/if}

<Canvas roomData={roomData} mode="client" bind:mouse={mouse}/>