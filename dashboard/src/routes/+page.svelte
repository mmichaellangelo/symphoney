<script lang="ts">
    import { onMount } from "svelte";
    import type { PageProps } from "./$types";
    import { invalidateAll } from "$app/navigation";

    /** Rate at which to fetch system stats */
    const UPDATE_TIMEOUT_MILLIS = 5000

    /** Page data */
    let { data }: PageProps = $props()

    /** True if fetching stats, otherwise false */
    let loading = $state(true)

    /** Periodically fetch stats */
    onMount(() => {
        loading = false

        const interval = setInterval(() => {
            loading = true
            invalidateAll()
                .then(() => loading = false)
        }, UPDATE_TIMEOUT_MILLIS)
        
        // Unmount >> cleanup
        return () => clearInterval(interval)
    })
</script>

<h2>Stats</h2>
<span>{loading ? "loading" : "---"}</span>
<p>{data.projectCount} projects</p>