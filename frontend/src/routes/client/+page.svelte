<script lang="ts">
    let roomCode = $state("")
    $effect(() => {
        roomCode = roomCode.toUpperCase()
    })

    async function handleSubmit(e: SubmitEvent) {
        e.preventDefault()
    }
</script>

<div id ="frame_main">
    <h1>Join Room</h1>
    <h4>ENTER CODE</h4>
    <form onsubmit={handleSubmit}>
        <input type="text" bind:value={roomCode} maxlength="4" autocomplete="off" autocorrect="off" required>
        <button type="submit">SUBMIT</button>
    </form>
</div>

<style>
    #frame_main {
        display: flex;
        flex-direction: column;
        align-items: center;
        margin-top: 10vh;
    }

    form {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    input[type=text] {
        background-color: black;
        border: 2px solid white;
        padding: 0.5rem;
        color: white;
        text-align: center;
        font-family: 'Courier New', Courier, monospace;
        font-size: 22pt;
        width: 4.5rem;
    }

    input[type=text]:focus {
        content: "";
        outline: 1px solid white;
    }

    button {
        margin-top: 1rem;
        padding: 0.5rem;
        font-family: Arial, Helvetica, sans-serif;
        font-size: 10pt;
        background-color: black;
        color: white;
        border: 2px solid white;
    }

    button:hover, button:focus {
        outline: 1px solid white;
    }

    button:active {
        background-color: white;
        color: black;
    }

    

</style>