<div id="frame_main">
    <img src="symphoney.png">
    <div id="frame_links">
        <a href="/server">Create Room</a>
        <a href="/client" id="join_link">Join Room</a>
    </div>
</div>

<style>
    #frame_main {
        width: 50vw;
        max-width: 500px;
        display: flex;
        flex-direction: column;
        align-items: center;
        margin-left: auto;
        margin-right: auto;
        margin-top: 20vh;
    }

    #frame_links {
        display: flex;
        flex-direction: row;
        margin-left: auto;
        margin-right: auto;
        
    }

    img {
        width: 100%;
    }

    a {
        color: white;
        text-decoration: none;
        text-shadow: 0px 0px 10px white;
        font-size: calc();
        padding: 1rem;
    }

    a:hover {
        text-shadow: 0px 0px 20px white;
    }

</style>