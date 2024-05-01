<script>
	import { onMount } from "svelte";

    let voices = $state([]);
    let voice = $state(null);
    let speed = $state(1);

    onMount(() => {
        voices = window.speechSynthesis.getVoices().filter(v => v.lang === 'en-US');
    });

    function test() {
        const utterance = new SpeechSynthesisUtterance("Moving head lights pointing forwards green, PAR lights off.");
        utterance.voice = voices[voice];
        utterance.rate = speed;
        window.speechSynthesis.speak(utterance);
    };
</script>

<select bind:value={voice}>
    {#each voices as v, i}
        <option value={i}>{v.name}</option>
    {/each}
</select>

<input type="range" min="0.1" max="2" step="0.1" bind:value={speed} />

<button onclick={test}>Test</button>