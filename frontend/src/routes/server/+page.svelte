<script lang="ts">
    import { enhance } from '$app/forms';

    let { data, form } = $props()
</script>

<h2>Server</h2>
<form method="POST" use:enhance>
    <button formaction="?/create_room">Create Room</button>
</form>
{#if form?.roomID}
<p>Room number: {form.roomID}</p>
{/if}
