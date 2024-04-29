<script lang="ts">
	import { goto } from "$app/navigation";
	import { onMount } from "svelte";

    onMount(async () => {
        const response = await fetch('/api/showfile');
        goto(response.ok ? '/show' : '/welcome');
    });
</script>