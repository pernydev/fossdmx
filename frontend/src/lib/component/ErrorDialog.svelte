<script lang="ts">
	import { error } from "$lib/error.svelte";
	import Dialog from "./Dialog.svelte";

    $effect(() => {
        if (error.show) {
            dialog?.showModal();
        } else {
            dialog?.close();
        }
    });

    let dialog: HTMLDialogElement | undefined = $state();
</script>

<Dialog bind:dialog>
    <h2 class="mt-0">Error!</h2>
    <p>{error.message}</p>
    <code>{error.stack}</code>

    {#if error.reload_required}
        <button onclick={() => location.reload()}>Reload App</button>
    {/if}
    {#if error.restart_required}
        <button onclick={() => {}}>Restart Server</button>
    {/if}
</Dialog>