<script lang="ts">
    export let dialog: HTMLDialogElement | undefined;
</script>

<dialog bind:this={dialog}>
    <slot></slot>
</dialog>

<style>
    dialog {
        background: var(--background-900);
        border: 2px solid var(--background-950);
        border-radius: 0.5em;
        color: var(--text-50);
    }
    dialog::backdrop {
        backdrop-filter: blur(2px) grayscale(0.5);
    }
</style>