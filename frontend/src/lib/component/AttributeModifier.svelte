<script lang="ts">
    type Props = {
        attributes: FixtureOverwrite;
        features: FixtureFeature[];
    };
    let { attributes = $bindable(), features = $bindable() } = $props() as Props;
</script>

<p>Attributes</p>
{#if features.hasOwnProperty('dimmer')}
    <div>
        <label for="dimmer">Dimmer</label>
        <input type="range" id="dimmer" min="0" max="255" bind:value={attributes.dimmer} />
    </div>
{/if}
