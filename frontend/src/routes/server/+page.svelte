<script lang="ts">
    
    import { PUBLIC_SERVER_ADDRESS } from '$env/static/public'
    let { form } = $props()

    // Handle Create Room
    async function handleSubmit(e: SubmitEvent) {
        e.preventDefault();

        const res = await fetch(`${PUBLIC_SERVER_ADDRESS}/room/`,
            {
                method: "POST",
            }
        )
    }

</script>

<h2>Server</h2>
<form onsubmit={handleSubmit}>
    <button type="submit">Create Room</button>
</form>
{#if form?.roomID}
<p>Room number: {form.roomID}</p>
{/if}
