<script lang="ts">

    import QrCreator from 'qr-creator';
    import { onMount } from 'svelte';
    
    // import { PUBLIC_SERVER_ADDRESS } from '$env/static/public'
    let { form } = $props()

    // Handle Create Room
    async function handleSubmit(e: SubmitEvent) {
        e.preventDefault();

        // const res = await fetch(`${PUBLIC_SERVER_ADDRESS}/room/`,
        //     {
        //         method: "POST",
        //     }
        // )
    }

    let qrDiv = $state<HTMLDivElement | null>(null);

    $effect(() => {
        if (qrDiv) {
            QrCreator.render({
                text: 'some text',
                radius: 0.1, // 0.0 to 0.5
                ecLevel: 'H', // L, M, Q, H
                fill: 'white', // foreground color
                background: null, // color or null for transparent
                size: 256 // in pixels
            }, qrDiv);
        }
        
    })

</script>

<h2>Server</h2>
<form onsubmit={handleSubmit}>
    <button type="submit">Create Room</button>
</form>
{#if form?.roomID}
<p>Room number: {form.roomID}</p>
{/if}

<div bind:this={qrDiv} id="qr_div"></div>

<style>
    #qr_div {
        
    }
</style>

